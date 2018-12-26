package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
)

/*

1. Execute as is
2. Add a flag for specifying port number
3. Add a flag for specifying webroot

*/

func main() {

	//port := flag.String("port", "8080", "Default port is 8080")
	var port string
	flag.StringVar(&port, "port", "8080", "Default port is 8080")

	var webroot string
	flag.StringVar(&webroot, "webroot", "/var/webroot", "Default root is /var/webroot")

	flag.Parse()
	//port := "8080"
	//webroot := "/var/webroot"

	log.Printf("Serving on port %s", port)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		http.FileServer(http.Dir(webroot)),
	))
}
