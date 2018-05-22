package main

import (
	"log"
	"net/http"
)

func main() {
	props = loadProps()
	openConn(getConnString(props))
	router := NewRouter()
	log.Println("Server started on port 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
