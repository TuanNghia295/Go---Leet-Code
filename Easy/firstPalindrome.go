package main

/*
Given an array of strings words, return the first palindromic string in the array.
If there is no such string, return an empty string "".
A string is palindromic if it reads the same forward and backward.

Example 1:
Input: words = ["abc","car","ada","racecar","cool"]
Output: "ada"
Explanation: The first string that is palindromic is "ada".
Note that "racecar" is also palindromic, but it is not the first.

Example 2:
Input: words = ["notapalindrome","racecar"]
Output: "racecar"
Explanation: The first and only string that is palindromic is "racecar".
*/



func firstPalindrome(words []string) string {
	for _, word := range words{
		if (isPadlindrome(word)){
			return  word
		}
		
	}
	return  ""
}

func isPadlindrome (s string) bool{
	// Check first character = last character or not
	first := 0
	last := len(s) - 1

	for first < last{
		if(s[first] != s[last]){
			return false
		}

		// Ex: "ABCD" => first is A and last is D => A != D => false and break loop

		// Ex: "ABCA"
		// If First = Last => increase first by 1 and decrease last by 1
		// Now first is B and last is C => loop until first > last to break loop
		first ++
		last --
	}
	return  true

}