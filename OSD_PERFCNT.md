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
### OSD daemon
**Metric list was generated dynamically by plugin and can be different on another setup.**  
Prefix: `/intel/storage/ceph/osd/[osd_id]`

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
WBThrottle/bytes_dirtied | uint64 | Dirty data
WBThrottle/bytes_wb | uint64 | Written data
WBThrottle/inodes_dirtied | uint64 | Entries waiting for write
WBThrottle/inodes_wb | uint64 | Written entries
WBThrottle/ios_dirtied | uint64 | Dirty operations
WBThrottle/ios_wb | uint64 | Written operations
filestore/apply_latency | float | Apply latency
filestore/bytes | uint64 | Data written to store
filestore/commitcycle | uint64 | Commit cycles
filestore/commitcycle_interval | float | Average interval between commits
filestore/commitcycle_latency | float | Average latency of commit
filestore/committing | uint64 | Is currently committing
filestore/journal_bytes | uint64 | Active journal operation size to be applied
filestore/journal_full | uint64 | Journal writes while full
filestore/journal_latency | float | Average journal queue completing latency
filestore/journal_ops | uint64 | Active journal entries to be applied
filestore/journal_queue_bytes | uint64 | Size of journal queue
filestore/journal_queue_ops | uint64 | Operations in journal queue
filestore/journal_wr | uint64 | Journal write IOs
filestore/journal_wr_bytes | uint64 | Journal data written
filestore/op_queue_bytes | uint64 | Size of writing to FS queue
filestore/op_queue_max_bytes | uint64 | Max data in writing to FS queue
filestore/op_queue_max_ops | uint64 | Max operations in writing to FS queue
filestore/op_queue_ops | uint64 | Operations in writing to FS queue
filestore/ops | uint64 | Operations written to store
filestore/queue_transaction_latency_avg | float | Store operation queue latency
finisher-JournalObjectStore/complete_latency | float | 
finisher-JournalObjectStore/queue_len | uint64 | 
finisher-filestore-apply-0/complete_latency | float | 
finisher-filestore-apply-0/queue_len | uint64 | 
finisher-filestore-ondisk-0/complete_latency | float | 
finisher-filestore-ondisk-0/queue_len | uint64 | 
leveldb/leveldb_compact | uint64 | Compactions
leveldb/leveldb_compact_queue_len | uint64 | Length of compaction queue
leveldb/leveldb_compact_queue_merge | uint64 | Mergings of ranges in compaction queue
leveldb/leveldb_compact_range | uint64 | Compactions by range
leveldb/leveldb_get | uint64 | Gets
leveldb/leveldb_get_latency | float | Get Latency
leveldb/leveldb_submit_latency | float | Submit Latency
leveldb/leveldb_submit_sync_latency | float | Submit Sync Latency
leveldb/leveldb_transaction | uint64 | Transactions
mutex-FileJournal::completions_lock/wait | float | Average time of mutex in locked state
mutex-FileJournal::finisher_lock/wait | float | Average time of mutex in locked state
mutex-FileJournal::write_lock/wait | float | Average time of mutex in locked state
mutex-FileJournal::writeq_lock/wait | float | Average time of mutex in locked state
mutex-JOS::ApplyManager::apply_lock/wait | float | Average time of mutex in locked state
mutex-JOS::ApplyManager::com_lock/wait | float | Average time of mutex in locked state
mutex-JOS::SubmitManager::lock/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.0/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.1/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.2/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.3/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.4/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.0/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.1/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.2/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.3/wait | float | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.4/wait | float | Average time of mutex in locked state
mutex-WBThrottle::lock/wait | float | Average time of mutex in locked state
objecter/command_active | uint64 | Active commands
objecter/command_resend | uint64 | Resent commands
objecter/command_send | uint64 | Sent commands
objecter/linger_active | uint64 | Active lingering operations
objecter/linger_ping | uint64 | Sent pings to lingering operations
objecter/linger_resend | uint64 | Resent lingering operations
objecter/linger_send | uint64 | Sent lingering operations
objecter/map_epoch | uint64 | OSD map epoch
objecter/map_full | uint64 | Full OSD maps received
objecter/map_inc | uint64 | Incremental OSD maps received
objecter/omap_del | uint64 | OSD OMAP delete operations
objecter/omap_rd | uint64 | OSD OMAP read operations
objecter/omap_wr | uint64 | OSD OMAP write operations
objecter/op | uint64 | Operations
objecter/op_ack | uint64 | Commit callbacks
objecter/op_active | uint64 | Operations active
objecter/op_commit | uint64 | Operation commits
objecter/op_laggy | uint64 | Laggy operations
objecter/op_pg | uint64 | PG operation
objecter/op_r | uint64 | Read operations
objecter/op_resend | uint64 | Resent operations
objecter/op_rmw | uint64 | Read-modify-write operations
objecter/op_send | uint64 | Sent operations
objecter/op_send_bytes | uint64 | Sent data
objecter/op_w | uint64 | Write operations
objecter/osd_laggy | uint64 | Laggy OSD sessions
objecter/osd_session_close | uint64 | Sessions closed
objecter/osd_session_open | uint64 | Sessions opened
objecter/osd_sessions | uint64 | Open sessions
objecter/osdop_append | uint64 | Append operation
objecter/osdop_call | uint64 | Call (execute) operations
objecter/osdop_clonerange | uint64 | Clone range operations
objecter/osdop_cmpxattr | uint64 | Xattr comparison operations
objecter/osdop_create | uint64 | Create object operations
objecter/osdop_delete | uint64 | Delete object operations
objecter/osdop_getxattr | uint64 | Get xattr operations
objecter/osdop_mapext | uint64 | Map extent operations
objecter/osdop_notify | uint64 | Notify about object operations
objecter/osdop_other | uint64 | Other operations
objecter/osdop_pgls | uint64 | 
objecter/osdop_pgls_filter | uint64 | 
objecter/osdop_read | uint64 | Read operations
objecter/osdop_resetxattrs | uint64 | Reset xattr operations
objecter/osdop_rmxattr | uint64 | Remove xattr operations
objecter/osdop_setxattr | uint64 | Set xattr operations
objecter/osdop_sparse_read | uint64 | Sparse read operations
objecter/osdop_src_cmpxattr | uint64 | Extended attribute comparison in multi operations
objecter/osdop_stat | uint64 | Stat operations
objecter/osdop_tmap_get | uint64 | TMAP get operations
objecter/osdop_tmap_put | uint64 | TMAP put operations
objecter/osdop_tmap_up | uint64 | TMAP update operations
objecter/osdop_truncate | uint64 | Truncate object operations
objecter/osdop_watch | uint64 | Watch by object operations
objecter/osdop_write | uint64 | Write operations
objecter/osdop_writefull | uint64 | Write full object operations
objecter/osdop_writesame | uint64 | Write same operations
objecter/osdop_zero | uint64 | Set object to zero operations
objecter/poolop_active | uint64 | Active pool operations
objecter/poolop_resend | uint64 | Resent pool operations
objecter/poolop_send | uint64 | Sent pool operations
objecter/poolstat_active | uint64 | Active get pool stat operations
objecter/poolstat_resend | uint64 | Resent pool stats
objecter/poolstat_send | uint64 | Pool stat operations sent
objecter/statfs_active | uint64 | Statfs operations
objecter/statfs_resend | uint64 | Resent FS stats
objecter/statfs_send | uint64 | Sent FS stats
osd/agent_evict | uint64 | Tiering agent evictions
osd/agent_flush | uint64 | Tiering agent flushes
osd/agent_skip | uint64 | Objects skipped by agent
osd/agent_wake | uint64 | Tiering agent wake up
osd/buffer_bytes | uint64 | Total allocated buffer size
osd/cached_crc | uint64 | Total number getting crc from crc_cache
osd/cached_crc_adjusted | uint64 | Total number getting crc from crc_cache with adjusting
osd/copyfrom | uint64 | Rados "copy-from" operations
osd/heartbeat_to_peers | uint64 | Heartbeat (ping) peers we send to
osd/history_alloc_Mbytes | uint64 | 
osd/history_alloc_num | uint64 | 
osd/loadavg | uint64 | CPU load
osd/map_message_epoch_dups | uint64 | OSD map duplicates
osd/map_message_epochs | uint64 | OSD map epochs
osd/map_messages | uint64 | OSD map messages
osd/messages_delayed_for_map | uint64 | Operations waiting for OSD map
osd/numpg | uint64 | Placement groups
osd/numpg_primary | uint64 | Placement groups for which this osd is primary
osd/numpg_replica | uint64 | Placement groups for which this osd is replica
osd/numpg_stray | uint64 | Placement groups ready to be deleted from this osd
osd/object_ctx_cache_hit | uint64 | Object context cache hits
osd/object_ctx_cache_total | uint64 | Object context cache lookups
osd/op | uint64 | Client operations
osd/op_cache_hit | uint64 | 
osd/op_in_bytes | uint64 | Client operations total write size
osd/op_latency | float | Latency of client operations (including queue time)
osd/op_out_bytes | uint64 | Client operations total read size
osd/op_prepare_latency | float | Latency of client operations (excluding queue time and wait for finished)
osd/op_process_latency | float | Latency of client operations (excluding queue time)
osd/op_r | uint64 | Client read operations
osd/op_r_latency | float | Latency of read operation (including queue time)
osd/op_r_out_bytes | uint64 | Client data read
osd/op_r_prepare_latency | float | Latency of read operations (excluding queue time and wait for finished)
osd/op_r_process_latency | float | Latency of read operation (excluding queue time)
osd/op_rw | uint64 | Client read-modify-write operations
osd/op_rw_in_bytes | uint64 | Client read-modify-write operations write in
osd/op_rw_latency | float | Latency of read-modify-write operation (including queue time)
osd/op_rw_out_bytes | uint64 | Client read-modify-write operations read out 
osd/op_rw_prepare_latency | float | Latency of read-modify-write operations (excluding queue time and wait for finished)
osd/op_rw_process_latency | float | Latency of read-modify-write operation (excluding queue time)
osd/op_rw_rlat | float | Client read-modify-write operation readable/applied latency
osd/op_w | uint64 | Client write operations
osd/op_w_in_bytes | uint64 | Client data written
osd/op_w_latency | float | Latency of write operation (including queue time)
osd/op_w_prepare_latency | float | Latency of write operations (excluding queue time and wait for finished)
osd/op_w_process_latency | float | Latency of write operation (excluding queue time)
osd/op_w_rlat | float | Client write operation readable/applied latency
osd/op_wip | uint64 | Replication operations currently being processed (primary)
osd/osd_map_cache_hit | uint64 | osdmap cache hit
osd/osd_map_cache_miss | uint64 | osdmap cache miss
osd/osd_map_cache_miss_low | uint64 | osdmap cache miss below cache lower bound
osd/osd_map_cache_miss_low_avg | uint64 | osdmap cache miss, avg distance below cache lower bound
osd/osd_pg_biginfo | uint64 | PG updated its biginfo attr
osd/osd_pg_fastinfo | uint64 | PG updated its info using fastinfo attr
osd/osd_pg_info | uint64 | PG updated its info (using any method)
osd/osd_tier_flush_lat | float | Object flush latency
osd/osd_tier_promote_lat | float | Object promote latency
osd/osd_tier_r_lat | float | Object proxy read latency
osd/pull | uint64 | Pull requests sent
osd/push | uint64 | Push messages sent
osd/push_out_bytes | uint64 | Pushed size
osd/recovery_ops | uint64 | Started recovery operations
osd/stat_bytes | uint64 | OSD size
osd/stat_bytes_avail | uint64 | Available space
osd/stat_bytes_used | uint64 | Used space
osd/subop | uint64 | Suboperations
osd/subop_in_bytes | uint64 | Suboperations total size
osd/subop_latency | float | Suboperations latency
osd/subop_pull | uint64 | Suboperations pull requests
osd/subop_pull_latency | float | Suboperations pull latency
osd/subop_push | uint64 | Suboperations push messages
osd/subop_push_in_bytes | uint64 | Suboperations pushed size
osd/subop_push_latency | float | Suboperations push latency
osd/subop_w | uint64 | Replicated writes
osd/subop_w_in_bytes | uint64 | Replicated written data size
osd/subop_w_latency | float | Replicated writes latency
osd/tier_clean | uint64 | Dirty tier flag cleaned
osd/tier_delay | uint64 | Tier delays (agent waiting)
osd/tier_dirty | uint64 | Dirty tier flag set
osd/tier_evict | uint64 | Tier evictions
osd/tier_flush | uint64 | Tier flushes
osd/tier_flush_fail | uint64 | Failed tier flushes
osd/tier_promote | uint64 | Tier promotions
osd/tier_proxy_read | uint64 | Tier proxy reads
osd/tier_proxy_write | uint64 | Tier proxy writes
osd/tier_try_flush | uint64 | Tier flush attempts
osd/tier_try_flush_fail | uint64 | Failed tier flush attempts
osd/tier_whiteout | uint64 | Tier whiteouts
recoverystate_perf/activating_latency | float | Activating recovery state latency
recoverystate_perf/active_latency | float | Active recovery state latency
recoverystate_perf/backfilling_latency | float | Backfilling recovery state latency
recoverystate_perf/clean_latency | float | Clean recovery state latency
recoverystate_perf/getinfo_latency | float | Getinfo recovery state latency
recoverystate_perf/getlog_latency | float | Getlog recovery state latency
recoverystate_perf/getmissing_latency | float | Getmissing recovery state latency
recoverystate_perf/incomplete_latency | float | Incomplete recovery state latency
recoverystate_perf/initial_latency | float | Initial recovery state latency
recoverystate_perf/notbackfilling_latency | float | Notbackfilling recovery state latency
recoverystate_perf/peering_latency | float | Peering recovery state latency
recoverystate_perf/primary_latency | float | Primary recovery state latency
recoverystate_perf/recovered_latency | float | Recovered recovery state latency
recoverystate_perf/recovering_latency | float | Recovering recovery state latency
recoverystate_perf/replicaactive_latency | float | Replicaactive recovery state latency
recoverystate_perf/repnotrecovering_latency | float | Repnotrecovering recovery state latency
recoverystate_perf/reprecovering_latency | float | RepRecovering recovery state latency
recoverystate_perf/repwaitbackfillreserved_latency | float | Rep wait backfill reserved recovery state latency
recoverystate_perf/repwaitrecoveryreserved_latency | float | Rep wait recovery reserved recovery state latency
recoverystate_perf/reset_latency | float | Reset recovery state latency
recoverystate_perf/start_latency | float | Start recovery state latency
recoverystate_perf/started_latency | float | Started recovery state latency
recoverystate_perf/stray_latency | float | Stray recovery state latency
recoverystate_perf/waitactingchange_latency | float | Waitactingchange recovery state latency
recoverystate_perf/waitlocalbackfillreserved_latency | float | Wait local backfill reserved recovery state latency
recoverystate_perf/waitlocalrecoveryreserved_latency | float | Wait local recovery reserved recovery state latency
recoverystate_perf/waitremotebackfillreserved_latency | float | Wait remote backfill reserved recovery state latency
recoverystate_perf/waitremoterecoveryreserved_latency | float | Wait remote recovery reserved recovery state latency
recoverystate_perf/waitupthru_latency | float | Waitupthru recovery state latency
throttle-msgr_dispatch_throttler-client/get | uint64 | Gets
throttle-msgr_dispatch_throttler-client/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-client/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-client/get_started | uint64 | Number of get calls, increased before wait
throttle-msgr_dispatch_throttler-client/get_sum | uint64 | Got data
throttle-msgr_dispatch_throttler-client/max | uint64 | Max value for throttle
throttle-msgr_dispatch_throttler-client/put | uint64 | Puts
throttle-msgr_dispatch_throttler-client/put_sum | uint64 | Put data
throttle-msgr_dispatch_throttler-client/take | uint64 | Takes
throttle-msgr_dispatch_throttler-client/take_sum | uint64 | Taken data
throttle-msgr_dispatch_throttler-client/val | uint64 | Currently available throttle
throttle-msgr_dispatch_throttler-client/wait | float | Waiting latency
throttle-msgr_dispatch_throttler-cluster/get | uint64 | Gets
throttle-msgr_dispatch_throttler-cluster/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-cluster/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-cluster/get_started | uint64 | Number of get calls, increased before wait
throttle-msgr_dispatch_throttler-cluster/get_sum | uint64 | Got data
throttle-msgr_dispatch_throttler-cluster/max | uint64 | Max value for throttle
throttle-msgr_dispatch_throttler-cluster/put | uint64 | Puts
throttle-msgr_dispatch_throttler-cluster/put_sum | uint64 | Put data
throttle-msgr_dispatch_throttler-cluster/take | uint64 | Takes
throttle-msgr_dispatch_throttler-cluster/take_sum | uint64 | Taken data
throttle-msgr_dispatch_throttler-cluster/val | uint64 | Currently available throttle
throttle-msgr_dispatch_throttler-cluster/wait | float | Waiting latency
throttle-msgr_dispatch_throttler-hb_back_server/get | uint64 | Gets
throttle-msgr_dispatch_throttler-hb_back_server/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-hb_back_server/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-hb_back_server/get_started | uint64 | Number of get calls, increased before wait
throttle-msgr_dispatch_throttler-hb_back_server/get_sum | uint64 | Got data
throttle-msgr_dispatch_throttler-hb_back_server/max | uint64 | Max value for throttle
throttle-msgr_dispatch_throttler-hb_back_server/put | uint64 | Puts
throttle-msgr_dispatch_throttler-hb_back_server/put_sum | uint64 | Put data
throttle-msgr_dispatch_throttler-hb_back_server/take | uint64 | Takes
throttle-msgr_dispatch_throttler-hb_back_server/take_sum | uint64 | Taken data
throttle-msgr_dispatch_throttler-hb_back_server/val | uint64 | Currently available throttle
throttle-msgr_dispatch_throttler-hb_back_server/wait | float | Waiting latency
throttle-msgr_dispatch_throttler-hb_front_server/get | uint64 | Gets
throttle-msgr_dispatch_throttler-hb_front_server/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-hb_front_server/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-hb_front_server/get_started | uint64 | Number of get calls, increased before wait
throttle-msgr_dispatch_throttler-hb_front_server/get_sum | uint64 | Got data
throttle-msgr_dispatch_throttler-hb_front_server/max | uint64 | Max value for throttle
throttle-msgr_dispatch_throttler-hb_front_server/put | uint64 | Puts
throttle-msgr_dispatch_throttler-hb_front_server/put_sum | uint64 | Put data
throttle-msgr_dispatch_throttler-hb_front_server/take | uint64 | Takes
throttle-msgr_dispatch_throttler-hb_front_server/take_sum | uint64 | Taken data
throttle-msgr_dispatch_throttler-hb_front_server/val | uint64 | Currently available throttle
throttle-msgr_dispatch_throttler-hb_front_server/wait | float | Waiting latency
throttle-msgr_dispatch_throttler-hbclient/get | uint64 | Gets
throttle-msgr_dispatch_throttler-hbclient/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-hbclient/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-hbclient/get_started | uint64 | Number of get calls, increased before wait
throttle-msgr_dispatch_throttler-hbclient/get_sum | uint64 | Got data
throttle-msgr_dispatch_throttler-hbclient/max | uint64 | Max value for throttle
throttle-msgr_dispatch_throttler-hbclient/put | uint64 | Puts
throttle-msgr_dispatch_throttler-hbclient/put_sum | uint64 | Put data
throttle-msgr_dispatch_throttler-hbclient/take | uint64 | Takes
throttle-msgr_dispatch_throttler-hbclient/take_sum | uint64 | Taken data
throttle-msgr_dispatch_throttler-hbclient/val | uint64 | Currently available throttle
throttle-msgr_dispatch_throttler-hbclient/wait | float | Waiting latency
throttle-msgr_dispatch_throttler-ms_objecter/get | uint64 | Gets
throttle-msgr_dispatch_throttler-ms_objecter/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-ms_objecter/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-ms_objecter/get_started | uint64 | Number of get calls, increased before wait
throttle-msgr_dispatch_throttler-ms_objecter/get_sum | uint64 | Got data
throttle-msgr_dispatch_throttler-ms_objecter/max | uint64 | Max value for throttle
throttle-msgr_dispatch_throttler-ms_objecter/put | uint64 | Puts
throttle-msgr_dispatch_throttler-ms_objecter/put_sum | uint64 | Put data
throttle-msgr_dispatch_throttler-ms_objecter/take | uint64 | Takes
throttle-msgr_dispatch_throttler-ms_objecter/take_sum | uint64 | Taken data
throttle-msgr_dispatch_throttler-ms_objecter/val | uint64 | Currently available throttle
throttle-msgr_dispatch_throttler-ms_objecter/wait | float | Waiting latency
throttle-objecter_bytes/get | uint64 | Gets
throttle-objecter_bytes/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-objecter_bytes/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-objecter_bytes/get_started | uint64 | Number of get calls, increased before wait
throttle-objecter_bytes/get_sum | uint64 | Got data
throttle-objecter_bytes/max | uint64 | Max value for throttle
throttle-objecter_bytes/put | uint64 | Puts
throttle-objecter_bytes/put_sum | uint64 | Put data
throttle-objecter_bytes/take | uint64 | Takes
throttle-objecter_bytes/take_sum | uint64 | Taken data
throttle-objecter_bytes/val | uint64 | Currently available throttle
throttle-objecter_bytes/wait | float | Waiting latency
throttle-objecter_ops/get | uint64 | Gets
throttle-objecter_ops/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-objecter_ops/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-objecter_ops/get_started | uint64 | Number of get calls, increased before wait
throttle-objecter_ops/get_sum | uint64 | Got data
throttle-objecter_ops/max | uint64 | Max value for throttle
throttle-objecter_ops/put | uint64 | Puts
throttle-objecter_ops/put_sum | uint64 | Put data
throttle-objecter_ops/take | uint64 | Takes
throttle-objecter_ops/take_sum | uint64 | Taken data
throttle-objecter_ops/val | uint64 | Currently available throttle
throttle-objecter_ops/wait | float | Waiting latency
throttle-osd_client_bytes/get | uint64 | Gets
throttle-osd_client_bytes/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-osd_client_bytes/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-osd_client_bytes/get_started | uint64 | Number of get calls, increased before wait
throttle-osd_client_bytes/get_sum | uint64 | Got data
throttle-osd_client_bytes/max | uint64 | Max value for throttle
throttle-osd_client_bytes/put | uint64 | Puts
throttle-osd_client_bytes/put_sum | uint64 | Put data
throttle-osd_client_bytes/take | uint64 | Takes
throttle-osd_client_bytes/take_sum | uint64 | Taken data
throttle-osd_client_bytes/val | uint64 | Currently available throttle
throttle-osd_client_bytes/wait | float | Waiting latency
throttle-osd_client_messages/get | uint64 | Gets
throttle-osd_client_messages/get_or_fail_fail | uint64 | Get blocked during get_or_fail
throttle-osd_client_messages/get_or_fail_success | uint64 | Successful get during get_or_fail
throttle-osd_client_messages/get_started | uint64 | Number of get calls, increased before wait
throttle-osd_client_messages/get_sum | uint64 | Got data
throttle-osd_client_messages/max | uint64 | Max value for throttle
throttle-osd_client_messages/put | uint64 | Puts
throttle-osd_client_messages/put_sum | uint64 | Put data
throttle-osd_client_messages/take | uint64 | Takes
throttle-osd_client_messages/take_sum | uint64 | Taken data
throttle-osd_client_messages/val | uint64 | Currently available throttle
throttle-osd_client_messages/wait | float | Waiting latency
