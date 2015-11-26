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

# snap Collector Plugin- CEPH

Collect Ceph performance counters from the Ceph Storage System for:
* MON (Ceph Monitor Daemon)
* OSD (Ceph Object Storage Daemon)
* MDS (Ceph Metadata Server Daemon)

The perf counters data are accessed via the Ceph admin socket.
The intention is that data will be collected, aggregated and fed into another tool for graphing and analysis.

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

### System Requirements

Include:

- Ceph Storage Cluster [http://ceph.com/]
- Ceph Administration Tool
- Root privileges might be needed

### Installation

### Configuration and Usage

Set proper snap Global Config field(s) to customize Ceph's path

## Documentation

### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type | Description (optional)
----------|-----------|-----------------------
**path** |  | Path to "ceph" executable "/usr/bin/ceph"
**socket_path** | | The location of the ceph monitoring sockets "/var/run/ceph"
**socket_prefix** | | The first part of all socket names "ceph-"
**socket_ext** | | Extension for socket filenames "asok"
If sockets do not have prefix, set *socket_prefix="none"*

By default metrics are gathered once per second.

Collect all available Ceph perf counters from:
* MON [see more...](MON_PERFCNT.md)
* MDS [see more...](MDS_PERFCNT.md)
* OSD [see more...](OSD_PERFCNT.md)

### Examples

### Roadmap

## Community Support
This repository is one of **many** plugins in the **snap Framework**: a powerful telemetry agent framework. To reach out on other use cases, visit:

* snap Gitter channel (@TODO Link)
* Our Google Group (@TODO Link)

The full project is at http://github.com:intelsdi-x/snap.

## Contributing
We love contributions! :heart_eyes:

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
snap, along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
List authors, co-authors and anyone you'd like to mention

* Author: [Izabella Raulin](https://github.com/IzabellaRaulin)

**Thank you!** Your contribution is incredibly important to us.
