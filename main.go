package main

import (
	"fmt"
)

var waterLevel, milkLevel, beansLevel, moneyLevel, cups int
var loop bool

func begin() {
	waterLevel, milkLevel, beansLevel, moneyLevel, cups = 400, 540, 120, 550, 9
	loop = true
}

func levelsPrinter() {
	fmt.Println("The coffee machine has:")
	fmt.Println(waterLevel, "of water")
	fmt.Println(milkLevel, "of milk")
	fmt.Println(beansLevel, "of coffee beans")
	fmt.Println(cups, "of disposable cups")
	fmt.Printf("$%d of money\n", moneyLevel)
}

func pourEspresso() {
	const waterXCup, beansXCup int = 250, 16
	const price int = 4
	if waterLevel < waterXCup || beansLevel < beansXCup {
		fmt.Println("Not enough ingredients to pour an espresso")
	} else {
		waterLevel -= waterXCup
		beansLevel -= beansXCup
		cups -= 1
		moneyLevel += 4
		fmt.Println("I have enough resources, making you a coffee!")
	}
}

func pourLatte() {
	const waterXCup, milkXCup, beansXCup int = 350, 75, 20
	const price int = 7

	if waterLevel < waterXCup || milkLevel < milkXCup || beansLevel < beansXCup {
		fmt.Println("Not enough ingredients to pour a latte")
	} else {
		waterLevel -= waterXCup
		milkLevel -= milkXCup
		beansLevel -= beansXCup
		cups -= 1
		moneyLevel += 7
		fmt.Println("I have enough resources, making you a coffee!")
	}
}

func pourCappuccino() {
	const waterXCup, milkXCup, beansXCup int = 200, 100, 12
	const price int = 6

	if waterLevel < waterXCup || milkLevel < milkXCup || beansLevel < beansXCup {
		fmt.Println("Not enough ingredients to pour a latte")
	} else {
		waterLevel -= waterXCup
		milkLevel -= milkXCup
		beansLevel -= beansXCup
		cups -= 1
		moneyLevel += 6
		fmt.Println("I have enough resources, making you a coffee!")
	}
}

func makeCoffee() {
	var choice int
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		pourEspresso()
	case 2:
		pourLatte()
	case 3:
		pourCappuccino()
	}

}

func fillMachine() {
	var water, milk, beans int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&beans)
	fmt.Println("Write how many disposable coffee cups you want to add:")
	waterLevel += water
	milkLevel += milk
	beansLevel += beans
}

func takeMoney() {
	fmt.Println("I gave you", moneyLevel)
	moneyLevel -= moneyLevel
}

func main() {
	// write your code here
	begin()
	for loop {
		var action string
		fmt.Println("Write action (buy, fill, take):")
		fmt.Scan(&action)

		switch action {
		case "buy":
			makeCoffee()
		case "fill":
			fillMachine()
		case "take":
			takeMoney()
		case "remaining":
			levelsPrinter()
		case "exit":
			loop = false
		}
	}
}
