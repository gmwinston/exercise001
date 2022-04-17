package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"net/http"
	"strings"
)

func main() {
	rout := mux.NewRouter().StrictSlash(true)
	rout.HandleFunc("/", handleFunc)
	rout.HandleFunc("/art/{id:[0-9]+}", articleshowhandler).Methods("GET").Name("article.show....")
	rout.HandleFunc("/bbb1/", articlepost).Methods("POST").Name("article.showP....")
	rout.NotFoundHandler = http.HandlerFunc(n404)
	rout.Use(middleware)
	http.ListenAndServe(":3000", removeTrailingSlash(rout))

}

func removeTrailingSlash(rout *mux.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		rout.ServeHTTP(w,r)
	})
}

func middleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
//		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
//		next.ServeHTTP(w, r)
     return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
	 //return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		  w.Header().Set("Content-Type", "text/html; charset=utf-8")
		  w.Write([]byte("middle................."))
		  next.ServeHTTP(w, r)

	})

}

func articlepost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post..........."))
	fmt.Fprintf(w,"posttttt......")
}

func n404(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("404................"))
}

func articleshowhandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
		fmt.Fprintf(w, "id:"+id+"/n")

}

func handleFunc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "<h1>Hello, goblog</h1>")
	writer.Write([]byte("path:" + request.URL.Path + "; "+ request.Method))
}
