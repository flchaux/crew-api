package main

import (
	"dal"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"model"
	"net/http"

	"github.com/joho/godotenv"
)

func handleTalentEndPoint(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		handleGetTalentList(writer, request)
	} else if request.Method == http.MethodPost {
		handleAddTalent(writer, request)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func handleGetTalentList(writer http.ResponseWriter, request *http.Request) {
	talents := dal.GetAllTalents()
	stringified, _ := json.Marshal(talents)
	io.WriteString(writer, string(stringified))
}

func handleAddTalent(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err == nil {
		var talent model.Talent
		parseError := json.Unmarshal([]byte(body), &talent)
		if parseError == nil {
			dbAddError := dal.AddTalent(talent)
			if dbAddError == nil {
				writer.WriteHeader(http.StatusCreated)
				fmt.Printf("new talent: %s", talent.FirstName)
			} else {
				writer.WriteHeader(http.StatusBadRequest)
				fmt.Print(dbAddError)
			}
		} else {
			fmt.Print("Error parsing talent data: ", parseError)
			writer.WriteHeader(http.StatusBadRequest)
		}

	} else {
		fmt.Print("Error reading request body: ", err)
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	godotenv.Load()
	http.HandleFunc("/talent", handleTalentEndPoint)
	http.ListenAndServe(":8080", nil)
}
