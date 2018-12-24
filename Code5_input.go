package main

import "fmt"

func main(){

	fmt.Printf("Enter the name: \n")

	var name string

	fmt.Scan(&name)

	fmt.Printf("Name is: %s \n", name)
}