package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var todo Todo
	json.Unmarshal(bodyBytes, &todo) //sunucudan gelen biti okuyor ve struck olan todo ya atıyor
	fmt.Println(todo)
}

func Demo2() {
	todo := Todo{1, 2, "alışveriş yapılacak", false}
	jsonTodo, err := json.Marshal(todo) // struc ı json a çevirir
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post(
		"https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonTodo), // post etmeden önce geçiçi süre paketliyor bardak gibi
	)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) //fake sunucudan gelen(yansıyan) veriyi oku-yazdır
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoresponse Todo
	json.Unmarshal(bodyBytes, &todoresponse)
	fmt.Println(todoresponse)

}
