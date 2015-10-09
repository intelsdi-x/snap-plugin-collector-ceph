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


# Ceph OSD Perf Counters

Prefix: /intel/storage/ceph/{osd_daemon_name}/

Metrics | Description
------------ | -------------
WBThrottle/bytes_dirtied | Dirty data 
WBThrottle/bytes_wb | Written data 
WBThrottle/ios_dirtied | Dirty operations 
WBThrottle/ios_wb | Written operations 
WBThrottle/inodes_dirtied | Entries waiting for write 
WBThrottle/inodes_wb | Written entries 
 | 
filestore/journal_queue_max_ops | Max operations in journal queue 
filestore/journal_queue_ops | Operations in journal queue 
filestore/journal_ops | Total journal entries written 
filestore/journal_queue_max_bytes | Max data in journal queue 
filestore/journal_queue_bytes | Size of journal queue 
filestore/journal_bytes | Total operations size in journal 
filestore/journal_latency/avgcount | Average journal queue completing latency 
filestore/journal_latency/sum | Average journal queue completing latency 
filestore/journal_wr | Journal write IOs 
filestore/journal_wr_bytes/avgcount | Journal data written 
filestore/journal_wr_bytes/sum | Journal data written 
filestore/journal_full | Journal writes while full 
filestore/committing | Is currently committing 
filestore/commitcycle | Commit cycles 
filestore/commitcycle_interval/avgcount | Average interval between commits 
filestore/commitcycle_interval/sum | Average interval between commits 
filestore/commitcycle_latency/avgcount | Average latency of commit 
filestore/commitcycle_latency/sum | Average latency of commit 
filestore/op_queue_max_ops | Max operations in writing to FS queue 
filestore/op_queue_ops | Operations in writing to FS queue 
filestore/ops | Operations written to store 
filestore/op_queue_max_bytes | Max data in writing to FS queue 
filestore/op_queue_bytes | Size of writing to FS queue 
filestore/bytes | Data written to store 
filestore/apply_latency/avgcount | Apply latency 
filestore/apply_latency/sum | Apply latency 
filestore/queue_transaction_latency_avg/avgcount | Store operation queue latency 
filestore/queue_transaction_latency_avg/sum | Store operation queue latency 
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
mutex-FileJournal::completions_lock/wait/avgcount | Average time of mutex in locked state 
mutex-FileJournal::completions_lock/wait/sum | Average time of mutex in locked state 
mutex-FileJournal::finisher_lock/wait/avgcount | Average time of mutex in locked state 
mutex-FileJournal::finisher_lock/wait/sum | Average time of mutex in locked state 
mutex-FileJournal::write_lock/wait/avgcount | Average time of mutex in locked state 
mutex-FileJournal::write_lock/wait/sum | Average time of mutex in locked state 
mutex-FileJournal::writeq_lock/wait/avgcount | Average time of mutex in locked state 
mutex-FileJournal::writeq_lock/wait/sum | Average time of mutex in locked state 
mutex-JOS::ApplyManager::apply_lock/wait/avgcount | Average time of mutex in locked state 
mutex-JOS::ApplyManager::apply_lock/wait/sum | Average time of mutex in locked state 
mutex-JOS::ApplyManager::com_lock/wait/avgcount | Average time of mutex in locked state 
mutex-JOS::ApplyManager::com_lock/wait/sum | Average time of mutex in locked state 
mutex-JOS::SubmitManager::lock/wait/avgcount | Average time of mutex in locked state 
mutex-JOS::SubmitManager::lock/wait/sum | Average time of mutex in locked state 
mutex-WBThrottle::lock/wait/avgcount | Average time of mutex in locked state 
mutex-WBThrottle::lock/wait/sum | Average time of mutex in locked state 
 | 
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
objecter/osdop_pgls |  [No description available]
objecter/osdop_pgls_filter |  [No description available]
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
 | 
