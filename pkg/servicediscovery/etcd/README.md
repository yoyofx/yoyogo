## docker 
```bash
docker run -p 2379:2379 --name etcd quay.io/coreos/etcd:v3.1.0 /usr/local/bin/etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379 --initial-advertise-peer-urls http://0.0.0.0:2380 --listen-peer-urls http://0.0.0.0:2380
```

```bash
docker exec -it etcd bin/sh

export ETCDCTL_API=3

etcdctl get /public/yoyogo_demo_dev/8ca0e6c8-1896-4fdc-a775-ab565a84f166

```


```bash
etcdctl version

etcdctl version: 3.1.0
API version: 3
```

```bash
etcdctl role add root
etcdctl user add root
etcdctl user  grant-role  root root
etcdctl auth enable
```