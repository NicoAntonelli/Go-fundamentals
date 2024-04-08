package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Constants for channel examples utilization
const PHONE_MAX_PRICE = 900
const LAPTOP_MAX_PRICE = 1500

// Work with Go routines using channels
func goRoutinesChannels() {
	fmt.Println("Go routines with channels!")
	var channel = make(chan int)
	var bufferChannel = make(chan int, 5)

	fmt.Println("Unbuffered channel capacity:", cap(channel))
	go process(channel)
	for value := range channel {
		// Read the channel value
		fmt.Println("Channel value:", value)
		// Simulate slow processing time from other tasks
		time.Sleep(time.Second/2)
	}
	
	fmt.Println("Buffered channel capacity:", cap(bufferChannel))
	go process(bufferChannel)
	for value := range bufferChannel {
		// Read the channel value
		fmt.Println("Buffered channel value:", value)
		// Simulate slow processing time from other tasks
		time.Sleep(time.Second/2)
	}
	fmt.Println()

	// Real-life example with channels
	var phoneChannel = make(chan string)
	var laptopChannel = make(chan string)
	var websites = []string{"www.amazon.com", "www.ebay.com", "www.mercadolibre.com"}

	for i := range websites {
		// Create a goroutine for each website
		go checkPhonePrices(websites[i], phoneChannel)
		go checkLaptopPrices(websites[i], laptopChannel)
	}

	sendMessage(phoneChannel, laptopChannel)
	fmt.Println()
}

// Go routine that access to a channel
func process(channel chan int) {
	// Defer stack this function and executes it at the end
	// Closes the channel
	defer close(channel)
	for i := 0; i < 5; i++ {
		// Set a value in the channel
		channel <- i
	}

	// Show message at process exit
	if cap(channel) >= 1 {
		fmt.Println("Process with unbuffered channel done")
		} else {
		fmt.Println("Process with buffered channel done")
	}
}

// Check if a website has a nice phone offer
func checkPhonePrices(website string, channel chan string) {
	for {
		time.Sleep(time.Second+1)
		var phonePrice = rand.Float32()*20
		if phonePrice < PHONE_MAX_PRICE {
			channel <- website
			break
		}
	}
}

// Check if a website has a nice laptop offer
func checkLaptopPrices(website string, channel chan string) {
	for {
		time.Sleep(time.Second+1)
		var laptopPrice = rand.Float32()*20
		if laptopPrice < LAPTOP_MAX_PRICE {
			channel <- website
			break
		}
	}
}

// Inform user if a deal of phone or laptop is found
func sendMessage(phoneChannel, laptopChannel chan string) {
	select {
		case phone := <-phoneChannel:
			fmt.Println("SMS | Phone offer found:", phone)
		case laptop := <-laptopChannel:
			fmt.Println("Email | Laptop offer found:", laptop)
	}
}
