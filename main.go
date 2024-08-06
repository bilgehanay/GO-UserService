package main

import "net/http"

func main() {
	port := ":8080"
	err := http.ListenAndServe(port, RouterGroup())
	if err != nil {
		panic(err)
	}
}
