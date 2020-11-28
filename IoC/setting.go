package ioc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	// RequestHeaders means
	RequestHeaders []RequestHeader

	// URL means
	URL string
)

// New means
func New() API {
	return &RequestHTTP{
		Request: &http.Request{},
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

// AppendNewHeaders means
func AppendNewHeaders(key string, value string) {
	if len(key) == 0 || len(value) == 0 {
		return
	}

	RequestHeaders = append(RequestHeaders, RequestHeader{
		Key:   key,
		Value: value,
	})
}

// NewRequest means
func (req *RequestHTTP) newRequest(method string, path string, requestbody interface{}) (err error) {
	url := fmt.Sprintf("%s%s", URL, path)

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

// newHeader means
func (req *RequestHTTP) newHeader() {
	if RequestHeaders == nil {
		return
	}

	for _, header := range RequestHeaders {
		req.Request.Header.Set(header.Key, header.Value)
	}
}

// executeAPI means
func (req *RequestHTTP) executeAPI() (response interface{}, err error) {
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
