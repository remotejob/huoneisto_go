package initfunc

import (
	"os"

	"github.com/remotejob/huoneisto_go/domains"
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
func GetPar() domains.InitStruct {

	var initstruct domains.InitStruct

	initstruct.Username = os.Getenv("USERNAME")

	initstruct.Password = os.Getenv("PASSWORD")
	initstruct.Themes = os.Getenv("THEMES")
	initstruct.Locale = os.Getenv("LOCALE")
	initstruct.Dbadmin = os.Getenv("DBADMIN")
	initstruct.Mechanism = os.Getenv("MECHANISM")
	initstruct.Addrs = []string{os.Getenv("ADDRS")}
	initstruct.Mainroute = os.Getenv("MAINROUTE")
	initstruct.Mobile = false
	initstruct.Analytics = os.Getenv("ANALYTICS")

	return initstruct
}
