package main

import (
	"fmt"
	"os"
)

func main() {

	introduction()
	menu()

	command := command()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("showing Logs...")
	case 0:
		fmt.Println("Exiting, good bye.")
		os.Exit(0)
	default:
		fmt.Println("I don't know this command")
		os.Exit(-1)
	}
}

func introduction() {
	name := "Kevin"
	var version float32 = 1.1
	fmt.Println("Hi, mr.", name)
	fmt.Println("This program is in version", version)
}

func menu() {
	fmt.Println("0- Exit program")
	fmt.Println("1- Start monitorating")
	fmt.Println("2- Show Logs")
}

func command() int {
	var commandSelect int
	fmt.Scan(&commandSelect)
	fmt.Println("The command select has:", commandSelect)
	return commandSelect
}
