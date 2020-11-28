package main

import (
	"fmt"
	userapi "golang-restful/userAPI"
	"log"
	"net/http"
)

func main() {

	response, err := userapi.Call(http.MethodGet, "todos/1", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response)

	// api.DomainName = "https://jsonplaceholder.typicode.com/"
	// restful := api.New()

	// resp, err := restful.Get("todos/1")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(resp)
}
