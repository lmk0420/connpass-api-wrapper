package main

import (
	"fmt"
	"encoding/json"
	"log"
	"strconv"
	"net/url"
	"./httpwrapper"
	"./types"
)

func main() {
	const endpoint = "https://connpass.com/api/v1/event/"
	keywordor := []string{"JavaScript", "TypeScript", "Go", "Java", "Kotlin",
						  "Vue", "React", "Nuxt", "Next", "Spring Boot"}
	keyword := []string{"東京"}
	count := 30
	param := types.RequestParam{Keyword: keyword, KeywordOR: keywordor, Count: count}
	values := url.Values{"keyword": keyword, "keywordor": param.KeywordOR, "count": {strconv.Itoa(count)}}

	data, err := httpwrapper.DoRequest("GET", endpoint, values, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	var rs types.ResultSet
	if err := json.Unmarshal(data, &rs); err != nil {
		log.Fatal(err)
	}

	for i, event := range rs.Events {
		fmt.Println(strconv.Itoa(i + 1) + " Title: " + event.Series.Title)
	}
}