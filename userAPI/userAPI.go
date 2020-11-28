package userapi

import (
	"flag"
	ioc "golang-restful/IoC"
)

var (
	domainAPI = flag.String("API", "https://jsonplaceholder.typicode.com/", "api server domain")
)

func init() {
	ioc.URL = *domainAPI
}

// Call means
func Call(httpMethod, path string, requestBody interface{}) (response interface{}, err error) {

	restful := ioc.New()

	defer func() {
		ioc.RequestHeaders = nil
		ioc.URL = ""
	}()

	ioc.AppendNewHeaders("Content-Type", "application/json; charset=utf-8")

	return restful.Call(httpMethod, path, requestBody)
}
