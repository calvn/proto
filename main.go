package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/calvn/proto/handlers"
)

const version = "1.0.1"

func main() {
	var httpAddr = flag.String("http", "0.0.0.0:8000", "HTTP service address")
	flag.Parse()

	log.Println("Starting server")

	errChan := make(chan error)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HelloHandler)
	mux.HandleFunc("/aws/identity", handlers.IdentityHandler)
	mux.Handle("/version", handlers.VersionHandler(version))

	httpServer := manners.NewServer()
	httpServer.Addr = *httpAddr
	httpServer.Handler = handlers.LoggingHandler(mux)

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
