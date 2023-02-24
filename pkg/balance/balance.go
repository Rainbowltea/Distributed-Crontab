package balance

import (
	"os/exec"
	"strconv"

	"Distributed-Crontab/pkg/log"
)

//获取节点CPU和内存使用信息
func GetNodeUsage() int {

	cpuUsage, err := exec.Command("top", "-b", "-n1").Output()
	if err != nil {
		log.LogrusObj.Error(err)
	}
	cpuUsageFloat, err := strconv.ParseFloat(string(cpuUsage), 64)
	memoryUsage, err := exec.Command("free", "-m").Output()
	if err != nil {
		log.LogrusObj.Error(err)
	}
	memoryUsageFloat, err := strconv.ParseFloat(string(memoryUsage), 64)
	weightedUsage := (cpuUsageFloat * 0.5) + (memoryUsageFloat * 0.5)
	return int(weightedUsage)
}
