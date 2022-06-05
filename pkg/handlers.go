package groupie

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
		return
	}
	Templates.ExecuteTemplate(w, "mainPage.html", SearchArtist)
}
