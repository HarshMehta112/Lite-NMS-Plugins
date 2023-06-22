package Polling

import (
	"MotadataPugins/Utils"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"sync"
	"time"
)

func errorDisplay(res map[string]interface{}) {

	bytes, _ := json.Marshal(res)

	fmt.Println(string(bytes))

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

		Utils.GetSystemInfo(sshClient, errorList, result)

		Utils.GetMemoryInfo(sshClient, errorList, result)

		Utils.GetCpuInfo(sshClient, errorList, result)

		Utils.GetDiskInfo(sshClient, errorList, result)

		Utils.GetUptimeInfo(sshClient, errorList, result)

		Utils.GetIfConfigInfo(sshClient, errorList, result)

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
