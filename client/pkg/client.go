package main

import (
	"net/http"
)

type Page struct {
    Title string
    Body  []byte
}

func main() {

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web.html")
	})

	http.HandleFunc("/cert", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web-cert.html")
	})
//	err:= http.ListenAndServeTLS(":8080", "client.cert.pem", "client.key.pem", nil)
	err:= http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}



func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
