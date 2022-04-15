package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)

}

func handleFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Hello, goblog</h1>")
}
