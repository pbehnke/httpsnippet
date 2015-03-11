package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {
	client := &http.Client{}
	url := "http://mockbin.com/har"
	payload := strings.NewReader("{\"foo\": \"bar\"}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}
