package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/azavorotnii/grpc-example/example"
	"google.golang.org/grpc"
)

type server struct{}

func (server) Add(ctx context.Context, args *example.ComplexArgs) (*example.Complex, error) {

	result := complex(0, 0)
	for _, arg := range args.Arg {
		result += complex(arg.Real, arg.Imag)
	}
	return &example.Complex{Real: real(result), Imag: imag(result)}, nil
}

func main() {
	var address string
	flag.StringVar(&address, "addr", ":8080", "")
	flag.Parse()

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	s := grpc.NewServer()
	example.RegisterCalculatorServer(s, server{})
	if err := s.Serve(l); err != nil {
		log.Panic(err)
	}
}
