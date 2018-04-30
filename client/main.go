package main

import (
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	url := "http://www.baidu.com"

	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error")
	}
	response, _ := client.Do(reqest)
	status := response.StatusCode

	log.Println(status)
}
