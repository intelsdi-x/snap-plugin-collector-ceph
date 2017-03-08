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
### MON daemon
**Metric list was generated dynamically by plugin and can be different on another setup.**  
Prefix: `/intel/storage/ceph/mon/[mon_id]`

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
cluster/mds_epoch | uint64 | Current epoch of MDS map
cluster/num_bytes | uint64 | Size of all objects
cluster/num_mds_failed | uint64 | Failed MDS
cluster/num_mds_in | uint64 | MDS in state "in" (they are in cluster)
cluster/num_mds_up | uint64 | MDSs that are up
cluster/num_mon | uint64 | Monitors
cluster/num_mon_quorum | uint64 | Monitors in quorum
cluster/num_object | uint64 | Objects
cluster/num_object_degraded | uint64 | Degraded (missing replicas) objects
cluster/num_object_misplaced | uint64 | Misplaced (wrong location in the cluster) objects
cluster/num_object_unfound | uint64 | Unfound objects
cluster/num_osd | uint64 | OSDs
cluster/num_osd_in | uint64 | OSD in state "in" (they are in cluster)
cluster/num_osd_up | uint64 | OSDs that are up
cluster/num_pg | uint64 | Placement groups
cluster/num_pg_active | uint64 | Placement groups in active state
cluster/num_pg_active_clean | uint64 | Placement groups in active+clean state
cluster/num_pg_peering | uint64 | Placement groups in peering state
cluster/num_pool | uint64 | Pools
cluster/osd_bytes | uint64 | Total capacity of cluster
cluster/osd_bytes_avail | uint64 | Available space
cluster/osd_bytes_used | uint64 | Used space
cluster/osd_epoch | uint64 | Current epoch of OSD map
finisher-monstore/complete_latency | float | 
finisher-monstore/queue_len | uint64 | 
leveldb/leveldb_compact | uint64 | Compactions
leveldb/leveldb_compact_queue_len | uint64 | Length of compaction queue
leveldb/leveldb_compact_queue_merge | uint64 | Mergings of ranges in compaction queue
leveldb/leveldb_compact_range | uint64 | Compactions by range
leveldb/leveldb_get | uint64 | Gets
leveldb/leveldb_get_latency | float | Get Latency
leveldb/leveldb_submit_latency | float | Submit Latency
leveldb/leveldb_submit_sync_latency | float | Submit Sync Latency
leveldb/leveldb_transaction | uint64 | Transactions
mon/election_call | uint64 | Elections started
mon/election_lose | uint64 | Elections lost
mon/election_win | uint64 | Elections won
mon/num_elections | uint64 | Elections participated in
mon/num_sessions | uint64 | Open sessions
mon/session_add | uint64 | Created sessions
mon/session_rm | uint64 | Removed sessions
mon/session_trim | uint64 | Trimmed sessions
paxos/accept_timeout | uint64 | Accept timeouts
paxos/begin | uint64 | Started and handled begins
paxos/begin_bytes | uint64 | Data in transaction on begin
paxos/begin_keys | uint64 | Keys in transaction on begin
paxos/begin_latency | float | Latency of begin operation
paxos/collect | uint64 | Peon collects
paxos/collect_bytes | uint64 | Data in transaction on peon collect
paxos/collect_keys | uint64 | Keys in transaction on peon collect
paxos/collect_latency | float | Peon collect latency
paxos/collect_timeout | uint64 | Collect timeouts
paxos/collect_uncommitted | uint64 | Uncommitted values in started and handled collects
paxos/commit | uint64 | Commits
paxos/commit_bytes | uint64 | Data in transaction on commit
paxos/commit_keys | uint64 | Keys in transaction on commit
paxos/commit_latency | float | Commit latency
paxos/lease_ack_timeout | uint64 | Lease acknowledgement timeouts
paxos/lease_timeout | uint64 | Lease timeouts
paxos/new_pn | uint64 | New proposal number queries
paxos/new_pn_latency | float | New proposal number getting latency
paxos/refresh | uint64 | Refreshes
paxos/refresh_latency | float | Refresh latency
paxos/restart | uint64 | Restarts
paxos/share_state | uint64 | Sharings of state
paxos/share_state_bytes | uint64 | Data in shared state
paxos/share_state_keys | uint64 | Keys in shared state
paxos/start_leader | uint64 | Starts in leader role
paxos/start_peon | uint64 | Starts in peon role
paxos/store_state | uint64 | Store a shared state on disk
paxos/store_state_bytes | uint64 | Data in transaction in stored state
paxos/store_state_keys | uint64 | Keys in transaction in stored state
paxos/store_state_latency | float | Storing state latency
throttle-mon_client_bytes/get | uint64 | Gets
throttle-mon_client_bytes/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-mon_client_bytes/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-mon_client_bytes/get_started | uint64 | Number of get calls, increased before wait
throttle-mon_client_bytes/get_sum | uint64 | Got data
throttle-mon_client_bytes/max | uint64 | Max value for throttle
throttle-mon_client_bytes/put | uint64 | Puts
throttle-mon_client_bytes/put_sum | uint64 | Put data
throttle-mon_client_bytes/take | uint64 | Takes
throttle-mon_client_bytes/take_sum | uint64 | Taken data
throttle-mon_client_bytes/val | uint64 | Currently available throttle
throttle-mon_client_bytes/wait | float | Waiting latency
throttle-mon_daemon_bytes/get | uint64 | Gets
throttle-mon_daemon_bytes/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-mon_daemon_bytes/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-mon_daemon_bytes/get_started | uint64 | Number of get calls, increased before wait
throttle-mon_daemon_bytes/get_sum | uint64 | Got data
throttle-mon_daemon_bytes/max | uint64 | Max value for throttle
throttle-mon_daemon_bytes/put | uint64 | Puts
throttle-mon_daemon_bytes/put_sum | uint64 | Put data
throttle-mon_daemon_bytes/take | uint64 | Takes
throttle-mon_daemon_bytes/take_sum | uint64 | Taken data
throttle-mon_daemon_bytes/val | uint64 | Currently available throttle
throttle-mon_daemon_bytes/wait | float | Waiting latency
throttle-msgr_dispatch_throttler-mon/get | uint64 | Gets
throttle-msgr_dispatch_throttler-mon/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-mon/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-mon/get_started | uint64 | Number of get calls, increased before wait
throttle-msgr_dispatch_throttler-mon/get_sum | uint64 | Got data
throttle-msgr_dispatch_throttler-mon/max | uint64 | Max value for throttle
throttle-msgr_dispatch_throttler-mon/put | uint64 | Puts
throttle-msgr_dispatch_throttler-mon/put_sum | uint64 | Put data
throttle-msgr_dispatch_throttler-mon/take | uint64 | Takes
throttle-msgr_dispatch_throttler-mon/take_sum | uint64 | Taken data
throttle-msgr_dispatch_throttler-mon/val | uint64 | Currently available throttle
throttle-msgr_dispatch_throttler-mon/wait | float | Waiting latency
