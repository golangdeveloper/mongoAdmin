package main

import (
	"net/http"

	"github.com/golangframework/file"
	"github.com/golangframework/httpmongo"
	"github.com/golangframework/moeregexp"
	"github.com/golangframework/moetemplate"
)

func router(w http.ResponseWriter, r *http.Request) {
	urlpath := r.URL.Path
	if moeregexp.IsMatch("^/admin$", urlpath) {
		handle_admin(w, r)
	} else {
		if moeregexp.IsMatch(httpmongo.Mongo_path, urlpath) {
			httpmongo.Httphandler_mongo("127.0.0.1:27017", w, r)
		} else {
			out := ""
			w.Write([]byte(out))
		}
	}

}
func handle_admin(w http.ResponseWriter, r *http.Request) {
	var mainpage string = file.ReadAllText(root + "/views/index.html")

	var hp map[string]string
	hp = make(map[string]string)
	hp["header"] = file.ReadAllText(root + "/views/header.part")
	hp["footer"] = file.ReadAllText(root + "/views/footer.part")

	var cs map[string]string
	cs = make(map[string]string)
	cs["title"] = "MongoAdmin"
	var out string

	out = moetemplate.Render(mainpage, hp, cs)

	w.Write([]byte(out))
}
