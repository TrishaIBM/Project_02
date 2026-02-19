package main

import "fmt"

func main() {
	var list []string
	var cmd, input string

	fmt.Println("Commands: add, remove, check, list, quit")

	for {
		fmt.Print("Command: ")
		fmt.Scan(&cmd)

		if cmd == "quit" {
			break
		}

		if cmd == "list" {
			fmt.Println("Current list:", list)
			continue
		}

		fmt.Print("Input: ")
		fmt.Scan(&input)

		switch cmd {
		case "add":
			list = append(list, input)
			fmt.Println("Added!")

		case "check":
			found := false
			for _, item := range list {
				if item == input {
					found = true
					break
				}
			}
			if found {
				fmt.Println("Exists!")
			} else {
				fmt.Println("Not found.")
			}

		case "remove":
			for i, item := range list {
				if item == input {

					list = append(list[:i], list[i+1:]...)
					fmt.Println("Removed.")
					break
				}
			}

		default:
			fmt.Println("Unknown command")
		}
	}
}
