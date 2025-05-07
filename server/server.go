package server

import (
	"fmt"
	"net/http"
)

func Serve() {
	http.HandleFunc("/api/search", Search)
	http.ListenAndServe(":6942", nil)
}

func Search(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hoyyyaa")
}
