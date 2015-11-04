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
	"errors"
	"fmt"
	"os"
	//	"strings"
	"testing"

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/core/cdata"
	"github.com/intelsdi-x/pulse/core/ctypes"
	. "github.com/smartystreets/goconvey/convey"
)

var ns_prefix = []string{ns_vendor, ns_class, ns_type}
var mockOut = []byte(`{
			"a": {
				"aa": 1,
				"ab": 2
			},
			"b": {
				"ba": 0
			},
			"c": {
				"ca": 1,
				"cb": {
					"cba": 22,
					"cbb": 22.2
				},
				"cc": 3 
			} 
		}`)

var mockOutInvalid = []byte(`{
			"a": {
				"aa": 1,
				"ab": 2
			}`)

var mockInUnmarshal = [][]byte{
	[]byte(`{ "a": { "aa": 1, "ab": 2 }, "b": { "ba": 0 } }`),  //correct
	[]byte(`{ {"a": { "aa": 1, "ab": 2 }, "b": { "ba": 0 } }`), //invalid character '{'
	[]byte(`{ "a": { "aa": 1 "ab": 2 }, "b": { "ba": 0 } }`),   //lack of a comma after value
	[]byte(`{ "a": { "aa": 1, "ab": 2 } "b": { "ba": 0 } }`),   //lack of a comma after }
	[]byte(`{ "a" { "aa" 1, "ab" 2 } "b" { "ba" 0 } }`),        //lack of ":"
}

// mocking exec command
type TestCmd struct {
	out               []byte
	err               error
	look_path         string
	look_err          error
	env_ceph_path     string
	env_socket_path   string
	env_socket_prefix string
	env_socket_ext    string
}

type TestSocket struct {
	path   string
	prefix string
	ext    string
}

func (t *TestCmd) perfDump(command string, args ...string) ([]byte, error) {
	return t.out, t.err
}

func (t *TestCmd) lookPath(file string) (string, error) {
	return t.look_path, t.look_err
}

func Test_trimPrefixAndSuffix(t *testing.T) {
	Convey("trimming", t, func() {
		So(trimPrefixAndSuffix("", "", ""), ShouldEqual, "")
		So(trimPrefixAndSuffix(" ", " ", ""), ShouldEqual, "")
		So(trimPrefixAndSuffix(" ", "  ", ""), ShouldEqual, " ")
		So(trimPrefixAndSuffix(" ", "  ", " "), ShouldEqual, "")

		So(trimPrefixAndSuffix("ceph-mon.a.asok", "ceph-", ".asok"), ShouldEqual, "mon.a")
		So(trimPrefixAndSuffix("ceph-osd.0.asok", "ceph-", ".asok"), ShouldEqual, "osd.0")
		So(trimPrefixAndSuffix("mon.a.asok", "ceph-", ".asok"), ShouldEqual, "mon.a")
		So(trimPrefixAndSuffix("mon.a.asok", "", ".asok"), ShouldEqual, "mon.a")
		So(trimPrefixAndSuffix("ceph-mon.a.asok", "", ".asok"), ShouldEqual, "ceph-mon.a")
		So(trimPrefixAndSuffix("ceph-mon.a.asok", "", ""), ShouldEqual, "ceph-mon.a.asok")
		So(trimPrefixAndSuffix("ceph-mon.a.asok", "Ceph-", ".asok"), ShouldEqual, "ceph-mon.a")
		So(trimPrefixAndSuffix("Ceph-Mon.a.asok", "Ceph-", ".asok"), ShouldEqual, "Mon.a")
		So(trimPrefixAndSuffix("ceph-Mon.a.ASOK", "ceph-", ".ASOK"), ShouldEqual, "Mon.a")

		So(trimPrefixAndSuffix("", "ABC", "XYZ"), ShouldEqual, "")
		So(trimPrefixAndSuffix("A CtestX_Z", "A C", "X_Z"), ShouldEqual, "test")
		So(trimPrefixAndSuffix("ABCtestX_Z", "CBA", "X_Z"), ShouldEqual, "ABCtest")
		So(trimPrefixAndSuffix("   testX_Z", "A C", "X_Z"), ShouldEqual, "   test")
		So(trimPrefixAndSuffix("   testX_Z", "A C", "X_Z"), ShouldEqual, "   test")
		So(trimPrefixAndSuffix("12Atest3  ", "12A", "3  "), ShouldEqual, "test")
	})
}

