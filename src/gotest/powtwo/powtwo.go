package main

import "fmt"

func main() {
	var value int = 128
	fmt.Printf("Is %d a power of 2 (RECURSIVE)? %b\n", value, IsPowerOfTwoRecursive(value));	
	value = 130
	fmt.Printf("Is %d a power of 2 (RECURSIVE)? %b\n", value, IsPowerOfTwoRecursive(value));	

	value = 128
	fmt.Printf("Is %d a power of 2 (BINARY)? %b\n", value, IsPowerOfTwoBoolean(value));	
	value = 130
	fmt.Printf("Is %d a power of 2 (BINARY)? %b\n", value, IsPowerOfTwoBoolean(value));
}

func IsPowerOfTwoRecursive(value int) bool {
	even := (value % 2) == 0

	if even {
		value /= 2
		if value == 1 {
			return true
		}else{
			return IsPowerOfTwoRecursive(value)
		}
	}

	return false
}

func IsPowerOfTwoBoolean(value int) bool{
	if value != 0 && (value & (value -1)) == 0 {
		return true
	}
	return false
}