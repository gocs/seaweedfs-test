package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gocs/seaweedfs-test/pkg/store"
	"github.com/gorilla/mux"
)

//go:embed templates/*
var assetData embed.FS

func main() {
	r := mux.NewRouter()

	tmpl, err := template.ParseFS(assetData, "templates/layout.html")
	if err != nil {
		log.Fatal("ParseFS err: ", err)
	}
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mf, _, err := r.FormFile("filepath")
		if err != nil {
			log.Fatal("FormFile err: ", err)
		}
		defer mf.Close()

		assignResp, err := store.Assign("http://seaweedfs:9333")
		if err != nil {
			log.Fatal("Assign err: ", err)
		}

		fidURL := fmt.Sprintf("%s/%s", "http://seaweedfs:8080", assignResp.Fid)
		form := map[string]io.Reader{"filepath": mf}
		_, err = store.Upload(fidURL, form)
		if err != nil {
			log.Fatal("Upload err: ", err)
		}

		log.Println("url:", fidURL)
		http.Redirect(w, r, "/", http.StatusFound)
	}).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
