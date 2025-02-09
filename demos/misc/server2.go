//go:build ignore

package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/cep/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}
