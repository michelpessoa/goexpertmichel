package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Com função anônima"))
	})

	http.HandleFunc("/2", BuscaCEP)
	http.ListenAndServe(":8080", nil)

}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo Michel2!"))
}
