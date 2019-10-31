# Making the appication run on any operating system

Golang is a [_compiled_ language](https://en.wikipedia.org/wiki/Compiled_language), this means that the code gets compiled (converted) in to [machine code](https://en.wikipedia.org/wiki/Machine_code), saved to disk, and then this binary file can be executed as the CPU understands it.

The alternative is called an [_interpreted language_](https://en.wikipedia.org/wiki/Interpreted_language), rather than converting to a binary file which is then executed, another program called an intepreter reads in the source code, compiles it into machine code on the fly and then executes what it just compiled.  This makes running the code slower compared to a compiled language as each time the code is run, it needs to be compiled before running.

There is also a hybrid of the two methods.

What is cool about Golang is that you can easily compile a program in linux to run in windows, osx or even the raspberry pi.  You can do this with other languages, but golang has all the capability built in and it is very easy to do.  

By default with you run the `go build` command it will compile the program for current operating system, that is what we have been doing so far.  Lets try compiling an application for another operating system.

```
cd slide-03
```

- Linux:
	`GOOS=linux GOARCH=amd64 go build -o something-linux something.go`
- Raspberry Pi:
	`GOOS=linux GOARCH=arm GOARM=5 go build -o something-rpi something.go`
- OSX:
	`GOOS=darwin GOARCH=amd64 go build -o something-osx something.go`
- Windows:
	`GOOS=windows GOARCH=386 go build -o something-windows.exe something.go`

The `-o something` tells the compiler to output the compiled binary to a file of that name.

## Lets look at what we have
```
ls -la
```

It will look similar to this:

```
root@9a0566ecc346:/go/tutorial/slide-03# ls -la
total 7708
drwxr-xr-x 8 root root     256 Oct 31 16:45 .
drwxr-xr-x 9 root root     288 Oct 31 16:43 ..
-rw-r--r-- 1 root root      47 Oct 31 16:44 README.md
-rwxr-xr-x 1 root root 2008801 Oct 31 16:44 something-linux
-rwxr-xr-x 1 root root 2120320 Oct 31 16:45 something-osx
-rwxr-xr-x 1 root root 1879818 Oct 31 16:44 something-rpi
-rwxr-xr-x 1 root root 1870848 Oct 31 16:45 something-windows.exe
-rw-r--r-- 1 root root      79 Oct 31 16:44 something.go
root@9a0566ecc346:/go/tutorial/slide-03#
```

## Lets run these

```
./something-linux
./something-osx
./something-rpi
./something-windows.exe
```

To see more about these binary executables install the `file` command, this tells us what a file is

```
apt-get install file -y
file something-linux
file something-osx
file something-rpi
file something-windows.exe
```

If you are feeling daring, you can now exit out of the container back to your laptop, move to the tutorial/slide-03 directory and try and run the binary that is suitable for your operating system and it should work! 
