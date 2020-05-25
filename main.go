package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/vittico/g-s-ca20201-micro/models"
)

type Env struct {
	CTX *gorm.DB
}

func main() {

	log.Println("Inicializando g-s-ca20201-micro...")

	// Get values from environment
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
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDB, pgTLSMode)

	/*
	 Dependency inject the ctx pool
	*/
	ctx := models.ConnectDB(psqlInfo)
	env := &Env{CTX: ctx}
	/**/

	// endpoint
	http.HandleFunc("/", env.homePage)
	http.HandleFunc("/user/details", env.userDetails)

	log.Fatal(http.ListenAndServe(":10000", nil))

}

func (env *Env) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Globers! Bienvenidos a Cloud Academy - G-S-CA20201! - / endpooint ")
	// Browsers usually hit twice, / endpooint and ico, let's ignore anythiing but the actual endpoint deal?
	if r.RequestURI == "/" {
		models.RecordIt(env.CTX)
		log.Println("Endpoint Hit: homePage")
	}
}

func (env *Env) userDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Globers! Bienvenidos a Cloud Academy - G-S-CA20201! - /user/details")
	// Browsers usually hit twice, / endpooint and ico, let's ignore anythiing but the actual endpoint deal?
	if r.RequestURI == "/" {
		models.RecordIt(env.CTX)
		log.Println("Endpoint Hit: userDetails")
	}
}
