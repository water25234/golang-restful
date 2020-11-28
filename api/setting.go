package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// DomainName means
var (
	DomainName string
)

// New means
func New() HTTPClient {
	return &RequestHTTP{
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// RequestHTTP means
type RequestHTTP struct {
	Request    *http.Request
	HTTPClient *http.Client
}

// RequestHeader means
type RequestHeader struct {
	Key   string
	Value string
}

// NewRequest means
func (req *RequestHTTP) NewRequest(method string, path string, requestbody interface{}) (err error) {
	url := fmt.Sprintf("%s%s", DomainName, path)

	body, err := requestBody(requestbody)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	req.Request, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

// NewHeader means
func (req *RequestHTTP) NewHeader(headers *[]RequestHeader) {
	if headers == nil {
		return
	}

	for _, header := range *headers {
		req.Request.Header.Set(header.Key, header.Value)
	}
}

// ExecuteAPI means
func (req *RequestHTTP) ExecuteAPI() (response interface{}, err error) {
	resp, err := req.HTTPClient.Do(req.Request)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer resp.Body.Close()
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	err = json.Unmarshal(respBodyBytes, &response)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return response, nil
}

// requestBody means
func requestBody(data interface{}) (body []byte, err error) {
	switch data := data.(type) {
	case nil:
		return nil, nil
	case string:
		return []byte(data), nil
	default:
		body, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
}
