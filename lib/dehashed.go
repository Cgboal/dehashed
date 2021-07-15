package dehashed

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Results struct {
	Entries []Entry `json:"entries"`
	Success bool    `json:"success"`
	Message string  `json:"message"`
}

type Entry struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Hash     string `json:"hashed_password"`
	Name     string `json:"name"`
	Source   string `json:"obtained_from"`
}

func ParseDehashedJson(json_data []byte) *Results {
	results := &Results{}
	err := json.Unmarshal(json_data, results)
	if err != nil {
		log.Fatal(err)
	}
	return results

}

func FetchResults(query string) ([]Entry, error) {
	var entries []Entry
	page_json := QueryDehashed(query)
	results := ParseDehashedJson(page_json)
	if !results.Success {
		return entries, errors.New(results.Message)
	}

	entries = results.Entries
	return entries, nil
}

func FilterHasPassword(entries []Entry) []Entry {
	var filtered_entries []Entry
	for _, entry := range entries {
		if entry.Password != "" {
			filtered_entries = append(filtered_entries, entry)
		}
	}

	return filtered_entries
}

func QueryDehashed(query string) []byte {
	username, api_key := getCredentials()

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.dehashed.com/search?query="+query, nil)
	req.SetBasicAuth(username, api_key)
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body

}

func getCredentials() (string, string) {
	username := os.Getenv("DEHASHED_USERNAME")
	key := os.Getenv("DEHASHED_API_KEY")
	return username, key
}
