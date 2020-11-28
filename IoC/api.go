package ioc

import (
	"fmt"
	"log"
	"net/http"
)

// HTTPMethodGet means
type HTTPMethodGet struct {
	*RequestHTTP
	// Path means
	Path string

	// HTTPMethod means
	HTTPMethod string
}

// HTTPMethod means
type HTTPMethod struct {
	*RequestHTTP
	// Path means
	Path string

	// HTTPMethod means
	HTTPMethod string

	// RequestBody means
	RequestBody interface{}
}

// Call means
func (req *RequestHTTP) Call(httpMethod, path string, requestBody interface{}) (response interface{}, err error) {
	if len(httpMethod) == 0 {
		return nil, fmt.Errorf("http method is empty")
	}

	var httpMethodIoc Handle

	switch httpMethod {
	case http.MethodGet:
		httpMethodIoc = &HTTPMethodGet{
			RequestHTTP: req,
			Path:        path,
			HTTPMethod:  httpMethod,
		}
	default:
		httpMethodIoc = &HTTPMethod{
			RequestHTTP: req,
			Path:        path,
			HTTPMethod:  httpMethod,
			RequestBody: requestBody,
		}
	}

	response, err = httpMethodIoc.Handle()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return response, err
}

// Handle means
func (req *HTTPMethodGet) Handle() (response interface{}, err error) {
	err = req.NewRequest(req.HTTPMethod, req.Path, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	req.NewHeader()

	response, err = req.executeAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return response, nil
}

// Handle means
func (req *HTTPMethod) Handle() (response interface{}, err error) {
	err = req.NewRequest(req.HTTPMethod, req.Path, req.RequestBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	req.NewHeader()

	response, err = req.executeAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return response, nil
}
