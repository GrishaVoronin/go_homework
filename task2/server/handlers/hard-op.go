package handlers

import (
	"math/rand"
	"net/http"
	"time"
)

func HardOp(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(11)+10) * time.Second)
	w.WriteHeader(http.StatusOK)
}
