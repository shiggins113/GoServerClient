package main

import (
	"fmt"
	"log"
	"net/http"
    "os"
    "io"
   // "crypto/tls"
)

func main() {
//comment
   // cert, err1 := tls.LoadX509KeyPair("server.cert.pem", "server.key.pem")
   // if err1 != nil {
	//	log.Fatal(err1)
//	}

    //cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
    mux.HandleFunc("/upload", uploadFile)
	//err:= http.ListenAndServe(":3000", mux)
    err:= http.ListenAndServeTLS(":8443","server.cert.pem", "server.key.pem", mux)
	log.Fatal(err)
	
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.\
  //  r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()

    err = os.MkdirAll("./uploads", os.ModePerm)
  //  dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename)))
    dst, err := os.Create(fmt.Sprintf("./uploads/%s", handler.Filename))

    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    defer dst.Close()

    _, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    fmt.Fprintf(w, "Successfully Uploaded File\n")
}


func index(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    // Common code for all requests can go here...

    switch r.Method {
    case http.MethodGet:
	fmt.Fprintf(w, "This is a GET")

    case http.MethodPost:
	fmt.Fprintf(w, "This is a POST")

    case http.MethodOptions:
        w.Header().Set("Allow", "GET, POST, OPTIONS")
        w.WriteHeader(http.StatusNoContent)

    default:
        w.Header().Set("Allow", "GET, POST, OPTIONS")
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
	}