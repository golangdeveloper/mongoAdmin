package main

import (
	"io/ioutil"
	"net/http"

	"github.com/golangdeveloper/moetemplate"
)

func handle_index(w http.ResponseWriter, r *http.Request) {
	var mainpage string
	buff, err := ioutil.ReadFile(root + "/views/" + "index.html")
	if err == nil {
		mainpage = string(buff)
	}
	var hp map[string]string
	hp["header"] = "<p>nihao</>"

	var cs map[string]string
	cs["title"] = "你好"

	var out string
	out = moetemplate.Render(mainpage, hp, cs)
	w.Write([]byte(out))
}
