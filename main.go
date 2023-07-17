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
			fmt.Println("Fetch the time while polling")
			timestamp := sshconnection.ExtractTimestamp()
			fmt.Printf("Timestamp: %s\n", timestamp)
			fmt.Println("__________________________________")

			fmt.Println("Extract CPU Percentage")
			client, err := sshconnection.EstablishSSHConnection(host,port,username,password)
			if err != nil {
				fmt.Printf("Failed to establish SSH connection: %s\n", err)
		        return	
			}
			defer client.Close()

			// Extract cpu percentage
			commandcpu := "top -bn1 | grep '%Cpu'"
			outputcpu, err := sshconnection.ExecuteCommand(client, commandcpu)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
				return
			}

			cpuPercentage, err := sshconnection.ExtractCPUPercentage(outputcpu)
			if err != nil {
				fmt.Printf("Failed to extract CPU percentage: %s\n", err)
				return
			}
			fmt.Printf("CPU Percentage: %s\n", cpuPercentage)
            fmt.Println("____________________________________________")

			// Fetch Host Name
			fmt.Println("Fetch the host name")
			commandHost := "hostname"
			out, err := sshconnection.ExecuteCommand(client,commandHost)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
			}
			fmt.Printf("Host Name: %s\n", out)
			fmt.Println("__________________________________")

			// Extract os.name
			fmt.Println("Extract osName")
			commandos := "cat /etc/os-release"
			outputos, err := sshconnection.ExecuteCommand(client, commandos)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
				return
			}

			osName, err := sshconnection.ExtractOSName(outputos)
			if err != nil {
				fmt.Printf("Failed to extract OS Name: %s\n", err)
				return
			}
			fmt.Printf("OS Name: %s\n", osName)
			fmt.Println("_________________________________________________")

			// Extract cpu.cores
			fmt.Println("Extract the no. of cpu cores")
			commandCpuCores := "cat /proc/cpuinfo | grep processor"
			outputCpuCores, err := sshconnection.ExecuteCommand(client, commandCpuCores)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
				return
			}

			noOfCpuCores, err := sshconnection.ExtractCPUCoreCount(outputCpuCores)
			if err != nil {
				fmt.Printf("Failed to extract CPU Core Count: %s\n", err)
				return
			}
			fmt.Printf("CPU Core Count: %d\n", noOfCpuCores)
			fmt.Println("_____________________________________")

			// Error : Failed to execute command: Process exited with status 1
			// Extract cpu.io.percent
			// fmt.Println("Extract the CPU IO Percentage")
			// commandIO := "iostat -c | grep '^avg-cpu'"
			// outputIO, err := sshconnection.ExecuteCommand(client, commandIO)
			// if err != nil {
			// 	fmt.Printf("Failed to execute command: %s\n", err)
			// 	return
			// }

			// cpuIOPercentage, err := sshconnection.ExtractCPUIOPercentage(outputIO)
			// if err != nil {
			// 	fmt.Printf("Failed to extract CPU IO Percentage: %s\n", err)
			// 	return
			// }
			// fmt.Printf("CPU IO Percentage: %s\n", cpuIOPercentage)
			// fmt.Println("___________________________________________")

			// Extract memory used bytes
			fmt.Println("Extract memory used in bytes")
			commandMU := "free -b | grep Mem"
			outputMU, err := sshconnection.ExecuteCommand(client, commandMU)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
				return
			}
			
			memoryUsed, err := sshconnection.ExtractMemoryUsed(outputMU)
			if err != nil {
				fmt.Printf("Failed to extract memory used: %s\n", err)
				return
			}
			fmt.Printf("Memory Used: %d bytes\n", memoryUsed)
			fmt.Println("___________________________________________")

			
			//Extract total memory in bytes
			fmt.Println("Extract total memory in bytes")
			//command is same that we use to fetch memory used
			memoryTotal, err := sshconnection.ExtractMemoryTotal(outputMU)
			if err != nil {
				fmt.Printf("Failed to extract memory total: %s\n", err)
				return
			}
			fmt.Printf("Memory Total: %d bytes\n", memoryTotal)
			fmt.Println("___________________________________________")

			// Extract memory used in percentage
			fmt.Println("Extract memory used in percentage")
			commandMUP := "free -m | grep Mem"
			outputMUP, err := sshconnection.ExecuteCommand(client, commandMUP)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
				return
			}
			
			memoryUsedPercentage, err := sshconnection.ExtractMemoryUsedPercentage(outputMUP)
			if err != nil {
				fmt.Printf("Failed to extract memory used percentage: %s\n", err)
				return
			}
			fmt.Printf("Memory Used Percentage: %s \n", memoryUsedPercentage)
			fmt.Println("___________________________________________")

			// Extract system.memory.used.percentage
			fmt.Println("Extract system.memory.used.percentage")
			command := "cat /proc/meminfo"
			output, err := sshconnection.ExecuteCommand(client, command)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
				return
			}
			
			systemMemUsedPercentage, err := sshconnection.ExtractSystemMemoryUsedPercentage(output)
			if err != nil {
				fmt.Printf("Failed to extract system memory used percentage: %s\n", err)
				return
			}
			fmt.Printf("System Memory Used Percentage: %s \n", systemMemUsedPercentage)
			fmt.Println("___________________________________________")

			//Fetch the interface name

			fmt.Println("Fetch the interface name")
			commInter := "ip link show"
			//commInter := "ifconfig | grep 'interface.name'"
			outputInterface, err := sshconnection.ExecuteCommand(client,commInter)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
				return
			}
			interfaceName, err := sshconnection.ExtractInterfaceName(outputInterface)
			if err != nil {
				fmt.Printf("Failed to extract interface name : %s\n", err)
				return
			}
			fmt.Printf("Interface Name: %s\n", interfaceName)
			fmt.Println("________________________________________________")

			// Fetch interface input speed in byte
			fmt.Println("Fetch interface input speed in bytes")
			commandIIS := "ip -s link show lo"
			outputIIS, err := sshconnection.ExecuteCommand(client,commandIIS)
			if err != nil {
				fmt.Printf("Failed to execute command: %s\n", err)
				return
			}
			inputSpeed, err := sshconnection.ExtractInterfaceInputSpeed(outputIIS)
			if err != nil {
				fmt.Printf("Failed to extract interface input speed: %s\n", err)
				return
			}
			fmt.Printf("Interface Input Speed: %s\n", inputSpeed)

	    case "snmp":
			
		case "windows":	
	}
}





