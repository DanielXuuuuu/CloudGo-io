package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID             string `json:"id"`
			Content string `json:"content"`
		}{ID: "17343130", Content: "Daniel Xu"})
	}
}