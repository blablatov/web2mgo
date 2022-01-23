package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"web2mgo"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// this handler is returning component path of URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// parsing of string
	i := strings.Index(r.URL.Path, ":") // get index of symbol ":"
	tp := &Mongo{}
	tp.Name = r.URL.Path[:i]   // get slice of string before symbol ":"
	tp.Data = r.URL.Path[i+1:] // get slice of string after symbol ":"
	fmt.Println("Name:", strings.TrimPrefix(tp.Name, "/"), "\nData:", tp.Data)

	go web2mgo.ODatareq(strings.TrimPrefix(tp.Name, "/"), tp.Data)
}
