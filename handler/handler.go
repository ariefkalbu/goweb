package handler

import (
	"goweb/model"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "__layout.html"))

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

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// data := map[string]interface{}{
	// 	"content": idNumb,
	// 	"title":   "kenapa",
	// }
	// data := model.Product{Id: 3, Name: "Ayla", Price: 210000, Stock: 6}
	data := []model.Product{
		{Id: 1, Name: "Ayla", Price: 210000, Stock: 5},
		{Id: 2, Name: "Pajero", Price: 420000, Stock: 3},
		{Id: 3, Name: "Xenia", Price: 100000, Stock: 1},
	}

	log.Println(data)

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "__layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "GET":
		w.Write([]byte("ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "__layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)

}

func Process(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		nama := r.Form.Get("nama")
		nomor := r.Form.Get("nomor")

		data := map[string]interface{}{
			"nama":  nama,
			"nomor": nomor,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "detail.html"), path.Join("views", "__layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)

}
