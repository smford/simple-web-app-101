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
	  }
	  fmt.Println("my name is:", myname)
	  fmt.Println("MY NAME IS:", myfirstfunction(myname))
	  os.Exit(0)
	}

	```
	- Line 2: using the Print function from the package log, display "main is starting"
	- Line 3 - 5: if the length of the variable myname is 0, display a Fatal message
	- Line 6 - 8: using the Println function from the package fmt, print out my name in lowercase and uppercase, then exit happily with the errorcode 0

1. Lets compile and run the code:
		
	```
	cd /go/tutorial/slide-02
	go build 2.go
	`./2`
	echo $?
	```
	- The output should be:
	```
	2019/10/28 19:35:42 starting init function
	2019/10/28 19:35:42 main is starting
	my name is: Boaty McBoatface
	MY NAME IS: BOATY MCBOATFACE
	0
	```
	- notice the time stamps?
	- notice the 0 at the end?  That is the errorcode returned from the application.  The command `echo $?` tells the OS to print out the last return code it recieved.

1. Let us see what happens when we change a couple things.
	- Change this line: `myname = "Boaty McBoatface"`
	- To: `//myname = "Boaty McBoatface"`

	The double / at the beginning of a line comments out the line.  By commenting out this line the variable myname will remain empty.
	
	Then lets compile and run the application again: 
	```
	cd /go/tutorial/slide-02
	go build 2.go
	./2
	echo $?
	```

1.  This time the output will be different:
	```
	2019/10/28 19:50:36 starting init function
	2019/10/28 19:50:36 main is starting
	2019/10/28 19:50:36 oops, name is too short! exiting
	1
	```
	
	- Here the return code was 1, indicating there was an error, but where did it come from? We didn't use os.Exit(1).  In this example the `log.Fatal` function tells the application to exit and return a 1 after printing "oops, name is too short!"

1. Error or return codes are very useful, they allow scripts and other tools to know whether any applications that were run exited happily, and if not happy why the exited.  In the web application world, this is very important as these error or return codes are one way used to monitor an applications health.
