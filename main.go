package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var pizzasMade, PizzasFailed, total int

type producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch

}
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order #%d!\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			PizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making Pizza #%d.It will take %d second...!\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready", pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *producer) {
	// keep track of which pizza we are trying to make
	var currPizzaNumber = 0
	// run forever or until we recieve a quit notification
	for {
		currentPizza := makePizza(currPizzaNumber)
		if currentPizza != nil {
			currPizzaNumber = currentPizza.pizzaNumber
			select {
			// tried to make a pizza
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}

	//try to make pizzas

}

func main() {

	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a starting message

	color.Cyan("The pizzeria is open for business <O><P><E><N>")

	// create a producer
	pizzaJob := &producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is mad!")
			}
		} else {
			color.Cyan("Done making Pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel", err)
			}
		}
	}
	// print out a ending message
	color.Cyan("SHOP IS CLOSED <C><L><O><S><E><D>")
}
