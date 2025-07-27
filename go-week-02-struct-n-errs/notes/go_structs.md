# Go Structs

Go structs provid a way to group related data in a managable way 


When defining a go struct, the "type" keyward has a significance

"Type" makes the struct global and "type" makes the stuct private

## Go Struct defination

We can define structs as a name data type and as anonymous data

1. type Employee struct {
name := string
age := int
} - private
2. Type Customer struct {
name := string
phone := string
} - public/global
3. job := struct{ 
   id:= int, 
   itle:= string
   }{ 
    id: 1234, 
    title: "Lead Engineer"
    } - anonymous


## Struct Pointers

When writing a struct function, kind of like a struct method, we need to parse the struct reference, but it needs to be an acurate reference and not a copy, so that mutation works as expected.

Pointers refer to a memory location and not a copy of the variable pointed so using a point taps from the actual source and not just the temporary copy.

You can use a pointer by adding a * before the struct name

## Nested Structs

This, just like it sounds is having structs in structs

We can use a struct inside another struct and access the individual parts of the embbeded struct without acessing it bit by bit, it is called composition in go

## Go Tags

Go tags are used to add some metadata to a struct for serializing data, maybe from an http respost data, the tags are used for making the key values of the json key value pairs when data is serialized. Also when data are serialized, the struct feilds have to be public ie start with capital letter

