package db

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"indexador/model"
	"net/http"
)

var (
	_url  = "http://localhost:4080/api/_bulkv2"
	_user = "admin"
	_pass = "admin111"
)

type ZincSearch struct {
}

func (z ZincSearch) InsertRecords(mails model.BulkEmail) error {
	method := "POST"
	jsonData, err := json.Marshal(mails)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(method, _url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(_user + ":" + _pass))
	request.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		err = errors.New("error al insertar registros en la db")
		return err
	}

	return nil
}
