package sdk

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// params := r.URL.Query()
	// fmt.Println(params["a"])
	// fmt.Fprintf(w, "<html>Hello World!<p>aaa</p><html>")
	http.Redirect(w, r, "/help", http.StatusFound)
}

func Help(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./help.html")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "<html>Help Page Not Found<html>")
	} else {
		t.Execute(w, nil)
	}
}
