package main

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	fmt.Println("starting init function")
}

func main() {
	fmt.Println("main is starting")
	var myname string
	myname = "Boaty McBoatface"
	fmt.Println("My name is:", myname)
	fmt.Println("MY NAME IS:", myfirstfunction(myname))
	os.Exit(0)
}

func myfirstfunction(mynewstring string) string {
	return strings.ToUpper(mynewstring)
}
