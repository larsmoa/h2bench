package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var serverAddress = ":8000"
var tlsCertFile = ""
var tlsKeyFile = ""

func parseFlags() {
	flag.StringVar(&serverAddress, "serverAddress", serverAddress,
		"The address to bind to")
	flag.StringVar(&tlsCertFile, "cert", tlsCertFile,
		"TLS certificate to use to secure the HTTP link.")
	flag.StringVar(&tlsKeyFile, "key", tlsKeyFile,
		"TLS private key to use to secure the HTTP link.")
	flag.Parse()
}

func handleRandomBytes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	byteCount, _ := strconv.Atoi(vars["byteCount"])

	fmt.Printf("")

	buf := make([]byte, byteCount)
	rand.Read(buf)

	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(buf)
}

func startServer() error {
	r := mux.NewRouter()
	r.Handle("/random/{byteCount:[0-9]+}",
		handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handleRandomBytes)))

	server := http.Server{
		Addr:    serverAddress,
		Handler: r,
	}

	fmt.Println("Supported requests:")
	fmt.Println(" - GET /random/{byteCount} returns random bytes back.")
	fmt.Println()

	if tlsCertFile != "" && tlsKeyFile != "" {
		fmt.Printf("Serving at %s using HTTPS...\n", serverAddress)
		return server.ListenAndServeTLS(tlsCertFile, tlsKeyFile)
	}
	fmt.Printf("Serving at %s using HTTP...\n", serverAddress)
	return server.ListenAndServe()
}

func main() {
	parseFlags()
	if err := startServer(); err != nil {
		log.Fatal(err)
	}
}
