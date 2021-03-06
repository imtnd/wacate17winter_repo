package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type price_inc_tax struct {
	Price int32  `json:"price"`
	Tax   int32  `json:"tax"`
	Err   string `json:"error"`
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var p int
	//パラメータを取得
	queryValues := r.URL.Query()
	p,_ = strconv.Atoi(queryValues.Get("price"))
	// パラメータの値が0のときはエラー
	if p == 0 {
		e := price_inc_tax{Err: "値段をパラメータに入れてね"}
		json.NewEncoder(w).Encode(e)
		return
	}
	var full_p int32
	var tax int32
	full_p = (int32) (p * 108 / 100)
	tax = (int32) (p * 8 / 100)
	price_include_tax := price_inc_tax{Price: full_p, Tax: tax}

	json.NewEncoder(w).Encode(price_include_tax)
}

func main() {
	router := httprouter.New()
	//indexにアクセスさせるときはindex関数の処理を行う
	router.GET("/index", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
