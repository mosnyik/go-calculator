# Go Variable Types

## Number Types

## Integers

### 1. Unsigned Integers - all positive value

1. uint8/byte (0-255) 8 bits
2. uint16 (0-65535) 16 bits
3. uint32 (0-4294967295) 32 bits
4. uint64 (0-18446744073709551615) 64 bits

### 2. Signed Integers - range from nagative to positive numbers

5. int8 (-128 - 127) 8 bits
6. int16 (-32768 - 32767) 16 bits
7. int32/rune (-2147483648 - 2147483647)
8. int64 (-9223372036854775808 - 9223372036854775807)
   
### 3. Machine Dependent Types default types based on the OS

1. uint (32 / 64 bits) this depends on the cpu architecture of the computer
2. int (32 / 64 bits) this depends on the cpu architecture of the computer
3. uintptr (an unsigned-positive integer to store the uninterpreted bits of a pointer value)

## Floating Point Numbers

### Floats

1. flaot32 (IEEE-754 32-bits flaoting-point numbers) - represented using 32-bit
2. flaot64 (IEEE-754 64-bits flaoting-point numbers)- represented using 64-bit

### Complex(Imaginary parts)

1. complex32 (complex numbers with flaot32 real and imaginary parts)
2. complex128 (complex numbers with flaot64 real and imaginary parts)

## String
- "Anything in double quotation mark"

## Characters
- 'Anything in single quotation mark'

## Boolean
- true
- false