package DB

import (
	"encoding/json"
	"fmt"
	"../SimpleDb"
	_ "github.com/lib/pq"
)

type TrendInfo struct {
	Dt         string
	Ord        int
	Code       string
	Trade_name string
	Up_per     float32
	Core_m     float32
	Core_p     float32
	Huge_m     float32
	Huge_p     float32
	Big_m      float32
	Big_p      float32
	Mid_m      float32
	Mid_p      float32
	Small_m    float32
	Small_p    float32
	Bname      string
	Bcode      string
	Sh         float32
	Sz         float32
}

func InitDB() {
	db, _ = SimpleDb.NewDb("postgres", "port=5432 user=ott password=ott dbname=buck sslmode=disable")
}

var db *SimpleDb.MyDb = nil

func GetBuckTrendList(dt string, asc string, limit string) string {
	fmt.Println("GetBuckTrendList dt", dt, "asc", asc, "limit", limit)
	var data []SimpleDb.DataRow
	var sql string
	if (dt == "") {
		sql = "select * from buck_trend where dt=(select max(dt) from buck_trend) order by core_m " + asc + " limit " + limit
	} else {
		sql = "select * from buck_trend where dt='" + dt + "' order by core_m " + asc + " limit " + limit
	}
	fmt.Println(sql);
	data, _ = db.QueryDataRows(sql)

	len := len(data)
	var arr [100]*TrendInfo
	for i := 0; i < len; i++ {
		t := new(TrendInfo)
		data[i].GetValue("dt", &t.Dt)
		data[i].GetValue("ord", &t.Ord)
		data[i].GetValue("code", &t.Code)
		data[i].GetValue("trade_name", &t.Trade_name)
		data[i].GetValue("up_per", &t.Up_per)
		data[i].GetValue("core_m", &t.Core_m)
		data[i].GetValue("core_p", &t.Core_p)
		data[i].GetValue("huge_m", &t.Huge_m)
		data[i].GetValue("huge_p", &t.Huge_p)
		data[i].GetValue("big_m", &t.Big_m)
		data[i].GetValue("big_p", &t.Big_p)
		data[i].GetValue("mid_m", &t.Mid_m)
		data[i].GetValue("mid_p", &t.Mid_p)
		data[i].GetValue("small_m", &t.Small_m)
		data[i].GetValue("small_p", &t.Small_p)
		data[i].GetValue("bname", &t.Bname)
		data[i].GetValue("bcode", &t.Bcode)
		data[i].GetValue("sh", &t.Sh)
		data[i].GetValue("sz", &t.Sz)
		arr[i] = t
	}
	//list := list.New()
	//list.PushBack(t)
	//for v := list.Front(); v != nil; v = v.Next() {
	//	log(v.Value.(*TrendInfo))
	//}
	//for i, t := range arr {
	//	if (i >= len) {
	//		break;
	//	}
	//	log(t)
	//}
	s, _ := json.Marshal(arr[0:len])
	return string(s)
}

func GetTradeTrend(trade string, start string, end string) string {
	fmt.Println("GetTradeTrend trade", trade, "start", start, "end", end)
	var data []SimpleDb.DataRow
	sql := "with t as(" +
		"select * from buck_trend where trade_name = '" + trade + "' and dt >= '" + start + "' and dt<='" + end + "'" +
		")" +
		"select * from t where dt in (" +
		"  select max(dt) from t group by substr(dt, 1,8)" +
		") order by  dt"
	fmt.Println(sql);
	data, _ = db.QueryDataRows(sql)

	len := len(data)
	var arr [100]*TrendInfo
	for i := 0; i < len; i++ {
		t := new(TrendInfo)
		data[i].GetValue("dt", &t.Dt)
		data[i].GetValue("ord", &t.Ord)
		data[i].GetValue("code", &t.Code)
		data[i].GetValue("trade_name", &t.Trade_name)
		data[i].GetValue("up_per", &t.Up_per)
		data[i].GetValue("core_m", &t.Core_m)
		data[i].GetValue("core_p", &t.Core_p)
		data[i].GetValue("huge_m", &t.Huge_m)
		data[i].GetValue("huge_p", &t.Huge_p)
		data[i].GetValue("big_m", &t.Big_m)
		data[i].GetValue("big_p", &t.Big_p)
		data[i].GetValue("mid_m", &t.Mid_m)
		data[i].GetValue("mid_p", &t.Mid_p)
		data[i].GetValue("small_m", &t.Small_m)
		data[i].GetValue("small_p", &t.Small_p)
		data[i].GetValue("bname", &t.Bname)
		data[i].GetValue("bcode", &t.Bcode)
		data[i].GetValue("sh", &t.Sh)
		data[i].GetValue("sz", &t.Sz)
		arr[i] = t
	}
	s, _ := json.Marshal(arr[0:len])
	return string(s)
}

