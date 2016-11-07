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


# Ceph MDS Perf Counters

Prefix: /intel/storage/ceph/{mds_daemon_name}/{mds_daemon_id}

Metrics | Description
------------ | -------------
mds/request | Requests
mds/reply | Replies
mds/reply_latency | Reply latency
mds/forward | Forwarding request
mds/dir_fetch | Directory fetch
mds/dir_commit | Directory commit
mds/dir_split | Directory split
mds/inode_max | Max inodes, cache size
mds/inodes | Inodes
mds/inodes_top | Inodes on top
mds/inodes_bottom | Inodes on bottom
mds/inodes_pin_tail | Inodes on pin tail
mds/inodes_pinned | Inodes pinned
mds/inodes_expired | Inodes expired
mds/inodes_with_caps | Inodes with capabilities
mds/caps | Capabilities
mds/subtrees | Subtrees
mds/traverse | Traverses
mds/traverse_hit | Traverse hits
mds/traverse_forward | Traverse forwards
mds/traverse_discover | Traverse directory discovers
mds/traverse_dir_fetch | Traverse incomplete directory content fetchings
mds/traverse_remote_ino | Traverse remote dentries
mds/traverse_lock | Traverse locks
mds/load_cent | Load per cent
mds/q | Dispatch queue length
mds/exported | Exports
mds/exported_inodes | Exported inodes
mds/imported | Imports
mds/imported_inodes | Imported inodes
mds_cache/num_strays | Stray dentries
mds_cache/num_strays_purging | Stray dentries purging
mds_cache/num_strays_delayed | Stray dentries delayed
mds_cache/num_purge_ops | Purge operations
mds_cache/strays_created | Stray dentries created
mds_cache/strays_purged | Stray dentries purged
mds_cache/strays_reintegrated | Stray dentries reintegrated
mds_cache/strays_migrated | Stray dentries migrated
mds_cache/num_recovering_processing | Files currently being recovered
mds_cache/num_recovering_enqueued | Files waiting for recovery
mds_cache/num_recovering_prioritized | Files waiting for recovery with elevated priority
mds_cache/recovery_started | File recoveries started
mds_cache/recovery_completed | File recoveries completed
mds_log/evadd | Events submitted
mds_log/evex | Total expired events
mds_log/evtrm | Trimmed events
mds_log/ev | Events
mds_log/evexg | Expiring events
mds_log/evexd | Current expired events
mds_log/segadd | Segments added
mds_log/segex | Total expired segments
mds_log/segtrm | Trimmed segments
mds_log/seg | Segments
mds_log/segexg | Expiring segments
mds_log/segexd | Current expired segments
mds_log/expos | Journaler xpire position
mds_log/wrpos | Journaler  write position
mds_log/rdpos | Journaler  read position
mds_log/jlat | Journaler flush latency
mds_mem/ino | Inodes
mds_mem/ino+ | Inodes opened
mds_mem/ino- | Inodes closed
mds_mem/dir | Directories
mds_mem/dir+ | Directories opened
mds_mem/dir- | Directories closed
mds_mem/dn | Dentries
mds_mem/dn+ | Dentries opened
mds_mem/dn- | Dentries closed
mds_mem/cap | Capabilities
mds_mem/cap+ | Capabilities added
mds_mem/cap- | Capabilities removed
mds_mem/rss | RSS
mds_mem/heap | Heap size
mds_mem/malloc | Malloc size
mds_mem/buf | Buffer size
mds_server/handle_client_request | Client requests
mds_server/handle_slave_request | Slave requests
mds_server/handle_client_session | Client session messages
mds_server/dispatch_client_request | Client requests dispatched
mds_server/dispatch_server_request | Server requests dispatched
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
throttle-msgr_dispatch_throttler-mds/val | Currently available throttle
throttle-msgr_dispatch_throttler-mds/max | Max value for throttle
throttle-msgr_dispatch_throttler-mds/get | Gets
throttle-msgr_dispatch_throttler-mds/get_sum | Got data
throttle-msgr_dispatch_throttler-mds/get_or_fail_fail | Get blocked during get_or_fail
throttle-msgr_dispatch_throttler-mds/get_or_fail_success | Successful get during get_or_fail
throttle-msgr_dispatch_throttler-mds/take | Takes
throttle-msgr_dispatch_throttler-mds/take_sum | Taken data
throttle-msgr_dispatch_throttler-mds/put | Puts
throttle-msgr_dispatch_throttler-mds/put_sum | Put data
throttle-msgr_dispatch_throttler-mds/wait | Waiting latency
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
