package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
)

// cfg := elasticsearch.Config{
// 	Addresses: []string{
// 		"http://127.0.0.1:9200",
// 		"http://127.0.0.1:9200",
// 	},
// 	Transport: &http.Transport{
// 		MaxIdleConnsPerHost:   10,
// 		ResponseHeaderTimeout: time.Second,
// 	}}
// es, err := elasticsearch.NewClient(cfg)

type esClient struct {
	client *elasticsearch.Client
	err    error
	mutex  sync.Mutex
}

var client esClient

func init() {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.client == nil {
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			client.err = err
			return
		}
		client.client = es
	}
}

func GetClient() (*elasticsearch.Client, error) {
	if client.err != nil {
		return nil, client.err
	}
	return client.client, nil
}

func main() {
	c, err := GetClient()
	if err != nil {
		println("1111", err)
		return
	}
	res, err := c.Search(c.Search.WithIndex("test"))
	if res.IsError() {
		println("2222", res.String())
		return
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		println("3333", err)
		return
	}
	for k, v := range r {
		println("------------------------")
		println("------------------------")
		fmt.Printf("%v", k)
		fmt.Printf("%v", v)
		println("------------------------")
		println("------------------------")
	}
	// var buf bytes.Buffer
	// query := map[string]interface{}{
	// "query": map[string]interface{}{
	// "match": map[string]interface{}{
	// 	"title": "test",
	// },
	// },
	// }
	// if err := json.NewEncoder(&buf).Encode(query); err != nil {
	// 	println("2", err)
	// 	return
	// }
	// res, err := es.Search(
	// 	es.Search.WithContext(context.Background()),
	// 	es.Search.WithIndex("user"),
	// 	// es.Search.WithBody(&buf),
	// 	es.Search.WithTrackTotalHits(true),
	// 	// es.Search.WithPretty(),
	// )
	// if err != nil {
	// 	println("3", err)
	// 	return
	// }
	// res, err := es.Info()
	// if err != nil {
	// 	println("a", err)
	// 	return
	// }
	// println(res)
	// defer res.Body.Close()
	// var r map[string]interface{}
	// if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	// 	log.Fatalf("Error parsing the response body: %s", err)
	// }
	// // Print client and server version numbers.
	// log.Printf("Client: %s", elasticsearch.Version)
	// log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	// log.Println(strings.Repeat("~", 37))
	// println("res:", res)
	// defer res.Body.Close()
}


