package helpers

import (
	"log"
	"net/http"
	"text/template"

	"artweb/model"
)

func ErrorExecute(data model.Data, w http.ResponseWriter, errStr string, status int) {
	tmp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	errExec := tmp.Execute(w, data)
	if errExec != nil {
		log.Println(errExec)
	}
}
