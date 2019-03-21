package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	c, err := statsd.New("ipdatadog")
	if err != nil {
		log.Fatal(err)
	}

	// prefix every metric with the app name
	c.Namespace = "test."
	url := "http://yourexample.com/path?yourparameters=value"
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal("Get: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.Tags = append(c.Tags, "failed")
		log.Println("Status Code is not 200")
		err = c.Count("service.count.nakama.academy", 1, nil, 1)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
		return
	}

	log.Println(string(bodyBytes))
	log.Println("Success")
	c.Tags = append(c.Tags, "success")
	err = c.Count("service.count.nakama.academy", 1, nil, 1)
}
