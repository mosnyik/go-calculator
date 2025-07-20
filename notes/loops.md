# Loops In Go

Go typically supports onlyfor loops, but this can be used in an amazing way

## Normal Loop
We can write a traditional looking for loop

for i:=0; i<5; i++{
    fmt.Println(i)
}

## Loop over a range
we can also loop over a range
for i:= range 5{
    fmt.Println(i)
}

## Loop over an array
We can also loop over an array

var names = [3] string { "Moze", "mosnyik", "moses"}
for k,v:= range name{
    fmt.Println(k,v)
}

