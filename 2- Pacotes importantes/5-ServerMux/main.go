package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Olá mundo"))
	// })

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Meu blog"})
	http.ListenAndServe(":80", mux)

	// mux2 := http.NewServeMux()
	// mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Olá Michel"))
	// })
	// http.ListenAndServe(":8080", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundão!"))

}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))

}
