package helpers

import (
	"net/http"
	"strings"
	"text/template"

	"artweb/model"
)

func ParseAndExecute(temp string, data model.Data, status int, w http.ResponseWriter, errNotInASCII bool) {
	// Handle other error messages via status code and content of data.Error
	if !strings.HasPrefix(data.Error, "The string includes characters outside the ASCII range...") {
		switch status {
		case 400:
			ErrorExecute(data, w, "400 Status Bad Request", http.StatusBadRequest)
			return
		case 404:
			ErrorExecute(data, w, "404 Page Not Found", http.StatusNotFound)
			return
		case 500:
			ErrorExecute(data, w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	if errNotInASCII {
		data = model.Data{
			Error: "The string contain characters outside the ASCII range...",
		}
	}

	// Render the template normally
	tmp, err := template.ParseFiles(temp)
	if err != nil {
		ErrorExecute(data, w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	errExec := tmp.Execute(w, data)
	if errExec != nil {
		ErrorExecute(data, w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
