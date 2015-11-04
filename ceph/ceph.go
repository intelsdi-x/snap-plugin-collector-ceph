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

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/control/plugin/cpolicy"
	"github.com/intelsdi-x/pulse/core/ctypes"
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

	// Default path to ceph executable
	cephBinPathDefault = "/usr/bin/ceph"

	// Default Ceph's socket config
	socketPathDefault   = "/var/run/ceph"
	socketPrefixDefault = "ceph-"
	socketExtDefault    = "asok"
)

// Ceph
type Ceph struct {
	path        string // path to ceph executable
	keys        []string
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

// execute command
func (c *RealCmd) perfDump(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}

func (c *RealCmd) lookPath(file string) (string, error) {
	return exec.LookPath(file)
}

// trimPrefixAndSuffix returns 's' without the provided prefix and suffix strings.
// If 's' neither starts with prefix, nor ends with suffix, 's' is returned unchanged.
func trimPrefixAndSuffix(s string, prefix string, suffix string) string {
	s = strings.TrimPrefix(s, prefix)
	s = strings.TrimSuffix(s, suffix)
	return s
}

// getCephDaemonMetrics executes "ceph --admin-daemon perf dump" command for defined daemon-socket and returns its metrics
func (c *Ceph) getCephDaemonMetrics(mts []plugin.PluginMetricType, daemon string) ([]plugin.PluginMetricType, error) {
	out, err := cmd.perfDump(filepath.Join(c.path, "ceph"), "--admin-daemon", filepath.Join(c.socket.path, daemon),
		"perf", "dump")
	timestamp := time.Now()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Ceph perf dump command execution failed for socket %+v, err=%+v\n",
			filepath.Join(c.socket.path, daemon), err)
		return nil, err
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(out, &dat); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot unmarshal JSON object from ceph-daemon socket, err=%+v\n", err)
		return nil, err
	}

	metrics := []plugin.PluginMetricType{}

	for _, m := range mts {
		nlen := len(m.Namespace())
		daemonName := trimPrefixAndSuffix(daemon, c.socket.prefix, "."+c.socket.ext)

		// compare daemonName with proper component of metric's namespace [intel storage ceph daemonName]
		if daemonName == m.Namespace()[daemonNameIndex] {
			// get metrics defined in task for this daemon
			dat_r := dat
			// skip the const components of metrics namespace
			for _, name := range m.Namespace()[daemonNameIndex+1 : nlen-1] {
				if dat_r[name] == nil {
					break
				}
				dat_r = dat_r[name].(map[string]interface{}) // get metric
			}

			hostname, _ := os.Hostname()
			metric := plugin.PluginMetricType{
				Namespace_: m.Namespace(),
				Data_:      dat_r[m.Namespace()[nlen-1]], // get value of metric
				Source_:    hostname + "/" + daemonName,
				Timestamp_: timestamp,
			}

			metrics = append(metrics, metric)
		}
	}

	if len(metrics) <= 0 {
		// do not find desired metrics
		return nil, errors.New("Do not find desired metrics for ceph-deamon")
	}

	return metrics, nil
}

// CollectMetrics returns all desired Ceph metrics defined in task manifest
func (ceph *Ceph) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {

	if len(mts) <= 0 {
		return nil, errors.New("No metrics defined to collect")
	}
	metrics := []plugin.PluginMetricType{}

	// init ceph plugin with Config settings (only once)
	if ceph.initialized == false {
		if err := ceph.init(mts[0].Config().Table()); err != nil {
			return nil, err
		}
		ceph.initialized = true
	}

	for _, daemon := range ceph.daemons {
		if dmetric, err := ceph.getCephDaemonMetrics(mts, daemon); err == nil {
			metrics = append(metrics, dmetric...)
		}
	}

	return metrics, nil
}

// GetMetricTypes returns the metric types exposed by ceph-daemon sockets
func (ceph *Ceph) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	// init ceph plugin with Global Config params
	if err := ceph.init(cfg.Table()); err != nil {
		return nil, err
	}

	mts := make([]plugin.PluginMetricType, len(ceph.keys))
	for i, k := range ceph.keys {
		mts[i] = plugin.PluginMetricType{Namespace_: strings.Split(strings.TrimPrefix(k, "/"), "/")}
	}

	return mts, nil
}

// GetConfigPolicy returns a ConfigPolicy
func (c *Ceph) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
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

// Init() initalizes ceph plugin, gets information about running ceph-daemons and available metrics
func (ceph *Ceph) init(config map[string]ctypes.ConfigValue) error {
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

	ceph.keys = []string{}

	for _, daemon := range ceph.daemons {
		for _, key := range dkeys[daemon] {
			daemonName := trimPrefixAndSuffix(daemon, ceph.socket.prefix, "."+ceph.socket.ext)
			ceph.keys = append(ceph.keys, strings.Join(createNamespace(daemonName+key), "/"))
		}
	}
	return nil
}

// New() returns Pulse-Plugin-Collector-Ceph instance
func New() *Ceph {
	ceph := &Ceph{initialized: false}
	cmd = &RealCmd{}

	return ceph
}

// createNamespace returns namespace slice of strings composed from: vendor, class, type and ceph-daemon name
func createNamespace(name string) []string {
	return []string{ns_vendor, ns_class, ns_type, name}
}
