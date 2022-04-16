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
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(writer, "<h1>Hello, goblog</h1>")
	writer.Write([]byte("path:" + request.URL.Path + "; "+ request.Method))
}
