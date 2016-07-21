package sqrl

import (
	"encoding/json"
	"log"
	"log/syslog"
)

type Report struct {
	Ram     `json:"Ram"`
	Swap    `json:"Swap"`
	OS      `json:"OS"`
	CPUInfo `json:"CPUInfo"`
	Inter   []Interface `json:"Interface"`
}

func MakeReport() (Report, error) {
	r, o, s, err := RamOSInfo()
	if err != nil {
		return Report{}, err
	}
	i, err := GetInterfaces()
	if err != nil {
		return Report{}, err
	}
	c, err := GetCPUInfo()
	if err != nil {
		return Report{}, err
	}
	rp := Report{
		Ram:     r,
		Swap:    s,
		OS:      o,
		CPUInfo: c,
		Inter:   i,
	}

	return rp, err
}

func (r Report) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

func (r *Report) LogWriter() {
	l, e := syslog.New(syslog.LOG_NOTICE, "Report")
	if e == nil {
		log.SetOutput(l)
	}

	log.Print(r)
}
