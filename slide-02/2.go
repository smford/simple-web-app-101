package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var myname string

const slidenumber = "2"

func init() {
	log.Print("starting init function")
	myname = "Boaty McBoatface"
}

func main() {
	log.Print("main is starting")
	if len(myname) == 0 {
		log.Fatal("oops, name is too short! exiting")
	}
	fmt.Println("my name is:", myname)
	fmt.Println("MY NAME IS:", myfirstfunction(myname))
	os.Exit(0)
}

func myfirstfunction(mynewstring string) string {
	return strings.ToUpper(mynewstring)
}
