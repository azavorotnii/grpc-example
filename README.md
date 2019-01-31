Requires following steps to init work:

* Load dependencies:
```bash
dep ensure
```
* install the protoc compiler that is used to generate gRPC service code. The simplest way to do this is to download pre-compiled binaries for your platform(protoc-<version>-<platform>.zip) from here: https://github.com/google/protobuf/releases
  Unzip this file.
  Update the environment variable PATH to include the path to the protoc binary file.
* install the protoc plugin for Go
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$GOPATH/bin
```