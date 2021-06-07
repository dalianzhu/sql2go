package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dalianzhu/sql2go/svc"
	"github.com/sirupsen/logrus"
)

func main() {
	runHttp()
}
func runHttp() {
	mux := http.NewServeMux()
	server :=
		http.Server{
			Addr:         fmt.Sprintf(":%d", 8080),
			Handler:      mux,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 5 * time.Second,
		}

	// 开始添加路由
	mux.HandleFunc("/sql2go/xorm", svc.XormIndex)
	mux.HandleFunc("/sql2go/xormConvert", svc.XormConvert)

	logrus.Infof("run http:%v", 8080)
	logrus.Infof("end:%v", server.ListenAndServe())
}
