package main

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"net/http"
	"time"
)

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func (c *Comment) GetValue(url string, id int) {

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
		fmt.Println("error call ", url, " ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("id = %d", id)
	if id <= 0 || id >= 500 {
		var respData []Comment
		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			fmt.Println("error comment json decoder1")
			return
		}
		writeXlsxComment("comments", respData)
	} else {
		var respData Comment
		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			fmt.Println("error comment json decoder2")
			return
		}

		comArr := []Comment{respData}
		writeXlsxComment("comments", comArr)
	}

}

type Photo struct {
	albumId      int    `json:"albumId"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

func writeXlsxComment(sheetName string, commentArr []Comment) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	//var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet(sheetName)
	if err != nil {
		fmt.Println("error cannot add sheet name")
	}

	for _, cm := range commentArr {
		row = sheet.AddRow()
		row.AddCell().SetValue(cm.PostId)
		row.AddCell().SetValue(cm.Id)
		row.AddCell().SetValue(cm.Name)
		row.AddCell().SetValue(cm.Email)
		row.AddCell().SetValue(cm.Body)

		err = file.Save("comments.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

}
