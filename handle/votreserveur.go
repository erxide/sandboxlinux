package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
)

func VotreServeur(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/votreserveur.html"))
	pageco := template.Must(template.ParseFiles("./templates/votreserveurco.html"))
	pagecoavecserv := template.Must(template.ParseFiles("./templates/votreserveurcoavecserv.html"))
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pseudo, ok := session.Values["pseudo"].(string)
	utilisateurs := forum.Utilisateurs{
		Pseudo: pseudo,
	}
	envoie := forum.Envoie{
		User: utilisateurs,
	}
	var serveurExiste int
	err = forum.Bd.QueryRow("SELECT COUNT(*) FROM McServeur WHERE proprio = ?", pseudo).Scan(&serveurExiste)
	if err != nil {
		fmt.Println(err)
	}
	if serveurExiste > 0 {
		// Si l'utilisateur a déjà un serveur, rediriger vers la page de gestion de serveur
		pagecoavecserv.Execute(w, envoie)
	} else {
		if ok {
			pageco.Execute(w, envoie)
		} else {
			page.Execute(w, r)
		}
	}
}
