package main

import (
	"net/http"
	"moetemplate"
)

func handle_index(w http.ResponseWriter, r *http.Request){
	var out=moetemplate.Ok() 
	w.Write([]byte(out))
}