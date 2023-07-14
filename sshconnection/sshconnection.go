package sshconnection

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func ExtractTimestamp() string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return timestamp
}

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
		fmt.Printf("Failed to execute command: %s. Error: %s\n", err, output)
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


func ExtractOSName(output string) (string, error) {
	pattern := `NAME="(.+)"` // Assuming the OS name is present in the NAME field of the output
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(output)

	fmt.Printf("Output: %s\n", output)

	if len(match) != 2 {
		return "", fmt.Errorf("unable to extract OS name")
	}
	osName := strings.TrimSpace(match[1])
	return osName, nil
}

func ExtractCPUCoreCount(output string) (int, error) {
	pattern := `processor\s+:\s+(\d+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(output, -1)

	fmt.Printf("Output: %s\n", output)

	if len(matches) == 0 {
		return 0, fmt.Errorf("unable to extract CPU core count")
	}
	coreCount := len(matches)
	return coreCount, nil
}

func ExtractCPUIOPercentage(output string) (string, error) {
	pattern := `IO:\s+(\d+\.\d+)\s+us`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(output)

	fmt.Printf("Output: %s\n", output)

	if len(match) != 2 {
		return "", fmt.Errorf("unable to extract CPU IO percentage")
	}
	cpuIOPercentageStr := match[1]
	cpuIOPercentage, err := strconv.ParseFloat(cpuIOPercentageStr, 64)

	if err != nil {
		return "", fmt.Errorf("failed to convert CPU IO percentage")
	}
	return fmt.Sprintf("%.2f", cpuIOPercentage), nil
}


func ExtractMemoryUsed(output string) (int64, error) {
	pattern := `Mem:\s+\d+\s+(\d+)\s+\d+`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(output)

	fmt.Printf("Output: %s\n", output)
	// fmt.Println(len(match))
	// fmt.Println(match[0])
	// fmt.Println(match[1])

	if len(match) != 2 {
		return 0, fmt.Errorf("unable to extract memory used")
	}

	memoryUsedStr := match[1]
	memoryUsed, err := strconv.ParseInt(memoryUsedStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert memory used")
	}
	return memoryUsed, nil
}

func ExtractMemoryTotal(output string) (int64,error)  {
	pattern := `Mem:\s+(\d+)\s+`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(output)

	fmt.Printf("Output: %s\n", output)
	// fmt.Println(len(match))
	// fmt.Println(match[0])
	// fmt.Println(match[1])
	
	if len(match) != 2 {
		return 0, fmt.Errorf("unable to extract memory total")
	}
	memoryTotalStr := match[1]
	memoryTotal, err := strconv.ParseInt(memoryTotalStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert memory total")
	}
	return memoryTotal, nil
}

func ExtractMemoryUsedPercentage(output string) (string, error) {
	pattern := `Mem:\s+\d+\s+(\d+)\s+\d+`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(output)

	fmt.Printf("Output: %s\n", output)

	if len(match) != 2 {
		return "", fmt.Errorf("unable to extract memory used percentage")
	}

	memoryUsedStr := strings.TrimSpace(match[1])
	memoryUsed, err := strconv.ParseFloat(memoryUsedStr, 64)
	if err != nil {
		return "", fmt.Errorf("failed to convert memory used percentage")
	}

	return fmt.Sprintf("%.2f", memoryUsed), nil
}

func ExtractSystemMemoryUsedPercentage(output string) (string, error) {
	pattern := `MemTotal:\s+(\d+) kB\n.*\nMemAvailable:\s+(\d+) kB`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(output)

	fmt.Printf("Output: %s\n", output)
	fmt.Println(len(match))
	fmt.Println(match[0])
	fmt.Println(match[1])
	fmt.Println(match[2])

	if len(match) != 3 {
		return "", fmt.Errorf("unable to extract memory used percentage")
	}

	memTotalStr := match[1]
	memAvailableStr := match[2]

	memTotal, err := strconv.ParseFloat(memTotalStr, 64)
	if err != nil {
		return "", fmt.Errorf("failed to convert memory total")
	}

	memAvailable, err := strconv.ParseFloat(memAvailableStr, 64)
	if err != nil {
		return "", fmt.Errorf("failed to convert memory available")
	}

	memUsedPercentage := ((memTotal - memAvailable) / memTotal) * 100

	return fmt.Sprintf("%.2f", memUsedPercentage), nil
}

func ExtractInterfaceName(output string) (string, error) {
	pattern := `^\d+:\s+([^:@]+)`
	//pattern := `^\d+:\s+([^\s:]+)`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(output)

	fmt.Printf("Output: %s\n", output)
	fmt.Println(match)
	// fmt.Println(match[0])

	if len(match) != 2 {
		return "", fmt.Errorf("unable to extract interface name")
	}
	interfaceName := match[1]
	return interfaceName, nil
}

func ExtractInterfaceInputSpeed(output string) (string, error) {
	fmt.Println(output)

	lines := strings.Split(output, "\n")
	for i, line := range lines {
		if strings.Contains(line, "RX:  bytes") {
			// Check if the next line contains the input speed
			if i+1 < len(lines) {
				fields := strings.Fields(lines[i+1])
				if len(fields) >= 1 {
					inputSpeed := fields[0]
					return inputSpeed, nil
				}
			}
		}
	}

	return "", fmt.Errorf("unable to extract interface input speed")
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