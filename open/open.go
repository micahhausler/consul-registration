package open

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func FindAddress(containerName string) string {
	data, err := ioutil.ReadFile("/etc/hosts")
	check(err)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, containerName) {
			parts := strings.Fields(line)
			return parts[0]
		}
	}
	return ""
}
