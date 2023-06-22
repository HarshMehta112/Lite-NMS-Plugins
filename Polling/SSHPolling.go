package Polling

import (
	"MotadataPugins/Utils"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
	"sync"
	"time"
)

func errorDisplay(res map[string]interface{}) {

	bytes, _ := json.Marshal(res)

	fmt.Println(string(bytes))

}

func getSystemInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForSystemInfo, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputOfSystemnfo, _ := sessionForSystemInfo.CombinedOutput(Utils.COMMAND_FOR_SYSTEM_INFO)

	outputOfSystemInfo := string(commandOutputOfSystemnfo)

	systemInfo := strings.Split(outputOfSystemInfo, "\n")

	result["system.name"] = systemInfo[0]

	result["operating.system.name"] = systemInfo[1]

	result["operating.system.version"] = systemInfo[2]
}

func getMemoryInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForMemory, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputOfMemory, _ := sessionForMemory.CombinedOutput(Utils.COMMAND_FOR_MEMORY)

	outputOfMemory := string(commandOutputOfMemory)

	MemoryInfo := strings.Split(outputOfMemory, "\n")

	result["memory.used.percentage"] = MemoryInfo[0]

	result["memory.free.percentage"] = MemoryInfo[1]
}

func getCpuInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForCPUInfo, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputForCpuInfo, _ := sessionForCPUInfo.CombinedOutput(Utils.COMMAND_FOR_CPU)

	outputForCpuInfo := string(commandOutputForCpuInfo)

	splitedNewLineCpuInfo := strings.Split(outputForCpuInfo, "\n")

	splitedBySpaceCpuInfo := strings.Split(splitedNewLineCpuInfo[2], " ")

	result["cpu.user.percentage"] = splitedBySpaceCpuInfo[0]

	result["cpu.system.percentage"] = splitedBySpaceCpuInfo[1]

	result["cpu.idle.percentage"] = splitedBySpaceCpuInfo[3]
}

func getDiskInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForDiskInfo, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputForDisk, _ := sessionForDiskInfo.CombinedOutput(Utils.COMMAND_FOR_DISK)

	outputForDisk := string(commandOutputForDisk)

	splitedByPercentageDiskInfo := strings.Split(outputForDisk, "%")

	result["disk.used.percentage"] = splitedByPercentageDiskInfo[0]

}

func getUptimeInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForUptimeInfo, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputForUptimeInfo, _ := sessionForUptimeInfo.CombinedOutput(Utils.COMMAND_FOR_UPTIME)

	outputForUptimeInfo := string(commandOutputForUptimeInfo)

	result["uptime"] = strings.Split(outputForUptimeInfo, "\n")[0]

}

func getIfConfigInfo(sshClient *ssh.Client, errorList []string, result map[string]interface{}) {

	sessionForIfConfig, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutputForIfConfig, _ := sessionForIfConfig.CombinedOutput(Utils.COMMAND_FOR_IFCONFIG)

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

func PollingSSH(data map[string]interface{}, wg *sync.WaitGroup) {

	defer wg.Done()

	currentTime := time.Now()

	defer func() {

		if r := recover(); r != nil {

			res := make(map[string]interface{})

			res["error"] = r

			errorDisplay(res)
		}
	}()

	sshUser := (data["USERNAME"]).(string)

	sshPassword := (data["PASSWORD"]).(string)

	sshHost := (data["IPADDRESS"]).(string)

	deviceId := data["DEVICEID"]

	config := &ssh.ClientConfig{

		Timeout: 10 * time.Second,

		User: sshUser,

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		Config: ssh.Config{Ciphers: []string{

			"aes256-ctr",
		}},
	}

	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	address := fmt.Sprintf("%s:%d", sshHost, 22)

	sshClient, err := ssh.Dial("tcp", address, config)

	var errorList []string

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	defer sshClient.Close()

	result := make(map[string]interface{})

	if len(errorList) == 0 {

		getSystemInfo(sshClient, errorList, result)

		getMemoryInfo(sshClient, errorList, result)

		getCpuInfo(sshClient, errorList, result)

		getDiskInfo(sshClient, errorList, result)

		getUptimeInfo(sshClient, errorList, result)

		getIfConfigInfo(sshClient, errorList, result)

		result["timestamp"] = currentTime.Format("2006-01-02 15:04:05")

		result["id"] = deviceId

		result["ip"] = sshHost

		var ans []map[string]interface{}

		ans = append(ans, result)

		bytes, _ := json.Marshal(ans)

		fmt.Println(string(bytes))

	} else {

		response := make(map[string]interface{})

		response["error"] = errorList

		errorDisplay(response)

	}

}
