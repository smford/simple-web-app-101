package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	myname   string
	yourname string
)

const slidenumber = "1"

func init() {
	fmt.Println("starting init function")
	myname = "Boaty McBoatface"
	yourname = "Jane Smith"
}

func main() {
	fmt.Println("main is starting")
	fmt.Println("This is code for slide number:", slidenumber)
	fmt.Println("My name is:", myname)
	fmt.Println("MY NAME IS:", myfirstfunction(myname))
	fmt.Println("your name is:", yourname)
	fmt.Println("MY NAME IS:", myfirstfunction(yourname))
	os.Exit(0)
}

func myfirstfunction(mynewstring string) string {
	return strings.ToUpper(mynewstring)
}
