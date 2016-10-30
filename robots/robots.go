package robots

import (
	"bytes"
	"log"
	"net/http"

	"github.com/remotejob/huoneisto_go/domains"
)

//Generate create robots.txt
func Generate(w http.ResponseWriter, r *http.Request) {

	initstruct := r.Context().Value("init").(domains.InitStruct)
	log.Println("themes", initstruct.Addrs[0])

	var buffer bytes.Buffer

	// sitefull := r.Host
	// site := strings.Split(sitefull, ":")[0]

	site := initstruct.Site

	buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://" + site + "/sitemap.xml\n")

	w.Header().Add("Content-type", "text/plain")
	w.Write(buffer.Bytes())

}
