package svc

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func XormIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("MainPage run")
	baseTmp := template.New("main")
	path := "./template/main.gtpl"
	logrus.Infof("MainPage path:%v", path)
	baseTmp, err := baseTmp.ParseFiles(path)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error:%v", err)))
		return
	}

	err = baseTmp.ExecuteTemplate(w, "main.gtpl", nil)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error:%v", err)))
		return
	}
}
