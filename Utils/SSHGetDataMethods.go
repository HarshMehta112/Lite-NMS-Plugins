package Utils

import (
	"golang.org/x/crypto/ssh"
	"strings"
)

func GetSystemInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForSystemInfo, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputOfSystemnfo, _ := sessionForSystemInfo.CombinedOutput(COMMAND_FOR_SYSTEM_INFO)

	outputOfSystemInfo := string(commandOutputOfSystemnfo)

	systemInfo := strings.Split(outputOfSystemInfo, "\n")

	result["system.name"] = systemInfo[0]

	result["operating.system.name"] = systemInfo[1]

	result["operating.system.version"] = systemInfo[2]
}

func GetMemoryInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForMemory, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputOfMemory, _ := sessionForMemory.CombinedOutput(COMMAND_FOR_MEMORY)

	outputOfMemory := string(commandOutputOfMemory)

	MemoryInfo := strings.Split(outputOfMemory, "\n")

	result["memory.used.percentage"] = MemoryInfo[0]

	result["memory.free.percentage"] = MemoryInfo[1]
}

func GetCpuInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForCPUInfo, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputForCpuInfo, _ := sessionForCPUInfo.CombinedOutput(COMMAND_FOR_CPU)

	outputForCpuInfo := string(commandOutputForCpuInfo)

	splitedNewLineCpuInfo := strings.Split(outputForCpuInfo, "\n")

	splitedBySpaceCpuInfo := strings.Split(splitedNewLineCpuInfo[2], " ")

	result["cpu.user.percentage"] = splitedBySpaceCpuInfo[0]

	result["cpu.system.percentage"] = splitedBySpaceCpuInfo[1]

	result["cpu.idle.percentage"] = splitedBySpaceCpuInfo[3]
}

func GetDiskInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForDiskInfo, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputForDisk, _ := sessionForDiskInfo.CombinedOutput(COMMAND_FOR_DISK)

	outputForDisk := string(commandOutputForDisk)

	splitedByPercentageDiskInfo := strings.Split(outputForDisk, "%")

	result["disk.used.percentage"] = splitedByPercentageDiskInfo[0]

}

func GetUptimeInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForUptimeInfo, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputForUptimeInfo, _ := sessionForUptimeInfo.CombinedOutput(COMMAND_FOR_UPTIME)

	outputForUptimeInfo := string(commandOutputForUptimeInfo)

	result["uptime"] = strings.Split(outputForUptimeInfo, "\n")[0]

}

func GetIfConfigInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForIfConfig, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputForIfConfig, _ := sessionForIfConfig.CombinedOutput(COMMAND_FOR_IFCONFIG)

	outputForUIfConfigInfo := string(commandOutputForIfConfig)

	splitedByNewLineForIfConfig := strings.Split(outputForUIfConfigInfo, "\n")

	splitedBySpaceForIfConfig := strings.Split(splitedByNewLineForIfConfig[0], " ")

	result["lo.name"] = splitedBySpaceForIfConfig[0]

	result["lo.RX.bytes"] = splitedBySpaceForIfConfig[1]

	result["lo.TX.bytes"] = splitedBySpaceForIfConfig[2]

	splitedBySpaceForIfConfig = strings.Split(splitedByNewLineForIfConfig[1], " ")

	result["en.name"] = splitedBySpaceForIfConfig[0]

	result["en.RX.bytes"] = splitedBySpaceForIfConfig[1]

	result["en.TX.bytes"] = splitedBySpaceForIfConfig[2]

	splitedBySpaceForIfConfig = strings.Split(splitedByNewLineForIfConfig[2], " ")

	result["wl.name"] = splitedBySpaceForIfConfig[0]

	result["wl.RX.bytes"] = splitedBySpaceForIfConfig[1]

	result["wl.TX.bytes"] = splitedBySpaceForIfConfig[2]

}
