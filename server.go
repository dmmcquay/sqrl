package sqrl

import (
	"log"
	"net/http"
)

func com(w http.ResponseWriter, req *http.Request) {
	c, err := GetCPUInfo()
	if err != nil {
		log.Fatal(err)
	}

	c.LogWriter()

}

func ram(w http.ResponseWriter, req *http.Request) {
	r, _, _, err := RamOSInfo()
	if err != nil {
		log.Fatal(err)
	}

	r.LogWriter()

}

func ops(w http.ResponseWriter, req *http.Request) {
	_, o, _, err := RamOSInfo()
	if err != nil {
		log.Fatal(err)
	}

	o.LogWriter()

}

func swap(w http.ResponseWriter, req *http.Request) {
	_, _, s, err := RamOSInfo()
	if err != nil {
		log.Fatal(err)
	}

	s.LogWriter()

}

//func network(w http.ResponseWriter, req *http.Request) {
//	i, err := GetInterfaces()
//	if err != nil {
//		log.Fatal(err)
//	}

//}

func all(w http.ResponseWriter, req *http.Request) {
	c, err := GetCPUInfo()
	if err != nil {
		log.Fatal(err)
	}
	r, o, s, err := RamOSInfo()
	if err != nil {
		log.Fatal(err)
	}
	//i, err := GetInterfaces()
	//if err != nil {
	//	log.Fatal(err)
	//}

	c.LogWriter()
	r.LogWriter()
	o.LogWriter()
	s.LogWriter()

}
