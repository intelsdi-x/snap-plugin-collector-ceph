/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ceph

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/ctypes"
)

const (
	// parts of returned namescape
	ns_vendor = "intel"
	ns_class  = "storage"
	ns_type   = "ceph"

	// Plugin namespace prefix
	namespacePrefix = "/intel/storage/ceph"

	// Daemon name index in metrics namespace: [intel storage ceph daemonName]
	daemonNameIndex = 3
	daemonIDIndex   = 4

	// Default path to ceph executable
	//cephBinPathDefault = "/usr/bin/ceph"
	cephBinPathDefault = "/home/mkleina/ceph/src"

	// Default Ceph's socket config
	//socketPathDefault   = "/var/run/ceph"
	socketPathDefault = "/home/mkleina/ceph/src/out"
	//socketPrefixDefault = "ceph-"
	socketPrefixDefault = ""
	socketExtDefault    = "asok"
)

// Ceph
type Ceph struct {
	path        string // path to ceph executable
	keys        []core.Namespace
	daemons     []string
	socket      Socket
	initialized bool // after init() plugin with Config set true to avoid reinitalization
}

// Ceph Socket
type Socket struct {
	path   string
	prefix string
	ext    string
}

// PerfDumper interfaces needed for mocking exec command in ceph_test.go
type Command interface {
	perfDump(string, ...string) ([]byte, error)
	lookPath(string) (string, error)
}

type RealCmd struct{}

var cmd Command

// New() returns Snap-Plugin-Collector-Ceph instance
func New() *Ceph {
	ceph := &Ceph{initialized: false}
	cmd = &RealCmd{}

	return ceph
}

// Init() initalizes ceph plugin, gets information about running ceph-daemons and available metrics
func (ceph *Ceph) Init(config map[string]ctypes.ConfigValue) error {
	if ceph.initialized {
		return nil
	}

	// set ceph conf param
	ceph.path = getCephBinaryPath(config)
	ceph.socket = getCephSocketConf(config)

	// get ceph daemon names based on socket details
	ceph.daemons = ceph.socket.getCephDaemonNames()
	if len(ceph.daemons) <= 0 {
		return fmt.Errorf("Can not get Ceph Daemon Name(s) - check if any Ceph Daemon is running")
	}

	dkeys := make(map[string][]string)

	for _, daemon := range ceph.daemons {
		socket := filepath.Join(ceph.socket.path, daemon)

		// perf dump command is `/path/to/ceph/bin --admin-daemon /path/to/exemplary/socket/osd.0.asok perf dump`
		out, err := cmd.perfDump(filepath.Join(ceph.path, "ceph"), "--admin-daemon", socket, "perf", "dump")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error execution of ceph-daemon perf dump command for socket %+v, err=%+v\n",
				socket, err)
			return err
		}

		if keys, err := parsePerfDumpOut(out); err == nil {
			dkeys[daemon] = keys
		} else {
			fmt.Fprintf(os.Stderr, "Error parsing output of ceph-daemon perf dump command for socket %+v, err=%+v\n",
				socket, err)
		}

	}

	if len(dkeys) == 0 {
		return fmt.Errorf("No Ceph metrics available")
	}

	ceph.keys = []core.Namespace{}

	for _, daemon := range ceph.daemons {
		for _, key := range dkeys[daemon] {
			daemonName := trimPrefixAndSuffix(daemon, ceph.socket.prefix, "."+ceph.socket.ext)
			cephNs := daemonName + key
			ceph.keys = append(ceph.keys, createNamespace(cephNs))
		}
	}

	ceph.initialized = true
	return nil
}

