package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

func service(key string) *customsearch.Service {
	client := &http.Client{Transport: &transport.APIKey{Key: key}}

	svc, err := customsearch.New(client)
	if err != nil {
		panic(err)
	}
	return svc
}

func search(svc *customsearch.Service, cx, query string) []map[string]interface{} {
	resp, err := svc.Cse.Siterestrict.List(query).Cx(cx).Do()
	if err != nil {
		panic(err)
	}
	results := make([]map[string]interface{}, 0)
	for _, result := range resp.Items {
		u, err := url.Parse(result.Link)
		if err != nil {
			panic(err)
		}
		p, _ := url.ParseQuery(u.RawQuery)
		id := p["id"][0]
		if stringInSlice(PlayURL+id, "url", results) {
			continue
		}
		if array, ok := result.Pagemap["mobileapplication"]; ok {
			var dec map[string]interface{}
			for _, v := range array {
				err := json.Unmarshal(v, &dec)
				if err != nil {
					panic(err)
				}
			}
			deleteKeys(dec, []string{
				"applicationcategory",
				"operatingsystem",
				"contentrating",
				"image",
			})
			dec["url"] = PlayURL + id
			results = append(results, dec)
		}
	}
	return results
}

/*
func getPkgs(svc *customsearch.Service, cx, query string) (urls []string) {
	results := search(query, svc, cx)
	for _, v := range results {
		u, err := url.Parse(v)
		if err != nil {
			panic(err)
		}
		p, _ := url.ParseQuery(u.RawQuery)
		id := p["id"][0]
		if !stringInSlice(id, urls) {
			urls = append(urls, id)
		}
	}
	return urls
}*/
