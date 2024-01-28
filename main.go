package main

import "fmt"

func main() {
	fmt.Println("Welcome to go lang")
	server := makeServer(":3000")
	server.run()
}
