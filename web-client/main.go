package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	var addr, backendAddr string
	flag.StringVar(&addr, "addr", ":8081", "")
	flag.StringVar(&backendAddr, "backendAddr", ":8080", "")
	flag.Parse()

	ctx := context.Background()
	backendConn, err := grpc.DialContext(ctx, backendAddr, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		outCtx, _ := context.WithCancel(ctx)
		outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())
		return outCtx, backendConn, nil
	}
	s := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()), // needed for proxy to function.
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
	)
	px := grpcweb.WrapServer(s,
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(string) bool { return true }),
	)
	fileServer := http.FileServer(
		&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "assets"},
	)

	err = http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if px.IsAcceptableGrpcCorsRequest(r) || px.IsGrpcWebRequest(r) || px.IsGrpcWebSocketRequest(r) {
			px.ServeHTTP(w, r)
			return
		}
		fileServer.ServeHTTP(w, r)
	}))
	if err != nil {
		log.Println(err)
	}
}
