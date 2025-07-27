# GO POINTERS

A pointer like the name sounds points to a memory address of a variable.

So we can have a variable name saved in a memory address 0xfft234e3e 

- The variable name is "name" which points to the value of name
- The value of name could be "Moses"
- The pointer to "Moses" is "&name" which returns "0xfft234e3e" the memory address
- To get the value of the variable from the pointer, you dereference it like so "*name" and it gives "Moses"

Summary, if you call 
name -> Moses
&name -> 0xfft234e3e
*name -> Moses

## Pointers and references
When we create a pointer, we can copy its address into anoter variable, even though the new variable would have an address of its own but it would hold the value of the address copied into it, and modification in the one variable would reflect is all the references. To change the change the variable that is referenced, you just derefernce it and then modify the value

## Uses of pointers:
1. To parse a property by reference - to reference a http request value in httphandler function. The request variable is a pointer because it can be modified, as opposed to the response which only reads the response
2. To make a property optional. If we are returning a value that can be nill, it has to be a point because we can have an invalid memory address returned. For an array, it is better to return the empty array, because the array is already a reference to the chunck of memory for the array - When returning a struct, you can return a nill by referencing the pointer to the struct or return an empty struct - Note: If you are returning a pointer, then make sure you do a nill check to properly handle the errors. Otherwise, just return an empty struct and forget about nill check
3. The most important reason in my opinion is to save memory when referencing a large struct, so you reference the struct instead of copying it upandan
   
