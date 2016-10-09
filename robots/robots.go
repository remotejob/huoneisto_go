package robots

import (
	"bytes"
	"log"
	"net/http"
	"strings"
)

//Generate create robots.txt
func Generate(w http.ResponseWriter, r *http.Request) {

	log.Println("themes", r.Context().Value("themes"))

	var buffer bytes.Buffer

	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]

	buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://" + site + "/sitemap.xml\n")

	w.Header().Add("Content-type", "text/plain")
	w.Write(buffer.Bytes())

}
