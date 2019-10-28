# Structure of a golang application

We are looking at slide-01/1.go

It is a simple application that prints out my name and your name, in lower case and upper case.

```
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
  fmt.Println("my name is:", myname)
  fmt.Println("MY NAME IS:", myfirstfunction(myname))
  fmt.Println("your name is:", yourname)
  fmt.Println("MY NAME IS:", myfirstfunction(yourname))
  os.Exit(0)
}

func myfirstfunction(mynewstring string) string {
  return strings.ToUpper(mynewstring)
}
```

---

1. **Functions** are little snippets of code that do a specific task.  In this code we have three functions we make, they are the lines starting with **func**:
	- **init**: a specially named function that runs first when the program starts
	- **main**: another specially named function that runs after init, and is considered the main part of your program
	- **myfirstfunction**: a function that we have made, we could have called it anything.  It doesn't get run unless it is told to run.

1. **Importing Functions**:  Go comes with many functions already built in, but we often need to "import" other functions from a **package** to extend the language to make it do things we want.  Here we are importing three packages:
	- **fmt**: allows us to print stuff out
	- **os**: allows us to interact with the operating system
	- **strings**: allow us to work with strings easier
	```
	import (
	  "fmt"
	  "os"
	  "strings"
	)
	```
    
1. **Variables** are places to store information in your application, here we are using two types of variables:
	- **const**: fixed variable, that once created cannot be changed
	- **var**: standard variables that you can change and use

1. Lets look at the init function
	```
	func init() {
	  fmt.Println("starting init function")
	  myname = "Boaty McBoatface"
	  yourname = "Jane Smith"
	}
	```
	- Line 1: the name of the function and opening { to indicate functions code starts here
    - Line 2: fmt.Println = use the fmt packages Println function to print a string 
    - Line 3: set the myname variable to be "Boaty McBoatface"
    - Line 4: set the yourname variable to be "Jane Smith"
    - Line 5: closing } to indicate function ends here

1. So, what's up with myfirstfunction? - it looks slightly different to the init or main functions
	```
    func myfirstfunction(mynewstring string) string {
    func init() {
    func main() {
    ```
    
    The difference between myfirstfunction and init & main, is that myfirstfunction can accept a string variable (that it then names as mynewstring), and it also returns a string whereas init and main do not.
    
1. Lets work through the myfirstfunction so I can explain what it does:

	```
    func myfirstfunction(mynewstring string) string {
	  return strings.ToUpper(mynewstring)
	}
	```
	- Line 1:
		- create a new function called myfirstfunction
		- allow mynewfunction to be given a string of text called mynewstring
		- and when the function finishes return a string.
    - Line 2:
        - take mynewstring
        - convert to uppercase using the strings function called ToUpper
        - exit the function, returning the string mynewstring in uppercase
 	- Line 3: closing } to indicate function ends here
    
1.	Now lets work through the main function
	```
	func main() {
	  fmt.Println("main is starting")
	  fmt.Println("This is code for slide number:", slidenumber)
	  fmt.Println("my name is:", myname)
	  fmt.Println("MY NAME IS:", myfirstfunction(myname))
	  fmt.Println("your name is:", yourname)
	  fmt.Println("MY NAME IS:", myfirstfunction(yourname))
	  os.Exit(0)
	}
	```

	- Line 1: create a main function
	- Line 2: using the fmt packages Println function, print out "main is starting"
	- Line 3: print out "This is code for slide number: 1"
	- Line 4: print out "my name is: Boaty McBoatface"
	- Line 5: run the function myfirstfunction(myname) which makes myname into upper case and returns "BOATY MCBOATFACE", meaning we print out "MY NAME IS: BOATY MCBOATFACE"
	- Line 6 & 7: similar to lines 4 & 5, but with yourname rather than myname
	- Line 8: run the Exit function from the os package we imported earlier, passing it the number 0
	- Line 9: closing } to indicate function ends here


1.  So I've explained the components of a simple go application, lets step through it again:
	1. Import some extra packages
	1. Setup a couple variables
	1. Define a constant called slidenumber
	1. Create an init function that fires when the application first runs
	1. Create a main function that runs after init, and this is where the bulk of the application runs
	1. Create a myfirstfunction function that only runs when main or init calls it
 
1.  Ok cool, we've got the code, lets see if it works:
	1. We first need to compile the code, this converts it into a binary file that your computer understands and can run
	```
	cd simple-web-app-101/slide-01
	go build 1.go
	```
	1. Lets run the application
	`./`
	1. The output should be:
	```
	starting init function
	main is starting
	This is code for slide number: 1
	my name is: Boaty McBoatface
	MY NAME IS: BOATY MCBOATFACE
	your name is: Jane Smith
	MY NAME IS: JANE SMITH
	```
