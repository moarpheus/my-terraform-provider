package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	var data map[string]interface{}

	resp, err := client.Get("https://official-joke-api.appspot.com/random_joke")

	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		err = fmt.Errorf("Cannot unmarshal json of API response: %v", err)
		return
	} else if data["result"] == "" {
		err = fmt.Errorf("missing result key in API response: %v", err)
		return
	}

	outputs := make(map[string]interface{})
	outputs["id"] = data["id"]
	outputs["type"] = data["type"]
	outputs["setup"] = data["setup"]
	outputs["punchline"] = data["punchline"]

	fmt.Printf("%+v \n", outputs)
}
