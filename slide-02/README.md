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
