package main

import (
	"flag"
	"net/http"
	"fmt"
	"os"
	"syscall"
	"os/signal"
	"log"
)

func main() {
	path := flag.String("path", "./", "path of directory to serve")
	port := flag.Int("port", 8080, "port for file server")
	flag.Parse()

	stop:= make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fs := http.FileServer(http.Dir(*path))
		http.Handle("/", fs)

		log.Printf("Listening to localhost:%d...\n", *port)
		err := http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), nil)
		if err != nil {

		}
	}()

	<-stop
	log.Printf("Server was stopped gracefully.")
}
