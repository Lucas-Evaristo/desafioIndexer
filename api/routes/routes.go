package routes

import (
	dbconections "api-zincsearch-desafio/dbConections"
	"api-zincsearch-desafio/handlers"
	"net/http"
)

func GetEmails(response http.ResponseWriter, request *http.Request) {

	term := request.URL.Query().Get("term")

	db := new(dbconections.ZincSearch)

	s := handlers.ServiceZinc{DB: db}

	data, err := s.SearchQuery(term)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		response.Write(data)
	}
}
