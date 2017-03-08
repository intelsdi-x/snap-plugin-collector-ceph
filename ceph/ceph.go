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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	// parts of returned namescape
	nsVendor = "intel"
	nsClass  = "storage"
	nsType   = "ceph"

	// Daemon name index in metrics namespace: [intel storage ceph daemonName]
	daemonNameIndex = 3
	daemonIDIndex   = 4

	// Default path to ceph executable
	cephBinPathDefault = "/usr/bin"

	// Default Ceph's socket config
	socketPathDefault   = "/var/run/ceph"
	socketPrefixDefault = "ceph-"
	socketExtDefault    = "asok"

	// Data types for JSON walker function
	jsonData   = 0
	jsonSchema = 1

	// Ceph related
	cephTypeFloat   uint8 = 0
	cephTypeUint64  uint8 = 1
	cephTypeAverage uint8 = 2
)

// Ceph
type Ceph struct {
	path        string // path to ceph executable
	daemons     []Daemon
	socket      Socket
	initialized bool                              // after init() plugin with Config set true to avoid reinitalization
	schema      map[string]map[string]interface{} // schema for each daemon name
}

// Ceph Socket
type Socket struct {
	path   string
	prefix string
	ext    string
}

// Ceph daemon
type Daemon struct {
	id       string
	name     string
	fullName string
}

// PerfDumper interfaces needed for mocking exec command in ceph_test.go
type Command interface {
	perfDump(string, string) ([]byte, error)
	perfSchema(string, string) ([]byte, error)
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

// init() initalizes ceph plugin, gets information about running ceph-daemons and available metrics
func (ceph *Ceph) init(config plugin.Config) error {
	if ceph.initialized {
		return nil
	}

	var err error

	// set ceph conf param
	ceph.path, err = getCephBinaryPath(config)
	if err != nil {
		return err
	}
	ceph.socket, err = getCephSocketConf(config)
	if err != nil {
		return err
	}

	// get ceph daemon names based on socket details
	ceph.daemons, err = ceph.socket.getCephDaemons()
	if err != nil {
		return err
	}
	if len(ceph.daemons) == 0 {
		return fmt.Errorf("Can not get Ceph Daemon Name(s) - check if any Ceph Daemon is running")
	}

	// Get all daemon schemas
	for _, d := range ceph.daemons {
		ceph.getCephDaemonSchema(d)
	}

	ceph.initialized = true
	return nil
}

func (ceph *Ceph) getCephDaemonSchema(daemon Daemon) error {
	// Create ceph schema cache if not exists
	if ceph.schema == nil {
		ceph.schema = make(map[string]map[string]interface{})
	}

	// Save Ceph perf schema in cache
	if ceph.schema[daemon.name] == nil {

		ceph.schema[daemon.name] = make(map[string]interface{})

		out, err := cmd.perfSchema(filepath.Join(ceph.path, "ceph"), filepath.Join(ceph.socket.path, daemon.fullName))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ceph perf schema command execution failed for socket %+v, err=%+v\n",
				filepath.Join(ceph.socket.path, daemon.fullName), err)
			return err
		}
		perfSchemaData := make(map[string]interface{})
		if err = decodeJSON(out, &perfSchemaData); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot unmarshal JSON object from ceph-daemon socket schema, err=%+v\n", err)
			return err
		}
		ceph.getJSONPathsByNamespace(perfSchemaData, jsonSchema, []string{}, []string{}, ceph.schema[daemon.name])
	}

	return nil
}

// getDaemon returns daemon using name and id
func (ceph *Ceph) getDaemon(name, id string) *Daemon {
	for _, d := range ceph.daemons {
		if d.name == name && d.id == id {
			return &d
		}
	}
	return nil
}

// decodeJSON decodes JSON string
func decodeJSON(data []byte, out interface{}) error {
	jd := json.NewDecoder(bytes.NewReader(data))

	// Use Number instead float64 by default
	jd.UseNumber()

	return jd.Decode(&out)
}

