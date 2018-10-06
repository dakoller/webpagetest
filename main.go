package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
        "gopkg.in/ini.v1"
)

func main() {
	resp, err := http.Get(os.Getenv("URL"))
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(os.Getenv("URL"))
	fmt.Println(len(body))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
