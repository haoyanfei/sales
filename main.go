package main

import (
	"fmt"
	"strconv"
	"time"
	_ "web1/routers"

	"github.com/astaxie/beego"
)

func dataKV(data map[int]string, k string) string {
	intK, _ := strconv.Atoi(k)
	return data[intK]
}

func convertT(timestamp string) string {
	t, v := strconv.ParseInt(timestamp, 10, 0)
	fmt.Println(v)
	tm := time.Unix(t, 0)
	fmt.Println("----------", t)
	str := tm.Format("2006-01-02 15:04:05")
	return str
}
func main() {
	beego.AddFuncMap("convertt", convertT)
	beego.AddFuncMap("dataKV", dataKV)
	beego.Run()
}
