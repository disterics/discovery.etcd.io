# discovery.etcd.io

This code powers the public service at https://discovery.etcd.io. The API is
documented in the etcd clustering documentation:

https://github.com/coreos/etcd/blob/master/Documentation/clustering.md#public-etcd-discovery-service

## Docker Container

You may run the etcd cluster in a docker container for discovery service:

```
export DISCOVERY_URL=`curl -X PUT https://discovery.etcd.io/new?size=3`
docker pull quay.io/coreos/etcd:v2.0.10
docker run -d -p 4001:4001 -p 2379:2379 -p 2380:2380 -v /usr/share/ca-certificates:/etc/ssl/certs \
  --name etcd-01 quay.io/coreos/etcd:v2.0.10 -name etcd-01 \
  -advertise-client-urls http://172.17.42.1:2379 \
  -initial-advertise-peer-urls http://172.17.42.1:2380 \
  -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
  -listen-peer-urls http://0.0.0.0:2380 \
  -discovery $DISCOVERY_URL
docker run -d -p 5001:5001 -p 3379:3379 -p 3380:3380 -v /usr/share/ca-certificates:/etc/ssl/certs \
  --name etcd-02 quay.io/coreos/etcd:v2.0.10 -name etcd-02 \
  -advertise-client-urls http://172.17.42.1:3379 \
  -initial-advertise-peer-urls http://172.17.42.1:3380 \
  -listen-client-urls http://0.0.0.0:3379,http://0.0.0.0:5001 \
  -listen-peer-urls http://0.0.0.0:3380 \
  -discovery $DISCOVERY_URL
docker run -d -p 6001:6001 -p 4379:4379 -p 4380:4380 -v /usr/share/ca-certificates:/etc/ssl/certs \
  --name etcd-03 quay.io/coreos/etcd:v2.0.10 -name etcd-03 \
  -advertise-client-urls http://172.17.42.1:4379 \
  -initial-advertise-peer-urls http://172.17.42.1:4380 \
  -listen-client-urls http://0.0.0.0:4379,http://0.0.0.0:6001 \
  -listen-peer-urls http://0.0.0.0:4380 \
  -discovery $DISCOVERY_URL
```

You may run the etcd discovery in a docker container:
```
docker pull quay.io/coreos/discovery.etcd.io
docker run -d -p 8080:8087 --name etcd-discovery quay.io/coreos/discovery.etcd.io -addr :8087 -host http://localhost:8080 \
  -etcd http://172.17.42.1:4001,http://172.17.42.1:5001,http://172.17.42.1:6001
```

Request for a discovery url:
```
curl -X PUT http://localhost:8080/new?size=3
```

## Development

discovery.etcd.io uses devweb for easy development. It is simple to get started:

```
./devweb
curl --verbose -X PUT localhost:8087/new
```
