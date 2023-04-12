package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	data := bytes.NewReader([]byte("{\n\"data\":\"123\"\n}"))
	client := http.Client{}
	res, err := client.Post("http://localhost:8081/save", "application/json", data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Запрос выполнен со статусом: %s \n", res.StatusCode)
}
