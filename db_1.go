package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	for {
		// file = os.Open("data.txt") // assumes a file already exits
		file, _ := os.Open("data.txt") // assumes a file already exits
		fmt.Println("\nEnter your Full Command: ")
		// scanner := bufio.NewScanner(os.Stdin)
		// scanner.Split(bufio.ScanWords)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		var commands []string
		commands = strings.Split(text, " ")

		var firstCommand = commands[0]

		// need fix for when q is entered ( dont have second command, will complain)
		// city code and city name
		// var aa, an string
		var startcc, endcc string
		var connections int

		if firstCommand == "a" {
			var secondCommand = commands[1]
			var cc, cn string
			cc = commands[2]
			for i := 0; i < len(commands); i++ {
				cn = cn + " " + commands[i]
			}
			if strings.Contains(secondCommand, "c") {
				// AddC(cc, cn, file)
				fmt.Fprintln(file, "c: "+cc+" "+cn)
			} else if secondCommand == "a" {
				AddA(cc, cn, file)
			} else if secondCommand == "f" {
				// AddF()
			} else {
				fmt.Println("wrong second command, try again")
			}

		} else if firstCommand == "l" {
			fmt.Println("inside l")
			var secondCommand = commands[1]
			if secondCommand == "c" {
				LoadC(file)
			} else if secondCommand == "a" {
				LoadA(file)
			} else if secondCommand == "f" {
				LoadF(file)
			} else {
				fmt.Println("wrong second command, try again")
			}

		} else if firstCommand == "f" {
			Find(startcc, endcc, connections)
			return
		} else if firstCommand == "q" {
			// quit
			os.Exit(3)
		} else {
			fmt.Println("wrong first command, try again")
		}
		defer file.Close()
	}

}

// AddC adds a city
func AddC(cc string, cn string, f *os.File) {
	var cityCode = cc
	var cityName = cn
	fmt.Fprintln(f, "c: "+cityCode+" "+cityName)
}

// AddA adds an airport
func AddA(aa string, an string, f *os.File) {
	fmt.Fprintln(f, "a: "+aa+" "+an)
}

// AddF adds a flight
func AddF(aa, cc1, cc2, p string, f *os.File) {
	fmt.Fprintln(f, "f: "+aa+" "+cc1+" "+cc2+" "+p)
}

// LoadC loads cities
func LoadC(f *os.File) {
	os.Open("data.txt")
	b, _ := ioutil.ReadAll(f)
	s := string(b)
	var arr []string
	arr = strings.Split(s, "\n")
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "c: ") {
			fmt.Println(arr[i])
		}
	}
}

// LoadA loads airports
func LoadA(f *os.File) {

}

// LoadF loads flights
func LoadF(f *os.File) {

}

// Find finds flights according to entries of starting city and
// finish city and number of connections
func Find(startcc string, endcc string, connections int) {

}
