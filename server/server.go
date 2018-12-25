package main

import (
	"fmt"
	"log"
	"net/http"
)

/*

1. Execute as is
2. Add a flag for specifying port number
3. Add a flag for specifying webroot

*/

func main() {
	port := "8080"
	webroot := "/var/webroot"

	log.Printf("Serving on port %s", port)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		http.FileServer(http.Dir(webroot)),
	))
}
