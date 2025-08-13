package routes

import (
	"net/http"
	"strings"

	"artweb/app"
	"artweb/helpers"
	"artweb/model"
)

func ArtStyleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ParseAndExecute("templates/error.html", model.Data{Error: "405 Method Not Allowed"}, http.StatusMethodNotAllowed, w, false)
		return
	}

	input := r.FormValue("input")
	banner := r.FormValue("banner")

	ArtStyle, err := app.AsciiArt(input, banner)
	if err != nil {
		if strings.HasPrefix(err.Error(), "The string includes characters outside the ASCII range...") {
			helpers.ParseAndExecute("templates/index.html", model.Data{Error: err.Error()}, http.StatusBadRequest, w, true)
		} else if err.Error() == "input error" {
			helpers.ParseAndExecute("templates/error.html", model.Data{Error: "400 Status Bad Request"}, http.StatusBadRequest, w, false)
		} else {
			helpers.ParseAndExecute("templates/error.html", model.Data{Error: "500 Status Internal Server"}, http.StatusInternalServerError, w, false)
		}
		return
	}
	helpers.ParseAndExecute("templates/ascii-art.html", model.Data{Output: ArtStyle}, 200, w, false)
}
