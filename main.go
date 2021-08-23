package main

func main() {}

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // type EsReb struct {
// // 	Took    int64 `json:"took"`
// // 	TimeOut bool  `json:"time_out"`
// // 	Shards
// // }
// // type Shards struct {
// // 	Total int64 `json:"total"`
// // }

// func main() {
// 	c, err := GetClient()
// 	if err != nil {
// 		println("1111", err)
// 		return
// 	}
// 	res, err := c.Search(c.Search.WithIndex("shop"))
// 	if res.IsError() {
// 		println("2222", res.String())
// 		return
// 	}
// 	// var r map[string]interface{}
// 	var r EsReb
// 	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
// 		println("3333", err)
// 		return
// 	}
// 	for _, v := range r.Hits.Hits {
// 		fmt.Printf("%v\n", v)
// 	}
// }
