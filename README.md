DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.
# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**

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

The quickest way to get a Ceph cluster up and running is to follow the Getting Started guides available at http://ceph.com/resources/downloads/.                 It can be tested also on a fake local cluster on Your machine. Read more about [how to deploy fake local Ceph cluster](VCLUSTER.md).

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
**path** | string | Path to "ceph" executable. Defaults to"/usr/bin"
**socket_path** | string | The location of the ceph monitoring sockets. Defaults to "/var/run/ceph"
**socket_prefix** | string | The first part of socket names. Defaults to "ceph-"
**socket_ext** | string | Extension for socket filenames. Defaults to "asok"
If sockets do not have prefix, set *socket_prefix="none"*

Sample Global Config is available in folder /examples/configs.

## Documentation

To learn more about this plugin and ceph perf counters, visit:

* [ceph perf counters doc](http://ceph.com/docs/master/dev/perf_counters)
* [Snap ceph unit test](https://github.com/intelsdi-x/snap-plugin-collector-ceph/blob/master/ceph/ceph_test.go)
* [Snap ceph examples](#examples)

Resetting the perf counters before measurement is recommended. This feature was added in the Ceph 0.90 (Hammer release):
```
$ sudo ceph daemon <daemon-name> perf reset all | <perf_cnt_name>
```

### Collected Metrics
This plugin has the ability to dynamically gather all Ceph perf counters straight from your Ceph version:

* Example metric list for Ceph 11.0.2:
  - [MDS daemon](MDS_PERFCNT.md)
  - [MON daemon](MON_PERFCNT.md)
  - [OSD daemon](OSD_PERFCNT.md)

By default metrics are gathered once per second.


### Examples

Example of running Snap ceph perf counters collector and writing data to file.

Run snap daemon with default config values (see the table above):

```
$ $SNAP_PATH/snapteld -l 1 -t 0
```

Or with your custom Snap Global Config (see examples/configs/snap-config-sample.json):

```
$ $SNAP_PATH/snapteld -l 1 -t 0 --config $SNAP_CEPH_PLUGIN_DIR/examples/configs/snap-config-sample.json
```

Load ceph plugin for collecting:
```
$ $SNAP_PATH/snaptel plugin load $SNAP_CEPH_PLUGIN_DIR/build/linux/x86_64/snap-plugin-collector-ceph
Plugin loaded
Name: ceph
Version: 6
Type: collector
Signed: false
Loaded Time: Tue, 09 Feb 2017 06:19:48 EST
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
Loaded Time: Tue, 09 Feb 2017 07:45:58 EST
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
NAMESPACE                                                DATA                       TIMESTAMP
/intel/storage/ceph/mon/a/cluster/num_mon                2                          2017-02-09 17:13:39.474247945 +0100 CET
/intel/storage/ceph/mon/a/cluster/num_object             20                         2017-02-09 17:13:39.474247945 +0100 CET
/intel/storage/ceph/mon/a/cluster/num_osd                2                          2017-02-09 17:13:39.474247945 +0100 CET
/intel/storage/ceph/mon/a/cluster/num_pg                 24                         2017-02-09 17:13:39.474247945 +0100 CET
/intel/storage/ceph/mon/a/cluster/osd_bytes              9.01767725056e+11          2017-02-09 17:13:39.474247945 +0100 CET
/intel/storage/ceph/mon/a/cluster/osd_bytes_used         1.8365669376e+11           2017-02-09 17:13:39.474247945 +0100 CET
/intel/storage/ceph/mon/b/cluster/num_mon                2                          2017-02-09 17:13:39.549567355 +0100 CET
/intel/storage/ceph/mon/b/cluster/num_object             20                         2017-02-09 17:13:39.549567355 +0100 CET
/intel/storage/ceph/mon/b/cluster/num_osd                2                          2017-02-09 17:13:39.549567355 +0100 CET
/intel/storage/ceph/mon/b/cluster/num_pg                 24                         2017-02-09 17:13:39.549567355 +0100 CET
/intel/storage/ceph/mon/b/cluster/osd_bytes              9.01767725056e+11          2017-02-09 17:13:39.549567355 +0100 CET
/intel/storage/ceph/mon/b/cluster/osd_bytes_used         1.8365669376e+11           2017-02-09 17:13:39.549567355 +0100 CET
/intel/storage/ceph/osd/0/osd/op                         8                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_cache_hit               0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_in_bytes                0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_latency                 0.043267072875             2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_out_bytes               0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_process_latency         0.0179476545               2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_r                       0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_r_latency               0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_r_out_bytes             0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_r_process_latency       0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_rw                      0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_rw_in_bytes             0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_rw_latency              0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_rw_out_bytes            0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_rw_process_latency      0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/0/osd/op_rw_rlat                 0                          2017-02-09 17:13:39.608994105 +0100 CET
/intel/storage/ceph/osd/1/osd/op                         22                         2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_cache_hit               0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_in_bytes                2148                       2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_latency                 0.03500758777272727        2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_out_bytes               0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_process_latency         0.03330366313636363        2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_r                       0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_r_latency               0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_r_out_bytes             0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_r_process_latency       0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_rw                      0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_rw_in_bytes             0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_rw_latency              0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_rw_out_bytes            0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_rw_process_latency      0                          2017-02-09 17:13:39.715109149 +0100 CET
/intel/storage/ceph/osd/1/osd/op_rw_rlat                 0                          2017-02-09 17:13:39.715109149 +0100 CET
```
(Keys `ctrl+c` terminate task watcher)

These data are published to file and stored there (in this example in /tmp/published_ceph).

Stop task:
```
$ $SNAP_PATH/snaptel task stop 029cc837-ccd7-41b0-8103-949c0ba0070f
Task stopped:
ID: 029cc837-ccd7-41b0-8103-949c0ba0070f
```

Run plugin in diagnostic mode:
```
$ snap-plugin-collector-ceph --config '{"path": "/somepath/ceph/build/bin", "socket_path": "/somepath/ceph/src/out", "socket_ext": "asok", "socket_prefix": ""}'
```
*Note: Passing configuration is important, because this plugin builds metric catalog dynamically based on ceph perf schema output for all daemons.*

**Notice:**                                                                                                                                                                                                                                                                                        **Using the Snap tribe is recommended.** Administrators can control all Snap nodes in a tribe agreement by messaging just one of them what makes cluster configuration management simple. Read more about the Snap tribe at https://github.com/intelsdi-x/snap.

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
