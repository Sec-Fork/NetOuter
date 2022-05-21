package checktcp

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ChecktcpP(targets_ports_path string) {
	var targetslist []string
	file, err := os.Open(targets_ports_path)
	if err != nil {
		fmt.Println(targets_ports_path, " open err.")
	}
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			targetslist = append(targetslist, line)
		}
	}
	for _, targets := range targetslist {
		ip_ports := "45.79.204.144" + targets
		Checktcp(ip_ports)
	}
}