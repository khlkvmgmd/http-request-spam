package main

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var completeCount = 0
var errorCount = 0

type any interface{}

func main() {
	attackUrl := "http://localhost:8080"
	method := "GET"
	count := 10000

	var data url.Values

	if attackUrl != "" {


		rand.Seed(time.Now().UnixNano())
		for i := 0; i < count; i++ {
			if i%5 == 0 {
				fmt.Println("Sended count:", i, "Good:", completeCount, "Bad:", errorCount)
			}
			go startAttack(attackUrl, method, data)
			time.Sleep(time.Millisecond)
		}
		fmt.Println("Done.", "Good: ", completeCount, "Error: ", errorCount)
	} else {
		fmt.Println("Set variable -url")
	}
}

func startAttack(attackUrl string, method string, data url.Values) bool {
	resp, err := http.PostForm(attackUrl, data)

	if err != nil {
		fmt.Println("Site not available: ", attackUrl, "\nERROR:")
		errorCount++
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		req := string(body)
		_ = req

		if err != nil || resp.StatusCode != 200 {
			if err != nil {
				log(err)
			} else {
				log(req)
			}
			errorCount++
		} else {
			completeCount++
		}
	}
	return true
}

func log(data any) {
	fmt.Println(data)
}

func getData(method string, data string) url.Values {
	log(method)
	if method == "POST" || method == "post" {
		var body = []byte(data)
		return getFormatPostData(body)
	} else {
		return nil
	}
}

/**
 * Format JSON to url.Values
 * Example JSON:
 * {"email": "test@inbox.ru","password": "123test321"}
 */
func getFormatPostData(body []byte) url.Values {
	m := map[string]string{}
	if err := json.Unmarshal(body, &m); err != nil {
		panic(err)
	}
	_body := url.Values{}
	for key, val := range m {
		_body.Add(key, val)
	}

	return _body
}

func getFormatGetData() {}