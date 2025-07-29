# Go Format and Go File I/O

## Go Format

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