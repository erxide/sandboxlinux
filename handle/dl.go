package handle

import "net/http"

func DlConfServeurVpn(w http.ResponseWriter, r *http.Request) {
	referer := r.Header.Get("Referer")
	filePath := "./dl/test.txt"
	w.Header().Set("Content-Disposition", "attachment; filename=test.txt")
	http.ServeFile(w, r, filePath)
	http.Redirect(w, r, referer, http.StatusFound)
}
