package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(&priceFetcher{})

	srv := NewJSONAPIServer(*listenAddr, svc)
	srv.Run()

}
