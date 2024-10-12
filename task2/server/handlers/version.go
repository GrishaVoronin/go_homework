package handlers

import "net/http"

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("v1.0.0"))
}
