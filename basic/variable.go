package main

import "fmt"

func main() {
	// declaration variable with data type
	var firstName string = "john"
	// declaration variable without data type
	lastName := "wick"
	// declaration multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	one, two := "satu", "dua"
	// variable underscore
	_, _ = "satu", 2
	// declaration with keyword new function
	name := new(string)
	name = &firstName
	// declaration with keyword make function
	maps := make(map[int]string)
	maps[1] = "satu"
	channel := make(chan string)
	// send data to channel in other goroutine
	go func() {
		channel <- "channel"
	}()
	receiveChannel := <-channel
	slice := make([]int, 1)
	slice[0] = 0
	// if you assign slice in index 1 when you just make len slice is 1 it will error index out of range
	// index start from 0
	// len start from 1
	//slice[1] = 0
	// if you append or add element to slice it will fine
	slice = append(slice, 1)
	fmt.Println(firstName, lastName, first, second, third, one, two, name, maps, channel, receiveChannel, slice)
}
