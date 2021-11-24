package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Keep Calm, I'm fixing", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "I'm learning golang web",
		"content": "I'm learning golang web",
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Keep Calm, I'm fixing", http.StatusInternalServerError)
		return
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Arief"))
}

func AriefHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Arief, belajar Go"))
}