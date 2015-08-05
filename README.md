# discovery.etcd.io

This code powers the public service at https://discovery.etcd.io. The API is
documented in the etcd clustering documentation:

https://github.com/coreos/etcd/blob/master/Documentation/clustering.md#public-etcd-discovery-service

## Docker Container

You may run the etcd cluster in a docker container for discovery service:

```
TOKEN=`curl -X PUT https://discovery.etcd.io/new?size=3`
docker pull quay.io/coreos/etcd:v0.4.6
docker run -d -p 4001:4001 -p 7001:7001 -v /usr/share/ca-certificates:/etc/ssl/certs \
  --name etcd-01 quay.io/coreos/etcd:v0.4.6 \
  -addr 172.17.42.1:4001 -peer-addr 172.17.42.1:7001 -discovery $TOKEN
docker run -d -p 5001:5001 -p 8001:8001 -v /usr/share/ca-certificates:/etc/ssl/certs \
  --name etcd-02 quay.io/coreos/etcd:v0.4.6 \
  -addr 172.17.42.1:5001 -peer-addr 172.17.42.1:8001 -discovery $TOKEN
docker run -d -p 6001:6001 -p 9001:9001 -v /usr/share/ca-certificates:/etc/ssl/certs \
  --name etcd-03 quay.io/coreos/etcd:v0.4.6 \
  -addr 172.17.42.1:6001 -peer-addr 172.17.42.1:9001 -discovery $TOKEN
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
