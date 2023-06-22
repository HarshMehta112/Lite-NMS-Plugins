package SSHDiscovery

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
)

func Discovery(data map[string]interface{}) []map[string]interface{} {

	const command = "hostname"

	var array []map[string]interface{}

	defer func() {

		if r := recover(); r != nil {

			res := make(map[string]interface{})

			res["error"] = r

		}
	}()

	sshUser := (data["username"]).(string)

	sshPassword := (data["password"]).(string)

	sshHost := (data["ip"]).(string)

	config := &ssh.ClientConfig{

		Timeout: 5 * time.Second,

		User: sshUser,

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		Config: ssh.Config{Ciphers: []string{

			"aes256-ctr",
		}},
	}

	var result = make(map[string]interface{})

	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	address := fmt.Sprintf("%s:%d", sshHost, 22)

	sshClient, clientErr := ssh.Dial("tcp", address, config)

	var errorList []string

	if clientErr != nil {

		result[sshHost] = "fail"

		array = append(array, result)

		errorList = append(errorList, clientErr.Error())

		return array
	}

	defer sshClient.Close()

	session, err := sshClient.NewSession()

	if err != nil {

		result[sshHost] = "fail"

		array = append(array, result)

		errorList = append(errorList, err.Error())

		return array
	}

	commandOutput, errCommandFire := session.CombinedOutput(command)

	if errCommandFire != nil {

		result[sshHost] = "fail"

		array = append(array, result)

		errorList = append(errorList, err.Error())

		return array
	}

	output := string(commandOutput)

	if output == "" {
		result[sshHost] = "fail"
	} else {
		result[sshHost] = "success"
	}

	array = append(array, result)

	return array
}
