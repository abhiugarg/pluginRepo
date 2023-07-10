package main

import (
	"encoding/json"
	"fmt"
	"tasking/sshconnection"
)

func main() {

	jsonData := "{\"host\": \"95.216.211.180\",\"port\": 22,\"username\": \"fcaps\",\"password\": \"RevDau@123\",\"timeout\": 500,\"device.type\": \"linux\"}"

	data := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonData),&data)
	if err != nil {
		fmt.Println("Failed to convert JSON ",err)
		return
	}
	// fmt.Println(data)
	// fmt.Printf("%T",data["port"])
	// fmt.Printf("%T",data["device.type"])
	// fmt.Printf("%T",data["host"])
	// fmt.Printf("%T",data["username"])
	// fmt.Printf("%T",data["password"])
	
	port := int(data["port"].(float64))
	host, _ := data["host"].(string)
	username, _ := data["username"].(string)
	password, _ := data["password"].(string)

	switch deviceType := data["device.type"].(string); deviceType {
		case "linux":
			client, err := sshconnection.EstablishSSHConnection(host,port,username,password)
			if err != nil {
				fmt.Printf("Failed to establish SSH connection: %s\n", err)
		        return	
			}
			defer client.Close()

			command := "top -bn1 | grep '%Cpu'"
			output, err := sshconnection.ExecuteCommand(client, command)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
				return
			}

			cpuPercentage, err := sshconnection.ExtractCPUPercentage(output)
			if err != nil {
				fmt.Printf("Failed to extract CPU percentage: %s\n", err)
				return
			}
			fmt.Printf("CPU Percentage: %s\n", cpuPercentage)

	    case "snmp":
			


		case "windows":	
	}
}





