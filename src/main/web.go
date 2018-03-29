package main

import (
	"fmt"
	//"html/template"
	"log"
	"net/http"
	//"strings"
	_ "github.com/lib/pq"
	"../SimpleDb"
	"../DB"
	//"time"
)

var db *SimpleDb.MyDb

func main() {
	DB.InitDB()
	http.HandleFunc("/buckInfo", GetBuckInfo)       //设置访问的路由
	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GetBuckInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("coming")
	r.ParseForm()
	dt := r.Form["dt"]
	//dt := time.Now().Format("200601021504")
	var result string
	if (len(dt) == 0) {
		result = DB.GetBuckTrendList("",r.Form["asc"][0],r.Form["limit"][0])
	} else {
		result = DB.GetBuckTrendList(dt[0],r.Form["asc"][0],r.Form["limit"][0])
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("charset", "utf-8")
	fmt.Fprintf(w, result) //这个写入到 w 的是输出到客户端的
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}