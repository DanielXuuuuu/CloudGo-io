package service

import (
	"net/http"
)

func NotImplemented (rw http.ResponseWriter, r *http.Request) { 
	http.Error(rw, "501 Not Implement", http.StatusNotImplemented)
}

func NotImplementedHandler() http.HandlerFunc { 
	return http.HandlerFunc(NotImplemented)
 }