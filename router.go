package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golangframework/File"
	"github.com/golangframework/JSON"
	"github.com/golangframework/Object"
	"github.com/golangframework/httpmongo"
	"github.com/golangframework/moeregexp"
	"github.com/golangframework/moetemplate"
)

var mongodb = ""

func getmongo() string {
	if mongodb == "" {
		conbs, err := File.ReadAllBytes(root + "/config.json")
		if err != nil {
			log.Print("无法读取配置文件")
			panic("无法读取配置文件")
		} else {
			var config JSON.JSON
			json.Unmarshal(conbs, &config)
			log.Print(string(conbs))
			mongodb = Object.Tostring(config["mongodb"])
			return mongodb
		}
	} else {
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
			fallthrough
		case "/admin":
			handle_admin(w, r)
		default:
			w.Write([]byte("404"))
		}
	}
}

var htmlparts = map[string]string{}

func Gethtmlparts() map[string]string {
	if htmlparts != nil && len(htmlparts) > 1 {
		return htmlparts
	} else {
		htmlparts = moetemplate.LoadPartFile(root + "/views/")
		return htmlparts
	}
}
func handle_admin(w http.ResponseWriter, r *http.Request) {
	mainpage, _ := File.ReadAllText(root + "/views/admin.html")
	var cs map[string]string = map[string]string{}
	cs["title"] = title + ""
	var out = moetemplate.Render(mainpage, Gethtmlparts(), cs)
	w.Write([]byte(out))
}
