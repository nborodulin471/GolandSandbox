package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

var DB = initDb()

type Data struct {
	Data string `json:"data"`
}

func main() {
	fmt.Println("Сервер стартует")
	http.HandleFunc("/save", saveData)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func initDb() *sql.DB {
	connStr := "user=postgres password=pgpwd4habr dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	handleError(err)
	return db
}

func saveData(w http.ResponseWriter, req *http.Request) {
	var data Data
	body, err := ioutil.ReadAll(req.Body)
	handleError(err)
	json.Unmarshal(body, &data)
	fmt.Printf("Было получено тело %s \n", string(body))
	_, err = DB.Exec(`INSERT INTO data(test) values($1)`, data.Data)
	handleError(err)
	fmt.Printf("Информация успешно записана")
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("Упали с ошибкой %v", err))
	}
}
