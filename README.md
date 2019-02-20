Very simple example, which can be extended into real project (I hope).

Requires following steps to init work:

* install the protoc compiler that is used to generate gRPC service code. The simplest way to do this is to download
  pre-compiled binaries for your platform(protoc-<version>-<platform>.zip) from here:

  https://github.com/google/protobuf/releases

  Unzip this file. Update the environment variable _PATH_ to include the path to the protoc binary file.
* install the protoc plugin for Go
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$GOPATH/bin
```

* Load dependencies:
```bash
dep ensure
```

* Install go packages used for assets package:
```bash
go get github.com/jteeuwen/go-bindata/...
```

* Generate sources:
```bash
go generate ./...
```

* Run server:
```bash
go run ./server/
```

* Run client or web-client:
```bash
go run ./client/
go run ./web-client/
```
Web-client will serve both assets (html/js) and grpc-web-proxy.
