package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/azavorotnii/grpc-example/example"
	"google.golang.org/grpc"
)

type BackoffTimeout struct {
	min, max, next time.Duration
}

func NewTimeout(min, max time.Duration) *BackoffTimeout {
	return &BackoffTimeout{min, max, min}
}

func (bt *BackoffTimeout) Next() time.Duration {
	current := bt.next
	next := current * 2
	if next > bt.max {
		next = bt.max
	}
	// adding jitter
	next = (next / 2) + time.Duration(rand.Intn(int(next)/2))
	return current
}

func (bt *BackoffTimeout) Reset() {
	bt.next = bt.min
}

func main() {
	var address string
	flag.StringVar(&address, "server", "localhost:8080", "")
	flag.Parse()

	timeout := NewTimeout(time.Second, time.Minute)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Println(err)
			t := timeout.Next()
			log.Println(timeout, t)
			time.Sleep(t)
			continue
		}
		defer func() { _ = conn.Close() }()

		log.Println("Connected.")
		timeout.Reset()

		client := example.NewCalculatorClient(conn)

		for {
			ctx := context.Background()

			var args []*example.Complex
			for i := rand.Intn(5) + 1; i >= 0; i-- {
				args = append(args, &example.Complex{Real: float64(rand.Intn(10)), Imag: 0})
			}

			result, err := client.Add(ctx, &example.ComplexArgs{Arg: args})
			if err != nil {
				log.Println(err)
				break
			}
			msg := fmt.Sprintf("(%v, %v)", args[0].Real, args[0].Imag)
			for i := 1; i < len(args); i++ {
				msg += fmt.Sprintf("+ (%v, %v)", args[i].Real, args[i].Imag)
			}
			msg += fmt.Sprintf(" = (%v, %v)", result.Real, result.Imag)
			fmt.Println(msg)

			time.Sleep(5 * time.Second)
		}
	}
}
