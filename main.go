package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vittico/g-s-ca20201-micro/dbhelper"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Globers! Bienvenidos a Cloud Academy - G-S-CA20201!")
	fmt.Println("Endpoint Hit: homePage")

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	log.Println("Inicializando g-s-ca20201-micro...")

	// Environment variables, secrets
	pgHost, errEnv := os.LookupEnv("PG_HOST")
	if errEnv != true {
		log.Fatal("No pude localizar PG_HOST")
	}
	pgPort, errEnv := os.LookupEnv("PG_PORT")
	if errEnv != true {
		log.Fatal("No pude localizar PG_PORT")
	}
	pgUser, errEnv := os.LookupEnv("PG_USER")
	if errEnv != true {
		log.Fatal("No pude localizar PG_USER")
	}
	pgPassword, errEnv := os.LookupEnv("PG_PASSWORD")
	if errEnv != true {
		log.Fatal("No pude localizar PG_PASSWORD")
	}
	pgDB, errEnv := os.LookupEnv("PG_DB")
	if errEnv != true {
		log.Fatal("No pude localizar PG_DB")
	}
	pgTLSMode, errEnv := os.LookupEnv("PG_TLSMode")
	if errEnv != true {
		log.Fatal("No pude localizar PG_TLSMode")
	}

	log.Println("PG Host -> " + pgHost)
	log.Println("PG Port -> " + pgPort)
	log.Println("PG User -> " + pgUser)
	log.Println("PG Password -> " + pgPassword)
	log.Print("PG DB -> " + pgDB)
	log.Print("PG TLS Mode -> " + pgTLSMode)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDB, pgTLSMode)

	d := dbhelper.Connection(psqlInfo)

	handleRequests()

}
