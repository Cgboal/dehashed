package dehashed

import (
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Results struct {
	Entries []Entry
}
type Entry struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Hash string `json:"hashed_password"`
	Name string `json:"name"`
	Source string `json:"obtained_from"`
}

func ParseDehashedJson(json_data []byte) []Entry {
	results := &Results{}
	err := json.Unmarshal(json_data, results)
	if err != nil {
		log.Fatal(err)
	}
	return results.Entries

}

func FetchPage(query string) []Entry {
	page_json := QueryDehashed(query)
	entries := ParseDehashedJson(page_json)
	return entries
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
	req, err := http.NewRequest("GET", "https://api.dehashed.com/search?query=" + query, nil)
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
