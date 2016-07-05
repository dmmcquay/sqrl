package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dmmcquay/sqrl"
)

func main() {
	c, err := sqrl.GetCPUInfo()
	if err != nil {
		log.Fatal(err)
	}
	j, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
