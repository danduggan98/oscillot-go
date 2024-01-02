package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "whaddup!")
}

func pingHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "PING!")
}

func main() {

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/ping", pingHandler)

	log.Fatal(http.ListenAndServe(":9000", nil))
}
