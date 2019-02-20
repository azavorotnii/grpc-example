package main

//go:generate protoc -I ../example example.proto3 --js_out=import_style=commonjs:../web-client/
//go:generate protoc -I ../example example.proto3 --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../web-client/
//go:generate npm install
//go:generate browserify index.js -o assets/bundle_gen.js
//go:generate go-bindata -o assets_gen.go assets/
