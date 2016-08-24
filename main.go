package googleshortener

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const MasterURL = "https://www.googleapis.com/urlshortener/v1/url"

type GoogleAPI struct {
	Key string
	Url string
}

func New(key string) (*GoogleAPI, error) {
	err := checkKey(key)
	if err != nil {
		return nil, err
	}

	return &GoogleAPI{
		Key: key,
		Url: MasterURL + "?key=" + key,
	}, nil
}

func checkKey(key string) error {
	url := MasterURL + "?key=" + key
	var jsonStr = []byte(`{'longUrl':'http://google.com'}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
