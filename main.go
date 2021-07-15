package main

import (
	"github.com/cgboal/dehashed/lib"
	"flag"
	"fmt"
	"log"
	"encoding/json"
)


func main() {

	output_json := flag.Bool("oJ", false, "Output in JSON format")
	show_all := flag.Bool("all", false, "Output entries without plaintext passwords")
	flag.Parse()

	query := flag.Arg(0)

	results, err := dehashed.FetchResults(query)
	if err != nil {
		log.Fatal(err)
	}
	if *show_all == false {
		results = dehashed.FilterHasPassword(results)
	}

	if *output_json == true {
		json_results, _ := json.Marshal(results)
		fmt.Println(string(json_results))
	} else {
		for _, result := range results {
			fmt.Printf("%s:%s\n", result.Email, result.Password)
		}
	}
}
