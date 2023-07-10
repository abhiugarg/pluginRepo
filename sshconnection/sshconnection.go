package sshconnection

import (
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/crypto/ssh"
)

func EstablishSSHConnection(serverAddr string, port int, username string, password string) (*ssh.Client, error) {

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d",serverAddr,port),config)
	if err != nil {
		return nil,err	
	}
	return client,nil
}


func ExecuteCommand(client *ssh.Client,command string) (string,error)  {
	session,err := client.NewSession()
	if err != nil {
		return "",err
	}
	defer session.Close()

	output,err := session.Output(command)
	if err != nil {
		return "",err
	}
	return string(output),nil
}

func ExtractCPUPercentage(output string) (string,error)  {

	pattern := `%Cpu\(s\):\s+(\d+\.\d+)\sus`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(output)

	fmt.Printf("Output: %s\n", output)

	if len(match) != 2 {
		//fmt.Printf("Output: %s\n", output)
		return "", fmt.Errorf("unable to extract CPU percentage")
	}
	cpuPercentageStr := match[1]
	cpuPercentage, err := strconv.ParseFloat(cpuPercentageStr, 64)

	if err != nil {
		return "", fmt.Errorf("failed to convert CPU percentage")
	}
	return fmt.Sprintf("%.2f", cpuPercentage), nil
	
}

// func main() {

// 	serverAddr := "95.216.211.180"
// 	port := 22
// 	username := "fcaps"
// 	password := "RevDau@123"

// 	client, err := EstablishSSHConnection(serverAddr, port, username, password)
// 	if err != nil {
// 		fmt.Printf("Failed to establish SSH connection: %s\n", err)
// 		return
// 	}
// 	defer client.Close()

// 	command := "top -bn1 | grep '%Cpu'"
// 	output, err := ExecuteCommand(client, command)
// 	if err != nil {
// 		fmt.Printf("Failed to execute command: %s\n", err)
// 		return
// 	}

// 	cpuPercentage, err := ExtractCPUPercentage(output)
// 	if err != nil {
// 		fmt.Printf("Failed to extract CPU percentage: %s\n", err)
// 		return
// 	}

// 	fmt.Printf("CPU Percentage: %s\n", cpuPercentage)




	// serverAddr := "95.216.211.180"
	// port := 22
	// username := "fcaps"
	// password := "RevDau@123"

	
	// ticker := time.NewTicker(5 * time.Minute)

	
	// go func() {
	// 	for range ticker.C {
	// 		client, err := establishSSHConnection(serverAddr, port, username, password)
	// 		if err != nil {
	// 			fmt.Printf("Failed to establish SSH connection: %s\n", err)
	// 			continue
	// 		}
	// 		defer client.Close()

	// 		command := "top -bn1 | grep '%Cpu'"
	// 		output, err := executeCommand(client, command)
	// 		if err != nil {
	// 			fmt.Printf("Failed to execute command: %s\n", err)
	// 			continue
	// 		}

	// 		cpuPercentage, err := extractCPUPercentage(output)
	// 		if err != nil {
	// 			fmt.Printf("Failed to extract CPU percentage: %s\n", err)
	// 			continue
	// 		}

	// 		fmt.Printf("CPU Percentage: %s\n", cpuPercentage)
	// 	}
	// }()

	// // Wait indefinitely
	// select {}
//}