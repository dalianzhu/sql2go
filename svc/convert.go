package svc

import (
	"go/format"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dalianzhu/myjson/v2"
	"github.com/dalianzhu/sql2go/xormtools"
)

// XormConvert ...
func XormConvert(w http.ResponseWriter, r *http.Request) {
	log.Printf("XormConvert run")
	err := r.ParseForm()
	if err != nil {
		log.Printf("XormConvert ParseForm error:%v", err)
		ErrJson(w, 1, err.Error())
		return
	}
	ret, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("XormConvert body is empty:%v", err)
		ErrJson(w, 1, err.Error())
		return
	}
	js := myjson.NewJson(ret)
	sql := js.Get("sql")
	if sql.IsErrOrNil() || sql.String() == "" {
		ErrJson(w, 1, "sql is empty")
		return
	}

	engineName := js.Get("engineName")
	if engineName.IsErrOrNil() || engineName.String() == "" {
		ErrJson(w, 1, "engineName is empty")
		return
	}

	parser := xormtools.NewParser(new(xormtools.XormConverter))

	xormRet, err := parser.Parse(sql.String(), "model", engineName.String())
	// log.Printf("ret:%v, err:%v", ret, err)
	if err != nil {
		ErrJson(w, 1, err.Error())
		return
	}

	fmtRet, err := format.Source([]byte(xormRet))
	if err != nil {
		ErrJson(w, 1, err.Error())
		return
	}
	rsp := myjson.NewJson("{}")
	rsp.Set("result", string(fmtRet))
	Json(w, rsp.Bytes())
}

func ErrJson(w http.ResponseWriter, errCode int, errStr string) error {
	w.Header().Set("Content-Type", "application/json")
	js := myjson.NewJson("{}")
	js.Set("err", errCode)
	js.Set("err_msg", errStr)
	_, err := w.Write(js.Bytes())
	return err
}

func Json(w http.ResponseWriter, jsonBytes []byte) error {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(jsonBytes)
	return err
}
