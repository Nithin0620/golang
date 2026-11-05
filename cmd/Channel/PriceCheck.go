package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Maximum price we are willing to pay for chicken
var MAX_CHICKEN_PRICE float32 = 5

func main() {

	// Create a channel to receive websites where chicken is available
	var chickenChannel = make(chan string)

	// List of websites we want to monitor
	var websites = []string{
		"walmart.com",
		"costco.com",
		"wholefoods.com",
	}

	// Start a separate goroutine for each website
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}

	// Wait for a deal and print the message
	sendMessage(chickenChannel)
}

// Continuously checks chicken prices on a particular website
func checkChickenPrices(website string, chickenChannel chan string) {

	for {
		// Simulate checking the website every second
		time.Sleep(time.Second * 1)

		// Generate a random chicken price between 0 and 20
		var chickenPrice = rand.Float32() * 20

		// Check if the price is within our budget
		if chickenPrice <= MAX_CHICKEN_PRICE {

			// Send the website name through the channel
			chickenChannel <- website

			// Stop checking this website after finding a deal
			break
		}
	}
}

// Receives the website name and prints the deal message
/*
func sendMessage(chickenChannel chan string) {

	// Receive data from the channel
	website := <-chickenChannel

	fmt.Printf("\nFound a deal on chicken at %s", website)
}
*/

func sendMessage(chickenChannel chan string, tofuChannel chan string) {

	select {

	// Executes when chickenChannel receives data
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Found deal on chicken at %v.", website)

	// Executes when tofuChannel receives data
	case website := <-tofuChannel:
		fmt.Printf("\nEmail Sent: Found deal on tofu at %v.", website)

	}
}