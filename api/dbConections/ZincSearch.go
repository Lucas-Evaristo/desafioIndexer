package dbconections

import (
	"api-zincsearch-desafio/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type ZincSearch struct {
}

var (
	_sortFields = []string{"-@timestamp"}
	_from       = 0
	_maxResults = 15
	_user       = "admin"
	_pass       = "admin111"
	_url        = "http://localhost:4080/api/Emails/_search"
)

func (z *ZincSearch) Search(term string) (*http.Response, error) {
	method := "POST"
	query := GetSearchQuery(term)

	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest(method, _url, buffer)
	if err != nil {
		return nil, err
	}

	auth := _user + ":" + _pass
	basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Set("Authorization", "Basic "+basicAuth)
	req.Header.Add("Content-Type", "text/plain")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetSearchQuery(term string) models.Search {
	query := models.Query{Term: term, Field: "_all"}

	search := models.Search{
		Search_type:  "matchphrase",
		Search_query: query,
		Sort_fields:  _sortFields,
		From:         _from,
		Max_results:  _maxResults,
		Source:       []string{},
	}
	return search
}
