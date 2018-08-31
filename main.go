package main

import (
	"./config"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//
	resp, e := http.Get(config.BaseUrl)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code is ", resp.StatusCode)
		return
	} else {
		all, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		fmt.Printf("%s\n", all)
	}
}
