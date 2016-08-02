package sqrl

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex = &sync.Mutex{}

func com(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	c, err := GetCPUInfo()
	if err != nil {
		log.Fatal(err)
	}
	mutex.Unlock()

	c.LogWriter()

	fmt.Println(c)
}

func ros(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	r, o, s, err := RamOSInfo()
	if err != nil {
		log.Fatal(err)
	}
	mutex.Unlock()

	r.LogWriter()
	o.LogWriter()
	s.LogWriter()

	fmt.Println(r)
	fmt.Println(o)
	fmt.Println(s)
}

func network(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	i, err := GetInterfaces()
	if err != nil {
		log.Fatal(err)
	}
	mutex.Unlock()

	fmt.Println(i)
}

func all(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	c, err := GetCPUInfo()
	if err != nil {
		log.Fatal(err)
	}
	r, o, s, err := RamOSInfo()
	if err != nil {
		log.Fatal(err)
	}
	i, err := GetInterfaces()
	if err != nil {
		log.Fatal(err)
	}
	mutex.Unlock()

	c.LogWriter()
	r.LogWriter()
	o.LogWriter()
	s.LogWriter()

	fmt.Println(c)
	fmt.Println(r)
	fmt.Println(o)
	fmt.Println(s)
	fmt.Println(i)
}
