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


# Ceph MON Perf Counters

Prefix: /intel/storage/ceph/{mon_daemon_name}/

Metrics | Description
------------ | -------------
cluster/num_mon | Monitors
cluster/num_mon_quorum | Monitors in quorum
cluster/num_osd | OSDs
cluster/num_osd_up | OSDs that are up
cluster/num_osd_in | OSD in state \in\ (they are in cluster)
cluster/osd_epoch | Current epoch of OSD map
cluster/osd_bytes | Total capacity of cluster
cluster/osd_bytes_used | Used space
cluster/osd_bytes_avail | Available space
cluster/num_pool | Pools
cluster/num_pg | Placement groups
cluster/num_pg_active_clean | Placement groups in active+clean state
cluster/num_pg_active | Placement groups in active state
cluster/num_pg_peering | Placement groups in peering state
cluster/num_object | Objects
cluster/num_object_degraded | Degraded (missing replicas) objects
cluster/num_object_misplaced | Misplaced (wrong location in the cluster) objects
cluster/num_object_unfound | Unfound objects
cluster/num_bytes | Size of all objects
cluster/num_mds_up | MDSs that are up
cluster/num_mds_in | MDS in state \in\ (they are in cluster)
cluster/num_mds_failed | Failed MDS
cluster/mds_epoch | Current epoch of MDS map
 | 
finisher-monstore/queue_len | Length of queue in finisher-monstore
 | 
leveldb/leveldb_get | Gets
leveldb/leveldb_transaction | Transactions
leveldb/leveldb_get_latency | Get Latency
leveldb/leveldb_submit_latency | Submit Latency
leveldb/leveldb_submit_sync_latency | Submit Sync Latency
leveldb/leveldb_compact | Compactions
leveldb/leveldb_compact_range | Compactions by range
leveldb/leveldb_compact_queue_merge | Mergings of ranges in compaction queue
leveldb/leveldb_compact_queue_len | Length of compaction queue
 | 
mon/num_sessions | Open sessions
mon/session_add | Created sessions
mon/session_rm | Removed sessions
mon/session_trim | Trimmed sessions
mon/num_elections | Elections participated in
mon/election_call | Elections started
mon/election_win | Elections won
mon/election_lose | Elections lost
 | 
paxos/start_leader | Starts in leader role
paxos/start_peon | Starts in peon role
paxos/restart | Restarts
paxos/refresh | Refreshes
paxos/refresh_latency/avgcount | Refresh latency
paxos/refresh_latency/sum | Refresh latency
paxos/begin | Started and handled begins
paxos/begin_keys/avgcount | Keys in transaction on begin
paxos/begin_keys/sum | Keys in transaction on begin
paxos/begin_bytes/avgcount | Data in transaction on begin
paxos/begin_bytes/sum | Data in transaction on begin
paxos/begin_latency/avgcount | Latency of begin operation
paxos/begin_latency/sum | Latency of begin operation
paxos/commit | Commits
paxos/commit_keys/avgcount | Keys in transaction on commit
paxos/commit_keys/sum | Keys in transaction on commit
paxos/commit_bytes/avgcount | Data in transaction on commit
paxos/commit_bytes/sum | Data in transaction on commit
paxos/commit_latency/avgcount | Commit latency
paxos/commit_latency/sum | Commit latency
paxos/collect | Peon collects
paxos/collect_keys/avgcount | Keys in transaction on peon collect
paxos/collect_keys/sum | Keys in transaction on peon collect
paxos/collect_bytes/avgcount | Data in transaction on peon collect
paxos/collect_bytes/sum | Data in transaction on peon collect
paxos/collect_latency/avgcount | Peon collect latency
paxos/collect_latency/sum | Peon collect latency
paxos/collect_uncommitted | Uncommitted values in started and handled collects
paxos/collect_timeout | Collect timeouts
paxos/accept_timeout | Accept timeouts
paxos/lease_ack_timeout | Lease acknowledgement timeouts
paxos/lease_timeout | Lease timeouts
paxos/store_state | Store a shared state on disk
paxos/store_state_keys/avgcount | Keys in transaction in stored state
paxos/store_state_keys/sum | Keys in transaction in stored state
paxos/store_state_bytes/avgcount | Data in transaction in stored state
paxos/store_state_bytes/sum | Data in transaction in stored state
paxos/store_state_latency/avgcount | Storing state latency
paxos/store_state_latency/sum | Storing state latency
paxos/share_state | Sharings of state
paxos/share_state_keys/avgcount | Keys in shared state
paxos/share_state_keys/sum | Keys in shared state
paxos/share_state_bytes/avgcount | Data in shared state
paxos/share_state_bytes/sum | Data in shared state
paxos/new_pn | New proposal number queries
paxos/new_pn_latency/avgcount | New proposal number getting latency
paxos/new_pn_latency/sum | New proposal number getting latency
 | 
throttle-mon_client_bytes/val | Currently available throttle
throttle-mon_client_bytes/max | Max value for throttle
throttle-mon_client_bytes/get | Gets
throttle-mon_client_bytes/get_sum | Got data
throttle-mon_client_bytes/get_or_fail_fail | Get blocked during get_or_fail
throttle-mon_client_bytes/get_or_fail_success | Successful get during get_or_fail
throttle-mon_client_bytes/take | Takes
throttle-mon_client_bytes/take_sum | Taken data
throttle-mon_client_bytes/put | Puts
throttle-mon_client_bytes/put_sum | Put data
throttle-mon_client_bytes/wait/avgcount | Waiting latency
throttle-mon_client_bytes/wait/sum | Waiting latency
 | 
throttle-mon_daemon_bytes/val | Currently available throttle
throttle-mon_daemon_bytes/max | Max value for throttle
throttle-mon_daemon_bytes/get | Gets
throttle-mon_daemon_bytes/get_sum | Got data
throttle-mon_daemon_bytes/get_or_fail_fail | Get blocked during get_or_fail
throttle-mon_daemon_bytes/get_or_fail_success | Successful get during get_or_fail
throttle-mon_daemon_bytes/take | Takes
throttle-mon_daemon_bytes/take_sum | Taken data
throttle-mon_daemon_bytes/put | Puts
throttle-mon_daemon_bytes/put_sum | Put data
throttle-mon_daemon_bytes/wait/avgcount | Waiting latency
throttle-mon_daemon_bytes/wait/sum | Waiting latency
 | 
throttle-msgr_dispatch_throttler-mon/val | Currently available throttle
throttle-msgr_dispatch_throttler-mon/max | Max value for throttle
throttle-msgr_dispatch_throttler-mon/get | Gets
throttle-msgr_dispatch_throttler-mon/get_sum | Got data
throttle-msgr_dispatch_throttler-mon/get_or_fail_fail | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-mon/get_or_fail_success | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-mon/take | Takes
throttle-msgr_dispatch_throttler-mon/take_sum | Taken data
throttle-msgr_dispatch_throttler-mon/put | Puts
throttle-msgr_dispatch_throttler-mon/put_sum | Put data
throttle-msgr_dispatch_throttler-mon/wait/avgcount | Waiting latency
throttle-msgr_dispatch_throttler-mon/wait/sum | Waiting latency
