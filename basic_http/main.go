package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	var helloRequest struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(req.Body).Decode(&helloRequest); err != nil {
		fmt.Fprintf(w, `{"message":"error decode json input"}`)
		return
	}

	defer req.Body.Close()

	var helloResponse struct {
		ServerStatus string `json:"server_status"`
		Message      string `json:"message"`
	}

	helloResponse.ServerStatus = "Normal"
	helloResponse.Message = fmt.Sprintf("Hello %s", helloRequest.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(helloResponse)

}

func queryUsers(w http.ResponseWriter, req *http.Request) {
	var request struct {
		PostIndex int `json:"post_index"`
	}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		fmt.Fprintf(w, `{"message":"error decode json input"}`)
		return
	}

	//fmt.Printf("request : %#v", request)

	defer req.Body.Close()

	url := "http://jsonplaceholder.typicode.com/posts/"

	type response struct {
		UserID int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	if request.PostIndex > 0 && request.PostIndex < 100 {
		url = fmt.Sprintf("%s%d", url, request.PostIndex)
	}

	fmt.Println("url : ", url)

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error json unmarshall : ", err)
		return
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("error call https://jsonplaceholder.typicode.com/comments : ", err)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if request.PostIndex > 0 && request.PostIndex < 100 {
		var respData response
		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			fmt.Println("error unmarshall https://jsonplaceholder.typicode.com/comments : ", err)
			return
		}
		json.NewEncoder(w).Encode(respData)
	} else {
		var reponseArr []response
		if err := json.NewDecoder(resp.Body).Decode(&reponseArr); err != nil {
			fmt.Println("error unmarshall https://jsonplaceholder.typicode.com/comments : ", err)
			return
		}
		json.NewEncoder(w).Encode(reponseArr)
	}

}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/query_users", queryUsers)
	http.ListenAndServe(":7777", nil)
}
