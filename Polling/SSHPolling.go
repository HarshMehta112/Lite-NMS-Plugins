package Polling

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
	"sync"
	"time"
)

func PollingSSH(data map[string]interface{}, wg *sync.WaitGroup) {

	defer wg.Done()

	const commandForCPU = "mpstat |awk  '{if ($4 != \"CPU\") print $5 \" \" $7 \" \" $8 \" \" $14}'"

	const commandForMemory = "free -m | grep \"Mem\" | awk '{print $3/$2*100}' && free -m | grep \"Mem\" | awk '{print $4/$2*100}'"

	const commandForDisk = "df . | awk 'NR==2 {print $5}'"

	const commandForUptime = "uptime -p"

	const commandForSystemInfo = "hostnamectl | grep \"Static\"| awk '{print $3}' && hostnamectl | grep \"Operating\"| awk '{print $3}' && hostnamectl |grep \"Operating\"| awk '{print $4}'"

	currentTime := time.Now()

	defer func() {

		if r := recover(); r != nil {

			res := make(map[string]interface{})

			res["error"] = r

		}
	}()

	sshUser := (data["USERNAME"]).(string)

	sshPassword := (data["PASSWORD"]).(string)

	sshHost := (data["IPADDRESS"]).(string)

	deviceId := (data["DEVICEID"])

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

	/*Firing first command and manipulating th output*/

	session1, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutput1, _ := session1.CombinedOutput(commandForSystemInfo)

	output := string(commandOutput1)

	system1 := strings.Split(output, "\n")

	result["system.name"] = system1[0]

	result["operating.system.name"] = system1[1]

	result["operating.system.version"] = system1[2]

	/*Ending of first Command*/

	/*============================================================================================*/

	/*Firing second Command*/

	session2, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutput2, _ := session2.CombinedOutput(commandForMemory)

	output2 := string(commandOutput2)

	system2 := strings.Split(output2, "\n")

	result["memory.used.percentage"] = system2[0]

	result["memory.free.percentage"] = system2[1]

	/*Ending of second Command*/

	/*============================================================================================*/

	/*Firing third Command*/

	session3, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutput3, _ := session3.CombinedOutput(commandForCPU)

	output3 := string(commandOutput3)

	res3 := strings.Split(output3, "\n")

	system3 := strings.Split(res3[2], " ")

	result["cpu.user.percentage"] = system3[0]

	result["cpu.system.percentage"] = system3[1]

	result["cpu.idle.percentage"] = system3[3]

	/*Ending of third Command*/

	/*============================================================================================*/

	/*Starting of fourth command*/

	session4, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutput4, _ := session4.CombinedOutput(commandForDisk)

	output4 := string(commandOutput4)

	res4 := strings.Split(output4, "%")

	result["disk.used.percentage"] = res4[0]

	/*Ending of fourth Command*/

	/*============================================================================================*/

	/*Starting of fifth command*/

	session5, err := sshClient.NewSession()

	if err != nil {

		errorList = append(errorList, err.Error())
	}

	commandOutput5, _ := session5.CombinedOutput(commandForUptime)

	output5 := string(commandOutput5)

	res5 := strings.Split(output5, "\n")

	result["uptime"] = res5[0]

	result["timestamp"] = currentTime.Format("2006-01-02 15:04:05")

	result["id"] = deviceId

	result["ip"] = sshHost

	/*============================================================================================*/

	/*Ending of fifth command*/

	var ans []map[string]interface{}

	ans = append(ans, result)

	bytes, _ := json.Marshal(ans)

	fmt.Println(string(bytes))

}
