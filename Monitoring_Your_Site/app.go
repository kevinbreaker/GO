package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorations = 3
const delay = 3

func main() {

	introduction()

	for {
		menu()

		command := command()

		switch command {
		case 1:
			monitoring()
		case 2:
			fmt.Println("showing Logs...")
			printLogs()
		case 0:
			fmt.Println("Exiting, good bye.")
			os.Exit(0)
		default:
			fmt.Println("I don't know this command")
			os.Exit(-1)
		}
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

func monitoring() {
	fmt.Println("Monitoring...")
	sites := sitesArchive()
	for i := 0; i < monitorations; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "loaded with success!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "with problem, status:", resp.Status)
		registerLog(site, false)
	}
}

func sitesArchive() []string {
	var sites []string
	archive, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(archive)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	archive.Close()
	return sites
}

func registerLog(site string, status bool) {

	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	archive.WriteString(time.Now().Format("02/01/2018 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	archive.Close()
}

func printLogs() {

	archive, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(archive))

}
