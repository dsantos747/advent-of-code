package handler

import (
	"fmt"
	"net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	var resp []byte = []byte(`HELLO MR WORLD HOW ARE YOU`)
		
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	w.Write(resp)
  	// fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
