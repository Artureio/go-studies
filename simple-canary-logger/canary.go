package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func main() {
	startMenu()
}

func greetings() {
	name := "artur"
	versao := 1.1

	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", versao)
}

func startMenu() {
	fmt.Println("1- start monitoring")
	fmt.Println("2- show logs")
	fmt.Println("0- exit")

	var readCommand int
	fmt.Scan(&readCommand)
	fmt.Println("O comando foi o numero:", readCommand)
	switch readCommand {
	case 1:
		for {
			iniciarMonitoramento()
		}
	case 2:
		showLogs()
	case 0:
		fmt.Println("exiting...")
		os.Exit(0)
	default:
		fmt.Println("Out of range")

	}
}

func readFiles() []string {
	sites := []string{}
	file, err := os.Open("sites.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sites = append(sites, scanner.Text())
	}

	checkError(err)

	fmt.Println(sites)
	file.Close()
	return sites
}

func iniciarMonitoramento() []int {
	siteResponse := []int{}
	sitenames := []string{}
	for _, site := range readFiles() {
		sitenames = append(sitenames, site)
	}
	fmt.Println("monitoring...")
	for {
		for i, site := range sitenames {
			resp, _ := http.Get(site)
			siteResponse = append(siteResponse, resp.StatusCode)
			defer resp.Body.Close()

			logger(site, siteResponse[i])
		}
	}
}

func logger(site string, resp int) {
	timestamp := time.Now().Format("02/01/2006 15:04:05")
	file, err := os.OpenFile("canary.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	checkError(err)

	if resp >= 500 {
		file.WriteString(timestamp + " | " + site + " is down" + " status code: " + fmt.Sprint(resp) + "\n")
		fmt.Println("Site", site, "is down", "status code: ", resp)
	} else {
		file.WriteString(timestamp + " | " + site + " is up" + " status code: " + fmt.Sprint(resp) + "\n")
		fmt.Println("Site", site, "is up", "status code: ", resp)
	}

	defer file.Close()
}

func showLogs() {
	file, err := os.Open("canary.log")
	checkError(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
