package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func persistMessage(message string) error {
	f, err := os.Create("./tmp/messages")
	if err != nil {
		return err
	}

	// write message to file
	bs := []byte(message + "\n")
	_, err = f.Write(bs)
	if err != nil {
		return err
	}

	return nil
}

func readMessages() ([]string, error) {
	m, err := ioutil.ReadFile("./tmp/messages")
	if err != nil {
		return nil, err
	}

	messages := strings.Split(string(m), "\n")

	return messages[:len(messages)-1], nil
}

func main() {

	var isPublishing bool
	// var messages []string
	for {

		// prompt user
		fmt.Println("publisher or subscriber? [p/s]")
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			t := scanner.Text()

			// print messages
			if t == "s" {

				// get messages from file
				messages, err := readMessages()
				if err != nil {
					log.Fatal(err)
				}
				for _, m := range messages {
					fmt.Println(m)
				}
			}

			// add messages to list if publishing mode active
			if isPublishing {
				switch t {
				case "":
				case "p":
				case "s":
					continue
				default:
					// messages = append(messages, t)
					persistMessage(t)
				}
			}

			// initialize publishing mode
			if t == "p" {
				isPublishing = true
				fmt.Println("you can now publish")
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
