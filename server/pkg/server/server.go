package server

import (
	"log"
	"net/http"
	"stephen/server/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

var gcfg *config.Config 

func Start(cfg *config.Config) {

	gcfg = cfg

	mux := http.NewServeMux()
	mux.HandleFunc("/cert", cert)
	mux.HandleFunc("/certData", certData)
	mux.HandleFunc("/certDelete", certDelete)
	err:= http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
