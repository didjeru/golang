package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func getRandomNumber(numberMin, numberMax int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("You are chiter!")
			os.Exit(3)
		}
	}()
	return numberMin + rand.Intn((numberMax)-(numberMin))
}

func inputData(msg string) string {
	var data string
	fmt.Print(msg)
	fmt.Scanln(&data)
	if data == "exit" {
		os.Exit(3)
	}
	return data
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Guessing game!\nThink of a number in the range of 0 to about 100")
	var numberMin, numberMax = 0, 100
	compNumber := getRandomNumber(numberMin, numberMax)
	fmt.Println("Your number", compNumber, "?")
	countTry := 0
	for {
		inputSymbol := inputData("Enter symbol < or = or > :")
		if inputSymbol == "<" || inputSymbol == "=" || inputSymbol == ">" {
			countTry++
		}
		switch inputSymbol {
		case "=":
			fmt.Println("Win witn", countTry, "try. And number -", compNumber)
			os.Exit(3)
		case ">":
			numberMin = compNumber + 1
			compNumber = getRandomNumber(numberMin, numberMax)
			fmt.Println("Your number", compNumber, "?")
		case "<":
			numberMax = compNumber
			compNumber = getRandomNumber(numberMin, numberMax)
			fmt.Println("Your number", compNumber, "?")
		default:
			fmt.Println("Wrong symbol")
		}
	}
}
