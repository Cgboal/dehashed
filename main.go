package main

import (
	"gitlab.com/cgboal/dehashed/lib"
	"flag"
	"fmt"
	"encoding/json"
)


func main() {

	output_json := flag.Bool("oJ", false, "Output in JSON format")
	show_all := flag.Bool("all", false, "Output entries without plaintext passwords")
	flag.Parse()

	query := flag.Arg(0)

	var results []dehashed.Entry

	//results := dehashed.FetchAll(query)
	for i := 20; i < 40; i++ {
		new_results := dehashed.FetchPage(query, i)
		results = append(results, new_results...)
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
