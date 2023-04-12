package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	go server("8080", "v1")
	go server("8080", "v2")
	port := fmt.Sprintf(":%s", "8080")
	log.Fatal(http.ListenAndServe(port, nil))

}
func server(port, version string) {
	route := fmt.Sprintf("/%s", version)
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to server version ", version)
		fmt.Fprintf(w, "API %s running", route)
	})
}
