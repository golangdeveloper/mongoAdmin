package main

import (
	"log"
	"net/http"
	"os"
)

var (
	root    = ""
	viewDir = "/views"
)

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix,
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			file := root + r.URL.Path
			log.Println(file)
			http.ServeFile(w, r, file)
		})
}
func main() {
	//检查根目录
	root, _ = os.Getwd()
	var mux = http.NewServeMux()
	staticDirHandler(mux, "/assets/", root+"/assets", 0)
	mux.HandleFunc("/", router)

	err := http.ListenAndServe(":80", mux)
	log.Println("http.ListenAndServe(:80)")
	if err != nil {
		log.Fatal("http.ListenAndServe:", err.Error())
	}
}