func GetCashFlow(start string, end string) string {
	fmt.Println("GetCashFlow", "start", start, "end", end)
	var data []SimpleDb.DataRow
	sql := "with t as(" +
		"select * from buck_trend where dt in (" +
		"  select max(dt) from buck_trend group by substr(dt, 1,8)" +
		")" +
		")" +
		"select * from t where (ord >0 and ord <=3) or (ord >=58 and ord <=60) " +
		"and dt >= '" + start + "' and dt<='" + end + "' order by dt,ord"
	fmt.Println(sql);
	data, _ = db.QueryDataRows(sql)

	len := len(data)
	var arr [100]*TrendInfo
	for i := 0; i < len; i++ {
		t := new(TrendInfo)
		data[i].GetValue("dt", &t.Dt)
		data[i].GetValue("ord", &t.Ord)
		data[i].GetValue("code", &t.Code)
		data[i].GetValue("trade_name", &t.Trade_name)
		data[i].GetValue("up_per", &t.Up_per)
		data[i].GetValue("core_m", &t.Core_m)
		data[i].GetValue("core_p", &t.Core_p)
		data[i].GetValue("huge_m", &t.Huge_m)
		data[i].GetValue("huge_p", &t.Huge_p)
		data[i].GetValue("big_m", &t.Big_m)
		data[i].GetValue("big_p", &t.Big_p)
		data[i].GetValue("mid_m", &t.Mid_m)
		data[i].GetValue("mid_p", &t.Mid_p)
		data[i].GetValue("small_m", &t.Small_m)
		data[i].GetValue("small_p", &t.Small_p)
		data[i].GetValue("bname", &t.Bname)
		data[i].GetValue("bcode", &t.Bcode)
		data[i].GetValue("sh", &t.Sh)
		data[i].GetValue("sz", &t.Sz)
		arr[i] = t
	}
	s, _ := json.Marshal(arr[0:len])
	return string(s)
}
func log(t *TrendInfo) {
	fmt.Print("\tDt:", t.Dt)
	fmt.Print("\tOrd:", t.Ord)
	fmt.Print("\tCode:", t.Code)
	fmt.Print("\tTrade_name:", t.Trade_name)
	fmt.Print("\tUp_per:", t.Up_per)
	fmt.Print("\tCore_m:", t.Core_m)
	fmt.Print("\tCore_p:", t.Core_p)
	fmt.Print("\tHuge_m:", t.Huge_m)
	fmt.Print("\tHuge_p:", t.Huge_p)
	fmt.Print("\tBig_m:", t.Big_m)
	fmt.Print("\tBig_p:", t.Big_p)
	fmt.Print("\tMid_m:", t.Mid_m)
	fmt.Print("\tMid_p:", t.Mid_p)
	fmt.Print("\tSmall_m:", t.Small_m)
	fmt.Print("\tSmall_p:", t.Small_p)
	fmt.Print("\tBname:", t.Bname)
	fmt.Print("\tBcode:", t.Bcode)
	fmt.Print("\tSh:", t.Sh)
	fmt.Println("\tSz:", t.Sz)
}
