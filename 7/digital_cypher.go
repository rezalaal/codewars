package kata7

import (
	"fmt"
	"strconv"
)

/*
Digital cypher
Introduction
Digital Cypher assigns to each letter of the alphabet unique number. For example:

 a  b  c  d  e  f  g  h  i  j  k  l  m
 1  2  3  4  5  6  7  8  9 10 11 12 13
 n  o  p  q  r  s  t  u  v  w  x  y  z
14 15 16 17 18 19 20 21 22 23 24 25 26
Instead of letters in encrypted word we write the corresponding number, eg. The word scout:

 s  c  o  u  t
19  3 15 21 20
Then we add to each obtained digit consecutive digits from the key. For example. In case of key equal to 1939 :

   s  c  o  u  t
  19  3 15 21 20
 + 1  9  3  9  1
 ---------------
  20 12 18 30 21

   m  a  s  t  e  r  p  i  e  c  e
  13  1 19 20  5 18 16  9  5  3  5
+  1  9  3  9  1  9  3  9  1  9  3
  --------------------------------
  14 10 22 29  6 27 19 18  6  12 8
Task
Write a function that accepts str string and key number and returns an array of integers representing encoded str.

Input / Output
The str input string consists of lowercase characters only.
The key input number is a positive integer.

Example
Encode("scout",1939);  ==>  [ 20, 12, 18, 30, 21]
Encode("masterpiece",1939);  ==>  [ 14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8]
*/
func Encode(str string, key int) []int {
	alphabets := map[string]int{"a":1,"b":2,"c":3,"d":4,"e":5,"f":6,"g":7,"h":8,"i":9.,"j":10,"k":11,"l":12,"m":13,"n":14,"o":15,"p":16,"q":17,"r":18,"s":19,"t":20,"u":21,"v":22,"w":23,"x":24,"y":25,"z":26}
	keyStr := strconv.Itoa(key)
	for len(keyStr) < len(str) {
		keyStr += keyStr
	}
	
	res := []int{}
	for index, char := range str {
		c, _ := strconv.Atoi(string(keyStr[index]))
		res = append(res,  alphabets[string(char)] + c)	
	}
	return res
}

func Encode2(str string, key int) []int {
	arr := []int{}
	keyStr := fmt.Sprint(key)
	
	for index ,i := range str {
	  arr = append(arr, int(i - 'a' + 1)+ int(keyStr[index % len(keyStr)]) - 48 )
	}
	
	  return arr
}

func Encode3(str string, key int) []int {
	var encoded = make([]int, len(str))
	var keyAsString = strconv.Itoa(key)
	for i := 0; i < len(str); i++ {
		encoded[i] = int(str[i] - 96 + keyAsString[i%len(keyAsString)] - 48)
	}
	return encoded
}