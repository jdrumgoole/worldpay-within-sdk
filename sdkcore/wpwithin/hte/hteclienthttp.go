package hte

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HTEClientHTTP interface {
	Get(url string) ([]byte, error)
	PostJSON(url string, postBody []byte) ([]byte, int, error)
}

type HTEClientHTTPImpl struct{}

func NewHTEClientHTTP() (HTEClientHTTP, error) {

	return &HTEClientHTTPImpl{}, nil
}

// Get Helper function to make a HTTP GET request
func (client *HTEClientHTTPImpl) Get(url string) ([]byte, error) {

	response, err := http.Get(url)

	if err != nil {

		return nil, err
	}

	byteResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {

		return nil, err
	}

	return byteResponse, nil
}

// PostJSON Helper function to make a http POST request
func (client *HTEClientHTTPImpl) PostJSON(url string, postBody []byte) ([]byte, int, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))

	if err != nil {

		return nil, 0, err
	}

	req.Header.Add("Content-Type", "application/json")

	_client := &http.Client{}

	resp, err := _client.Do(req)

	if err != nil {

		return nil, 0, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return nil, resp.StatusCode, err
	}

	return bodyBytes, resp.StatusCode, nil
}
