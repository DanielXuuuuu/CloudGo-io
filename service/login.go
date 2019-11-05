package service

import (
	"fmt"
	"net/http"
	"text/template"
	"github.com/unrolled/render"
)

func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("method:", r.Method) //获取请求的方法
		if r.Method == "GET" {
			t := template.Must(template.ParseFiles("templates/loginGET.html"))
			t.Execute(w, nil)
		} else {
			r.ParseForm()
			username :=  r.Form["username"][0]
			password :=  r.Form["password"][0]

			t := template.Must(template.ParseFiles("templates/loginPOST.html"))
			t.Execute(w, map[string]string{
				"username": username,
				"password": password,
			})
		}
	}
}