func Test_getCephDaemonNames(t *testing.T) {
	socketTestA := &Socket{path: "./test/", prefix: "prefix-", ext: "ext"}
	socketTestB := &Socket{path: "./test/", prefix: "", ext: "ext"}
	socketTestC := &Socket{path: "./test/", prefix: "prefix-", ext: "extB"}
	socketTestD := &Socket{path: "./test/", prefix: "", ext: ""}

	Convey("obtaining daemon names when there is no directory", t, func() {
		So(func() { socketTestA.getCephDaemonNames() }, ShouldPanic)
	})

	os.Mkdir("test", os.ModePerm)

	Convey("obtaining daemon names when there is no file in the indicated directory", t, func() {
		So(func() { socketTestA.getCephDaemonNames() }, ShouldNotPanic)
		result := socketTestA.getCephDaemonNames()
		So(result, ShouldBeEmpty)
	})

	os.Create("test/socket.aaa")
	os.Create("test/socket.bbb")

	Convey("obtaining daemon names when there is no proper file in the indicated directory", t, func() {
		So(func() { socketTestA.getCephDaemonNames() }, ShouldNotPanic)
		result := socketTestA.getCephDaemonNames()
		So(result, ShouldBeEmpty)
	})

	os.Create("test/prefix-socket1.ext")
	os.Create("test/prefix-socket2.ext")
	os.Create("test/socket3.ext")
	os.Create("test/prefix-socket4.extB")

	Convey("obtaining daemon names from existing files (1)", t, func() {
		result := socketTestA.getCephDaemonNames()
		So(result, ShouldContain, "prefix-socket1.ext")
		So(result, ShouldContain, "prefix-socket2.ext")
	})

	Convey("obtaining daemon names from existing files (2)", t, func() {
		result := socketTestB.getCephDaemonNames()
		So(result, ShouldContain, "socket3.ext")
	})

	Convey("obtaining daemon names from existing files (3)", t, func() {
		result := socketTestC.getCephDaemonNames()
		So(result, ShouldContain, "prefix-socket4.extB")
	})

	Convey("obtaining daemon names from existing files (4)", t, func() {
		result := socketTestD.getCephDaemonNames()
		So(result, ShouldBeEmpty)
	})

	os.RemoveAll("test/")

}

func createMockMetrics(count uint64, hostname string) []plugin.PluginMetricType {
	metrics := make([]plugin.PluginMetricType, count)
	for i, _ := range metrics {

		metrics[i] = plugin.PluginMetricType{
			Namespace_: []string{"test", "metric", "namespace"},
			Data_:      1.01 + float32(i),
			Source_:    hostname,
		}
	}

	return metrics
}

func checkMetricsResemblance(metrics []plugin.PluginMetricType, metricsMock []plugin.PluginMetricType) {
	for i, mock := range metricsMock {
		So(metrics[i].Namespace(), ShouldResemble, mock.Namespace())
		So(metrics[i].Source(), ShouldEqual, mock.Source())
		So(metrics[i].Data(), ShouldEqual, mock.Data())
	}
}

// desiredMtsCnt returns number of desired metrics to collect from specific daemon
func desiredMtsCnt(daemonName string, metrics []plugin.PluginMetricType) int {
	cnt := 0
	for _, m := range metrics {
		if m.Namespace()[daemonNameIndex] == daemonName {
			cnt++
		}
	}

	return cnt
}

