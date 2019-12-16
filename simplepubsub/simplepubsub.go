package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var isPublishing bool
	var messages []string
	for {
		fmt.Println("publisher or subscriber? [p/s]")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			t := scanner.Text()

			if t == "s" {
				for _, m := range messages {
					fmt.Println(m)
				}
			}

			if isPublishing {
				messages = append(messages, t)
			}

			if t == "p" {
				isPublishing = true
				fmt.Println("you can now publish")
			}

		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}
}
