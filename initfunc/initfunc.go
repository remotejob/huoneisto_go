package initfunc

import (
	"log"
	"os"
)

// var themes string
// var locale string

// var addrs []string
// var database string
// var username string
// var password string
// var mechanism string
// var mainroute string

//GetPar get start parameters
func GetPar() (string, string, []string, string, string, string, string, string) {

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	themes := os.Getenv("THEMES")
	locale := os.Getenv("LOCALE")
	dbadmin := os.Getenv("DBADMIN")
	mechanism := os.Getenv("MECHANISM")
	addrs := []string{os.Getenv("ADDRS")}
	mainroute := os.Getenv("MAINROUTE")
	log.Println("mongodbpass", password)

	return themes, locale, addrs, dbadmin, username, password, mechanism, mainroute
}
