
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/sayhello/{name}", Sayhello())
	log.Fatal(http.ListenAndServe(":8009", router))
}

func Sayhello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := http.StatusBadRequest // if req is not GET
		if r.Method == "GET" {
			code = http.StatusOK
			vars := mux.Vars(r)
			name := vars["name"]

			greet := fmt.Sprintf("Hello %s \n", name)
			w.Write([]byte(greet))
		} else {
			w.WriteHeader(code)
		}
	}
}