// getCephDaemonMetrics executes "ceph --admin-daemon perf dump" command for defined daemon-socket and returns metrics matching with mts
// If provided metric list is nil, function returns all available Ceph metrics using "ceph --admin-daemon perf schema" command
// Delta metrics (sum / avgcount) are calculated automatically
func (ceph *Ceph) getCephDaemonMetrics(mts []plugin.Metric, daemon Daemon) ([]plugin.Metric, error) {
	var perfDumpData map[string]interface{}
	metricsCollected := []plugin.Metric{}

	// Go throught schema and build metric list
	for schemaNs, schemaData := range ceph.schema[daemon.name] {
		description := ""
		if schemaDataParams, ok := schemaData.(map[string]interface{}); ok {
			if val, ok := schemaDataParams["description"].(string); ok {
				description = val
			}
		}

		metricMeta := plugin.Metric{
			Namespace:   plugin.NewNamespace(nsVendor, nsClass, nsType, daemon.name).AddDynamicElement(daemon.name+"_id", fmt.Sprintf("ID of %s daemon", daemon.name)).AddStaticElements(strings.Split(schemaNs, "/")...),
			Tags:        map[string]string{"daemon_source": fmt.Sprintf("%s.%s", daemon.name, daemon.id)},
			Timestamp:   time.Now(),
			Description: description,
		}

		for _, m := range mts {
			// Provided Snap namespace matches with currently checked schema namespace
			if matchSlice(m.Namespace.Strings()[daemonNameIndex:], append([]string{daemon.name, daemon.id}, strings.Split(schemaNs, "/")...)) {
				// Ceph perf dump
				if perfDumpData == nil {
					perfDumpData = make(map[string]interface{})
					out, err := cmd.perfDump(filepath.Join(ceph.path, "ceph"), filepath.Join(ceph.socket.path, daemon.fullName))
					if err != nil {
						fmt.Fprintf(os.Stderr, "Ceph perf dump command execution failed for socket %+v, err=%+v\n",
							filepath.Join(ceph.socket.path, daemon.fullName), err)
						return nil, err
					}
					if err := decodeJSON(out, &perfDumpData); err != nil {
						fmt.Fprintf(os.Stderr, "Cannot unmarshal JSON object from ceph-daemon socket data, err=%+v\n", err)
						return nil, err
					}
				}

				daemonMetrics := make(map[string]interface{})
				ceph.getJSONPathsByNamespace(perfDumpData, jsonData, m.Namespace.Strings()[daemonIDIndex+1:], []string{}, daemonMetrics)

				// Retrieve perf dump data based on schema namespace
				data := daemonMetrics[schemaNs]

				// Check if data type is average and join delta metrics, see: http://docs.ceph.com/docs/giant/dev/perf_counters/
				if check, err := isCephTypeAverage(schemaData); err != nil {
					return nil, err
				} else if check {
					if daemonMetrics[schemaNs+"/sum"] != nil && daemonMetrics[schemaNs+"/avgcount"] != nil {
						sum, err := daemonMetrics[schemaNs+"/sum"].(json.Number).Float64()
						if err != nil {
							fmt.Fprintf(os.Stderr, "Cannot get counter sum value: %s", err)
						}
						avgcount, err := daemonMetrics[schemaNs+"/avgcount"].(json.Number).Float64()
						if err != nil {
							fmt.Fprintf(os.Stderr, "Cannot get counter avgcount value: %s", err)
						}
						if avgcount != 0 {
							data = json.Number(strconv.FormatFloat(sum/avgcount, 'e', -1, 64))
						} else {
							data = json.Number("0")
						}
					}
				}

				// Return metric data converted to desired type
				metricMeta.Namespace[4].Value = daemon.id
				if data != nil {
					if check, err := isCephTypeUint64(schemaData); err != nil {
						return nil, err
					} else if check {
						if val, err := data.(json.Number).Int64(); err != nil {
							fmt.Fprintf(os.Stderr, "Cannot convert data to int64: %s", err)
						} else {
							metricMeta.Data = val
						}
					}

					if check, err := isCephTypeFloat(schemaData); err != nil {
						return nil, err
					} else if check {
						if val, err := data.(json.Number).Float64(); err != nil {
							fmt.Fprintf(os.Stderr, "Cannot convert data to float64: %s", err)
						} else {
							metricMeta.Data = val
						}
					}
				}
				metricsCollected = append(metricsCollected, metricMeta)
			}
		}

		// Return just perf schema metric
		if mts == nil {
			metricsCollected = append(metricsCollected, metricMeta)
		}
	}

	// No metrics found at all
	if len(metricsCollected) == 0 {
		return metricsCollected, errors.New("no metrics found")
	}

	return metricsCollected, nil
}

