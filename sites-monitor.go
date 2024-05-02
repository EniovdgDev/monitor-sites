package main // pacote principal e executavel
import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const name = "Admin"
const version = 1.1

const timesToMonitoring = 3
const delay = 5

func main() {
	showIntro()

	for {
		showMenu()

		input := readInput()

		switch input {
		case 1:
			startMonitoring()
		case 2:
			println("Exibindo os logs.")
		case 0:
			print("Exiting.")
			os.Exit(0)
		default:
			print("I don't know this function.")
			os.Exit(-1)
		}
	}

}

func showIntro() {
	fmt.Println("Hello ", name)
	fmt.Println("This software is running on version: ", version)
}

func showMenu() {
	fmt.Println("1- Start  monitoring.")
	fmt.Println("2- Show logs.")
	fmt.Println("0- Exit.")
}

func readInput() int {
	var selected int
	fmt.Scan(&selected)
	return selected
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readSitesFromFile()

	for i := 0; i < timesToMonitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		println()
	}
	fmt.Println() // just for visual adjustment
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred. Error: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("The site ", site, " was loaded successfully.")
		registerLog(site, true)
	} else {
		fmt.Println("The site ", site, " is in error. Status Code: ", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFromFile() []string {

	var sites []string

	file, err := os.Open("sitesList.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	leitor := bufio.NewReader(file)

	for {
		row, err := leitor.ReadString('\n')
		row = strings.TrimSpace((row))
		sites = append(sites, row)
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
