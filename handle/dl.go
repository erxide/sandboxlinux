package handle

import (
	"forum/forum"
	"net/http"
)

func DlConfServeurVpn(w http.ResponseWriter, r *http.Request) {
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := session.Values["pseudo"].(string)
	if !ok {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	referer := r.Header.Get("Referer")
	filePath := "./dl/test.txt"
	w.Header().Set("Content-Disposition", "attachment; filename=test.txt")
	http.ServeFile(w, r, filePath)
	http.Redirect(w, r, referer, http.StatusFound)
}
