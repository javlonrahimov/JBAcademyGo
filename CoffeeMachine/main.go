package main

import (
	"fmt"
)

type CoffeeMachine struct {
	water          int
	milk           int
	coffeeBeans    int
	disposableCups int
	money          int
}

const (
	ESPRESSO   = 1
	LATTE      = 2
	CAPPUCCINO = 3
)

func main() {
	coffeeMachine := CoffeeMachine{
		water:          400,
		milk:           540,
		coffeeBeans:    120,
		disposableCups: 9,
		money:          550,
	}

	for {
		showMenu()
		switch scanString() {
		case "buy":
			coffeeMachine.buy()
		case "fill":
			coffeeMachine.fill()
		case "take":
			coffeeMachine.take()
		case "remaining":
			coffeeMachine.showStatus()
		case "exit":
			return
		}
	}
}

func showMenu() {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
}

func (c CoffeeMachine) showStatus() {
	fmt.Println()
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d of water\n", c.water)
	fmt.Printf("%d of milk\n", c.milk)
	fmt.Printf("%d of coffee beans\n", c.coffeeBeans)
	fmt.Printf("%d of disposable cups\n", c.disposableCups)
	fmt.Printf("$%d of money\n", c.money)
	fmt.Println()
}

func (c *CoffeeMachine) fill() {
	fmt.Println("Write how many ml of water you want to add:")
	c.water += scanInt()
	fmt.Println("Write how many ml of milk you want to add:")
	c.milk += scanInt()
	fmt.Println("Write how many grams of coffee beans you want to add:")
	c.coffeeBeans += scanInt()
	fmt.Println("Write how many disposable coffee cups you want to add:")
	c.disposableCups += scanInt()
}

func (c *CoffeeMachine) buy() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	switch scanInt() {
	case ESPRESSO:
		if ok, insufficient := c.canMake(1, ESPRESSO); ok {
			c.water -= 250
			c.coffeeBeans -= 16
			c.disposableCups--
			c.money += 4
			fmt.Println("I have enough resources, making you a coffee!")
		} else {
			fmt.Printf("Sorry, not enough %s!", insufficient)
		}
	case LATTE:
		if ok, insufficient := c.canMake(1, LATTE); ok {
			c.water -= 350
			c.milk -= 75
			c.coffeeBeans -= 20
			c.disposableCups--
			c.money += 7
			fmt.Println("I have enough resources, making you a coffee!")
		} else {
			fmt.Printf("Sorry, not enough %s!", insufficient)
		}
	case CAPPUCCINO:
		if ok, insufficient := c.canMake(1, CAPPUCCINO); ok {
			c.water -= 200
			c.milk -= 100
			c.coffeeBeans -= 12
			c.disposableCups--
			c.money += 6
			fmt.Println("I have enough resources, making you a coffee!")
		} else {
			fmt.Printf("Sorry, not enough %s!", insufficient)
		}
	default:
		return
	}
}

func (c *CoffeeMachine) take() {
	fmt.Printf("I gave you $%d", c.money)
	c.money = 0
}

func (c CoffeeMachine) canMake(amount int, coffeeType int) (bool, string) {
	switch coffeeType {
	case ESPRESSO:
		if c.water < amount*250 {
			return false, "water"
		}
		if c.coffeeBeans < amount*16 {
			return false, "coffee beans"
		}
		if c.disposableCups < amount {
			return false, "disposable cups"
		}
		return true, ""
	case LATTE:
		if c.water < amount*350 {
			return false, "water"
		}
		if c.milk < amount*75 {
			return false, "milk"
		}
		if c.coffeeBeans < amount*20 {
			return false, "coffee beans"
		}
		if c.disposableCups < amount {
			return false, "disposable cups"
		}
		return true, ""
	case CAPPUCCINO:
		if c.water < amount*200 {
			return false, "water"
		}
		if c.milk < amount*100 {
			return false, "milk"
		}
		if c.coffeeBeans < amount*12 {
			return false, "coffee beans"
		}
		if c.disposableCups < amount {
			return false, "disposable cups"
		}
		return true, ""
	}
	return true, ""
}

func scanInt() int {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		return 0
	}
	return number
}

func scanString() string {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		return ""
	}
	return str
}
