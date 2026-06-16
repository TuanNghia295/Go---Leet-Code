package main

// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

/*
Example 1:

Input: num = 38
Output: 2
Explanation: The process is
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
Since 2 has only one digit, return it.
Example 2:

Input: num = 0
Output: 0
*/
func addDigits(num int) int {
	//    Ép kiểu sang string
	//  Split
	// + 2 số lại, nếu giá trị lớn hơn 10 thì làm lại như trên

	for num >= 10 {
		sum := 0

		for num > 0 {
			sum = sum + (num % 10) // Lấy số cuối
			num = num / 10         // Lấy số đầu
		}

		num = sum
	}
	return num
}
