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

	// t := &Todo{
	// 	UserID: 12345,
	// }

	// t.UserDemo()

	// api.DomainName = "https://jsonplaceholder.typicode.com/"
	// restful := api.New()

	// resp, err := restful.Get("todos/1")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(resp)
}

// Todo struct
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t *Todo) UserDemo() {
	fmt.Println(t.UserID)
}
