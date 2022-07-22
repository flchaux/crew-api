package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"model"
	"net/http"
)

func HandleTalentEndPoint(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		handleGetTalentList(writer, request)
	} else if request.Method == http.MethodPost {
		handleAddTalent(writer, request)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func handleGetTalentList(writer http.ResponseWriter, request *http.Request) {
	// TODO: read talents from DB
	talents := [...]model.Talent{{"TestName"}}
	stringified, _ := json.Marshal(talents)
	io.WriteString(writer, string(stringified))
}

func handleAddTalent(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err == nil {
		var talent model.Talent
		parseError := json.Unmarshal([]byte(body), &talent)
		if parseError == nil {
			// TODO: add talent to DB
			writer.WriteHeader(http.StatusCreated)
			fmt.Printf("new talent: %s", talent.FirstName)
		} else {
			fmt.Printf("Error parsing talent data: ", parseError)
			writer.WriteHeader(http.StatusBadRequest)
		}

	} else {
		fmt.Printf("Error reading request body: ", err)
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/talent", HandleTalentEndPoint)
	http.ListenAndServe(":8080", nil)
}
