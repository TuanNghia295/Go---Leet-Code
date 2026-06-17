package main

import "strings"

/*
Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.
Example 1:
Input: s = "Hello"
Output: "hello"
*/

func toLowerCase(s string) string {
	return strings.ToLower(s)
}