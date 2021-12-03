package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game !")
	
	var name string = "Paul"
	/*
	var age int = 21
	var unsignedint uint = 12
	var float64type float64 = -1234.56
	var float32type float32 = -1234.56
	var boolean bool = true
	*/
	fmt.Println(name)

	/* Implicit declaration of string */
	name2 := "Thomas"
	fmt.Println(name2)
	/* Implicit declaration of int */
	age := 123
	fmt.Println(age)


	fmt.Printf("Hello %v, how are you doing ?\n", name)
	fmt.Printf("Hello %v %v, how are you doing ?\n", name, name)
	fmt.Printf("Hello %v, you are %v\n", name, age)

	var userName string
	fmt.Scan(&userName)
	fmt.Println(userName)










}
