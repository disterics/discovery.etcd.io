# discovery.etcd.io

This code powers the public service at https://discovery.etcd.io. The API is
documented in the etcd clustering documentation:

https://github.com/coreos/etcd/blob/master/Documentation/clustering.md#public-etcd-discovery-service

## Docker Container

You need to run etcd cluster for etcd discovery, follow the instructions from link below to run etcd cluster in a docker container

https://github.com/coreos/etcd/blob/master/Documentation/docker_guide.md

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
