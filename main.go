package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mukeshmahato17/pricefetcher/client"
	"github.com/mukeshmahato17/pricefetcher/proto"
)

func main() {
	var (
		jsonAddr = flag.String("jsonaddr", ":3000", "listen address of the JSON service")
		grpcAddr = flag.String("grpcaddr", ":4000", "listen address of the gRPC service")
		svc      = NewLoggingService(&priceFetcher{})
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(*grpcAddr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(time.Second * 3)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v", resp)
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
