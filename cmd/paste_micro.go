package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/FBreuer2/pastebin-micro/db"
	"github.com/FBreuer2/pastebin-micro/entity"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

func main() {
	box := packr.NewBox("../assets")
	router := mux.NewRouter()

	pasteDB, _ := db.NewInMemoryPasteDB()

	paste, _ := entity.CreatePaste("anonymous", "int main() { return 0;}", "", "")
	pasteDB.SavePaste(paste)

	serveIndex := func(w http.ResponseWriter, _ *http.Request) {
		resp, err := box.FindString("index.html")
		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(w, resp)
	}

	pasteTemplateString, err := box.FindString("paste.html")
	if err != nil {
		log.Fatal(err)
	}

	pasteTemplate, err := template.New("Paste").Parse(pasteTemplateString)

	servePaste := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if pasteID := vars["id"]; pasteID == "" {
			log.Println("ID not found")
			http.Redirect(w, r, "/", 301)
			return
		}

		pasteID, err := strconv.ParseUint(vars["id"], 10, 64)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		paste, err := pasteDB.GetPaste(pasteID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = pasteTemplate.Execute(w, paste)
	}

	postPaste := func(w http.ResponseWriter, r *http.Request) {
		content := r.FormValue("paste_content")
		password := r.FormValue("password")

		newPaste, err := entity.CreatePaste("anonymous", content, "", password)

		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", 301)
			return
		}
		paste, err := pasteDB.SavePaste(newPaste)

		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", 301)
			return
		}

		http.Redirect(w, r, "/pastes/"+strconv.FormatUint(paste.ID, 10), 301)
	}

	router.HandleFunc("/", serveIndex)
	router.HandleFunc("/index.html", serveIndex)
	router.HandleFunc("/pastes/{id:[0-9]+}", servePaste)
	router.HandleFunc("/pastes", postPaste).Methods("POST")

	http.ListenAndServe(":3000", router)
}
