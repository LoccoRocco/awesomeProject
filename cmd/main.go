package main

import "fmt"

func main() {
	err := Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Works!")
}