// Matches single namespace entry, asterisk matches with everything
func matchEntry(a, b string) bool {
	return a == b || a == "*" || b == "*"
}

// matchSlice matches 2 slices with asterisk support
func matchSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !matchEntry(a[i], b[i]) {
			return false
		}
	}

	return true
}

// getJSONPathsByNamespace returns all JSON path - value entries for provided namespace
// If provided namespace is empty, all JSON paths are returned
// If dataType is jsonSchema, schema data will be returned as path -> map[description, type, nick] instead of separate path -> value entries
// If dataType is JSON_DATA, every JSON data value will be returned as separate path -> value entry
func (ceph *Ceph) getJSONPathsByNamespace(data map[string]interface{}, dataType int, namespace []string, resultNamespace []string, results map[string]interface{}) {
	// Go through all JSON data keys
	for key := range data {
		// Check if splitted JSON key matches provided namespace ("metric.0" will match both "metric/0" or "metric/*")
		// Skip match if no namespace provided - in this case match all JSON keys (used to gather all JSON metric paths)
		if len(namespace) == 0 || matchEntry(namespace[0], key) {
			// If key contain other keys, go deeper into JSON structure
			if reflect.ValueOf(data[key]).Kind() == reflect.Map {
				// For Ceph JSON schema, return map of schema values for each namespace
				// For instance, return just /somemetric/{map} instead of somemetric/description, somemetric/type etc.
				if dataType == jsonSchema && data[key].(map[string]interface{})["type"] != nil {
					results[strings.Join(resultNamespace, "/")+"/"+key] = data[key]
					continue
				}
				if len(namespace) != 0 {
					namespace = namespace[1:]
				}
				ceph.getJSONPathsByNamespace(data[key].(map[string]interface{}), dataType, namespace, append(resultNamespace, key), results)
			} else { // Dead end of JSON tree - add path to result map
				results[strings.Join(resultNamespace, "/")+"/"+key] = data[key]
			}
		}
	}
}

// CollectMetrics returns all desired Ceph metrics defined in task manifest
func (ceph *Ceph) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	if len(mts) == 0 {
		return nil, errors.New("No metrics defined to collect")
	}

	// init ceph plugin with Config settings (only once)
	if err := ceph.init(mts[0].Config); err != nil {
		return nil, err
	}

	outMetrics := []plugin.Metric{}
	for _, daemon := range ceph.daemons {
		// Retrieve daemon metric values
		if dmetric, err := ceph.getCephDaemonMetrics(mts, daemon); err == nil {
			outMetrics = append(outMetrics, dmetric...)
		}
	}

	return outMetrics, nil
}

// GetMetricTypes returns the metric types exposed by ceph-daemon sockets
func (ceph *Ceph) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	mts := []plugin.Metric{}

	// init ceph plugin with Config settings (only once)
	if err := ceph.init(cfg); err != nil {
		return nil, err
	}

	// Get metric types for each daemon (if more daemon IDs, only first found ID will be scanned for available metrics)
	daemonsCollected := make(map[string]bool)

	for _, daemon := range ceph.daemons {
		if !daemonsCollected[daemon.name] {
			if dmetric, err := ceph.getCephDaemonMetrics(nil, daemon); err == nil {
				mts = append(mts, dmetric...)
				daemonsCollected[daemon.name] = true
			}
		}
	}

	return mts, nil
}

// GetConfigPolicy returns a ConfigPolicy
func (ceph *Ceph) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{"intel", "storage", "ceph"}, "socket_path", false, plugin.SetDefaultString(socketPathDefault))
	policy.AddNewStringRule([]string{"intel", "storage", "ceph"}, "socket_prefix", false, plugin.SetDefaultString(socketPrefixDefault))
	policy.AddNewStringRule([]string{"intel", "storage", "ceph"}, "socket_ext", false, plugin.SetDefaultString(socketExtDefault))
	policy.AddNewStringRule([]string{"intel", "storage", "ceph"}, "path", false, plugin.SetDefaultString(cephBinPathDefault))
	return *policy, nil
}

