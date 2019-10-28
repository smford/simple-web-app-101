# Changing the simple golang application

We are looking at slide-02/2.go

It is a simple application, similar to slide-01, which will print your name, in lower case and upper case, with logging and an error message.

```
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var myname string

func init() {
	log.Print("starting init function")
	myname = "Boaty McBoatface"
}

func main() {
	log.Print("main is starting")
	if len(myname) == 0 {
		log.Fatal("oops, name is too short! exiting")
		os.Exit(1)
	}

	fmt.Println("my name is:", myname)
	fmt.Println("MY NAME IS:", myfirstfunction(myname))
	os.Exit(0)
}

func myfirstfunction(mynewstring string) string {
	return strings.ToUpper(mynewstring)
}
```

---

1. Lets look at the import, you will notice that it now imports a package called **log**, and you will now see that there are some lines like:
	`log.Print("main is starting")`
	In the previous slide we'd used `fmt.Println("main is starting")`
	Using the log packages Print function, it will add a time stamp to the begining of any text it prints out, this is useful for logging.

1. Lets look at the main function:
	```
	func main() {
	  log.Print("main is starting")
	  if len(myname) == 0 {
	    log.Fatal("oops, name is too short! exiting")
	    os.Exit(1)
	  }
	  fmt.Println("my name is:", myname)
	  fmt.Println("MY NAME IS:", myfirstfunction(myname))
	  os.Exit(0)
	}

	```
	- Line 2: using the Print function from the package log, display "main is starting"
	- Line 3 - 6: if the length of the variable myname is 0, display a Fatal message and exit the application with an errorcode of 1
	- Line 7 - 9: using the Println function from the package fmt, print out my name in lowercase and uppercase, then exit happily with the errorcode 0

1. Lets compile and run the code:
		
	```
  cd simple-web-app-101/slide-02
  go build 2.go
  ```
  `./2`
  - The output should be:
  ```
	2019/10/28 19:35:42 starting init function
	2019/10/28 19:35:42 main is starting
	my name is: Boaty McBoatface
	MY NAME IS: BOATY MCBOATFACE
	```
	- notice the time stamps?
