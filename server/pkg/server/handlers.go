package server

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"stephen/server/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)


func cert(w http.ResponseWriter, r *http.Request){
    
    file, header, err := r.FormFile("myFile")

    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }

    size := header.Size   
    b1 := make([]byte, size)
    n1, err := file.Read(b1)
    value := string(b1[:n1])

    defer file.Close()

    var pemData = []byte(value)

   for block, rest := pem.Decode(pemData); block !=nil; block, rest = pem.Decode(rest){ 
        cert, err := x509.ParseCertificates(block.Bytes)
        if err != nil {
            panic("failed to parse certificate: " + err.Error())
        }
        
       certIssuer := cert[0].Issuer.String()
       beforeDate := cert[0].NotBefore.String()
       afterDate := cert[0].NotAfter.String()

       certInsert(certIssuer, beforeDate, afterDate, gcfg)

    }

}

func certInsert(cert string, beforeDate string, afterDate string, cfg *config.Config) {

    query := fmt.Sprintf("INSERT INTO certs(cert,beforeDate,afterDate) VALUES ( '%s', '%s', '%s' )", cert, beforeDate, afterDate)

    _, err := cfg.DB.Exec(query)
	
    if err != nil {
        panic(err.Error())
    }

}
