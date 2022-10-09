package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseURL string = "https://kr.indeed.com/jobs?q=javascript+react&limit=50&vjk=bf7bf53f892557fb"

func main() {
	getPages()

}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)

	body, err := io.ReadAll(res.Body)
	checkErr(err)
	fmt.Printf("%s", body)
	defer res.Body.Close()
	return pages

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkStatusCode(resp *http.Response) {
	if resp.StatusCode > 200 {
		if resp.StatusCode != 403 {
			log.Fatalln(resp.Status)
		}
	}
}
