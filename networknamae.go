package networknamae

import (
	"os/exec"
	"runtime"
	"strings"
)

func SSID() string {
	os := runtime.GOOS
	var ssid string
	var output string
	var command string
	switch os {
	case "windows":
		ssid = ""
		break
	case "darwin":
		command = "networksetup"
		output = runCommand(command, "-listallhardwareports")
		lines := strings.Split(output, "\n")
		var stringFound string
		for index, line := range lines {
			if strings.Contains(line, "Wi-Fi") {
				stringFound = lines[index+1]
				break
			}
		}
		card := strings.Split(stringFound, " ")[1]
		output = runCommand(command, "-getairportnetwork", card)
		ssid = strings.Split(output, " ")[len(strings.Split(output, " "))-1]
		break
	case "linux":
		ssid = ""
		break
	default:
		ssid = ""
		break
	}
	return ssid
}

func runCommand(arguments ...string) string {
	var result []byte
	var err error
	switch len(arguments) {
	case 1:
		result, err = exec.Command(arguments[0]).Output()
		break
	case 2:
		result, err = exec.Command(arguments[0], arguments[1]).Output()
		break
	case 3:
		result, err = exec.Command(arguments[0], arguments[1], arguments[2]).Output()
	default:
		break
	}
	check(err, "Running cmd"+arguments[0])
	return string(result)
}

func check(err error, origin string) {
	if err != nil {
		panic(err.Error())
	}
}
