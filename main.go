package main

const numberOfPizzas = 10

var pizzasMade, piszzaFailed, total int

type producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {

	//seed the random number generator

	// print out a starting message

	// create a producer

	// run the producer in the background

	// create and run consumer

	// print out a ending message

}
