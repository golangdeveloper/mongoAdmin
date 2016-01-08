package main

import (
	"net/http"

	"github.com/golangframework/file"
	"github.com/golangframework/moetemplate"
)

func handle_index(w http.ResponseWriter, r *http.Request) {
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