osd/op_wip | Replication operations currently being processed (primary) 
osd/op | Client operations 
osd/op_in_bytes | Client operations total write size 
osd/op_out_bytes | Client operations total read size 
osd/op_latency/avgcount | Latency of client operations (including queue time) 
osd/op_latency/sum | Latency of client operations (including queue time) 
osd/op_process_latency/avgcount | Latency of client operations (excluding queue time) 
osd/op_process_latency/sum | Latency of client operations (excluding queue time) 
osd/op_r | Client read operations 
osd/op_r_out_bytes | Client data read 
osd/op_r_latency/avgcount | Latency of read operation (including queue time) 
osd/op_r_latency/sum | Latency of read operation (including queue time) 
osd/op_r_process_latency/avgcount | Latency of read operation (excluding queue time) 
osd/op_r_process_latency/sum | Latency of read operation (excluding queue time) 
osd/op_w | Client write operations 
osd/op_w_in_bytes | Client data written 
osd/op_w_rlat/avgcount | Client write operation readable\/applied latency 
osd/op_w_rlat/sum | Client write operation readable\/applied latency 
osd/op_w_latency/avgcount | Latency of write operation (including queue time) 
osd/op_w_latency/sum | Latency of write operation (including queue time) 
osd/op_w_process_latency/avgcount | Latency of write operation (excluding queue time) 
osd/op_w_process_latency/sum | Latency of write operation (excluding queue time) 
osd/op_rw | Client read-modify-write operations 
osd/op_rw_in_bytes | Client read-modify-write operations write in 
osd/op_rw_out_bytes | Client read-modify-write operations read out  
osd/op_rw_rlat/avgcount | Client read-modify-write operation readable\/applied latency 
osd/op_rw_rlat/sum | Client read-modify-write operation readable\/applied latency 
osd/op_rw_latency/avgcount | Latency of read-modify-write operation (including queue time) 
osd/op_rw_latency/sum | Latency of read-modify-write operation (including queue time) 
osd/op_rw_process_latency/avgcount | Latency of read-modify-write operation (excluding queue time) 
osd/op_rw_process_latency/sum | Latency of read-modify-write operation (excluding queue time) 
osd/subop | Suboperations 
osd/subop_in_bytes | Suboperations total size 
osd/subop_latency/avgcount | Suboperations latency 
osd/subop_latency/sum | Suboperations latency 
osd/subop_w | Replicated writes 
osd/subop_w_in_bytes | Replicated written data size 
osd/subop_w_latency/avgcount | Replicated writes latency 
osd/subop_w_latency/sum | Replicated writes latency 
osd/subop_pull | Suboperations pull requests 
osd/subop_pull_latency/avgcount | Suboperations pull latency 
osd/subop_pull_latency/sum | Suboperations pull latency 
osd/subop_push | Suboperations push messages 
osd/subop_push_in_bytes | Suboperations pushed size 
osd/subop_push_latency/avgcount | Suboperations push latency 
osd/subop_push_latency/sum | Suboperations push latency 
osd/pull | Pull requests sent 
osd/push | Push messages sent 
osd/push_out_bytes | Pushed size 
osd/push_in | Inbound push messages 
osd/push_in_bytes | Inbound pushed size 
osd/recovery_ops | Started recovery operations 
osd/loadavg | CPU load 
osd/buffer_bytes | Total allocated buffer size 
osd/numpg | Placement groups 
osd/numpg_primary | Placement groups for which this osd is primary 
osd/numpg_replica | Placement groups for which this osd is replica 
osd/numpg_stray | Placement groups ready to be deleted from this osd 
osd/heartbeat_to_peers | Heartbeat (ping) peers we send to 
osd/heartbeat_from_peers | Heartbeat (ping) peers we recv from 
osd/map_messages | OSD map messages 
osd/map_message_epochs | OSD map epochs 
osd/map_message_epoch_dups | OSD map duplicates 
osd/messages_delayed_for_map | Operations waiting for OSD map 
osd/stat_bytes | OSD size 
osd/stat_bytes_used | Used space 
osd/stat_bytes_avail | Available space 
osd/copyfrom | Rados \"copy-from\" operations
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
osd/op_cache_hit | 
osd/op_tier_flush_lat | Object flush latency 
osd/op_tier_promote_lat | Object promote latency 
osd/op_tier_r_lat | Object proxy read latency 
 | 
