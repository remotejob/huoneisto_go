package blog

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"path"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	"github.com/remotejob/huoneisto_go/dbhandler"
	"github.com/remotejob/huoneisto_go/domains"

	"github.com/remotejob/huoneisto_go/ldjsonhandler"
	shuffle "github.com/shogo82148/go-shuffle"
	mgo "gopkg.in/mgo.v2"

	"github.com/jpillora/go-ogle-analytics"
)

// func checkReq(w http.ResponseWriter, r *http.Request, mobile string) {

// 	now := time.Now()
// 	var record domains.LogRecord
// 	var ltype string
// 	var log string

// 	if strings.Contains(r.Referer(), "www.google") {

// 		ltype = "google"

// 	}

// 	if strings.Contains(mobile, "true") {

// 		ltype = "mobile"

// 	}

// 	if len(ltype) > 1 {
// 		log = r.Referer() + "," + r.RequestURI + "," + r.UserAgent()
// 		record = domains.LogRecord{Date: now, Log: log, Ltype: ltype}
// 		go insertlog.InsertIntoDB(record)
// 	}

// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//CreateArticelePage createPage
func CreateArticelePage(w http.ResponseWriter, r *http.Request) {

	initstruct := r.Context().Value("init").(domains.InitStruct)
	// log.Println(initstruct, r.RequestURI)

	// client, err := ga.NewClient(initstruct.Analytics)
	// if err != nil {
	// 	panic(err)
	// }

	// err = client.Send(ga.NewEvent("hit", r.RequestURI).Label("article"))
	// if err != nil {
	// 	panic(err)
	// }

	var lp string
	var headercommon string
	var newtemplate string

	// site := initstruct.Site
	mobile := initstruct.Mobile
	// log.Println("article site", site, "mobile", mobile)
	// checkReq(w, r, mobile)

	vars := mux.Vars(r)

	mtitle := vars["mtitle"]

	log.Println("mtitle", mtitle)

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

	if !mobile {
		lp = path.Join("templates", "layout.html")
		headercommon = path.Join("templates", "header_common.html")
		newtemplate = "layout.html"

	} else {
		lp = path.Join("templates", "m_layout.html")
		headercommon = path.Join("templates", "m_header_common.html")
		newtemplate = "m_layout.html"
	}

	article := dbhandler.GetOneArticle(*dbsession, initstruct.Site, mtitle)

	// var articles []domains.Articlefull

	if len(article) == 1 {

		funcMap := template.FuncMap{
			"marshal": func(objtoTemplate domains.ObjtoTemplate) template.JS {

				b := ldjsonhandler.Create(objtoTemplate.Articles, "Selected Article")

				return template.JS(b)
			},
			"title": func(a domains.ObjtoTemplate) string {

				return a.Articles[0].Title

			},
			"hasField": func(v interface{}, name string) bool {
				rv := reflect.ValueOf(v)
				if rv.Kind() == reflect.Ptr {
					rv = rv.Elem()
				}
				if rv.Kind() != reflect.Struct {
					return false
				}
				return rv.FieldByName(name).IsValid()
			},
		}
		t, err := template.New(newtemplate).Funcs(funcMap).ParseFiles(lp, headercommon)
		check(err)

		objtoInject := domains.ObjtoTemplate{article, initstruct.Assets, initstruct.Analytics}

		err = t.Execute(w, objtoInject)
		check(err)

	} else {
		http.NotFound(w, r)
		// http.Error(w, http.StatusText(404), 404)
	}

}

//CreateIndexPage create Index Page
func CreateIndexPage(w http.ResponseWriter, r *http.Request) {

	var lp string
	var headercommon string
	var newtemplate string

	initstruct := r.Context().Value("init").(domains.InitStruct)

	client, err := ga.NewClient(initstruct.Analytics)
	if err != nil {
		panic(err)
	}

	err = client.Send(ga.NewEvent("hit", r.RequestURI).Label("index"))
	if err != nil {
		panic(err)
	}
	site := initstruct.Site
	mobile := initstruct.Mobile
	// log.Println("article site", site, "mobile", mobile, "img", initstruct.Assets)
	// checkReq(w, r, mobile)

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

	if !mobile {
		lp = path.Join("templates", "home_page.html")
		headercommon = path.Join("templates", "header_common.html")
		newtemplate = "home_page.html"

	} else {
		lp = path.Join("templates", "m_home_page.html")
		headercommon = path.Join("templates", "m_header_common.html")
		newtemplate = "m_home_page.html"
	}

	funcMap := template.FuncMap{
		"marshal": func(objtoTemplate domains.ObjtoTemplate) template.JS {
			b := ldjsonhandler.Create(objtoTemplate.Articles, "Index Page")
			return template.JS(b)
		},
		"title": func(a domains.ObjtoTemplate) string {

			return a.Articles[0].Title

		},
		"hasField": func(v interface{}, name string) bool {
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			if rv.Kind() != reflect.Struct {
				return false
			}
			return rv.FieldByName(name).IsValid()
		},
	}

	t, err := template.New(newtemplate).Funcs(funcMap).ParseFiles(lp, headercommon)
	check(err)

	allarticles := dbhandler.GetAllForStatic(*dbsession, site)

	if len(allarticles) > 0 {

		var numberstoshuffle []int
		for num := range allarticles {

			numberstoshuffle = append(numberstoshuffle, num)

		}
		rand.Seed(time.Now().UTC().UnixNano())

		shuffle.Ints(numberstoshuffle)

		var atricleToInject []domains.Articlefull

		for c, i := range numberstoshuffle {

			if c < 10 {

				atricleToInject = append(atricleToInject, allarticles[i])
			}

		}

		if len(atricleToInject) > 0 {

			objtoInject := domains.ObjtoTemplate{atricleToInject, initstruct.Assets, initstruct.Analytics}

			err = t.Execute(w, objtoInject)
			check(err)
		}

	} else {
		http.Error(w, http.StatusText(404), 404)
	}

}