// GetCephDaemonMetrics executes "ceph --admin-daemon perf dump" command for defined daemon-socket and returns its metrics
func (ceph *Ceph) GetCephDaemonMetrics(mts []plugin.MetricType, daemon string) ([]plugin.MetricType, error) {
	out, err := cmd.perfDump(filepath.Join(ceph.path, "ceph"), "--admin-daemon", filepath.Join(ceph.socket.path, daemon),
		"perf", "dump")
	timestamp := time.Now()
	hostname, _ := os.Hostname()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Ceph perf dump command execution failed for socket %+v, err=%+v\n",
			filepath.Join(ceph.socket.path, daemon), err)
		return nil, err
	}

	var jsonData map[string]interface{}

	if err := json.Unmarshal(out, &jsonData); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot unmarshal JSON object from ceph-daemon socket, err=%+v\n", err)
		return nil, err
	}

	daemonNameSplit := strings.Split(trimPrefixAndSuffix(daemon, ceph.socket.prefix, "."+ceph.socket.ext), ".")
	daemonName := daemonNameSplit[0]
	daemonID := daemonNameSplit[1]

	metrics := []plugin.MetricType{}

	for _, m := range mts {
		// Get metrics defined in task for this daemon
		if matchSlice(m.Namespace().Strings()[daemonNameIndex:daemonIDIndex+1], daemonNameSplit) {
			daemonMetrics := make(map[string]interface{})
			ceph.getJSONDataByNamespace(jsonData, m.Namespace().Strings()[daemonIDIndex+1:], []string{}, daemonMetrics)

			// No metrics found for desired namespace
			if len(daemonMetrics) == 0 {
				daemonMetrics[strings.Join(m.Namespace().Strings()[daemonIDIndex+1:], "/")] = nil
			}

			for ns, data := range daemonMetrics {
				metric := plugin.MetricType{
					Namespace_: core.NewNamespace(m.Namespace().Strings()[:daemonIDIndex]...).AddStaticElement(daemonID).AddStaticElements(strings.Split(ns, "/")...),
					Data_:      data, // get value of metric
					Tags_:      map[string]string{core.STD_TAG_PLUGIN_RUNNING_ON: hostname + "/" + daemonName + "." + daemonID},
					Timestamp_: timestamp,
				}

				metrics = append(metrics, metric)
			}
		}
	}

	// No metrics fount at all
	if len(metrics) == 0 {
		return metrics, errors.New("No metrics found!")
	}

	return metrics, nil
}

// matchSlice matches 2 slices with asterisk support
func matchSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] && a[i] != "*" && b[i] != "*" {
			return false
		}
	}

	return true
}

func (ceph *Ceph) getJSONDataByNamespace(data map[string]interface{}, namespace []string, resultNamespace []string, results map[string]interface{}) {
	// Go through all JSON data keys
	for key, _ := range data {

		// Convert ceph key to namespace slice for comparsion
		keyNs := strings.Split(key, ".")

		if matchSlice(namespace[:len(keyNs)], keyNs) {
			if reflect.ValueOf(data[key]).Kind() == reflect.Map {
				// Go deeper into JSON structure
				ceph.getJSONDataByNamespace(data[key].(map[string]interface{}), namespace[len(keyNs):], append(resultNamespace, keyNs...), results)
			} else {
				results[strings.Join(resultNamespace, "/")+"/"+key] = data[key]
			}
		}
	}
}

// CollectMetrics returns all desired Ceph metrics defined in task manifest
func (ceph *Ceph) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {
	if len(mts) <= 0 {
		return nil, errors.New("No metrics defined to collect")
	}
	metrics := []plugin.MetricType{}

	// init ceph plugin with Config settings (only once)
	if err := ceph.Init(mts[0].Config().Table()); err != nil {
		return nil, err
	}

	for _, daemon := range ceph.daemons {
		if dmetric, err := ceph.GetCephDaemonMetrics(mts, daemon); err == nil {
			metrics = append(metrics, dmetric...)
		}
	}

	return metrics, nil
}

// GetMetricTypes returns the metric types exposed by ceph-daemon sockets
func (ceph *Ceph) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
	mts := []plugin.MetricType{}
	for _, metricMeta := range allMetrics {
		mts = append(mts, plugin.MetricType{Namespace_: createNamespace(metricMeta.ns)})
	}

	return mts, nil
}

// GetConfigPolicy returns a ConfigPolicy
func (ceph *Ceph) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	p := cpolicy.New()

	return p, nil
}

// getCephBinaryPath returns path to ceph executable
func getCephBinaryPath(config map[string]ctypes.ConfigValue) string {
	// check global config as a first
	if path, ok := config["path"]; ok {
		return path.(ctypes.ConfigValueStr).Value
	}

	// check PATH environment variable
	if path, err := cmd.lookPath("ceph"); err == nil {
		//command "LookPath" resolves the path to a complete name, so the "ceph" suffix needs to be trimmed
		return strings.TrimSuffix(path, "/ceph")
	}

	return cephBinPathDefault
}

