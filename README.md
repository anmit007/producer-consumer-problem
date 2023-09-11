# Producer-Consumer Pizza Order Processing

This Go program simulates a pizzeria's order processing system using a producer-consumer pattern. The pizzeria receives pizza orders, processes them, and notifies customers about the status of their orders.

## Overview

The program consists of two main components: the **producer** and the **consumer**.

- The **producer** represents the pizzeria kitchen. It generates pizza orders, simulates pizza-making processes, and sends these orders to the consumer for further processing.

- The **consumer** represents the order delivery system. It receives pizza orders, processes them, and notifies customers about their order status.

## Key Features

- The program uses channels for communication between the producer and the consumer.
- It simulates the making of pizzas with random delays and random success/failure conditions.
- It tracks the number of pizzas made and failed, as well as the total number of orders.
- The producer runs in the background as a goroutine, continuously generating pizza orders until a specified number of orders is reached.
- The consumer processes orders, prints success/failure messages, and exits when all orders are completed.
- The program uses the "github.com/fatih/color" package to display colorful messages in the console.

## How to Run

1. Make sure you have Go installed on your system.

2. Clone the repository containing the Go source code:

-- git clone <repository_url>
-- cd <repository_directory>

3. Build and run the program:
-- go run main.go


4. The program will start simulating a pizzeria's operations. You will see colorful console messages indicating the progress of pizza orders.

5. Once all pizza orders are processed, the program will display a closing message, and the pizzeria will be closed.

## Customization

- You can adjust the `numberOfPizzas` constant to change the total number of pizzas to be made.
- Modify the random conditions in the `makePizza` function to control the success/failure rates.
- Customize the delay times for making pizzas by modifying the `delay` variable in the `makePizza` function.
- Adjust the closing message and any other output messages to your liking.

## Error Handling

The program includes basic error handling for channel closures. If an error occurs when closing channels, it will be displayed as a red error message.

## License

This program is provided under an open-source license. Feel free to use, modify, and distribute it as needed, following the terms of the license.

---

Enjoy your pizza-order processing simulation! If you have any questions or encounter issues, please don't hesitate to contact the author or open an issue in the repository.