recoverystate_perf/initial_latency/avgcount | Initial recovery state latency 
recoverystate_perf/initial_latency/sum | Initial recovery state latency 
recoverystate_perf/started_latency/avgcount | Started recovery state latency 
recoverystate_perf/started_latency/sum | Started recovery state latency 
recoverystate_perf/reset_latency/avgcount | Reset recovery state latency 
recoverystate_perf/reset_latency/sum | Reset recovery state latency 
recoverystate_perf/start_latency/avgcount | Start recovery state latency 
recoverystate_perf/start_latency/sum | Start recovery state latency 
recoverystate_perf/primary_latency/avgcount | Primary recovery state latency 
recoverystate_perf/primary_latency/sum | Primary recovery state latency 
recoverystate_perf/peering_latency/avgcount | Peering recovery state latency 
recoverystate_perf/peering_latency/sum | Peering recovery state latency 
recoverystate_perf/backfilling_latency/avgcount | Backfilling recovery state latency 
recoverystate_perf/backfilling_latency/sum | Backfilling recovery state latency 
recoverystate_perf/waitremotebackfillreserved_latency/avgcount | Wait remote backfill reserved recovery state latency 
recoverystate_perf/waitremotebackfillreserved_latency/sum | Wait remote backfill reserved recovery state latency 
recoverystate_perf/waitlocalbackfillreserved_latency/avgcount | Wait local backfill reserved recovery state latency 
recoverystate_perf/waitlocalbackfillreserved_latency/sum | Wait local backfill reserved recovery state latency 
recoverystate_perf/notbackfilling_latency/avgcount | Notbackfilling recovery state latency 
recoverystate_perf/notbackfilling_latency/sum | Notbackfilling recovery state latency 
recoverystate_perf/repnotrecovering_latency/avgcount | Repnotrecovering recovery state latency 
recoverystate_perf/repnotrecovering_latency/sum | Repnotrecovering recovery state latency 
recoverystate_perf/repwaitrecoveryreserved_latency/avgcount | Rep wait recovery reserved recovery state latency 
recoverystate_perf/repwaitrecoveryreserved_latency/sum | Rep wait recovery reserved recovery state latency 
recoverystate_perf/repwaitbackfillreserved_latency/avgcount | Rep wait backfill reserved recovery state latency 
recoverystate_perf/repwaitbackfillreserved_latency/sum | Rep wait backfill reserved recovery state latency 
recoverystate_perf/RepRecovering_latency/avgcount | RepRecovering recovery state latency 
recoverystate_perf/RepRecovering_latency/sum | RepRecovering recovery state latency 
recoverystate_perf/activating_latency/avgcount | Activating recovery state latency 
recoverystate_perf/activating_latency/sum | Activating recovery state latency 
recoverystate_perf/waitlocalrecoveryreserved_latency/avgcount | Wait local recovery reserved recovery state latency 
recoverystate_perf/waitlocalrecoveryreserved_latency/sum | Wait local recovery reserved recovery state latency 
recoverystate_perf/waitremoterecoveryreserved_latency/avgcount | Wait remote recovery reserved recovery state latency 
recoverystate_perf/waitremoterecoveryreserved_latency/sum | Wait remote recovery reserved recovery state latency 
recoverystate_perf/recovering_latency/avgcount | Recovering recovery state latency 
recoverystate_perf/recovering_latency/sum | Recovering recovery state latency 
recoverystate_perf/recovered_latency/avgcount | Recovered recovery state latency 
recoverystate_perf/recovered_latency/sum | Recovered recovery state latency 
recoverystate_perf/clean_latency/avgcount | Clean recovery state latency 
recoverystate_perf/clean_latency/sum | Clean recovery state latency 
recoverystate_perf/active_latency/avgcount | Active recovery state latency 
recoverystate_perf/active_latency/sum | Active recovery state latency 
recoverystate_perf/replicaactive_latency/avgcount | Replicaactive recovery state latency 
recoverystate_perf/replicaactive_latency/sum | Replicaactive recovery state latency 
recoverystate_perf/stray_latency/avgcount | Stray recovery state latency 
recoverystate_perf/stray_latency/sum | Stray recovery state latency 
recoverystate_perf/getinfo_latency/avgcount | Getinfo recovery state latency 
recoverystate_perf/getinfo_latency/sum | Getinfo recovery state latency 
recoverystate_perf/getlog_latency/avgcount | Getlog recovery state latency 
recoverystate_perf/getlog_latency/sum | Getlog recovery state latency 
recoverystate_perf/waitactingchange_latency/avgcount | Waitactingchange recovery state latency 
recoverystate_perf/waitactingchange_latency/sum | Waitactingchange recovery state latency 
recoverystate_perf/incomplete_latency/avgcount | Incomplete recovery state latency 
recoverystate_perf/incomplete_latency/sum | Incomplete recovery state latency 
recoverystate_perf/getmissing_latency/avgcount | Getmissing recovery state latency 
recoverystate_perf/getmissing_latency/sum | Getmissing recovery state latency 
recoverystate_perf/waitupthru_latency/avgcount | Waitupthru recovery state latency 
recoverystate_perf/waitupthru_latency/sum | Waitupthru recovery state latency 
 | 
