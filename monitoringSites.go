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

const timesOfMonitoraments = 4
const delay = 5

func main() {
	showIntroduction()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
			printOutLogs()
		case 0:
			fmt.Println("Exit Software")
			os.Exit(0)
		default:
			fmt.Println("Command invalid.")
			os.Exit(-1)

		}
	}
}

func showIntroduction() {
	name := "D"
	version := 1.1
	fmt.Println("Hello, Sr", name)
	fmt.Println("This Software are in version", version)
}

func showMenu() {
	fmt.Println("1- Start monitoraments!")
	fmt.Println("2- Show Logs!")
	fmt.Println("0- Exit Software!")
}

func readCommand() int {
	var commandLooked int
	fmt.Scan(&commandLooked)
	fmt.Println("Command executed", commandLooked)
	fmt.Println("")

	return commandLooked
}

func startMonitoring() {
	fmt.Println("monitoring...")

	sites := readSitesArchive()

	for i := 0; i < timesOfMonitoraments; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ": ", site)
			testSites(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSites(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Occurred a error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "has been uploaded successfully!")
		registerLogs(site, true)
	} else {
		fmt.Println("Site:", site, "it has problems. StatusCode",
			resp.StatusCode)
		registerLogs(site, false)
	}

}

func readSitesArchive() []string {

	var sites []string

	archive, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Occurred a error:", err)
	}

	Reader := bufio.NewReader(archive)

	for {
		line, err := Reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	archive.Close()
	return sites
}

func registerLogs(site string, status bool) {

	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	archive.WriteString(time.Now().Format("01/02/2006 15:04:05") + " - " + site + "- online:" + strconv.FormatBool(status) + "\n")

	archive.Close()
}

func printOutLogs() {
	archive, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(archive))
}
