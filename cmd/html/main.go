// https://www-alexedwards-net.translate.goog/blog/form-validation-and-processing?_x_tr_sl=auto&_x_tr_tl=ru&_x_tr_hl=ru
/*
	Проверка формы и обработка в Go
*/
package main

import (
	"github.com/bmizerany/pat"
	"html/template"
	"log"
	"net/http"
)

const (
	homeHTML         = "cmd/html/templates/home.html"
	confirmationHTML = "cmd/html/templates/confirmation.html"
)

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(home))
	mux.Post("/", http.HandlerFunc(send))
	mux.Get("/confirmation", http.HandlerFunc(confirmation))

	log.Println("Listening... :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, homeHTML, nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	// Step 1: Validate form
	msg := &Message{
		Email:   r.PostFormValue("email"),
		Content: r.PostFormValue("content"),
	}

	if msg.Validate() == false {
		render(w, homeHTML, msg)
		return
	}

	// Step 2: Action
	if err := msg.Action(); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	// Step 3: Redirect to confirmation page
	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
	//http.Redirect(w, r, "/", http.StatusSeeOther)
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, confirmationHTML, nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
