package main

import "strings"

/*
A pangram is a sentence where every letter of the English alphabet appears at least once.
Given a string sentence containing only lowercase English letters,
return true if sentence is a pangram, or false otherwise.

Example 1:
Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true
Explanation: sentence contains at least one of every letter of the English alphabet.

Example 2:

Input: sentence = "leetcode"
Output: false
*/
func checkIfPangram(sentence string) bool {
	//  convert to lowercase
	//  loop to check alphabet

	ls := strings.ToLower(sentence)
	// Sử dụng map[rune]bool để lưu các ký tự đã xuất hiện trong chuỗi.
	result := make(map[rune]bool)

	for _, ch := range ls {
		// Thêm ký tự đó vào map với giá trị là true.
		result[ch] = true
	}

	// Trả về true nếu map có đủ 26 phần tử
	return len(result) == 26
}
