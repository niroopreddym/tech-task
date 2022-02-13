package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	noOfPages := 5
	respChan := make(chan string)
	errChan := make(chan error)
	for i := 1; i < noOfPages; i++ {
		url := "http://127.0.0.1:9295/bank?pagenumber=" + strconv.Itoa(i)
		go readResponse(url, respChan, errChan)
	}

	for val := range respChan {
		fmt.Println(val)
	}

	err := <-errChan
	fmt.Println(err)
}

func readResponse(url string, respChan chan string, errChan chan error) {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		errChan <- err
	}

	respChan <- string(res)
}
