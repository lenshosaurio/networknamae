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
		output = runCommand("cmd", "/C", "netsh", "wlan", "show", "interfaces")
		lines := strings.Split(output, "\r\n")
		_, line := findElement("Profile", lines)
		if len(strings.Split(line, ": ")) > 1 {
			ssid = strings.Split(line, ": ")[1]
		}
	case "darwin":
		command = "networksetup"
		output = runCommand(command, "-listallhardwareports")
		lines := strings.Split(output, "\n")
		position, _ := findElement("Wi-Fi", lines)
		stringFound := lines[position+1]
		card := strings.Split(stringFound, " ")[1]
		output = runCommand(command, "-getairportnetwork", card)
		ssid = strings.Split(output, " ")[len(strings.Split(output, " "))-1]
	case "linux":
		ssid = ""
	default:
		ssid = ""
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
		break
	case 6:
		result, err = exec.Command(arguments[0], arguments[1], arguments[2], arguments[3], arguments[4], arguments[5]).Output()
		break
	default:
		break
	}
	check(err, "Running cmd"+arguments[0])
	return string(result)
}
func findElement(element string, output []string) (int, string) {
	for index, row := range output {
		if strings.Contains(row, element) {
			return index, row
		}
	}
	return -1, ""
}
func check(err error, origin string) {
	if err != nil {
		panic(err.Error())
	}
}
