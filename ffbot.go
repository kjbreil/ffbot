package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kjbreil/ffbot/fantasy"
)

func main() {
	// for now read from console and send input to function printing output

	// open a reader from stdin
	r := bufio.NewReader(os.Stdin)

	// forever for loop
	for {
		// use fmt instead of log
		fmt.Print("::")
		// assign texxt when a newline is detected
		text, err := r.ReadString('\n')
		// drop the newline from text
		text = strings.Replace(text, "\n", "", -1)

		if err != nil {
			// use log since its a failure and timestamp is desired
			log.Fatal(err)
		}
		// hold me
		fmt.Println(fantasy.Holder(text))
	}
}
