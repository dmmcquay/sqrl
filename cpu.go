package sqrl

import (
	"encoding/json"
	"log"
	"log/syslog"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/cpu"
)

type CPUInfo struct {
	CPU   int    `json:"cpu"`
	Cores int32  `json:"cores"`
	Speed string `json:"speed"`
	Model string `json:"model"`
}

func GetCPUInfo() (CPUInfo, error) {
	info, err := cpu.Info()
	if err != nil {
		return CPUInfo{}, err
	}
	mInfo := strings.Split(info[0].ModelName, "@") // mInfo contains speed and model
	c := CPUInfo{
		CPU:   runtime.NumCPU(),
		Cores: info[0].Cores,
		Speed: strings.Trim(mInfo[1], " "),
		Model: strings.Trim(mInfo[0], " "),
	}
	return c, nil
}

func (c CPUInfo) String() string {
	j, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(j)
}

func (c *CPUInfo) LogWriter() {
	l, e := syslog.New(syslog.LOG_NOTICE, "sqrl-CPUInfo")
	if e == nil {
		log.SetOutput(l)
	}

	log.Print(c)
}