// getCephSocketConf returns path to folder contains daemon-sockets, prefix and extension of socket's name
func getCephSocketConf(config map[string]ctypes.ConfigValue) Socket {
	s := Socket{}

	// set path to socket, defaults to "/var/run/ceph"
	if path, ok := config["socket_path"]; ok {
		s.path = path.(ctypes.ConfigValueStr).Value
	} else {
		s.path = socketPathDefault
	}

	// set socket prefix, defaults to "ceph-"
	if prefix, ok := config["socket_prefix"]; ok {
		s.prefix = prefix.(ctypes.ConfigValueStr).Value
		// if equals "none", set empty socket prefix
		if strings.ToLower(s.prefix) == "none" {
			s.prefix = ""
		}
	} else {
		s.prefix = socketPrefixDefault
	}

	// set socket extension, defaults to "asok"
	if ext, ok := config["socket_ext"]; ok {
		s.ext = ext.(ctypes.ConfigValueStr).Value
	} else {
		s.ext = socketExtDefault
	}

	return s
}

// createNamespace returns namespace slice of strings composed from: vendor, class, type, ceph daemon name and ceph daemon id
func createNamespace(ns string) core.Namespace {
	result := core.NewNamespace()
	nsSplit := strings.Split(strings.TrimPrefix(ns, "/"), "/")

	for i, nsEntry := range nsSplit {
		if nsEntry == "*" {
			result = result.AddDynamicElement(nsSplit[i-1]+"_id", "ID of "+nsSplit[i-1])
		} else {
			result = result.AddStaticElement(nsEntry)
		}
	}
	return result
}

// getCephDaemonNames scans the path to ceph sockets in search of an instance which name contains specified prefix and extension.
// Returns names of available ceph-daemon sockets
func (s *Socket) getCephDaemonNames() []string {
	files, err := ioutil.ReadDir(s.path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Not found any ceph-daemon socket in %+v. Check if the given path is correct\n", s.path)
		panic(err)
	}
	names := []string{}
	suffix := "." + s.ext
	for _, f := range files {
		// read names of all items in socket path with specific prefix-name (defaults to "ceph-") and extension (defaults to .asok)
		if strings.HasPrefix(f.Name(), s.prefix) && strings.HasSuffix(f.Name(), suffix) {
			names = append(names, f.Name())
		}
	}

	return names
}

func (c *RealCmd) lookPath(file string) (string, error) {
	return exec.LookPath(file)
}

// parseMapToNamespace returns a slice contains metrics's namespace. As input there is map of strings to arbitrary data and current
// prefix of namespace
func parseMapToNamespace(stats map[string]interface{}, prefix string, out *[]string) {
	for key, val := range stats {
		switch reflect.TypeOf(val).Kind() {
		case reflect.Map:
			// get the next level of nested map
			c := filepath.Join(prefix, key)
			parseMapToNamespace(val.(map[string]interface{}), c, out)
			break

		default:
			// append namespace to output
			c := filepath.Join(prefix, key)
			*out = append(*out, c)
			break
		}
	}
}

// parsePerfDumpOut parses output of ceph-daemon perf dump command and returns string with name of available metrics
func parsePerfDumpOut(cmdOut []byte) ([]string, error) {
	prefix := "/"
	keys := []string{}

	var stats map[string]interface{}

	if err := json.Unmarshal(cmdOut, &stats); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot unmarshal JSON object from ceph-daemon socket, err=%+v\n", err)
		return nil, err
	}
	parseMapToNamespace(stats, prefix, &keys)

	return keys, nil
}

// execute command
func (c *RealCmd) perfDump(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}

// trimPrefixAndSuffix returns 's' without the provided prefix and suffix strings.
// If 's' neither starts with prefix, nor ends with suffix, 's' is returned unchanged.
func trimPrefixAndSuffix(s string, prefix string, suffix string) string {
	s = strings.TrimPrefix(s, prefix)
	s = strings.TrimSuffix(s, suffix)
	return s
}

// assignMetricMeta assigs metadata to metric using predefined metadata slice
func assignMetricMeta(mt *plugin.MetricType, allMetrics []metric) {
	for _, metricMeta := range allMetrics {
		fmt.Println("Match ", metricMeta.ns, "with", mt.Namespace().String())
		if matchSlice(strings.Split(metricMeta.ns, "/"), strings.Split(mt.Namespace().String(), "/")) {
			mt.Description_ = metricMeta.description
			mt.Unit_ = metricMeta.unit
			break
		}
	}
}
