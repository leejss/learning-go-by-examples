package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/")
	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// // Read body
	// body, err := io.ReadAll(resp.Body) // 모든 데이터를 읽어서 바이트 배열로 반환

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(body))
	jsonExample()
}

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func jsonExample() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result []Todo
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

}
