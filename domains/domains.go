package domains

import (
	"encoding/xml"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//InitStruct start param
type InitStruct struct {
	Themes    string
	Locale    string
	Addrs     []string
	Dbadmin   string
	Username  string
	Password  string
	Mechanism string
	Mainroute string
	Mobile    bool
	Site      string
	Analytics string
	Assets    map[int][]string
}

//LogRecord substitude Nginx log capacity
type LogRecord struct {
	Date  time.Time
	Log   string
	Ltype string
}

// Gphrase comment
type Gphrase struct {
	Phrase string `bson:"Phrase"`
	Rating int    `bson:"Rating"`
}

//Articlefull complite entity from DB
type Articlefull struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Title  string
	Stitle string
	// Tags      string
	Contents  string
	Mcontents string
	Site      string
	Author    string
	Created   time.Time
	Updated   time.Time
}

//Sitemap_from_db only info for Sitemap format
type Sitemap_from_db struct {
	Stitle  string
	Site    string
	Updated time.Time
}

//Pages struct keep sitemap obj
type Pages struct {
	//	Version string   `xml:"version,attr"`
	XMLName xml.Name `xml:"urlset"`
	XmlNS   string   `xml:"xmlns,attr"`
	//	XmlImageNS string   `xml:"xmlns:image,attr"`
	//	XmlNewsNS  string   `xml:"xmlns:news,attr"`
	Pages []*Page `xml:"url"`
}

//Page sitemap Page
type Page struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Lastmod    string   `xml:"lastmod"`
	Changefreq string   `xml:"changefreq"`
	//	Name       string   `xml:"news:news>news:publication>news:name"`
	//	Language   string   `xml:"news:news>news:publication>news:language"`
	//	Title      string   `xml:"news:news>news:title"`
	//	Keywords   string   `xml:"news:news>news:keywords"`
	//	Image      string   `xml:"image:image>image:loc"`
}
