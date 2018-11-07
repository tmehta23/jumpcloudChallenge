// main.go - contains main method that starts/ runs server,
// handles endpoint requests, and performs shutdown.
// See ../README.md for instructions on how to create exe.

package main

import "context"
import "log"
import "net/http"

// Main method that runs server and accepts requests by endpoint.
func main() {

  // Create new server and http request multiplexer.
  mux := http.NewServeMux()
  server := http.Server{Addr:":8080", Handler:mux}

  // For hash and stats endpoints, direct request to respective handler function.
  mux.HandleFunc("/hash", hashHandler)
  mux.HandleFunc("/stats", statsHandler)

  // For shutdown endpoint, perform graceful shutdown of server.
  mux.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
        server.Shutdown(context.Background())
    })

  // If error in shutdown, log.
  if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    log.Fatal(err)
  }

  // log "Shutdown success".
  log.Printf("\nShutdown has been successfully completed. The program will not take new requests at this time.")
}
