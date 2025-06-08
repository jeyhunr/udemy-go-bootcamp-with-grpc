package main

import "fmt"

func Recover() {

	process()
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			fmt.Println("Continuing execution...")
		}
	}()

	fmt.Println("Inside process()")
	panic("This is a panic!")
}
