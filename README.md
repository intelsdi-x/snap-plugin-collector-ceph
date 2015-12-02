<!--
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
-->

# Snap Collector Plugin - CEPH

Collect Ceph performance counters from the Ceph Storage System for:
* MON (Ceph Monitor Daemon)
* OSD (Ceph Object Storage Daemon)
* MDS (Ceph Metadata Server Daemon)

The perf counters data are accessed via the Ceph admin socket.
The intention is that data will be collected, aggregated and fed into another tool for graphing and analysis.

Project link: https://github.com/intelsdi-x/snap-plugin-collector-ceph

Documentation for ceph perf counters:	http://ceph.com/docs/master/dev/perf_counters/

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

In order to use this plugin you need ceph cluster running. 

It can be tested also on a fake local cluster on Your machine. Read more about [how to create virtual Ceph cluster](VCLUSTER.md).  

### System Requirements

Include:

- Ceph Storage Cluster [http://ceph.com/]
- Ceph Administration Tool
- Root privileges might be needed

### Installation

 - Plugin compilation
```
make
```

### Configuration and Usage

Set proper Snap Global Config field(s) to customize Ceph's path:

Namespace | Data Type | Description
----------|-----------|-----------------------
**path** | string | Path to "ceph" executable. Defaults to"/usr/bin/ceph"
**socket_path** | string | The location of the ceph monitoring sockets. Defaults to "/var/run/ceph"
**socket_prefix** | string | The first part of socket names. Defaults to "ceph-"
**socket_ext** | string | Extension for socket filenames. Defaults to "asok"
If sockets do not have prefix, set *socket_prefix="none"*

Sample Global Config is available in folder /examples/configs.
 
## Documentation

Documentation for ceph perf counters is available at http://ceph.com/docs/master/dev/perf_counters/

Resetting the perf counters before measurement is recommended. This feature was added in the Ceph 0.90 (Hammer release):
```
$ sudo ceph daemon <daemon-name> perf reset all | <perf_cnt_name>
```

### Collected Metrics
This plugin has the ability to gather the following Ceph perf counters from :

* MON [see more...](MON_PERFCNT.md)
* MDS [see more...](MDS_PERFCNT.md)
* OSD [see more...](OSD_PERFCNT.md)

By default metrics are gathered once per second.



### Examples

Example of running snap ceph perf counters collector and writing data to a file.

Run the snap daemon on each node with defaults settings:
```
$ snapd -l 1 -t 0
```
Or set custom settings in the Snap Global Config in Ceph section (see examples/configs/pulse-config-sample.json):
```
$ snapd -l 1 -t 0 --config $SNAP_CEPH_PLUGIN_DIR/examples/configs/pulse-config-sample.json
```

Load ceph plugin for collecting:
```
$ snapctl plugin load $SNAP_CEPH_PLUGIN_DIR/build/rootfs/snap-plugin-collector-ceph
Plugin loaded
Name: ceph
Version: 1
Type: collector
Signed: false
Loaded Time: Tue, 01 Dec 2015 06:19:48 EST
```
See available metrics for all ceph-daemon in cluster:
```
$ snapctl metric list
```

Or see available metrics only for OSDs:
```
$ snapctl metric list | grep ceph/osd
```

Load file plugin for publishing:
```
$ snapctl plugin load $SNAP_DIR/build/plugin/snap-publisher-file
Plugin loaded
Name: file
Version: 3
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
                "/intel/storage/ceph/mon.a/cluster/num_mon": {},
                "/intel/storage/ceph/mon.a/cluster/num_osd": {},
                "/intel/storage/ceph/mon.a/cluster/num_object": {},
                "/intel/storage/ceph/mon.a/cluster/num_pg": {},
                "/intel/storage/ceph/mon.a/cluster/osd_bytes_used": {},
                "/intel/storage/ceph/mon.a/cluster/osd_bytes": {},
                "/intel/storage/ceph/mon.b/cluster/osd_bytes_used": {},
                "/intel/storage/ceph/mon.b/cluster/osd_bytes": {},
                "/intel/storage/ceph/osd.0/filestore/bytes": {},
                "/intel/storage/ceph/osd.0/filestore/journal_bytes": {},
                "/intel/storage/ceph/osd.0/filestore/journal_latency/avgcount": {},
                "/intel/storage/ceph/osd.0/filestore/journal_latency/sum": {},
                "/intel/storage/ceph/osd.0/filestore/journal_queue_bytes": {},
                "/intel/storage/ceph/osd.0/filestore/journal_queue_max_bytes": {},
                "/intel/storage/ceph/osd.0/filestore/journal_wr": {},
                "/intel/storage/ceph/osd.0/filestore/op_queue_bytes": {},
                "/intel/storage/ceph/osd.0/filestore/op_queue_max_bytes": {},
                "/intel/storage/ceph/osd.0/filestore/queue_transaction_latency_avg/avgcount": {},
                "/intel/storage/ceph/osd.0/filestore/queue_transaction_latency_avg/sum": {},
                "/intel/storage/ceph/osd.0/osd/op": {},
                "/intel/storage/ceph/osd.0/osd/op_in_bytes": {},
                "/intel/storage/ceph/osd.0/osd/op_latency/avgcount": {},
                "/intel/storage/ceph/osd.0/osd/op_latency/sum": {},
                "/intel/storage/ceph/osd.0/osd/op_process_latency/avgcount": {},
                "/intel/storage/ceph/osd.0/osd/op_process_latency/sum": {},
                "/intel/storage/ceph/osd.0/osd/op_w_in_bytes": {},
                "/intel/storage/ceph/osd.0/osd/op_w_latency/avgcount": {},
                "/intel/storage/ceph/osd.0/osd/op_w_latency/sum": {},
                "/intel/storage/ceph/osd.0/osd/op_w_process_latency/avgcount": {},
                "/intel/storage/ceph/osd.0/osd/op_w_process_latency/sum": {},
                "/intel/storage/ceph/osd.1/filestore/bytes": {},
                "/intel/storage/ceph/osd.1/filestore/journal_bytes": {},
                "/intel/storage/ceph/osd.1/filestore/journal_latency/avgcount": {},
                "/intel/storage/ceph/mds.a/objecter/op_w": {},
                "/intel/storage/ceph/mds.b/objecter/op_w": {}
            },
            "config": {
                "/intel/storage/ceph": {
                    "user": "root",
                    "password": "secret"
                }
            },
            "process": null,
            "publish": [
                {
                    "plugin_name": "file",
                    "plugin_version": 3,
                    "config": {
                        "file": "/tmp/published_ceph"
                    }
                }
            ]
        }
    }
}
```

Create a task:
```
$ snapctl task create -t $SNAP_CEPH_PLUGIN_DIR/examples/tasks/ceph-file.json
Using task manifest to create task
Task created
ID: 029cc837-ccd7-41b0-8103-949c0ba0070f
Name: Task-029cc837-ccd7-41b0-8103-949c0ba0070f
State: Running
```

See sample output from `snapctl task watch <task_id>`

```
$ snapctl task watch 029cc837-ccd7-41b0-8103-949c0ba0070f

Watching Task (029cc837-ccd7-41b0-8103-949c0ba0070f):
NAMESPACE                                                                        DATA                    TIMESTAMP                                   SOURCE
/intel/storage/ceph/mds.a/objecter/op_w                                          36                      2015-12-01 07:48:49.942933001 -0500 EST         gklab-108-166/mds.a
/intel/storage/ceph/mds.b/objecter/op_w                                          33                      2015-12-01 07:48:50.000874606 -0500 EST         gklab-108-166/mds.b
/intel/storage/ceph/mon.a/cluster/num_mon                                        3                       2015-12-01 07:48:50.122093748 -0500 EST         gklab-108-166/mon.a
/intel/storage/ceph/mon.a/cluster/num_object                                     54                      2015-12-01 07:48:50.122093748 -0500 EST         gklab-108-166/mon.a
/intel/storage/ceph/mon.a/cluster/num_osd                                        3                       2015-12-01 07:48:50.122093748 -0500 EST         gklab-108-166/mon.a
/intel/storage/ceph/mon.a/cluster/num_pg                                         24                      2015-12-01 07:48:50.122093748 -0500 EST         gklab-108-166/mon.a
/intel/storage/ceph/mon.a/cluster/osd_bytes                                      4.35501797376e+11       2015-12-01 07:48:50.122093748 -0500 EST         gklab-108-166/mon.a
/intel/storage/ceph/mon.a/cluster/osd_bytes_used                                 1.21718276096e+11       2015-12-01 07:48:50.122093748 -0500 EST         gklab-108-166/mon.a
/intel/storage/ceph/mon.b/cluster/osd_bytes                                      4.35501797376e+11       2015-12-01 07:48:50.178708858 -0500 EST         gklab-108-166/mon.b
/intel/storage/ceph/mon.b/cluster/osd_bytes_used                                 1.21718276096e+11       2015-12-01 07:48:50.178708858 -0500 EST         gklab-108-166/mon.b
/intel/storage/ceph/osd.0/filestore/bytes                                        387402                  2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/journal_bytes                                0                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/journal_latency/avgcount                     0                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/journal_latency/sum                          0                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/journal_queue_bytes                          0                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/journal_queue_max_bytes                      0                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/journal_wr                                   0                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/op_queue_bytes                               0                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/op_queue_max_bytes                           1.048576e+08            2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/queue_transaction_latency_avg/avgcount       216                     2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/filestore/queue_transaction_latency_avg/sum            4.0477e-05              2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op                                                 4                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_in_bytes                                        444                     2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_latency/avgcount                                4                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_latency/sum                                     33.209642932            2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_process_latency/avgcount                        4                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_process_latency/sum                             20.001234107            2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_w_in_bytes                                      0                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_w_latency/avgcount                              2                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_w_latency/sum                                   21.607948808            2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_w_process_latency/avgcount                      2                       2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.0/osd/op_w_process_latency/sum                           10.006580525            2015-12-01 07:48:50.295831645 -0500 EST         gklab-108-166/osd.0
/intel/storage/ceph/osd.1/filestore/bytes                                        367731                  2015-12-01 07:48:50.374545867 -0500 EST         gklab-108-166/osd.1
/intel/storage/ceph/osd.1/filestore/journal_bytes                                0                       2015-12-01 07:48:50.374545867 -0500 EST         gklab-108-166/osd.1
/intel/storage/ceph/osd.1/filestore/journal_latency/avgcount                     0                       2015-12-01 07:48:50.374545867 -0500 EST         gklab-108-166/osd.1
```
(Keys `ctrl+c` terminate task watcher)

These data are published to file (in this example publishing to /tmp/published_ceph).


**Notice:**
**Using the snap tribe is recommended.** Administrators can control all snap nodes in a tribe agreement by messaging just one of them what makes cluster configuration management simple. Read more about the snap tribe at https://github.com/intelsdi-x/snap.

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-ceph/issues).


## Community Support
This repository is one of **many** plugins in the **snap Framework**: a powerful telemetry agent framework.
The full project is at http://github.com:intelsdi-x/snap.

## Contributing
We love contributions! :heart_eyes:

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
Snap, along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
List authors, co-authors and anyone you'd like to mention

* Author: [Izabella Raulin](https://github.com/IzabellaRaulin)

**Thank you!** Your contribution is incredibly important to us.
