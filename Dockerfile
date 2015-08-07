FROM golang:onbuild
MAINTAINER "CoreOS, Inc"
EXPOSE 8087
CMD ["go-wrapper", "run", "--addr=:8087"]
