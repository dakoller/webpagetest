package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
        "gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("secretfile")
        if err != nil {
            fmt.Printf("Fail to read file: %v", err)
            os.Exit(1)
        }

        // Classic read of values, default section can be represented as empty string
        fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
        fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())	


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
