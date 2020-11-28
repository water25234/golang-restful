package api

import (
	"log"
	"net/http"
)

// Get means
func (req *RequestHTTP) Get(path string) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	headers := &[]RequestHeader{}
	req.NewHeader(headers)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}

// Post means
func (req *RequestHTTP) Post(path string, requestBody interface{}) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodPost, path, requestBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	headers := &[]RequestHeader{
		{
			Key:   "Content-Type",
			Value: "application/json; charset=utf-8",
		},
	}
	req.NewHeader(headers)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}

// Put means
func (req *RequestHTTP) Put(path string, requestBody interface{}) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodPut, path, requestBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	headers := &[]RequestHeader{
		{
			Key:   "Content-Type",
			Value: "application/json; charset=utf-8",
		},
	}
	req.NewHeader(headers)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}

// Patch means
func (req *RequestHTTP) Patch(path string, requestBody interface{}) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodPatch, path, requestBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	headers := &[]RequestHeader{
		{
			Key:   "Content-Type",
			Value: "application/json; charset=utf-8",
		},
	}
	req.NewHeader(headers)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}

// Delete means
func (req *RequestHTTP) Delete(path string, requestBody interface{}) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodDelete, path, requestBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	headers := &[]RequestHeader{
		{
			Key:   "Content-Type",
			Value: "application/json; charset=utf-8",
		},
	}
	req.NewHeader(headers)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}
