package fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	response, getError := http.Get("https://hiring.crew.work/v1/talents")
	body, readError = ioutil.ReadAll(response.Body)
	var talents []Talent
	json.Unmarshal()
}
