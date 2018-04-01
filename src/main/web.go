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
	"time"
)

var db *SimpleDb.MyDb

func main() {
	DB.InitDB()
	http.HandleFunc("/buckInfo", GetBuckInfo)       //设置访问的路由
	http.HandleFunc("/tradeTrend", GetTradeTrend);
	http.HandleFunc("/cashflow", GetCashFlow);
	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func GetCashFlow(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	end := time.Now()
	start := end.Add(-14 * time.Hour * 24) //当前时间加24小时，即明天的这个时间
	var result string
	result = DB.GetCashFlow(start.Format("20060102"), end.Format("20060102"))
	//fmt.Println("result",result);
	writeToClient(w, result)
}
func GetTradeTrend(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	trade := r.Form["trade"]
	end := time.Now()
	start := end.Add(-14 * time.Hour * 24) //当前时间加24小时，即明天的这个时间
	var result string
	result = DB.GetTradeTrend(trade[0], start.Format("20060102"), end.Format("20060102"))
	//fmt.Println("result",result);
	writeToClient(w, result)
}

func GetBuckInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dt := r.Form["dt"]
	//dt := time.Now().Format("200601021504")
	t := r.Form["type"]
	if (len(t) > 0) {

	}
	var result string
	if (len(dt) == 0) {
		result = DB.GetBuckTrendList("", r.Form["asc"][0], r.Form["limit"][0])
	} else {
		result = DB.GetBuckTrendList(dt[0], r.Form["asc"][0], r.Form["limit"][0])
	}
	writeToClient(w, result)
}
func writeToClient(w http.ResponseWriter, result string) {
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