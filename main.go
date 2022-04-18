package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var rout = muxinit().StrictSlash(true)

func muxinit() *mux.Router {
	return mux.NewRouter()
}

func main() {

	//config := mysql.Config{
	//	User: "root",
	//	Passwd: "123456",
	//	Addr: "10.0.1.129:3306",
	//	DBName: "test001",
	//	AllowNativePasswords: true,
	//}
	//db, err := sql.Open("mysql", config.FormatDSN())
	db, err := sql.Open("mysql", "root:123456@tcp(10.0.1.129:3306)/test001")

	fmt.Println(db)
	if err != nil {
		fmt.Println("DB error...")
	}
	db.SetConnMaxLifetime(5*time.Second)
	db.SetMaxOpenConns(25)
	db.SetConnMaxIdleTime(25)

	err = db.Ping()
	if err != nil {
		fmt.Println("DB error 02....")
	}else {
		fmt.Println("DB conn ")
	}

	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body longtext COLLATE utf8mb4_unicode_ci
); `
	db.Exec(createArticlesSQL)
	if err != nil {
       fmt.Println("db exec error...")
	   fmt.Println(err)
	}else {
		fmt.Println("db ok")
	}

	stmt, err := db.Prepare("INSERT INTO articles(title, body) VALUES(?,?)")
	if err != nil {
		fmt.Println("stmt error...")
	}
	defer  db.Close()
    defer stmt.Close()
	 title1 := "1111111121"
	 body1 := "2222222222334422222222222"

	 rs, err := stmt.Exec(title1,body1)
	if err != nil {
		fmt.Println("stmt exec error")
	}
	if id, _ :=rs.LastInsertId(); id >0 {
		fmt.Println("id: %v",id)
	}

	rout.HandleFunc("/", handleFunc)
	rout.HandleFunc("/art/{id:[0-9]+}", articleshowhandler).Methods("GET").Name("article.show....")
	rout.HandleFunc("/art/create", articlepost).Methods("GET").Name("article.showP....")
	rout.HandleFunc("/art",artshowp).Methods("POST").Name("artshowP")
	rout.HandleFunc("/art",artshow).Methods("GET").Name("artshow")
	rout.NotFoundHandler = http.HandlerFunc(n404)
	rout.Use(middleware)
	http.ListenAndServe(":3000", removeTrailingSlash(rout))

}

func artshow(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("none.... data"))
}

func artshowp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if checkerr(w, err) {
		return
	}
	fmt.Fprintf(w, "POST PostForm: %v <br>", r.PostFormValue("body"))
	fmt.Fprintf(w, "POST Form: %v <br>", r.Form.Get("body"))
}

func checkerr(w http.ResponseWriter, err error) bool {
	if err != nil {
		fmt.Fprintf(w, "請提供正確資料")
		return true
	}
	return false
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

type ArticlesFormData struct {
	Title  string
	Body   string
	URL    interface{}
	Errors interface{}
}

func articlepost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println("r.ParseForm() error")
	}
	data := ArticlesFormData{
		Title:  r.PostFormValue("title"),
		Body:   r.PostFormValue("body"),
		URL:    "/art",
		Errors: nil,
	}

	htmla, err := template.ParseFiles("create.gohtml")
	if err != nil {
       fmt.Println("error001")
	}
	if htmla == nil {
		fmt.Println("htmla .... nil")
	}
    fmt.Println(htmla)
	fmt.Println("htmla")
	w.Write([]byte("post..........."))
	err = htmla.Execute(w,data)
	if err != nil {

	}
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
