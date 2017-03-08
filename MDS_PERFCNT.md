<!--
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2017 Intel Corporation

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

# Snap Ceph Perf Counters Collector Plugin
## Metric list for Ceph 11.0.2
### MDS daemon
**Metric list was generated dynamically by plugin and can be different on another setup.**  
Prefix: `/intel/storage/ceph/mds/[mds_id]`


NAMESPACE | UNIT | DESCRIPTION
----------|------|-------------
AsyncMessenger::Worker-0/msgr_active_connections | uint64 | Active connection number
AsyncMessenger::Worker-0/msgr_created_connections | uint64 | Created connection number
AsyncMessenger::Worker-0/msgr_recv_bytes | uint64 | Network received bytes
AsyncMessenger::Worker-0/msgr_recv_messages | uint64 | Network received messages
AsyncMessenger::Worker-0/msgr_send_bytes | uint64 | Network received bytes
AsyncMessenger::Worker-0/msgr_send_messages | uint64 | Network sent messages
AsyncMessenger::Worker-0/msgr_send_messages_inline | uint64 | Network sent inline messages
AsyncMessenger::Worker-1/msgr_active_connections | uint64 | Active connection number
AsyncMessenger::Worker-1/msgr_created_connections | uint64 | Created connection number
AsyncMessenger::Worker-1/msgr_recv_bytes | uint64 | Network received bytes
AsyncMessenger::Worker-1/msgr_recv_messages | uint64 | Network received messages
AsyncMessenger::Worker-1/msgr_send_bytes | uint64 | Network received bytes
AsyncMessenger::Worker-1/msgr_send_messages | uint64 | Network sent messages
AsyncMessenger::Worker-1/msgr_send_messages_inline | uint64 | Network sent inline messages
AsyncMessenger::Worker-2/msgr_active_connections | uint64 | Active connection number
AsyncMessenger::Worker-2/msgr_created_connections | uint64 | Created connection number
AsyncMessenger::Worker-2/msgr_recv_bytes | uint64 | Network received bytes
AsyncMessenger::Worker-2/msgr_recv_messages | uint64 | Network received messages
AsyncMessenger::Worker-2/msgr_send_bytes | uint64 | Network received bytes
AsyncMessenger::Worker-2/msgr_send_messages | uint64 | Network sent messages
AsyncMessenger::Worker-2/msgr_send_messages_inline | uint64 | Network sent inline messages
throttle-msgr_dispatch_throttler-mds/get | uint64 | Gets
throttle-msgr_dispatch_throttler-mds/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-mds/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-mds/get_started | uint64 | Number of get calls, increased before wait
throttle-msgr_dispatch_throttler-mds/get_sum | uint64 | Got data
throttle-msgr_dispatch_throttler-mds/max | uint64 | Max value for throttle
throttle-msgr_dispatch_throttler-mds/put | uint64 | Puts
throttle-msgr_dispatch_throttler-mds/put_sum | uint64 | Put data
throttle-msgr_dispatch_throttler-mds/take | uint64 | Takes
throttle-msgr_dispatch_throttler-mds/take_sum | uint64 | Taken data
throttle-msgr_dispatch_throttler-mds/val | uint64 | Currently available throttle
throttle-msgr_dispatch_throttler-mds/wait | float | Waiting latency
