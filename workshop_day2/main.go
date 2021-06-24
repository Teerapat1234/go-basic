package main

import (
	"flag"
	"fmt"
)

func main() {

	apiName := flag.String("api_name", "", "api url")
	id := flag.Int("id", 0, "id")

	flag.Parse()

	if *apiName != "comments" && *apiName != "photos" {
		fmt.Println("no api_name input")
		return
	}

	apiUrl := fmt.Sprintf("%s%s/", "http://jsonplaceholder.typicode.com/", *apiName)

	if *id > 0 && *id < 500 {
		apiUrl = fmt.Sprintf("%s%d", apiUrl, *id)
	}

	if *apiName == "comments" {
		var com Comment
		com.GetValue(apiUrl, *id)
	}
}
