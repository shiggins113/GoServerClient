package server

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"stephen/server/pkg/config"
	_ "github.com/go-sql-driver/mysql"
    "encoding/json"
    "database/sql"
    "strconv"
)

type CertData struct {
    CertID      int    `json:"certID"`
    Cert        string `json:"cert"`
    BeforeDate  string `json:"beforeDate"`
    AfterDate   string `json:"afterDate"`
}

func certDelete(w http.ResponseWriter, r *http.Request){

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    certID := r.URL.Query().Get("certID")
   // certID := r.URL.Query().Get("certID")
    fmt.Println(certID)
    certID1, err := strconv.Atoi(certID)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}
    fmt.Println(certID)
    deleteCert(gcfg, certID1)

}

func certData(w http.ResponseWriter, r *http.Request){

    fmt.Println("Select * from certs;")
    rows, err := certSelect(gcfg)

    var results []CertData
    
    for rows.Next() {
               var certData CertData
                err := rows.Scan(&certData.CertID, &certData.Cert, &certData.BeforeDate, &certData.AfterDate)
                if err != nil {
                    fmt.Println("Error scanning row:", err)
                    return
                }
                results = append(results, certData)
            }
    
        if err := rows.Err(); err != nil {
            fmt.Println("Error iterating over result rows:", err)
            return
        }
    
            // Convert the results to JSON
            jsonData, err := json.Marshal(results)
            if err != nil {
                fmt.Println("Error marshaling JSON:", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
            }
        
            // Set the response headers and write the JSON data
            w.Header().Set("Content-Type", "application/json")
            w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            w.Write(jsonData)
    
}


func cert(w http.ResponseWriter, r *http.Request){
    
    file, header, err := r.FormFile("myFile")

    if err != nil {
        http.Error(w, "Error retrieving the file", http.StatusBadRequest)
        return
    }

    defer file.Close()

    // Check if the file is empty
    if header.Size <= 0 {
        http.Error(w, "Empty file", http.StatusBadRequest)
        return
    }


    size := header.Size   
    b1 := make([]byte, size)
    n1, err := file.Read(b1)
    value := string(b1[:n1])

    var pemData = []byte(value)

   for block, rest := pem.Decode(pemData); block !=nil; block, rest = pem.Decode(rest){ 

        cert, err := x509.ParseCertificates(block.Bytes)
        if err != nil {
            http.Error(w, "Failed to parse certificate", http.StatusInternalServerError)
            return
        }
        
       certIssuer := cert[0].Issuer.String()
       beforeDate := cert[0].NotBefore.String()
       afterDate := cert[0].NotAfter.String()

       certInsert(certIssuer, beforeDate, afterDate, gcfg)

    }

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("File uploaded successfully"))

}

func certInsert(cert string, beforeDate string, afterDate string, cfg *config.Config) {

    query := fmt.Sprintf("INSERT INTO certs(cert,beforeDate,afterDate) VALUES ( '%s', '%s', '%s' )", cert, beforeDate, afterDate)

    _, err := cfg.DB.Exec(query)
	
    if err != nil {
        panic(err.Error())
    }

}

func certSelect(cfg *config.Config) (*sql.Rows, error) {

    query := "Select * from certs;"
    rows, err := cfg.DB.Query(query)

    if err != nil {
        fmt.Println("Error executing query:", err)
        return nil, err
    }

    return rows, nil
}



func deleteCert(cfg *config.Config, certID int) {

    

    query := fmt.Sprintf("Delete from certs where certID = '%d'", certID)
        _, err := cfg.DB.Exec(query)
    
        if err != nil {
            fmt.Println("Error executing query:", err)
        }    
    
    }

