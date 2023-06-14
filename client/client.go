package main

import (
	"fmt"
	"io/ioutil"
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
	err:= http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}

}

func getRequest() {

//	resp, err := http.Get("https://localhost:8443/")
	resp, err := http.Get("http://go-server.cfapps-01.slot-21.pez.vmware.com/")	
	checkError(err)
	
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	
	fmt.Println(string(bytes))

}



func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
