package service

import (
	"net/http"
	"os"
    "github.com/urfave/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
		Directory: "templates", // Specify what path to load the templates from.
		Extensions: []string{".html"}, 
		IndentJSON: true,
    })

    n := negroni.Classic()
    router := mux.NewRouter()

    initRoutes(router, formatter)

    n.UseHandler(router)
    return n
}

func initRoutes(router *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil{
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	router.HandleFunc("/", homePageHandler(formatter));
	router.HandleFunc("/login", loginHandler(formatter))
	router.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets"))))
	router.HandleFunc("/unknown", NotImplementedHandler()).Methods("GET")
}
