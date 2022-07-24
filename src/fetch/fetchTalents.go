package main

import (
	"dal"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"model"
	"net/http"

	"github.com/joho/godotenv"
)

func insertTalentsFromPage(page int) int {
	url := fmt.Sprintf("https://hiring.crew.work/v1/talents?limit=20&offset=%d", page)
	response, getError := http.Get(url)
	if getError != nil {
		panic(getError)
	}
	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		panic(readError)
	}
	var talents []model.Talent
	parseError := json.Unmarshal([]byte(body), &talents)
	if parseError != nil {
		panic(parseError)
	}
	for _, talent := range talents {
		dal.AddTalent(talent)
	}
	return len(talents)
}

func main() {
	godotenv.Load()
	page := 0
	for {
		nbInserted := insertTalentsFromPage(page)
		page++
		fmt.Printf("Add %d from page %d\n", nbInserted, page)
		if nbInserted < 20 {
			break
		}
	}
}
