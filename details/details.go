package details

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"path"
	"reflect"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/jpillora/go-ogle-analytics"
	"github.com/remotejob/huoneisto_go/dbhandler"
	"github.com/remotejob/huoneisto_go/domains"
	"github.com/remotejob/huoneisto_go/ldjsonhandler"
	shuffle "github.com/shogo82148/go-shuffle"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Show(w http.ResponseWriter, r *http.Request) {

	initstruct := r.Context().Value("init").(domains.InitStruct)
	log.Println(initstruct, r.RequestURI)

	client, err := ga.NewClient(initstruct.Analytics)
	if err != nil {
		panic(err)
	}

	err = client.Send(ga.NewEvent("details", r.RequestURI).Label("details button"))
	if err != nil {
		panic(err)
	}

	site := initstruct.Site
	// mobile := initstruct.Mobile
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

	details := path.Join("templates", "details.html")
	newtemplate := "details"

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

			t, err := template.New(newtemplate).Funcs(funcMap).ParseFiles(details)
			check(err)
			objtoInject := domains.ObjtoTemplate{atricleToInject, initstruct.Assets, initstruct.Analytics}

			err = t.Execute(w, objtoInject)
			check(err)
		} else {
			http.Error(w, http.StatusText(404), 404)
		}

	}
}
