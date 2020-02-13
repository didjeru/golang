package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandomNumber(numberMin, numberMax int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nYou are chiter!\n")
			menu()
		}
	}()
	return numberMin + rand.Intn((numberMax)-(numberMin))
}

func inputData(msg string) string {
	var data string
	fmt.Print(msg)
	fmt.Scanln(&data)
	if data == "q" {
		os.Exit(3)
	}
	if data == "m" {
		menu()
	}
	return data
}

func print(value string, countTry, compNumber int) {
	switch value {
	case "win":
		fmt.Println("\nWin witn", countTry, "try. And number -", compNumber, "\n")
	case "comp":
		fmt.Printf("\nYour number %d?\n", compNumber)
	case "userMinus":
		fmt.Println("\nYour number < computer number. Try again!")
	case "userPlus":
		fmt.Println("\nYour number > computer number. Try again!")
	case "err":
		fmt.Println("\nWRONG INPUT!")
	default:
		fmt.Println("error")
	}
}

func variant1(compNumber, numberMin, numberMax int) {
	fmt.Printf("\nThink of a number in the range of %d to about %d", numberMin, numberMax)
	countTry := 0
	print("comp", countTry, compNumber)
	for {
		inputSymbol := inputData("\nType 'm' for menu, 'q' for quit.\nEnter symbol < = >: ")
		if inputSymbol == "<" || inputSymbol == "=" || inputSymbol == ">" {
			countTry++
		}
		switch inputSymbol {
		case "=":
			print("win", countTry, compNumber)
			menu()
		case ">":
			numberMin = compNumber + 1
			compNumber = getRandomNumber(numberMin, numberMax)
			print("comp", countTry, compNumber)
		case "<":
			numberMax = compNumber
			compNumber = getRandomNumber(numberMin, numberMax)
			print("comp", countTry, compNumber)
		default:
			print("err", countTry, compNumber)
		}
	}
}

func variant2(compNumber, numberMin, numberMax int) {
	fmt.Printf("\nComputer think of a number in the range of %d to about %d", numberMin, numberMax)
	countTry := 0
	for {
		userNumber, err := strconv.Atoi(inputData("\nType 'm' for menu, 'q' for quit.\nEnter your numer: "))
		if err != nil {
			print("err", countTry, compNumber)
		}
		countTry++
		if compNumber == userNumber {
			print("win", countTry, compNumber)
			menu()
		}
		if compNumber > userNumber {
			print("userMinus", countTry, compNumber)
		}
		if compNumber < userNumber {
			print("userPlus", countTry, compNumber)
		}
	}
}

func menu() {
	numberMin, numberMax := 0, 100
	compNumber := getRandomNumber(numberMin, numberMax)
	for {
		inputVariant := inputData("Type 'q' for quit.\nEnter 1 if you think a number. Enter 2 if computer think a number: ")
		switch inputVariant {
		case "1":
			variant1(compNumber, numberMin, numberMax)
		case "2":
			variant2(compNumber, numberMin, numberMax)
		default:
			print("err", 0, 0)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Guessing game!")
	menu()
}
