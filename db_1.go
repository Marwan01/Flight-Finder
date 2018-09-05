package main

import "fmt"

func main() {
	readFirstCommand()
	fmt.Println()
}
func readFirstCommand() {
	var firstCommand = ""
	for {
		fmt.Scanln(&firstCommand)
		if firstCommand == "a" {
			add(readSecondCommand())
			return
		} else if firstCommand == "l" {
			// list
			return
		} else if firstCommand == "f" {
			// find
			return
		} else if firstCommand == "q" {
			// quit
			return
		} else {
			fmt.Println("wrong entry, try again")
		}
	}
}

func readSecondCommand() string {
	var secondCommand = ""
	for {
		fmt.Scanln(&secondCommand)
		if secondCommand == "c" {
			return "c"
		} else if secondCommand == "a" {
			return "a"
		} else if secondCommand == "f" {
			return "f"
		} else {
			fmt.Println("wrong entry, try again")
		}
	}

}

func add(secondCommand string) {
	if secondCommand == "c" {
		addC()
	} else if secondCommand == "a" {
		addA()
	} else {
		addF()
	}
}

func addC() {
	var cityCode = ""
	var cityName = ""
	fmt.Scan(&cityCode)
	fmt.Scanln(&cityName)
}

func addA() {
	fmt.Println("yay made it to a, a")

}

func addF() {

}
