// Ping-signaler
package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var ontvpackets1 int

func main() {
	for {
		cmd := exec.Command("netstat", "-s")
		pipe, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		var stukjes []string
		scanner := bufio.NewScanner(pipe)
		for scanner.Scan() {
			tekst := scanner.Text()
			if strings.Contains(tekst, "Messages") {
				tekst = strings.Trim(tekst, " ")
				stukjes = strings.Split(tekst, " ")
				break
			}
		}
		ontvicmpv4 := ""
		for _, stukje := range stukjes {
			if stukje != "Messages" && stukje != "" {
				ontvicmpv4 = stukje
				if ontvpackets1, err = strconv.Atoi(ontvicmpv4); err != nil {
					panic(err)
				}
				break
			}
		}
		fmt.Println("U bent", ontvpackets1, "keer gepingd")
		time.Sleep(3 * time.Second)
	}
}
