// +build small

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
	"os"
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	. "github.com/smartystreets/goconvey/convey"
)

var ns_prefix = []string{nsVendor, nsClass, nsType}
var mockOut = []byte(`{
			"a": {
				"aa": 1,
				"ab": 2
			},
			"b": {
				"bavg": {
					"avgcount": 6,
					"sum": 3.000000000
				}
			},
			"c": {
				"ca": 1,
				"cb": {
					"cba": 123456.1234567890,
					"cbb": 12345678901234567
				},
				"cc": 3
			}
		}`)

var mockOutSchema = []byte(`{
			"a": {
				"aa": {
					"description": "Some metric called aa",
					"nick": "",
					"type": 10
				},
				"ab": {
					"description": "Some metric called ab",
					"nick": "",
					"type": 10
				}
			},
			"b": {
				"bavg": {
					"description": "Some delta metric",
					"nick": "",
					"type": 5
				}
			},
			"c": {
				"ca": {
					"description": "Some metric called ca",
					"nick": "",
					"type": 10
				},
				"cb": {
					"cba": {
						"description": "Some metric called cba",
						"nick": "",
						"type": 1
					},
					"cbb": {
						"no_description": "",
						"nick": "",
						"type": 10
					}
				},
				"cc": {
					"description": "Some metric called cc",
					"nick": "",
					"type": 10
				}
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
	outDump           []byte
	outSchema         []byte
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

func (t *TestCmd) perfDump(cephPath string, socketPath string) ([]byte, error) {
	return t.outDump, t.err
}

func (t *TestCmd) perfSchema(cephPath string, socketPath string) ([]byte, error) {
	return t.outSchema, t.err
}

func (t *TestCmd) lookPath(file string) (string, error) {
	return t.look_path, t.look_err
}

func Test_checkBit(t *testing.T) {
	Convey("checking enabled bits for 170", t, func() {
		var number uint8 = 170
		So(checkBit(number, 0), ShouldBeFalse)
		So(checkBit(number, 1), ShouldBeTrue)
		So(checkBit(number, 2), ShouldBeFalse)
		So(checkBit(number, 3), ShouldBeTrue)
		So(checkBit(number, 4), ShouldBeFalse)
		So(checkBit(number, 5), ShouldBeTrue)
		So(checkBit(number, 6), ShouldBeFalse)
		So(checkBit(number, 7), ShouldBeTrue)
	})

	Convey("checking enabled bits for 85", t, func() {
		var number uint8 = 85
		So(checkBit(number, 0), ShouldBeTrue)
		So(checkBit(number, 1), ShouldBeFalse)
		So(checkBit(number, 2), ShouldBeTrue)
		So(checkBit(number, 3), ShouldBeFalse)
		So(checkBit(number, 4), ShouldBeTrue)
		So(checkBit(number, 5), ShouldBeFalse)
		So(checkBit(number, 6), ShouldBeTrue)
		So(checkBit(number, 7), ShouldBeFalse)
	})

	Convey("checking enabled bits for 255", t, func() {
		var number uint8 = 255
		So(checkBit(number, 0), ShouldBeTrue)
		So(checkBit(number, 1), ShouldBeTrue)
		So(checkBit(number, 2), ShouldBeTrue)
		So(checkBit(number, 3), ShouldBeTrue)
		So(checkBit(number, 4), ShouldBeTrue)
		So(checkBit(number, 5), ShouldBeTrue)
		So(checkBit(number, 6), ShouldBeTrue)
		So(checkBit(number, 7), ShouldBeTrue)
	})

	Convey("checking enabled bits for 0", t, func() {
		var number uint8 = 0
		So(checkBit(number, 0), ShouldBeFalse)
		So(checkBit(number, 1), ShouldBeFalse)
		So(checkBit(number, 2), ShouldBeFalse)
		So(checkBit(number, 3), ShouldBeFalse)
		So(checkBit(number, 4), ShouldBeFalse)
		So(checkBit(number, 5), ShouldBeFalse)
		So(checkBit(number, 6), ShouldBeFalse)
		So(checkBit(number, 7), ShouldBeFalse)
	})
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

func daemonsToStrings(daemons []Daemon) []string {
	result := []string{}

	for _, d := range daemons {
		result = append(result, d.name+"."+d.id)
	}

	return result
}

func Test_getCephDaemonNames(t *testing.T) {
	socketTestA := &Socket{path: "./test/", prefix: "prefix-", ext: "ext"}
	socketTestB := &Socket{path: "./test/", prefix: "", ext: "ext"}
	socketTestC := &Socket{path: "./test/", prefix: "prefix-", ext: "extB"}
	socketTestD := &Socket{path: "./test/", prefix: "", ext: ""}

	Convey("obtaining daemon names when there is no directory", t, func() {
		daemons, err := socketTestA.getCephDaemons()
		So(daemons, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	os.Mkdir("test", os.ModePerm)

	Convey("obtaining daemon names when there is no file in the indicated directory", t, func() {
		So(func() { socketTestA.getCephDaemons() }, ShouldNotPanic)
		result, err := socketTestA.getCephDaemons()
		So(err, ShouldBeNil)
		So(result, ShouldBeEmpty)
	})

	os.Create("test/socket.aaa")
	os.Create("test/socket.bbb")

	Convey("obtaining daemon names when there is no proper file in the indicated directory", t, func() {
		So(func() { socketTestA.getCephDaemons() }, ShouldNotPanic)
		result, err := socketTestA.getCephDaemons()
		So(err, ShouldBeNil)
		So(result, ShouldBeEmpty)
	})

	os.Create("test/prefix-socket.1.ext")
	os.Create("test/prefix-socket.2.ext")
	os.Create("test/socket.3.ext")
	os.Create("test/prefix-socket.4.extB")

	daemonName := make(map[string]Daemon)
	daemonName["socket3.ext"] = Daemon{id: "", name: "socket3", fullName: "socket3.ext"}

	Convey("obtaining daemon names from existing files (1)", t, func() {
		result, err := socketTestA.getCephDaemons()
		So(err, ShouldBeNil)
		daemonNames := daemonsToStrings(result)
		So(daemonNames, ShouldContain, "socket.1")
		So(daemonNames, ShouldContain, "socket.2")
	})

	Convey("obtaining daemon names from existing files (2)", t, func() {
		result, err := socketTestB.getCephDaemons()
		So(err, ShouldBeNil)
		daemonNames := daemonsToStrings(result)
		So(daemonNames, ShouldContain, "prefix-socket.1")
		So(daemonNames, ShouldContain, "prefix-socket.2")
		So(daemonNames, ShouldContain, "socket.3")
	})

	Convey("obtaining daemon names from existing files (3)", t, func() {
		result, err := socketTestC.getCephDaemons()
		So(err, ShouldBeNil)
		daemonNames := daemonsToStrings(result)
		So(daemonNames, ShouldContain, "socket.4")
	})

	Convey("obtaining daemon names from existing files (4)", t, func() {
		result, err := socketTestD.getCephDaemons()
		So(err, ShouldBeNil)
		daemonNames := daemonsToStrings(result)
		So(daemonNames, ShouldBeEmpty)
	})

	os.RemoveAll("test/")

}

func createMockMetrics(count uint64, hostname string) []plugin.Metric {
	metrics := make([]plugin.Metric, count)
	for i, _ := range metrics {

		metrics[i] = plugin.Metric{
			Namespace: plugin.NewNamespace("test", "metric", "namespace"),
			Data:      1.01 + float32(i),
			Tags:      map[string]string{"source": hostname},
		}
	}

	return metrics
}

func checkMetricsResemblance(metrics []plugin.Metric, metricsMock []plugin.Metric) {
	for i, mock := range metricsMock {
		So(metrics[i].Namespace, ShouldResemble, mock.Namespace)
		So(metrics[i].Tags, ShouldResemble, mock.Tags)
		So(metrics[i].Data, ShouldEqual, mock.Data)
	}
}

func Test_getCephDaemonMetrics(t *testing.T) {
	mts := []plugin.Metric{
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "mds", "a", "a", "aa"),
		},
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "mds", "a", "a", "cc"),
		},
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "mds", "a", "c", "cb", "cba"),
		},
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "mds", "a", "c", "cb", "cbb"),
		},
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "mds", "b", "b", "bavg"),
		},
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "mds", "c", "no", "nono"),
		},
	}

	mds := make(map[string]Daemon)
	mds["a"] = Daemon{id: "a", name: "mds", fullName: "mds.a.asok"}
	mds["b"] = Daemon{id: "b", name: "mds", fullName: "mds.b.asok"}
	mds["c"] = Daemon{id: "c", name: "mds", fullName: "mds.c.asok"}
	mds["d"] = Daemon{id: "d", name: "mds", fullName: "mds.d.asok"}

	Convey("invalid getting metrics, perf dump command execution error", t, func() {
		cmd = &TestCmd{err: errors.New("exit status 1")}
		testCeph := &Ceph{}
		So(func() { testCeph.getCephDaemonMetrics(mts, mds["a"]) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, mds["a"])
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("invalid getting metrics, empty perf dump output", t, func() {
		cmd = &TestCmd{}
		testCeph := &Ceph{}
		So(func() { testCeph.getCephDaemonMetrics(mts, mds["a"]) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, mds["a"])
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("unmarshal perf dump output", t, func() {
		cmd = &TestCmd{outDump: mockInUnmarshal[0], outSchema: mockOutSchema}
		testCeph := &Ceph{}
		testCeph.getCephDaemonSchema(mds["a"])
		//first mock in mockInUnmarshal is correct
		So(func() { testCeph.getCephDaemonMetrics(mts, mds["a"]) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, mds["a"])
		So(result, ShouldNotBeEmpty)
		So(err, ShouldBeNil)
	})

	Convey("invalid unmarshal perf dump output, incorrect json format", t, func() {
		testCeph := &Ceph{}
		for _, moi := range mockInUnmarshal[1:] {
			cmd = &TestCmd{outDump: moi, outSchema: mockOutSchema}
			testCeph.getCephDaemonSchema(mds["a"])
			So(func() { testCeph.getCephDaemonMetrics(mts, mds["a"]) }, ShouldNotPanic)
			result, err := testCeph.getCephDaemonMetrics(mts, mds["a"])
			So(result, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		}
	})

	Convey("no defined desired metrics for daemon", t, func() {
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		testCeph := &Ceph{}
		testCeph.getCephDaemonSchema(mds["d"])
		So(func() { testCeph.getCephDaemonMetrics(mts, mds["d"]) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, mds["d"])
		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	Convey("no available desired metrics for daemon", t, func() {
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		testCeph := &Ceph{}
		testCeph.getCephDaemonSchema(mds["c"])
		So(func() { testCeph.getCephDaemonMetrics(mts, mds["c"]) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, mds["c"])
		So(len(result), ShouldEqual, 0)
		So(err, ShouldNotBeNil)
	})

	Convey("get ceph-daemon metrics", t, func() {
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		testCeph := &Ceph{}
		testCeph.getCephDaemonSchema(mds["a"])
		So(func() { testCeph.getCephDaemonMetrics(mts, mds["a"]) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, mds["a"])
		So(result, ShouldNotBeEmpty)
		So(len(result), ShouldEqual, 3)
		for _, r := range result {
			So(r.Data, ShouldNotBeNil)
		}
		So(err, ShouldBeNil)
	})

	Convey("get ceph-daemon metrics (delta metrics)", t, func() {
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		testCeph := &Ceph{}
		testCeph.getCephDaemonSchema(mds["b"])
		So(func() { testCeph.getCephDaemonMetrics(mts, mds["b"]) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics(mts, mds["b"])
		So(result, ShouldNotBeEmpty)
		So(len(result), ShouldEqual, 1)
		So(result[0].Data, ShouldEqual, 0.5)
		So(result[0].Description, ShouldNotBeEmpty)
		So(err, ShouldBeNil)
	})

	Convey("check ceph-daemon float64 number metrics", t, func() {
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		testCeph := &Ceph{}
		testCeph.getCephDaemonSchema(mds["a"])
		So(func() { testCeph.getCephDaemonMetrics([]plugin.Metric{mts[2]}, mds["a"]) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics([]plugin.Metric{mts[2]}, mds["a"])
		So(result, ShouldNotBeEmpty)
		var checkData float64 = 123456.1234567890
		So(result[0].Data, ShouldHaveSameTypeAs, checkData)
		So(result[0].Data, ShouldEqual, checkData)
		So(result[0].Description, ShouldNotBeEmpty)
		So(err, ShouldBeNil)
	})

	Convey("check ceph-daemon large number metrics", t, func() {
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		testCeph := &Ceph{}
		testCeph.getCephDaemonSchema(mds["a"])
		So(func() { testCeph.getCephDaemonMetrics([]plugin.Metric{mts[3]}, mds["a"]) }, ShouldNotPanic)
		result, err := testCeph.getCephDaemonMetrics([]plugin.Metric{mts[3]}, mds["a"])
		So(result, ShouldNotBeEmpty)
		var checkData int64 = 12345678901234567
		So(result[0].Data, ShouldHaveSameTypeAs, checkData)
		So(result[0].Data, ShouldEqual, checkData)
		So(result[0].Description, ShouldBeEmpty)
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
	cfg := plugin.Config{}
	cfg["path"] = "/usr/bin"
	cfg["socket_path"] = "./test"
	cfg["socket_prefix"] = "ceph-"
	cfg["socket_ext"] = "asok"

	Convey("init ceph plugin and getting metric namespace", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		So(func() { ceph.GetMetricTypes(cfg) }, ShouldNotPanic)
		result, err := ceph.GetMetricTypes(cfg)
		So(result, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("invalid path to ceph-daemon sockets", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		cfg_invalid := plugin.Config{}
		cfg_invalid["socket_path"] = "./test_invalid"
		cfg_invalid["socket_prefix"] = "ceph-"
		cfg_invalid["socket_ext"] = "asok"
		err := ceph.init(cfg_invalid)
		So(err, ShouldNotBeNil)
	})

	Convey("no ceph-daemon sockets available with set prefix", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		cfg_inv_prefix := plugin.Config{}
		cfg_inv_prefix["socket_path"] = "./test"
		cfg_inv_prefix["socket_prefix"] = "ceph_inv-"
		So(func() { ceph.init(cfg_inv_prefix) }, ShouldNotPanic)
		err := ceph.init(cfg_inv_prefix)
		So(err, ShouldNotBeNil)
	})

	Convey("no ceph-daemon sockets available with set extension", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		cfg_inv_ext := plugin.Config{}
		cfg_inv_ext["socket_path"] = "./test"
		cfg_inv_ext["socket_ext"] = "asok_inv"
		So(func() { ceph.init(cfg_inv_ext) }, ShouldNotPanic)
		err := ceph.init(cfg_inv_ext)
		So(err, ShouldNotBeNil)
	})

	os.RemoveAll("test/")
}

func Test_getCephBinaryPath(t *testing.T) {
	path := "/path/to/ceph/bin"

	Convey("path Ceph executable in Snap Global Config", t, func() {
		cmd = &TestCmd{}
		cfg := plugin.Config{}
		cfg["path"] = path

		So(func() { getCephBinaryPath(cfg) }, ShouldNotPanic)
		result, err := getCephBinaryPath(cfg)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, path)
	})
}

func Test_getSocketConf(t *testing.T) {
	ceph := New()

	Convey("check Ceph socket default config", t, func() {
		cmd = &TestCmd{}

		testPolicy := plugin.NewConfigPolicy()
		testPolicy.AddNewStringRule([]string{"intel", "storage", "ceph"}, "socket_path", false, plugin.SetDefaultString(socketPathDefault))
		testPolicy.AddNewStringRule([]string{"intel", "storage", "ceph"}, "socket_prefix", false, plugin.SetDefaultString(socketPrefixDefault))
		testPolicy.AddNewStringRule([]string{"intel", "storage", "ceph"}, "socket_ext", false, plugin.SetDefaultString(socketExtDefault))
		testPolicy.AddNewStringRule([]string{"intel", "storage", "ceph"}, "path", false, plugin.SetDefaultString(cephBinPathDefault))

		cp, _ := ceph.GetConfigPolicy()

		So(cp, ShouldResemble, *testPolicy)
	})

	Convey("customize Ceph socket conf in Snap Global Conf", t, func() {
		path := "path/to/ceph/socket"
		prefix := "test-"
		extension := "asdf"
		cmd = &TestCmd{}
		cfg := plugin.Config{}
		cfg["socket_path"] = path
		cfg["socket_prefix"] = prefix
		cfg["socket_ext"] = extension

		So(func() { getCephSocketConf(cfg) }, ShouldNotPanic)
		result, err := getCephSocketConf(cfg)
		So(err, ShouldBeNil)
		So(result.path, ShouldEqual, path)
		So(result.prefix, ShouldEqual, prefix)
		So(result.ext, ShouldEqual, extension)
	})

	Convey("customize Ceph socket prefix, set none", t, func() {
		cmd = &TestCmd{}
		cfg := plugin.Config{}
		cfg["socket_path"] = ""
		cfg["socket_prefix"] = "none"
		cfg["socket_ext"] = ""
		So(func() { getCephSocketConf(cfg) }, ShouldNotPanic)
		result, err := getCephSocketConf(cfg)
		So(err, ShouldBeNil)
		So(result.prefix, ShouldBeEmpty)
	})
}

func Test_CollectMetrics(t *testing.T) {
	// folder for mocked ceph socket, empty in the beginning
	os.Mkdir("test", os.ModePerm)

	// set ceph socket conf value
	config := plugin.Config{}
	config["path"] = "/usr/bin"
	config["socket_path"] = "./test"
	config["socket_prefix"] = "none"
	config["socket_ext"] = "asok"

	// Test data
	mts := []plugin.Metric{
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "osd", "2", "a", "aa"), Config: config,
		},
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "osd", "3", "c", "cb", "cba"), Config: config,
		},
	}
	mtsDynamic := []plugin.Metric{
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "osd", "*", "a", "aa"), Config: config,
		},
	}
	mtsDynamicErr := []plugin.Metric{
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "storage", "ceph", "osd", "*", "mnopqrst", "aabbcc"), Config: config,
		},
	}

	Convey("no metrics defined to collect", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		ceph.daemons = []Daemon{}
		mts_empty := []plugin.Metric{}
		So(func() { ceph.CollectMetrics(mts_empty) }, ShouldNotPanic)
		result, err := ceph.CollectMetrics(mts_empty)
		So(result, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("no ceph daemon is running", t, func() {
		// none mocked ceph socket in ./test
		ceph := &Ceph{}
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		ceph.daemons = []Daemon{}
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
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
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
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		So(func() { ceph.CollectMetrics(mts) }, ShouldNotPanic)
		result, err := ceph.CollectMetrics(mts)

		So(len(result), ShouldEqual, len(mts))

		for _, r := range result {
			So(r.Data, ShouldNotBeNil)
		}
		So(err, ShouldBeNil)
	})

	Convey("collect dynamic metrics with no errors", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		So(func() { ceph.CollectMetrics(mtsDynamic) }, ShouldNotPanic)
		result, err := ceph.CollectMetrics(mtsDynamic)

		So(len(result), ShouldEqual, 4)

		for _, r := range result {
			So(r.Data, ShouldEqual, 1)
		}
		So(err, ShouldBeNil)
	})

	Convey("collect dynamic metrics for not existing metric", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		So(func() { ceph.CollectMetrics(mtsDynamicErr) }, ShouldNotPanic)
		result, err := ceph.CollectMetrics(mtsDynamicErr)
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 0)
	})

	os.Remove("test/osd.0.asok")
	os.Remove("test/osd.1.asok")
	os.Remove("test/osd.2.asok")
	os.Remove("test/osd.3.asok")

	Convey("collect dynamic metrics with no daemons available", t, func() {
		ceph := &Ceph{}
		cmd = &TestCmd{outDump: mockOut, outSchema: mockOutSchema}
		So(func() { ceph.CollectMetrics(mtsDynamic) }, ShouldNotPanic)
		result, err := ceph.CollectMetrics(mtsDynamic)

		So(result, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	os.Remove("test/")
}

func Test_New(t *testing.T) {
	Convey("create new Ceph structure", t, func() {
		So(func() { New() }, ShouldNotPanic)
		result := New()
		So(result, ShouldNotBeNil)
	})
}
