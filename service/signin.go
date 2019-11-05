package service

import (
	"net/http"
	"github.com/unrolled/render"
)

func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		formatter.HTML(w, http.StatusOK, "signin", nil)
	}
}