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
	err:= http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
