package sqrl

import (
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
		Speed: strings.Trim(mInfo[1], " "), //here
		Model: strings.Trim(mInfo[0], " "), //here
	}
	return c, nil
}
