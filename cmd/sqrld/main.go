package main

import (
	"fmt"
	"log"

	"github.com/dmmcquay/sqrl"
)

func main() {
	//c, err := sqrl.GetCPUInfo()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//c.LogWriter()

	//r, o, s, err := sqrl.RamOSInfo()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//r.LogWriter()
	//o.LogWriter()
	//s.LogWriter()

	i, err := sqrl.GetInterfaces()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)

	//r, err := sqrl.MakeReport()
	//if err != nil {
	//	log.Fatal(err)
	//}

	////a, e := sqrl.GetInterfaces()
	////if e != nil {
	////	log.Fatal(e)
	////}
	////fmt.Println(a)
	//r.LogWriter()
}
