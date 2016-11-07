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

## Snap Ceph Perf Counters Collector Plugin


# Ceph OSD Perf Counters

Prefix: /intel/storage/ceph/{osd_daemon_name}/{osd_daemon_id}

Metrics | Description
------------ | -------------
WBThrottle/bytes_dirtied | Dirty data
WBThrottle/bytes_wb | Written data
WBThrottle/ios_dirtied | Dirty operations
WBThrottle/ios_wb | Written operations
WBThrottle/inodes_dirtied | Entries waiting for write
WBThrottle/inodes_wb | Written entries
filestore/journal_queue_ops | Operations in journal queue
filestore/journal_queue_bytes | Size of journal queue
filestore/journal_ops | Total journal entries written
filestore/journal_bytes | Total operations size in journal
filestore/journal_latency | Average journal queue completing latency
filestore/journal_wr | Journal write IOs
filestore/journal_wr_bytes | Journal data written
filestore/journal_full | Journal writes while full
filestore/committing | Is currently committing
filestore/commitcycle | Commit cycles
filestore/commitcycle_interval | Average interval between commits
filestore/commitcycle_latency | Average latency of commit
filestore/op_queue_max_ops | Max operations in writing to FS queue
filestore/op_queue_ops | Operations in writing to FS queue
filestore/ops | Operations written to store
filestore/op_queue_max_bytes | Max data in writing to FS queue
filestore/op_queue_bytes | Size of writing to FS queue
filestore/bytes | Data written to store
filestore/apply_latency | Apply latency
filestore/queue_transaction_latency_avg | Store operation queue latency
finisher-JournalObjectStore/queue_len | [No description]
finisher-JournalObjectStore/complete_latency | [No description]
finisher-filestore-apply-0/queue_len | [No description]
finisher-filestore-apply-0/complete_latency | [No description]
finisher-filestore-ondisk-0/queue_len | [No description]
finisher-filestore-ondisk-0/complete_latency | [No description]
leveldb/leveldb_get | Gets
leveldb/leveldb_transaction | Transactions
leveldb/leveldb_get_latency | Get Latency
leveldb/leveldb_submit_latency | Submit Latency
leveldb/leveldb_submit_sync_latency | Submit Sync Latency
leveldb/leveldb_compact | Compactions
leveldb/leveldb_compact_range | Compactions by range
leveldb/leveldb_compact_queue_merge | Mergings of ranges in compaction queue
leveldb/leveldb_compact_queue_len | Length of compaction queue
mutex-FileJournal::completions_lock/wait | Average time of mutex in locked state
mutex-FileJournal::finisher_lock/wait | Average time of mutex in locked state
mutex-FileJournal::write_lock/wait | Average time of mutex in locked state
mutex-FileJournal::writeq_lock/wait | Average time of mutex in locked state
mutex-JOS::ApplyManager::apply_lock/wait | Average time of mutex in locked state
mutex-JOS::ApplyManager::com_lock/wait | Average time of mutex in locked state
mutex-JOS::SubmitManager::lock/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.0/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.1/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.2/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.3/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:.4/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.0/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.1/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.2/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.3/wait | Average time of mutex in locked state
mutex-OSD:ShardedOpWQ:order:.4/wait | Average time of mutex in locked state
mutex-WBThrottle::lock/wait | Average time of mutex in locked state
objecter/op_active | Operations active
objecter/op_laggy | Laggy operations
objecter/op_send | Sent operations
objecter/op_send_bytes | Sent data
objecter/op_resend | Resent operations
objecter/op_ack | Commit callbacks
objecter/op_commit | Operation commits
objecter/op | Operations
objecter/op_r | Read operations
objecter/op_w | Write operations
objecter/op_rmw | Read-modify-write operations
objecter/op_pg | PG operation
objecter/osdop_stat | Stat operations
objecter/osdop_create | Create object operations
objecter/osdop_read | Read operations
objecter/osdop_write | Write operations
objecter/osdop_writefull | Write full object operations
objecter/osdop_append | Append operation
objecter/osdop_zero | Set object to zero operations
objecter/osdop_truncate | Truncate object operations
objecter/osdop_delete | Delete object operations
objecter/osdop_mapext | Map extent operations
objecter/osdop_sparse_read | Sparse read operations
objecter/osdop_clonerange | Clone range operations
objecter/osdop_getxattr | Get xattr operations
objecter/osdop_setxattr | Set xattr operations
objecter/osdop_cmpxattr | Xattr comparison operations
objecter/osdop_rmxattr | Remove xattr operations
objecter/osdop_resetxattrs | Reset xattr operations
objecter/osdop_tmap_up | TMAP update operations
objecter/osdop_tmap_put | TMAP put operations
objecter/osdop_tmap_get | TMAP get operations
objecter/osdop_call | Call (execute) operations
objecter/osdop_watch | Watch by object operations
objecter/osdop_notify | Notify about object operations
objecter/osdop_src_cmpxattr | Extended attribute comparison in multi operations
objecter/osdop_pgls | [No description]
objecter/osdop_pgls_filter | [No description]
objecter/osdop_other | Other operations
objecter/linger_active | Active lingering operations
objecter/linger_send | Sent lingering operations
objecter/linger_resend | Resent lingering operations
objecter/linger_ping | Sent pings to lingering operations
objecter/poolop_active | Active pool operations
objecter/poolop_send | Sent pool operations
objecter/poolop_resend | Resent pool operations
objecter/poolstat_active | Active get pool stat operations
objecter/poolstat_send | Pool stat operations sent
objecter/poolstat_resend | Resent pool stats
objecter/statfs_active | Statfs operations
objecter/statfs_send | Sent FS stats
objecter/statfs_resend | Resent FS stats
objecter/command_active | Active commands
objecter/command_send | Sent commands
objecter/command_resend | Resent commands
objecter/map_epoch | OSD map epoch
objecter/map_full | Full OSD maps received
objecter/map_inc | Incremental OSD maps received
objecter/osd_sessions | Open sessions
objecter/osd_session_open | Sessions opened
objecter/osd_session_close | Sessions closed
objecter/osd_laggy | Laggy OSD sessions
objecter/omap_wr | OSD OMAP write operations
objecter/omap_rd | OSD OMAP read operations
objecter/omap_del | OSD OMAP delete operations
osd/op_wip | Replication operations currently being processed (primary)
osd/op | Client operations
osd/op_in_bytes | Client operations total write size
osd/op_out_bytes | Client operations total read size
osd/op_latency | Latency of client operations (including queue time)
osd/op_process_latency | Latency of client operations (excluding queue time)
osd/op_prepare_latency | Latency of client operations (excluding queue time and wait for finished)
osd/op_r | Client read operations
osd/op_r_out_bytes | Client data read
osd/op_r_latency | Latency of read operation (including queue time)
osd/op_r_process_latency | Latency of read operation (excluding queue time)
osd/op_r_prepare_latency | Latency of read operations (excluding queue time and wait for finished)
osd/op_w | Client write operations
osd/op_w_in_bytes | Client data written
osd/op_w_rlat | Client write operation readable/applied latency
osd/op_w_latency | Latency of write operation (including queue time)
osd/op_w_process_latency | Latency of write operation (excluding queue time)
osd/op_w_prepare_latency | Latency of write operations (excluding queue time and wait for finished)
osd/op_rw | Client read-modify-write operations
osd/op_rw_in_bytes | Client read-modify-write operations write in
osd/op_rw_out_bytes | Client read-modify-write operations read out 
osd/op_rw_rlat | Client read-modify-write operation readable/applied latency
osd/op_rw_latency | Latency of read-modify-write operation (including queue time)
osd/op_rw_process_latency | Latency of read-modify-write operation (excluding queue time)
osd/op_rw_prepare_latency | Latency of read-modify-write operations (excluding queue time and wait for finished)
osd/subop | Suboperations
osd/subop_in_bytes | Suboperations total size
osd/subop_latency | Suboperations latency
osd/subop_w | Replicated writes
osd/subop_w_in_bytes | Replicated written data size
osd/subop_w_latency | Replicated writes latency
osd/subop_pull | Suboperations pull requests
osd/subop_pull_latency | Suboperations pull latency
osd/subop_push | Suboperations push messages
osd/subop_push_in_bytes | Suboperations pushed size
osd/subop_push_latency | Suboperations push latency
osd/pull | Pull requests sent
osd/push | Push messages sent
osd/push_out_bytes | Pushed size
osd/push_in | Inbound push messages
osd/push_in_bytes | Inbound pushed size
osd/recovery_ops | Started recovery operations
osd/loadavg | CPU load
osd/buffer_bytes | Total allocated buffer size
osd/history_alloc_Mbytes | [No description]
osd/history_alloc_num | [No description]
osd/cached_crc | Total number getting crc from crc_cache
osd/cached_crc_adjusted | Total number getting crc from crc_cache with adjusting
osd/numpg | Placement groups
osd/numpg_primary | Placement groups for which this osd is primary
osd/numpg_replica | Placement groups for which this osd is replica
osd/numpg_stray | Placement groups ready to be deleted from this osd
osd/heartbeat_to_peers | Heartbeat (ping) peers we send to
osd/map_messages | OSD map messages
osd/map_message_epochs | OSD map epochs
osd/map_message_epoch_dups | OSD map duplicates
osd/messages_delayed_for_map | Operations waiting for OSD map
osd/stat_bytes | OSD size
osd/stat_bytes_used | Used space
osd/stat_bytes_avail | Available space
osd/copyfrom | Rados \copy-from\ operations
osd/tier_promote | Tier promotions
osd/tier_flush | Tier flushes
osd/tier_flush_fail | Failed tier flushes
osd/tier_try_flush | Tier flush attempts
osd/tier_try_flush_fail | Failed tier flush attempts
osd/tier_evict | Tier evictions
osd/tier_whiteout | Tier whiteouts
osd/tier_dirty | Dirty tier flag set
osd/tier_clean | Dirty tier flag cleaned
osd/tier_delay | Tier delays (agent waiting)
osd/tier_proxy_read | Tier proxy reads
osd/tier_proxy_write | Tier proxy writes
osd/agent_wake | Tiering agent wake up
osd/agent_skip | Objects skipped by agent
osd/agent_flush | Tiering agent flushes
osd/agent_evict | Tiering agent evictions
osd/object_ctx_cache_hit | Object context cache hits
osd/object_ctx_cache_total | Object context cache lookups
osd/op_cache_hit | [No description]
osd/osd_tier_flush_lat | Object flush latency
osd/osd_tier_promote_lat | Object promote latency
osd/osd_tier_r_lat | Object proxy read latency
recoverystate_perf/initial_latency | Initial recovery state latency
recoverystate_perf/started_latency | Started recovery state latency
recoverystate_perf/reset_latency | Reset recovery state latency
recoverystate_perf/start_latency | Start recovery state latency
recoverystate_perf/primary_latency | Primary recovery state latency
recoverystate_perf/peering_latency | Peering recovery state latency
recoverystate_perf/backfilling_latency | Backfilling recovery state latency
recoverystate_perf/waitremotebackfillreserved_latency | Wait remote backfill reserved recovery state latency
recoverystate_perf/waitlocalbackfillreserved_latency | Wait local backfill reserved recovery state latency
recoverystate_perf/notbackfilling_latency | Notbackfilling recovery state latency
recoverystate_perf/repnotrecovering_latency | Repnotrecovering recovery state latency
recoverystate_perf/repwaitrecoveryreserved_latency | Rep wait recovery reserved recovery state latency
recoverystate_perf/repwaitbackfillreserved_latency | Rep wait backfill reserved recovery state latency
recoverystate_perf/reprecovering_latency | RepRecovering recovery state latency
recoverystate_perf/activating_latency | Activating recovery state latency
recoverystate_perf/waitlocalrecoveryreserved_latency | Wait local recovery reserved recovery state latency
recoverystate_perf/waitremoterecoveryreserved_latency | Wait remote recovery reserved recovery state latency
recoverystate_perf/recovering_latency | Recovering recovery state latency
recoverystate_perf/recovered_latency | Recovered recovery state latency
recoverystate_perf/clean_latency | Clean recovery state latency
recoverystate_perf/active_latency | Active recovery state latency
recoverystate_perf/replicaactive_latency | Replicaactive recovery state latency
recoverystate_perf/stray_latency | Stray recovery state latency
recoverystate_perf/getinfo_latency | Getinfo recovery state latency
recoverystate_perf/getlog_latency | Getlog recovery state latency
recoverystate_perf/waitactingchange_latency | Waitactingchange recovery state latency
recoverystate_perf/incomplete_latency | Incomplete recovery state latency
recoverystate_perf/getmissing_latency | Getmissing recovery state latency
recoverystate_perf/waitupthru_latency | Waitupthru recovery state latency
throttle-msgr_dispatch_throttler-client/val | Currently available throttle
throttle-msgr_dispatch_throttler-client/max | Max value for throttle
throttle-msgr_dispatch_throttler-client/get | Gets
throttle-msgr_dispatch_throttler-client/get_sum | Got data
throttle-msgr_dispatch_throttler-client/get_or_fail_fail | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-client/get_or_fail_success | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-client/take | Takes
throttle-msgr_dispatch_throttler-client/take_sum | Taken data
throttle-msgr_dispatch_throttler-client/put | Puts
throttle-msgr_dispatch_throttler-client/put_sum | Put data
throttle-msgr_dispatch_throttler-client/wait | Waiting latency
throttle-msgr_dispatch_throttler-cluster/val | Currently available throttle
throttle-msgr_dispatch_throttler-cluster/max | Max value for throttle
throttle-msgr_dispatch_throttler-cluster/get | Gets
throttle-msgr_dispatch_throttler-cluster/get_sum | Got data
throttle-msgr_dispatch_throttler-cluster/get_or_fail_fail | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-cluster/get_or_fail_success | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-cluster/take | Takes
throttle-msgr_dispatch_throttler-cluster/take_sum | Taken data
throttle-msgr_dispatch_throttler-cluster/put | Puts
throttle-msgr_dispatch_throttler-cluster/put_sum | Put data
throttle-msgr_dispatch_throttler-cluster/wait | Waiting latency
throttle-msgr_dispatch_throttler-hb_back_server/val | Currently available throttle
throttle-msgr_dispatch_throttler-hb_back_server/max | Max value for throttle
throttle-msgr_dispatch_throttler-hb_back_server/get | Gets
throttle-msgr_dispatch_throttler-hb_back_server/get_sum | Got data
throttle-msgr_dispatch_throttler-hb_back_server/get_or_fail_fail | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-hb_back_server/get_or_fail_success | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-hb_back_server/take | Takes
throttle-msgr_dispatch_throttler-hb_back_server/take_sum | Taken data
throttle-msgr_dispatch_throttler-hb_back_server/put | Puts
throttle-msgr_dispatch_throttler-hb_back_server/put_sum | Put data
throttle-msgr_dispatch_throttler-hb_back_server/wait | Waiting latency
throttle-msgr_dispatch_throttler-hb_front_server/val | Currently available throttle
throttle-msgr_dispatch_throttler-hb_front_server/max | Max value for throttle
throttle-msgr_dispatch_throttler-hb_front_server/get | Gets
throttle-msgr_dispatch_throttler-hb_front_server/get_sum | Got data
throttle-msgr_dispatch_throttler-hb_front_server/get_or_fail_fail | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-hb_front_server/get_or_fail_success | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-hb_front_server/take | Takes
throttle-msgr_dispatch_throttler-hb_front_server/take_sum | Taken data
throttle-msgr_dispatch_throttler-hb_front_server/put | Puts
throttle-msgr_dispatch_throttler-hb_front_server/put_sum | Put data
throttle-msgr_dispatch_throttler-hb_front_server/wait | Waiting latency
throttle-msgr_dispatch_throttler-hbclient/val | Currently available throttle
throttle-msgr_dispatch_throttler-hbclient/max | Max value for throttle
throttle-msgr_dispatch_throttler-hbclient/get | Gets
throttle-msgr_dispatch_throttler-hbclient/get_sum | Got data
throttle-msgr_dispatch_throttler-hbclient/get_or_fail_fail | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-hbclient/get_or_fail_success | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-hbclient/take | Takes
throttle-msgr_dispatch_throttler-hbclient/take_sum | Taken data
throttle-msgr_dispatch_throttler-hbclient/put | Puts
throttle-msgr_dispatch_throttler-hbclient/put_sum | Put data
throttle-msgr_dispatch_throttler-hbclient/wait | Waiting latency
throttle-msgr_dispatch_throttler-ms_objecter/val | Currently available throttle
throttle-msgr_dispatch_throttler-ms_objecter/max | Max value for throttle
throttle-msgr_dispatch_throttler-ms_objecter/get | Gets
throttle-msgr_dispatch_throttler-ms_objecter/get_sum | Got data
throttle-msgr_dispatch_throttler-ms_objecter/get_or_fail_fail | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-ms_objecter/get_or_fail_success | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-ms_objecter/take | Takes
throttle-msgr_dispatch_throttler-ms_objecter/take_sum | Taken data
throttle-msgr_dispatch_throttler-ms_objecter/put | Puts
throttle-msgr_dispatch_throttler-ms_objecter/put_sum | Put data
throttle-msgr_dispatch_throttler-ms_objecter/wait | Waiting latency
throttle-objecter_bytes/val | Currently available throttle
throttle-objecter_bytes/max | Max value for throttle
throttle-objecter_bytes/get | Gets
throttle-objecter_bytes/get_sum | Got data
throttle-objecter_bytes/get_or_fail_fail | Get blocked during get_or_fail
throttle-objecter_bytes/get_or_fail_success | Successful get during get_or_fail
throttle-objecter_bytes/take | Takes
throttle-objecter_bytes/take_sum | Taken data
throttle-objecter_bytes/put | Puts
throttle-objecter_bytes/put_sum | Put data
throttle-objecter_bytes/wait | Waiting latency
throttle-objecter_ops/val | Currently available throttle
throttle-objecter_ops/max | Max value for throttle
throttle-objecter_ops/get | Gets
throttle-objecter_ops/get_sum | Got data
throttle-objecter_ops/get_or_fail_fail | Get blocked during get_or_fail
throttle-objecter_ops/get_or_fail_success | Successful get during get_or_fail
throttle-objecter_ops/take | Takes
throttle-objecter_ops/take_sum | Taken data
throttle-objecter_ops/put | Puts
throttle-objecter_ops/put_sum | Put data
throttle-objecter_ops/wait | Waiting latency
throttle-osd_client_bytes/val | Currently available throttle
throttle-osd_client_bytes/max | Max value for throttle
throttle-osd_client_bytes/get | Gets
throttle-osd_client_bytes/get_sum | Got data
throttle-osd_client_bytes/get_or_fail_fail | Get blocked during get_or_fail
throttle-osd_client_bytes/get_or_fail_success | Successful get during get_or_fail
throttle-osd_client_bytes/take | Takes
throttle-osd_client_bytes/take_sum | Taken data
throttle-osd_client_bytes/put | Puts
throttle-osd_client_bytes/put_sum | Put data
throttle-osd_client_bytes/wait | Waiting latency
throttle-osd_client_messages/val | Currently available throttle
throttle-osd_client_messages/max | Max value for throttle
throttle-osd_client_messages/get | Gets
throttle-osd_client_messages/get_sum | Got data
throttle-osd_client_messages/get_or_fail_fail | Get blocked during get_or_fail
throttle-osd_client_messages/get_or_fail_success | Successful get during get_or_fail
throttle-osd_client_messages/take | Takes
throttle-osd_client_messages/take_sum | Taken data
throttle-osd_client_messages/put | Puts
throttle-osd_client_messages/put_sum | Put data
throttle-osd_client_messages/wait | Waiting latency

