package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) showChampion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.NotFound(w, r)
		return
	}
	name = strings.Title(strings.ToLower(name))
	fmt.Println(name)
	champion, err := app.riotClient.DataDragon.GetChampion(name)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(champion.Lore))
}
