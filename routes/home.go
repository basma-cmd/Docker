package routes

import (
	"net/http"

	"artweb/helpers"
	"artweb/model"
)

// handle the "/" route
func ArtWebHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		helpers.ParseAndExecute("templates/error.html", model.Data{Error: "404 Page Not Found"}, http.StatusNotFound, w, false)
		return
	}

	if r.Method != http.MethodGet {
		helpers.ParseAndExecute("templates/error.html", model.Data{Error: "405 Method Not Allowed"}, http.StatusMethodNotAllowed, w, false)
		return
	}

	helpers.ParseAndExecute("templates/index.html", model.Data{}, 200, w, false)
}
