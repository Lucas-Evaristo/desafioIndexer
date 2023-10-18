package models

import "net/http"

type DBHandler interface {
	Search(term string) (*http.Response, error)
}
