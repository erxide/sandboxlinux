package handle

import (
	"html/template"
	"net/http"
)

func VotreServeur(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/votreserveur.html"))
	page.Execute(w, r)
}
