package kata8

import "fmt"

/*
#Returning Strings

Make a function that will return a greeting statement that uses an input; your program should return, 
"Hello, <name> how are you doing today?".

[Make sure you type the exact thing I wrote or the program may not execute properly]
*/
func Greeting(name string) string {
	return "Hello, "+name+" how are you doing today?"
}

func Greeting1(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}