// https://github.com/cloudnativedevops/demo/tree/main/hello (with some modifications)

/*

TO-DO:

[*] Environment variables
 > CONTAINER_ID: Display container's ID
 > HTTP_PORT: Setup HTTP port (default is 8888)
[*] Logging
 > Log to STDOUT every HTTP request

*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from simple-go-app" )
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running server. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
