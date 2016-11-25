# Snap collector plugin - Ceph

This  plugin collects Ceph performance counters from the Ceph Storage System for:
* MON (Ceph Monitor Daemon)
* OSD (Ceph Object Storage Daemon)
* MDS (Ceph Metadata Server Daemon)

The perf counters data are accessed via the Ceph admin socket.
The intention is that data will be collected, aggregated and fed into another tool for graphing and analysis.

This plugin is used in the [Snap framework] (http://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

In order to use this plugin you need Ceph cluster running.

### System Requirements

* [Ceph Storage Cluster] (http://ceph.com/)
* [Ceph Administration Tool] (http://docs.ceph.com/docs/v9.0.2/man/8/ceph/)
* Root privileges might be needed

### Installation

#### To deploy the Ceph cluster

The quickest way to get a Ceph cluster up and running is to follow the Getting Started guides available at http://ceph.com/resources/downloads/. 		It can be tested also on a fake local cluster on Your machine. Read more about [how to deploy fake local Ceph cluster](VCLUSTER.md).

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-ceph
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-ceph.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/${GOOS}/${GOARCH}`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/linux/x86_64`

* Set proper Snap Global Config field(s) to customize Ceph's path:

Namespace | Data Type | Description
----------|-----------|-----------------------
**path** | string | Path to "ceph" executable. Defaults to"/usr/bin/ceph"
**socket_path** | string | The location of the ceph monitoring sockets. Defaults to "/var/run/ceph"
**socket_prefix** | string | The first part of socket names. Defaults to "ceph-"
**socket_ext** | string | Extension for socket filenames. Defaults to "asok"
If sockets do not have prefix, set *socket_prefix="none"*

Sample Global Config is available in folder /examples/configs.

## Documentation

To learn more about this plugin and ceph perf counters, visit:

* [ceph perf counters doc] (http://ceph.com/docs/master/dev/perf_counters)
* [Snap ceph unit test](https://github.com/intelsdi-x/snap-plugin-collector-ceph/blob/master/ceph/ceph_test.go)
* [Snap ceph examples](#examples)

Resetting the perf counters before measurement is recommended. This feature was added in the Ceph 0.90 (Hammer release):
```
$ sudo ceph daemon <daemon-name> perf reset all | <perf_cnt_name>
```

### Collected Metrics
This plugin has the ability to gather the following Ceph perf counters from:

* MON [see more...](MON_PERFCNT.md)
* MDS [see more...](MDS_PERFCNT.md)
* OSD [see more...](OSD_PERFCNT.md)

By default metrics are gathered once per second.


### Examples

Example of running Snap ceph perf counters collector and writing data to file.

Run the Snap daemon on each node with defaults settings:
```
$ $SNAP_PATH/snapteld -l 1 -t 0
```
Or set custom settings in Snap Global Config in Ceph section (see examples/configs/snap-config-sample.json):
```
$ $SNAP_PATH/snapteld -l 1 -t 0 --config $SNAP_CEPH_PLUGIN_DIR/examples/configs/snap-config-sample.json
```

Load ceph plugin for collecting:
```
$ $SNAP_PATH/snaptel plugin load $SNAP_CEPH_PLUGIN_DIR/build/linux/x86_64/snap-plugin-collector-ceph
Plugin loaded
Name: ceph
Version: 4
Type: collector
Signed: false
Loaded Time: Tue, 01 Dec 2015 06:19:48 EST
```
See available metrics for all ceph-daemon in cluster:
```
$ $SNAP_PATH/snaptel metric list
```

Or see available metrics only for OSDs:
```
$ $SNAP_PATH/snaptel metric list | grep ceph/osd
```
Download desired publisher plugin eg.
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
```
Load file plugin for publishing:
```
$ $SNAP_PATH/snaptel plugin load snap-plugin-publisher-file
Plugin loaded
Name: file
Version: 4
Type: publisher
Signed: false
Loaded Time: Tue, 01 Dec 2015 07:45:58 EST
```

Create a task JSON file (exemplary file in examples/tasks/ceph-file.json):
```json
{
  "version": 1,
  "schedule": {
      "type": "simple",
      "interval": "1s"
  },
  "workflow": {
    "collect": {
      "metrics": {
        "/intel/storage/ceph/mon/*/cluster/num_mon": {},
        "/intel/storage/ceph/mon/*/cluster/num_osd": {},
        "/intel/storage/ceph/mon/*/cluster/num_object": {},
        "/intel/storage/ceph/mon/*/cluster/num_pg": {},
        "/intel/storage/ceph/mon/*/cluster/osd_bytes_used": {},
        "/intel/storage/ceph/mon/*/cluster/osd_bytes": {},
        "/intel/storage/ceph/mon/*/cluster/osd_bytes_used": {},
        "/intel/storage/ceph/mon/*/cluster/osd_bytes": {},
        "/intel/storage/ceph/osd/*/osd/op": {},
        "/intel/storage/ceph/osd/*/osd/op_cache_hit": {},
        "/intel/storage/ceph/osd/*/osd/op_in_bytes": {},
        "/intel/storage/ceph/osd/*/osd/op_latency": {},
        "/intel/storage/ceph/osd/*/osd/op_out_bytes": {},
        "/intel/storage/ceph/osd/*/osd/op_process_latency": {},
        "/intel/storage/ceph/osd/*/osd/op_r": {},
        "/intel/storage/ceph/osd/*/osd/op_r_latency": {},
        "/intel/storage/ceph/osd/*/osd/op_r_out_bytes": {},
        "/intel/storage/ceph/osd/*/osd/op_r_process_latency": {},
        "/intel/storage/ceph/osd/*/osd/op_rw": {},
        "/intel/storage/ceph/osd/*/osd/op_rw_in_bytes": {},
        "/intel/storage/ceph/osd/*/osd/op_rw_latency": {},
        "/intel/storage/ceph/osd/*/osd/op_rw_out_bytes": {},
        "/intel/storage/ceph/osd/*/osd/op_rw_process_latency": {},
        "/intel/storage/ceph/osd/*/osd/op_rw_rlat": {}
      },
      "config": {
          "/intel/storage/ceph": {
              "user": "root",
              "password": "secret"
          }
      },
      "publish": [
          {
              "plugin_name": "file",
              "config": {
                  "file": "/tmp/snap-ceph.file.log"
              }
          }
      ]
    }
  }
}
```

Create a task:
```
$ $SNAP_PATH/snaptel task create -t $SNAP_CEPH_PLUGIN_DIR/examples/tasks/ceph-file.json
Using task manifest to create task
Task created
ID: 029cc837-ccd7-41b0-8103-949c0ba0070f
Name: Task-029cc837-ccd7-41b0-8103-949c0ba0070f
State: Running
```

See sample output from `snaptel task watch <task_id>`

```
$ $SNAP_PATH/snaptel task watch 029cc837-ccd7-41b0-8103-949c0ba0070f

Watching Task (029cc837-ccd7-41b0-8103-949c0ba0070f):
NAMESPACE                                                        DATA                    TIMESTAMP
/intel/storage/ceph/mon/a/cluster/num_mon                        1                       2016-05-25 12:46:25.008948694 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_object                     20                      2016-05-25 12:46:25.008948694 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_osd                        2                       2016-05-25 12:46:25.008948694 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_pg                         24                      2016-05-25 12:46:25.008948694 +0200 CEST
/intel/storage/ceph/mon/a/cluster/osd_bytes                      2.147483648e+09         2016-05-25 12:46:25.008948694 +0200 CEST
/intel/storage/ceph/mon/a/cluster/osd_bytes_used                 40960                   2016-05-25 12:46:25.008948694 +0200 CEST
/intel/storage/ceph/osd/0/osd/op                                 8                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_cache_hit                       0                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_in_bytes                        1032                    2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_latency/avgcount                8                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_latency/sum                     0.082809936             2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_out_bytes                       0                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_process_latency/avgcount        8                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_process_latency/sum             0.036669336             2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r                               0                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_latency/avgcount              0                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_latency/sum                   0                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_out_bytes                     0                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_process_latency/avgcount      0                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_process_latency/sum           0                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw                              4                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_in_bytes                     1032                    2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_latency/avgcount             4                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_latency/sum                  0.03521104              2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_out_bytes                    0                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_process_latency/avgcount     4                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_process_latency/sum          0.023461402             2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_rlat/avgcount                4                       2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_rlat/sum                     0.034607209             2016-05-25 12:46:25.146922352 +0200 CEST
/intel/storage/ceph/osd/1/osd/op                                 22                      2016-05-25 12:46:25.28233031 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_cache_hit                       0                       2016-05-25 12:46:25.28233031 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_mon                        1                       2016-05-25 12:46:26.014505366 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_object                     20                      2016-05-25 12:46:26.014505366 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_osd                        2                       2016-05-25 12:46:26.014505366 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_pg                         24                      2016-05-25 12:46:26.014505366 +0200 CEST
/intel/storage/ceph/mon/a/cluster/osd_bytes                      2.147483648e+09         2016-05-25 12:46:26.014505366 +0200 CEST
/intel/storage/ceph/mon/a/cluster/osd_bytes_used                 40960                   2016-05-25 12:46:26.014505366 +0200 CEST
/intel/storage/ceph/osd/0/osd/op                                 8                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_cache_hit                       0                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_in_bytes                        1032                    2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_latency/avgcount                8                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_latency/sum                     0.082809936             2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_out_bytes                       0                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_process_latency/avgcount        8                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_process_latency/sum             0.036669336             2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r                               0                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_latency/avgcount              0                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_latency/sum                   0                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_out_bytes                     0                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_process_latency/avgcount      0                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_process_latency/sum           0                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw                              4                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_in_bytes                     1032                    2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_latency/avgcount             4                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_latency/sum                  0.03521104              2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_out_bytes                    0                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_process_latency/avgcount     4                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_process_latency/sum          0.023461402             2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_rlat/avgcount                4                       2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_rlat/sum                     0.034607209             2016-05-25 12:46:26.153578696 +0200 CEST
/intel/storage/ceph/osd/1/osd/op                                 22                      2016-05-25 12:46:26.289102505 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_cache_hit                       0                       2016-05-25 12:46:26.289102505 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_mon                        1                       2016-05-25 12:46:27.026023314 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_object                     20                      2016-05-25 12:46:27.026023314 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_osd                        2                       2016-05-25 12:46:27.026023314 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_pg                         24                      2016-05-25 12:46:27.026023314 +0200 CEST
/intel/storage/ceph/mon/a/cluster/osd_bytes                      2.147483648e+09         2016-05-25 12:46:27.026023314 +0200 CEST
/intel/storage/ceph/mon/a/cluster/osd_bytes_used                 40960                   2016-05-25 12:46:27.026023314 +0200 CEST
/intel/storage/ceph/osd/0/osd/op                                 8                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_cache_hit                       0                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_in_bytes                        1032                    2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_latency/avgcount                8                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_latency/sum                     0.082809936             2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_out_bytes                       0                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_process_latency/avgcount        8                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_process_latency/sum             0.036669336             2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r                               0                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_latency/avgcount              0                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_latency/sum                   0                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_out_bytes                     0                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_process_latency/avgcount      0                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_process_latency/sum           0                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw                              4                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_in_bytes                     1032                    2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_latency/avgcount             4                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_latency/sum                  0.03521104              2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_out_bytes                    0                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_process_latency/avgcount     4                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_process_latency/sum          0.023461402             2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_rlat/avgcount                4                       2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_rlat/sum                     0.034607209             2016-05-25 12:46:27.162922457 +0200 CEST
/intel/storage/ceph/osd/1/osd/op                                 22                      2016-05-25 12:46:27.316945185 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_cache_hit                       0                       2016-05-25 12:46:27.316945185 +0200 CEST
^Cntel/storage/ceph/mon/a/cluster/num_mon                        1                       2016-05-25 12:46:49.008879389 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_object                     20                      2016-05-25 12:46:49.008879389 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_osd                        2                       2016-05-25 12:46:49.008879389 +0200 CEST
/intel/storage/ceph/mon/a/cluster/num_pg                         24                      2016-05-25 12:46:49.008879389 +0200 CEST
/intel/storage/ceph/mon/a/cluster/osd_bytes                      2.147483648e+09         2016-05-25 12:46:49.008879389 +0200 CEST
/intel/storage/ceph/mon/a/cluster/osd_bytes_used                 40960                   2016-05-25 12:46:49.008879389 +0200 CEST
/intel/storage/ceph/osd/0/osd/op                                 8                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_cache_hit                       0                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_in_bytes                        1032                    2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_latency/avgcount                8                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_latency/sum                     0.082809936             2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_out_bytes                       0                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_process_latency/avgcount        8                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_process_latency/sum             0.036669336             2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r                               0                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_latency/avgcount              0                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_latency/sum                   0                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_out_bytes                     0                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_process_latency/avgcount      0                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_r_process_latency/sum           0                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw                              4                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_in_bytes                     1032                    2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_latency/avgcount             4                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_latency/sum                  0.03521104              2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_out_bytes                    0                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_process_latency/avgcount     4                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_process_latency/sum          0.023461402             2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_rlat/avgcount                4                       2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/0/osd/op_rw_rlat/sum                     0.034607209             2016-05-25 12:46:49.148291065 +0200 CEST
/intel/storage/ceph/osd/1/osd/op                                 22                      2016-05-25 12:46:49.283742514 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_cache_hit                       0                       2016-05-25 12:46:49.283742514 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_in_bytes                        8680                    2016-05-25 12:46:49.283742514 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_latency/avgcount                22                      2016-05-25 12:46:49.283742514 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_latency/sum                     0.197345214             2016-05-25 12:46:49.283742514 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_out_bytes                       0                       2016-05-25 12:46:49.283742514 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_process_latency/avgcount        22                      2016-05-25 12:46:49.283742514 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_process_latency/sum             0.139076264             2016-05-25 12:46:49.283742514 +0200 CEST
/intel/storage/ceph/osd/1/osd/op_r                               0                       2016-05-25 12:46:49.283742514 +0200 CEST
```
(Keys `ctrl+c` terminate task watcher)

These data are published to file and stored there (in this example in /tmp/published_ceph).

Stop task:
```
$ $SNAP_PATH/snaptel task stop 029cc837-ccd7-41b0-8103-949c0ba0070f
Task stopped:
ID: 029cc837-ccd7-41b0-8103-949c0ba0070f
```

**Notice:**																																			**Using the Snap tribe is recommended.** Administrators can control all Snap nodes in a tribe agreement by messaging just one of them what makes cluster configuration management simple. Read more about the Snap tribe at https://github.com/intelsdi-x/snap.

### Roadmap
This plugin is in active development. As we launch this plugin, we have a few items in mind for the next release:
- [ ] Concurrency execution of collecting metrics from Ceph sockets

As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-ceph/issues).

## Community Support
This repository is one of **many** plugins in the **Snap**, a powerful telemetry agent framework. See the full project at
http://github.com/intelsdi-x/snap. To reach out on other use cases, visit [Slack](http://slack.snap-telemetry.io).


## Contributing
We love contributions! :heart_eyes:

There is more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).


## Acknowledgements

* Author: [Izabella Raulin](https://github.com/IzabellaRaulin)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
