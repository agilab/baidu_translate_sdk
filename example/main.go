package main

import (
	"flag"
	"log"

	bts "github.com/agilab/baidu_translate_sdk"
)

var (
	appid  = flag.String("appid", "", "appid")
	appkey = flag.String("appkey", "", "appkey")
	query  = flag.String("query", "", "query")
)

func main() {
	flag.Parse()
	bt := bts.CreateBaiduTranslator(*appid, *appkey)
	result, err := bt.Translate(*query)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	log.Printf("%s", result)
}