// getCephBinaryPath returns path to ceph executable
func getCephBinaryPath(config plugin.Config) (string, error) {
	// check global config as a first
	path, err := config.GetString("path")
	if err != nil {
		return "", err
	}

	return path, nil
}

// getCephSocketConf returns path to folder contains daemon-sockets, prefix and extension of socket's name
func getCephSocketConf(config plugin.Config) (Socket, error) {
	s := Socket{}

	// set path to socket, defaults to "/var/run/ceph"
	if path, err := config.GetString("socket_path"); err != nil {
		return Socket{}, err
	} else {
		s.path = path
	}

	// set socket prefix, defaults to "ceph-"
	if prefix, err := config.GetString("socket_prefix"); err != nil {
		return Socket{}, err
	} else {
		s.prefix = prefix
		// if equals "none", set empty socket prefix
		if strings.ToLower(s.prefix) == "none" {
			s.prefix = ""
		}
	}

	// set socket extension, defaults to "asok"
	if ext, err := config.GetString("socket_ext"); err != nil {
		return Socket{}, err
	} else {
		s.ext = ext
	}

	return s, nil
}

// checkBit checks if idx bit of number is on
func checkBit(number, idx uint8) bool {
	return (number & (1 << idx)) != 0
}

// checkCephType is checking ceph data type based on data type bits
func checkCephType(schemaData interface{}, cephType uint8) (bool, error) {
	if schemaMap, ok := schemaData.(map[string]interface{}); ok {
		if typeVal, err := schemaMap["type"].(json.Number).Int64(); err == nil {
			return checkBit(uint8(typeVal), cephType), nil
		}
	}
	return false, fmt.Errorf("cannot get ceph data type")
}

func isCephTypeFloat(schemaData interface{}) (bool, error) { // NOTE: Not used for now, will be used after changing to use json.Decoder instead of standard json.Unmarshal
	return checkCephType(schemaData, cephTypeFloat)
}
func isCephTypeUint64(schemaData interface{}) (bool, error) {
	return checkCephType(schemaData, cephTypeUint64)
}
func isCephTypeAverage(schemaData interface{}) (bool, error) {
	return checkCephType(schemaData, cephTypeAverage)
}

// getCephDaemonNames scans the path to ceph sockets in search of an instance which name contains specified prefix and extension.
// Returns names of available ceph-daemon sockets
func (s *Socket) getCephDaemons() ([]Daemon, error) {
	files, err := ioutil.ReadDir(s.path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Not found any ceph-daemon socket in %+v. Check if the given path is correct\n", s.path)
		return nil, err
	}
	daemons := []Daemon{}
	suffix := "." + s.ext
	for _, f := range files {
		// read names of all items in socket path with specific prefix-name (defaults to "ceph-") and extension (defaults to .asok)
		if strings.HasPrefix(f.Name(), s.prefix) && strings.HasSuffix(f.Name(), suffix) {
			daemonNameSplit := strings.Split(trimPrefixAndSuffix(f.Name(), s.prefix, suffix), ".")
			if len(daemonNameSplit) < 2 {
				return nil, fmt.Errorf("Invalid socket file name: %s. Expected: %s<daemon_name>%s", f.Name(), s.prefix, suffix)
			}
			daemons = append(daemons, Daemon{id: daemonNameSplit[1], name: daemonNameSplit[0], fullName: f.Name()})
		}
	}

	return daemons, nil
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

// execute Ceph perf dump command
func perf(cephPath string, socketPath string, operation string) ([]byte, error) {
	return exec.Command(cephPath, "--admin-daemon", socketPath, "perf", operation).Output()
}

func (c *RealCmd) perfDump(cephPath string, socketPath string) ([]byte, error) {
	return perf(cephPath, socketPath, "dump")
}

// execute Ceph perf schema command
func (c *RealCmd) perfSchema(cephPath string, socketPath string) ([]byte, error) {
	return perf(cephPath, socketPath, "schema")
}

// trimPrefixAndSuffix returns 's' without the provided prefix and suffix strings.
// If 's' neither starts with prefix, nor ends with suffix, 's' is returned unchanged.
func trimPrefixAndSuffix(s string, prefix string, suffix string) string {
	s = strings.TrimPrefix(s, prefix)
	s = strings.TrimSuffix(s, suffix)
	return s
}
