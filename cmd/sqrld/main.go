package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dmmcquay/sqrl"
)

func main() {
	r, err := sqrl.MakeReport()
	if err != nil {
		log.Fatal(err)
	}
	j, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
