package initfunc

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/remotejob/huoneisto_go/domains"
)

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

	var imgmap map[int][]string
	var i int

	imgmap = make(map[int][]string)

	f, _ := os.Open("images.csv")
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		log.Println(record)
		imgmap[i] = record
		i++

	}
	initstruct.Assets = imgmap

	return initstruct
}
