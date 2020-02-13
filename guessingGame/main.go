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
	case "InputErr":
		fmt.Println("\nWRONG INPUT!")
	default:
		fmt.Println("error")
	}
}

func settings() {
	numberMin, numberMax := 0, 100
	compNumber := getRandomNumber(numberMin, numberMax)
	countTry := 0
}

func variant1() {
	settings()
	fmt.Printf("\nThink of a number in the range of %d to about %d", numberMin, numberMax)
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
			print("InputErr", 0, 0)
			print("comp", countTry, compNumber)
		}
	}
}

func variant2() {
	settings()
	fmt.Printf("\nComputer think of a number in the range of %d to about %d", numberMin, numberMax)
	for {
		userNumber, err := strconv.Atoi(inputData("\nType 'm' for menu, 'q' for quit.\nEnter your numer: "))
		if err != nil {
			print("InputErr", countTry, compNumber)
			countTry--
		}
		countTry++
		if compNumber == userNumber {
			print("win", countTry, compNumber)
			menu()
		}
		if userNumber != 0 && compNumber > userNumber {
			print("userMinus", countTry, compNumber)
		}
		if userNumber != 0 && compNumber < userNumber {
			print("userPlus", countTry, compNumber)
		}
	}
}

func menu() {
	for {
		inputVariant := inputData("Type 'q' for quit.\nEnter 1 if you think a number. Enter 2 if computer think a number: ")
		switch inputVariant {
		case "1":
			variant1()
		case "2":
			variant2()
		default:
			print("InputErr", 0, 0)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Guessing game!")
	menu()
}
