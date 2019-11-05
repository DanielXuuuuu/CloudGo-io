package service

import (
	"net/http"
	"text/template"
	"github.com/unrolled/render"
)

func userInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username :=  r.Form["username"][0]
		password :=  r.Form["password"][0]
		email := r.Form["email"][0]

		t := template.Must(template.ParseFiles("templates/userInfo.html"))
		t.Execute(w, map[string]string{
			"username": username,
			"email": email,
			"password": password,
		})
	}
}