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
	http.ListenAndServe(":8080", nil)
	//getRequest()
}

func getRequest() {

	resp, err := http.Get("http://localhost:3000/")
	checkError(err)
	
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	
	fmt.Println(string(bytes))

}

/*func postRequest() {

	//resp, err := http.Post("http://localhost:3000/", "POST", &buf)
	//resp, err := http.NewRequest("http://localhost:3000/")
	checkError(err)
	
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	
	fmt.Println(string(bytes))

}*/

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
