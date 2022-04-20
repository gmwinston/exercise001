package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8080", nil) //設置監聽的埠
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello................"))
}

// 處理/upload 邏輯
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //獲取請求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.gohtml")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(1024)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		log.Println("copy.....")
	}
}
