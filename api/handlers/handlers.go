package handlers

import (
	"api-zincsearch-desafio/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServiceZinc struct {
	DB models.DBHandler
}

func (s ServiceZinc) SearchQuery(term string) ([]byte, error) {
	var resZinc models.Response
	var emails []models.Source
	response, err := s.DB.Search(term)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &resZinc)
		if err != nil {
			return nil, err
		}

		for _, src := range resZinc.Hits.Hits {
			emails = append(emails, src.Source)
		}

		emailsJson, err := json.Marshal(emails)
		if err != nil {
			return nil, err
		}

		return emailsJson, nil
	} else {
		err := fmt.Errorf("error en la respuesta de la base de datos")
		return nil, err
	}

}
