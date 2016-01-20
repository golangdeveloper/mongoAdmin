package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golangframework/File"
	"github.com/golangframework/JSON"
	"github.com/golangframework/Object"
	"github.com/golangframework/htmlparts"
	"github.com/golangframework/httpmongo"
	"github.com/golangframework/moeregexp"
)

var mongodb = ""
var port = ""

func getmongo() string {
	if mongodb == "" {
		conbs, err := File.ReadAllBytes(root + "/config.json")
		if err != nil {
			log.Print("无法读取配置文件")
			panic("无法读取配置文件")
		} else {
			var config JSON.JSON = JSON.JSON{}
			json.Unmarshal(conbs, &config)
			log.Print(config)
			mongodb = Object.Tostring(config["mongodb"])
			port = Object.Tostring(config["port"])
			return mongodb
		}
	} else {
		log.Print(mongodb)
		return mongodb
	}
}

var (
	title = "mongoAdmin"
)

func router(w http.ResponseWriter, r *http.Request) {
	urlpath := r.URL.Path

	if moeregexp.IsMatch(httpmongo.Mongo_path, urlpath) {
		httpmongo.Httphandler_mongo(getmongo(), w, r)
	} else {
		switch urlpath[0:] {
		case "/":
			handle_admin(w, r)
		default:
			w.Write([]byte("404"))
		}
	}
}

var hps = map[string]string{}

func handle_admin(w http.ResponseWriter, r *http.Request) {
	hps = htmlparts.LoadPartFile(root + "/views/")
	mainpage, _ := File.ReadAllText(root + "/views/admin.html")
	var cs map[string]string = map[string]string{}
	cs["title"] = title + ""
	var out = htmlparts.Render(mainpage, hps, cs)
	w.Write([]byte(out))
}
