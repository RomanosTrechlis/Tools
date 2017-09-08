package main

import (
	"flag"
	"net/http"
	"fmt"
)

func main() {
	path := flag.String("path", "./", "path of directory to serve")
	port := flag.Int("port", 8080, "port for file server")
	fs := http.FileServer(http.Dir(*path))
	http.Handle("/", fs)

	fmt.Printf("Listening to localhost:%d", *port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), nil)
}
