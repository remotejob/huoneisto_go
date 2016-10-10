package sitemap

import (
	"encoding/xml"
	"log"
	"net/http"
	"time"

	"github.com/remotejob/huoneisto_go/domains"
	"github.com/remotejob/kaukotyoeu/dbhandler"
	mgo "gopkg.in/mgo.v2"
)

var resultXML []byte

//CheckServeSitemap create dinamic sitemap.xml file
func CheckServeSitemap(w http.ResponseWriter, r *http.Request) {

	initstruct := r.Context().Value("init").(domains.InitStruct)
	log.Println(initstruct)

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     initstruct.Addrs,
		Timeout:   60 * time.Second,
		Database:  initstruct.Dbadmin,
		Username:  initstruct.Username,
		Password:  initstruct.Password,
		Mechanism: initstruct.Mechanism,
	}

	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	allsitemaplinks := dbhandler.GetAllSitemaplinks(*dbsession, initstruct.Site)

	docList := new(domains.Pages)
	docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

	for _, sitemaplink := range allsitemaplinks {

		if sitemaplink.Site == initstruct.Site {

			doc := new(domains.Page)
			doc.Loc = "http://" + initstruct.Site + "/" + initstruct.Themes + "/" + initstruct.Locale + "/" + initstruct.Mainroute + "/" + sitemaplink.Stitle + ".html"
			doc.Lastmod = sitemaplink.Updated.Format(time.RFC3339)
			doc.Changefreq = "monthly"
			docList.Pages = append(docList.Pages, doc)

		}

	}

	resultXML, err = xml.MarshalIndent(docList, "", "  ")
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Add("Content-type", "application/xml")
	_, err = w.Write(resultXML)
	if err != nil {
		log.Println(err.Error())
	}

}
