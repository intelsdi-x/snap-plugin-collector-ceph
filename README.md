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

## Pulse Ceph Perf Counters Collector Plugin

# Description
Collect Ceph performance counters from the Ceph Storage System for:
* MON (Ceph Monitor Daemon)
* OSD (Ceph Object Storage Daemon)
* MDS (Ceph Metadata Server Daemon)

The perf counters data are accessed via the Ceph admin socket.
The intention is that data will be collected, aggregated and fed into another tool for graphing and analysis.

Documentation for ceph perf counters:	http://ceph.com/docs/master/dev/perf_counters/

# Dependencies
* Ceph Storage Cluster [http://ceph.com/]
* Ceph Administration Tool


# Configuration
Set proper Pulse Global Config field(s) to customize Ceph's path

Ceph Config field | Description | Default
------------ | ------------- | -------------
**path** | Path to "ceph" executable | "/usr/bin/ceph"
**socket_path** | The location of the ceph monitoring sockets | "/var/run/ceph"
**socket_prefix** | The first part of all socket names | "ceph-"
**socket_ext** | Extension for socket filenames| "asok"
If sockets do not have prefix, set *socket_prefix="none"*

By default metrics are gathered once per second.
# Limitations
Root privileges might be needed

# Metrics
Collect all available Ceph perf counters from:
* MON [see more...](MON_PERFCNT.md)
* MDS [see more...](MDS_PERFCNT.md)
* OSD [see more...](OSD_PERFCNT.md)