func Test_getCephDaemonMetrics(t *testing.T) {
	mts := []plugin.PluginMetricType{
		plugin.PluginMetricType{
			Namespace_: []string{"intel", "storage", "ceph", "mds.a", "a", "aa"},
		},
		plugin.PluginMetricType{
			Namespace_: []string{"intel", "storage", "ceph", "mds.a", "b", "ba"},
		},
		plugin.PluginMetricType{
			Namespace_: []string{"intel", "storage", "ceph", "mds.a", "c", "cb", "cba"},
		},
		plugin.PluginMetricType{
			Namespace_: []string{"intel", "storage", "ceph", "mds.a", "c", "cb", "cbb"},
		},
		plugin.PluginMetricType{
			Namespace_: []string{"intel", "storage", "ceph", "mds.b", "a", "aa"},
		},
		plugin.PluginMetricType{
			Namespace_: []string{"intel", "storage", "ceph", "mds.b", "a", "cc"},
		},
		plugin.PluginMetricType{
			Namespace_: []string{"intel", "storage", "ceph", "mds.c", "no", "nono"},
		},
	}

	testCeph := &Ceph{}

	Convey("invalid getting metrics, perf dump command execution error", t, func() {
		cmd = &TestCmd{err: errors.New("exit status 1")}

		So(func() { testCeph.getCephDaemonMetrics(mts, "mds.a") }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, "mds.a")
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("invalid getting metrics, empty perf dump output", t, func() {
		cmd = &TestCmd{}
		So(func() { testCeph.getCephDaemonMetrics(mts, "mds.a") }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, "mds.a")
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("unmarshal perf dump output", t, func() {
		//first mock in mockInUnmarshal is correct
		cmd = &TestCmd{out: mockInUnmarshal[0]}
		dName := "mds.a"
		So(func() { testCeph.getCephDaemonMetrics(mts, dName) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, dName)
		So(result, ShouldNotBeEmpty)
		So(err, ShouldBeNil)
	})

	Convey("invalid unmarshal perf dump output, incorrect json format", t, func() {
		for _, moi := range mockInUnmarshal[1:] {
			cmd = &TestCmd{out: moi}
			dName := "mds.a"
			So(func() { testCeph.getCephDaemonMetrics(mts, dName) }, ShouldNotPanic)
			result, err := testCeph.getCephDaemonMetrics(mts, dName)
			So(result, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		}
	})

	Convey("no defined desired metrics for daemon", t, func() {
		cmd = &TestCmd{out: mockOut}
		dName := "mds.d"
		So(func() { testCeph.getCephDaemonMetrics(mts, dName) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, dName)
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("no available desired metrics for daemon", t, func() {
		cmd = &TestCmd{out: mockOut}
		dName := "mds.c"
		So(func() { testCeph.getCephDaemonMetrics(mts, dName) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, dName)
		So(len(result), ShouldEqual, desiredMtsCnt(dName, mts))
		So(result[0].Data(), ShouldBeNil)
		So(err, ShouldBeNil)
	})

	Convey("get ceph-daemon metrics (1)", t, func() {
		cmd = &TestCmd{out: mockOut}
		dName := "mds.a"
		So(func() { testCeph.getCephDaemonMetrics(mts, dName) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, dName)
		So(result, ShouldNotBeEmpty)
		So(len(result), ShouldEqual, desiredMtsCnt(dName, mts))
		So(result[0].Data(), ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("get ceph-daemon metrics (2)", t, func() {
		cmd = &TestCmd{out: mockOut}
		dName := "mds.b"
		So(func() { testCeph.getCephDaemonMetrics(mts, dName) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, dName)
		So(result, ShouldNotBeEmpty)
		So(len(result), ShouldEqual, desiredMtsCnt(dName, mts))
		So(result[0].Data(), ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func TestGetMetricTypes(t *testing.T) {
	// Testing GetMetricTypes function covers checking initialization of Ceph plugin

	// mock ceph socket
	os.Mkdir("test", os.ModePerm)
	os.Create("test/ceph-mds.a.asok")
	os.Create("test/ceph-mds.b.asok")
	os.Create("test/ceph-mds.c.asok")

	// set ceph socket conf value
	cfg := plugin.NewPluginConfigType()
	cfg.AddItem("socket_path", ctypes.ConfigValueStr{Value: "./test"})
	cfg.AddItem("socket_prefix", ctypes.ConfigValueStr{Value: "ceph-"})
	cfg.AddItem("socket_ext", ctypes.ConfigValueStr{Value: "asok"})

	Convey("init ceph plugin and getting metric namespace", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{out: mockOut}
		So(func() { ceph.GetMetricTypes(cfg) }, ShouldNotPanic)
		result, err := ceph.GetMetricTypes(cfg)
		So(result, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("getting metric namespace with error - perf dump output is not valid", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{out: mockOutInvalid}
		So(func() { ceph.GetMetricTypes(cfg) }, ShouldNotPanic)
		result, err := ceph.GetMetricTypes(cfg)
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("getting metric namespace with error - perf dump execution error", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{err: errors.New("execution error")}
		So(func() { ceph.GetMetricTypes(cfg) }, ShouldNotPanic)
		result, err := ceph.GetMetricTypes(cfg)
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("getting metric namespace with error - perf dump output is empty", t, func() {
		ceph := &Ceph{}
		moi := []byte(fmt.Sprintf("\n\n\n\n"))
		cmd = &TestCmd{out: moi}
		So(func() { ceph.GetMetricTypes(cfg) }, ShouldNotPanic)
		result, err := ceph.GetMetricTypes(cfg)
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("getting metric namespace with error - none perf dump output", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{}
		So(func() { ceph.GetMetricTypes(cfg) }, ShouldNotPanic)
		result, err := ceph.GetMetricTypes(cfg)
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("invalid path to ceph-daemon sockets", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{out: mockOut}
		cfg_invalid := plugin.NewPluginConfigType()
		cfg_invalid.AddItem("socket_path", ctypes.ConfigValueStr{Value: "./test_invalid"})
		cfg_invalid.AddItem("socket_prefix", ctypes.ConfigValueStr{Value: "ceph-"})
		cfg_invalid.AddItem("socket_ext", ctypes.ConfigValueStr{Value: "asok"})
		So(func() { ceph.GetMetricTypes(cfg_invalid) }, ShouldPanic)
	})

	Convey("no ceph-daemon sockets available with set prefix", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{out: mockOut}
		cfg_inv_prefix := plugin.NewPluginConfigType()
		cfg_inv_prefix.AddItem("socket_path", ctypes.ConfigValueStr{Value: "./test"})
		cfg_inv_prefix.AddItem("socket_prefix", ctypes.ConfigValueStr{Value: "ceph_inv-"})
		So(func() { ceph.GetMetricTypes(cfg_inv_prefix) }, ShouldNotPanic)
		result, err := ceph.GetMetricTypes(cfg_inv_prefix)
		So(result, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("no ceph-daemon sockets available with set extension", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{out: mockOut}
		cfg_inv_ext := plugin.NewPluginConfigType()
		cfg_inv_ext.AddItem("socket_path", ctypes.ConfigValueStr{Value: "./test"})
		cfg_inv_ext.AddItem("socket_ext", ctypes.ConfigValueStr{Value: "asok_inv"})
		So(func() { ceph.GetMetricTypes(cfg_inv_ext) }, ShouldNotPanic)
		result, err := ceph.GetMetricTypes(cfg_inv_ext)
		So(result, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	os.RemoveAll("test/")
}

func Test_parsePerfDumpOut(t *testing.T) {
	Convey("parse perf dump output when output is empty", t, func() {
		out := []byte{}
		So(func() { parsePerfDumpOut(out) }, ShouldNotPanic)
		result, err := parsePerfDumpOut(out)
		So(result, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("parse perf dump returns no error", t, func() {
		So(func() { parsePerfDumpOut(mockOut) }, ShouldNotPanic)
		result, err := parsePerfDumpOut(mockOut)
		So(result, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("parse perf dump returns error", t, func() {
		So(func() { parsePerfDumpOut(mockOutInvalid) }, ShouldNotPanic)
		result, err := parsePerfDumpOut(mockOutInvalid)
		So(result, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

func Test_getCephBinaryPath(t *testing.T) {
	path := "/path/to/ceph/bin"

	Convey("path Ceph executable in Pulse Global Config", t, func() {
		cmd = &TestCmd{}
		cfg := plugin.NewPluginConfigType()
		cfg.AddItem("path", ctypes.ConfigValueStr{Value: path})

		So(func() { getCephBinaryPath(cfg.Table()) }, ShouldNotPanic)
		result := getCephBinaryPath(cfg.Table())
		So(result, ShouldEqual, path)
	})

	Convey("looking for Ceph executable in $PATH with no error", t, func() {
		cfg := plugin.NewPluginConfigType()
		cmd = &TestCmd{look_path: path}
		So(func() { getCephBinaryPath(cfg.Table()) }, ShouldNotPanic)
		result := getCephBinaryPath(cfg.Table())
		So(result, ShouldEqual, path)
	})

	Convey("looking for Ceph executable in $PATH with error", t, func() {
		cfg := plugin.NewPluginConfigType()
		cmd = &TestCmd{look_path: path, look_err: errors.New("error")}
		So(func() { getCephBinaryPath(cfg.Table()) }, ShouldNotPanic)
		result := getCephBinaryPath(cfg.Table())
		// getting default path to ceph executable
		So(result, ShouldEqual, cephBinPathDefault)
	})
}

func Test_getSocketConf(t *testing.T) {
	Convey("defaults Ceph socket details", t, func() {
		cmd = &TestCmd{}
		cfg := plugin.NewPluginConfigType()

		So(func() { getCephSocketConf(cfg.Table()) }, ShouldNotPanic)
		result := getCephSocketConf(cfg.Table())
		So(result.path, ShouldEqual, socketPathDefault)
		So(result.prefix, ShouldEqual, socketPrefixDefault)
		So(result.ext, ShouldEqual, socketExtDefault)
	})

	Convey("customize Ceph socket conf in Pulse Global Conf", t, func() {
		path := "path/to/ceph/socket"
		prefix := "test-"
		extension := "asdf"
		cmd = &TestCmd{}
		cfg := plugin.NewPluginConfigType()
		cfg.AddItem("socket_path", ctypes.ConfigValueStr{Value: path})
		cfg.AddItem("socket_prefix", ctypes.ConfigValueStr{Value: prefix})
		cfg.AddItem("socket_ext", ctypes.ConfigValueStr{Value: extension})

		So(func() { getCephSocketConf(cfg.Table()) }, ShouldNotPanic)
		result := getCephSocketConf(cfg.Table())
		So(result.path, ShouldEqual, path)
		So(result.prefix, ShouldEqual, prefix)
		So(result.ext, ShouldEqual, extension)
	})

	Convey("customize Ceph socket prefix, set none", t, func() {
		cmd = &TestCmd{}
		cfg := plugin.NewPluginConfigType()
		cfg.AddItem("socket_prefix", ctypes.ConfigValueStr{Value: "none"})
		So(func() { getCephSocketConf(cfg.Table()) }, ShouldNotPanic)
		result := getCephSocketConf(cfg.Table())
		So(result.prefix, ShouldBeEmpty)
	})
}

func Test_CollectMetrics(t *testing.T) {
	// folder for mocked ceph socket, empty in the beginning
	os.Mkdir("test", os.ModePerm)

	mts := []plugin.PluginMetricType{
		plugin.PluginMetricType{
			Namespace_: []string{"intel", "storage", "ceph", "osd.2", "a", "aa"},
		},
		plugin.PluginMetricType{
			Namespace_: []string{"intel", "storage", "ceph", "osd.3", "c", "cb", "cba"},
		},
	}

	config := cdata.NewNode()
	for i, _ := range mts {
		mts[i].Config_ = config
	}

	// set ceph socket conf value
	config.AddItem("socket_path", ctypes.ConfigValueStr{Value: "./test"})
	config.AddItem("socket_prefix", ctypes.ConfigValueStr{Value: "none"})
	config.AddItem("socket_ext", ctypes.ConfigValueStr{Value: "asok"})

	Convey("no metrics defined to collect", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{out: mockOut}
		ceph.daemons = []string{}
		mts_empty := []plugin.PluginMetricType{}
		So(func() { ceph.CollectMetrics(mts_empty) }, ShouldNotPanic)
		result, err := ceph.CollectMetrics(mts_empty)
		So(result, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("no ceph daemon is running", t, func() {
		// none mocked ceph socket in ./test
		ceph := &Ceph{}
		cmd = &TestCmd{out: mockOut}
		ceph.daemons = []string{}
		So(func() { ceph.CollectMetrics(mts) }, ShouldNotPanic)
		result, err := ceph.CollectMetrics(mts)
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	// create ceph socket
	os.Create("test/osd.0.asok")
	os.Create("test/osd.1.asok")

	Convey("try collecting metrics for not available socket", t, func() {
		// metrics for osd.2 and osd.3 are not available
		ceph := &Ceph{}
		cmd = &TestCmd{out: mockOut}
		So(func() { ceph.CollectMetrics(mts) }, ShouldNotPanic)
		result, err := ceph.CollectMetrics(mts)
		So(result, ShouldBeEmpty)
		So(err, ShouldBeNil)
	})

	// create additional ceph socket
	os.Create("test/osd.2.asok")
	os.Create("test/osd.3.asok")

	Convey("collect metrics with no errors", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{out: mockOut}
		So(func() { ceph.CollectMetrics(mts) }, ShouldNotPanic)
		result, err := ceph.CollectMetrics(mts)
		So(len(result), ShouldEqual, len(mts))

		for _, r := range result {
			So(r.Data(), ShouldNotBeNil)
		}
		So(err, ShouldBeNil)
	})
	os.RemoveAll("test/")
}

func Test_New(t *testing.T) {
	Convey("create new Ceph structure", t, func() {
		So(func() { New() }, ShouldNotPanic)
		result := New()
		So(result, ShouldNotBeNil)
	})
}
