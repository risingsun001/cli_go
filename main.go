package main

import "fmt"



type cliCommand struct{
	name string
	description string
	callback func() error
}




Commandlist=map[string]cliCommand{
	"help":
	{
		name: "help"
		description: "Displays a help message"
	}
}






func main() {

	fmt.Printf("Hello, World!")
}
