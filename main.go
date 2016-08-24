package googleshortener

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func (b *GoogleAPI) MakeShort(longurl string) (Response, error) {
	url := MasterURL + "?key=" + b.Key
	var jsonStr = []byte(`{'longUrl':'` + longurl + `'}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return Response{}, err
	}

	if !strings.Contains(res.LongUrl, longurl) {
		return Response{}, fmt.Errorf("Invalid key")
	}

	return res, nil
}

func checkKey(key string) error {
	url := MasterURL + "?key=" + key
	LongUrl := "http://google.com"
	var jsonStr = []byte(`{'longUrl':'` + LongUrl + `'}`)
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

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if !strings.Contains(res.LongUrl, LongUrl) {
		return fmt.Errorf("Invalid key")
	}

	return nil
}