throttle-filestore_bytes/val | Currently available throttle 
throttle-filestore_bytes/max | Max value for throttle 
throttle-filestore_bytes/get | Gets 
throttle-filestore_bytes/get_sum | Got data 
throttle-filestore_bytes/get_or_fail_fail | Get blocked during get_or_fail 
throttle-filestore_bytes/get_or_fail_success | Successful get during get_or_fail 
throttle-filestore_bytes/take | Takes 
throttle-filestore_bytes/take_sum | Taken data 
throttle-filestore_bytes/put | Puts 
throttle-filestore_bytes/put_sum | Put data 
throttle-filestore_bytes/wait/avgcount | Waiting latency 
throttle-filestore_bytes/wait/sum | Waiting latency 
 | 
throttle-filestore_ops/val | Currently available throttle 
throttle-filestore_ops/max | Max value for throttle 
throttle-filestore_ops/get | Gets 
throttle-filestore_ops/get_sum | Got data 
throttle-filestore_ops/get_or_fail_fail | Get blocked during get_or_fail 
throttle-filestore_ops/get_or_fail_success | Successful get during get_or_fail 
throttle-filestore_ops/take | Takes 
throttle-filestore_ops/take_sum | Taken data 
throttle-filestore_ops/put | Puts 
throttle-filestore_ops/put_sum | Put data 
throttle-filestore_ops/wait/avgcount | Waiting latency 
throttle-filestore_ops/wait/sum | Waiting latency 
 | 
throttle-journal_bytes/val | Currently available throttle 
throttle-journal_bytes/max | Max value for throttle 
throttle-journal_bytes/get | Gets 
throttle-journal_bytes/get_sum | Got data 
throttle-journal_bytes/get_or_fail_fail | Get blocked during get_or_fail 
throttle-journal_bytes/get_or_fail_success | Successful get during get_or_fail 
throttle-journal_bytes/take | Takes 
throttle-journal_bytes/take_sum | Taken data 
throttle-journal_bytes/put | Puts 
throttle-journal_bytes/put_sum | Put data 
throttle-journal_bytes/wait/avgcount | Waiting latency 
throttle-journal_bytes/wait/sum | Waiting latency 
 | 
throttle-journal_ops/val | Currently available throttle 
throttle-journal_ops/max | Max value for throttle 
throttle-journal_ops/get | Gets 
throttle-journal_ops/get_sum | Got data 
throttle-journal_ops/get_or_fail_fail | Get blocked during get_or_fail 
throttle-journal_ops/get_or_fail_success | Successful get during get_or_fail 
throttle-journal_ops/take | Takes 
throttle-journal_ops/take_sum | Taken data 
throttle-journal_ops/put | Puts 
throttle-journal_ops/put_sum | Put data 
throttle-journal_ops/wait/avgcount | Waiting latency 
throttle-journal_ops/wait/sum | Waiting latency 
 | 
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
throttle-msgr_dispatch_throttler-client/wait/avgcount | Waiting latency 
throttle-msgr_dispatch_throttler-client/wait/sum | Waiting latency 
 | 
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
throttle-msgr_dispatch_throttler-cluster/wait/avgcount | Waiting latency 
throttle-msgr_dispatch_throttler-cluster/wait/sum | Waiting latency 
 | 
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
throttle-msgr_dispatch_throttler-hb_back_server/wait/avgcount | Waiting latency 
throttle-msgr_dispatch_throttler-hb_back_server/wait/sum | Waiting latency 
 | 
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
throttle-msgr_dispatch_throttler-hb_front_server/wait/avgcount | Waiting latency 
throttle-msgr_dispatch_throttler-hb_front_server/wait/sum | Waiting latency 
 | 
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
throttle-msgr_dispatch_throttler-hbclient/wait/avgcount | Waiting latency 
throttle-msgr_dispatch_throttler-hbclient/wait/sum | Waiting latency 
 | 
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
throttle-msgr_dispatch_throttler-ms_objecter/wait/avgcount | Waiting latency 
throttle-msgr_dispatch_throttler-ms_objecter/wait/sum | Waiting latency 
 | 
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
throttle-objecter_bytes/wait/avgcount | Waiting latency 
throttle-objecter_bytes/wait/sum | Waiting latency 
 | 
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
throttle-objecter_ops/wait/avgcount | Waiting latency 
throttle-objecter_ops/wait/sum | Waiting latency 
 | 
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
throttle-osd_client_bytes/wait/avgcount | Waiting latency 
throttle-osd_client_bytes/wait/sum | Waiting latency 
 | 
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
throttle-osd_client_messages/wait/avgcount | Waiting latency 
throttle-osd_client_messages/wait/sum | Waiting latency 
