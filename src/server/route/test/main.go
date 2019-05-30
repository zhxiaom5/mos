package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
