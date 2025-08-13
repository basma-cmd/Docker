package routes

import (
	"net/http"
	"os"

	"artweb/helpers"
	"artweb/model"
)

func StaticsRouter(w http.ResponseWriter, r *http.Request) {
	info, err := os.Stat(r.URL.Path[1:])
	if err != nil || info.IsDir() {
		helpers.ParseAndExecute("templates/error.html", model.Data{Error: "403 Status Forbidden"}, http.StatusForbidden, w, false)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
