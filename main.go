package main

import (
	"log"
	"net/http"
	_ "webt/routes"
)

func main() {

	// public views

	log.Fatal(http.ListenAndServe(":8080", nil))
}
