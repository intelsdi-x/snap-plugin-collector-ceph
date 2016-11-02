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

## How to deploy fake local Ceph cluster

If You would like to check snap with Ceph collector plugin in local machine, clone Ceph's source from https://github.com/ceph/ceph and following the instruction to build it. 

Next, go to $CEPH_DIR/src and create virtual cluster:
```
$ MON=3 OSD=3 MDS=3 ./vstart.sh -n -J -d
```

Check status of created cluster, should be HEALTH_OK:
```
$ ./ceph -s

*** DEVELOPER MODE: setting PATH, PYTHONPATH and LD_LIBRARY_PATH ***
    cluster f6d89e76-2a7f-4538-beed-4118ae3cf342
     health HEALTH_OK
     monmap e1: 3 mons at {a=10.102.108.166:6789/0,b=10.102.108.166:6790/0,c=10.102.108.166:6791/0}
            election epoch 6, quorum 0,1,2 a,b,c
     mdsmap e7: 3/3/3 up {0=a=up:creating,1=c=up:creating,2=b=up:creating}
     osdmap e10: 3 osds: 3 up, 3 in
      pgmap v14: 24 pgs, 3 pools, 4230 bytes data, 54 objects
            113 GB used, 271 GB / 405 GB avail
                  24 active+clean
  client io 4084 B/s wr, 11 op/s
```
In the $CEPH_DIR/src/out directory should occure ceph-daemon asok. Set path to them in  snap Global Config, also customize the socket prefix and extension might be required.

Now You are ready to use the snap to collect ceph perf counters from local cluster!

Create test pool and make some writes to see that values of perf counters will be changed:

```
./ceph osd pool create test_pool 128 128
./rados bench -p test_pool 10 write
```

To delete a pool, execute:
```
ceph osd pool delete test_pool test_pool --yes-i-really-really-mean-it
```

To stop virtual cluster, execute:
```
./stop.sh
```

More info can be found at:
 - http://docs.ceph.com/docs/jewel/dev/dev_cluster_deployement/
 - http://docs.ceph.com/docs/jewel/rados/operations/pools/
 - http://docs.ceph.com/docs/jewel/man/8/rados/
