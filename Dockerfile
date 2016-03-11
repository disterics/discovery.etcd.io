FROM golang:onbuild
MAINTAINER "CoreOS, Inc"
EXPOSE 8087
ENTRYPOINT ["go-wrapper", "run"]
CMD ["--addr=:8087"]
