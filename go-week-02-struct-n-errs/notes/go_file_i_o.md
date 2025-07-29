# Go Format and Go File I/O

## Go Format

Like in linux, go has a ways to expose formats for I/O streams

- Standard input
- Standard output
- Standard error

These are exposed through the os.strin
So by default, fmt.Println() goes by default to standard output
But if we use fmt.Fprintln(), it can be redirected to os.Strerr 

So we can have the fmt.Println() which takes a variable number of args and format them in a straight line, seperated by space

Also we have the fmt.Printf() which also takes a variable number of args and the first arg is a format string that tells how the rest should be formated

When we use a fmt.Fprintln() or fmt.Fprintf(), what it does is allow for providing and output identifier for redirecting the output eg fmt.Fprintf(os.Strerr, "the error goes here")

The last variation is the fmt.Sprintln() and the fmt.Sprintf(); 

These do not return to an output, rather they return a formated string that can be used somewhere in the program, just like a function that returns a value

For formating, look at https://pkg.go.dev/fmt

## Go I/O
Go has a package os, which have a couple of usefull functions like the 
- os in
- os out
- os err
- os creat
- os open

The os package is used to work with files in go. Anything in working with files like creating a file, editing a file, deleting a file etc

Package io has utilities for doing some input and output like read and write

Package bufferio utilities for doing some buffered input and output like the buffered I/O scanner

Package io/ioutil has some amzing utilities that can enable a you to read or write an entire file in one shot

Package strconv has utilites that can parse numbers from text that is can convert to/from string representations. The utils are like parse int, parse boolean, parse float etc that can read a string rep of a type and correctly parse it into the desired type

If a function is going to return an error, the error would be the last thing it would return if it is retutning many things at once