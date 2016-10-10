//Package ldjsonhandler for ld-Json
package ldjsonhandler

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/kazarena/json-gold/ld"
	"github.com/remotejob/huoneisto_go/domains"
)

//Create ld-json comments on web page
func Create(articles []domains.Articlefull, pgType string) []byte {

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	var doc map[string]interface{}

	var itemListElement []interface{}
	for pos, article := range articles {

		createdstr := article.Created.Format("2006-01-02")
		updatedstr := article.Updated.Format("2006-01-02")
		pagelink := "http://" + article.Site + "/job/fi_FI/blogi/" + article.Stitle + ".html"
		publisher := map[string]interface{}{"@type": "Organization", "name": "Remote Job Finland OY", "logo": map[string]interface{}{"@type": "ImageObject", "url": "http://mazurov.eu/img/mazurovopt.jpg", "height": "200px", "width": "300px"}}
		image := map[string]interface{}{"@type": "ImageObject", "url": "http://" + article.Site + "/assets/img/free_for_job.png", "height": "256px", "width": "256px"}
		mainEntityOfPage := map[string]interface{}{"@type": "WebPage", "@id": "http://" + article.Site}

		var headline string

		runes := bytes.Runes([]byte(article.Title))
		if len(runes) > 109 {

			headline = string(runes[:108]) + "."

		} else {

			headline = article.Title + "."

		}

		listItem := map[string]interface{}{"@type": "ListItem", "position": pos, "item": map[string]interface{}{"@type": "Article", "author": article.Author,
			"headline":         headline,
			"publisher":        publisher,
			"image":            image,
			"datepublished":    createdstr,
			"datemodified":     updatedstr,
			"mainEntityOfPage": mainEntityOfPage,
			"url":              pagelink,
			"articleSection":   "job",
			"articleBody":      article.Contents}, "name": headline}

		itemListElement = append(itemListElement, listItem)
	}

	doc = map[string]interface{}{
		"@context":        "http://schema.org",
		"@type":           "BreadcrumbList",
		"itemListElement": itemListElement,
		"name":            pgType,
	}

	comp, err := proc.Compact(doc, nil, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)

	}

	b, err := json.Marshal(comp)
	if err != nil {
		log.Println(err.Error())

	}

	return b
}
