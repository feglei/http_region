package main

import (
	"syscall"
	"log"
	"net/http"
	"fmt"
	"./models"
	"encoding/json"
)



func FindRegion(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	ip := r.Form.Get("ip")
	region := models.FindRegionModel( ip )

	body, err := json.Marshal(region)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, string(body))
}



func main() {

	// 修改文件数
	ulimit()

	// 初始化数据
	models.InitRegionModel()
	fmt.Println( "初始化数据完毕!" )

	// 初始化 路由
	http.HandleFunc("/r", FindRegion )
	fmt.Println( "初始化路由完毕!" )

	fmt.Println( "启动服务器 23910端口" )
	err := http.ListenAndServe(":23910", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func ulimit() {

	var rlimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		log.Panic("can't modify ulimit", err)
	}
	rlimit.Cur = 655350
	rlimit.Max = 655350
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		log.Panic("can't modify ulimit", err)
	}
}