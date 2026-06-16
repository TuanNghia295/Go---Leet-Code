package main

import (
	"fmt"
	"sort"
	"strings"
)

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


func main()  {
	// fmt.Println(32 / 10)
	// fmt.Println(addDigits(32))
	// fmt.Println(numberGame([]int{2,5,3,2}))
	// fmt.Println(checkIfPangram("Nghia"))

	fmt.Println(numberOfEmployeesWhoMetTarget([]int{0,1,2,3,4},2))
}


func addDigits(num int)int{
//    Ép kiểu sang string
//  Split
// + 2 số lại, nếu giá trị lớn hơn 10 thì làm lại như trên

	for num >= 10{
		sum := 0

		for num > 0{
			sum = sum + (num % 10) // Lấy số cuối
			num = num / 10 // Lấy số đầu
		}

		num = sum
	}
	return num
}





/*
You are given a 0-indexed integer array nums of even length and there is also an empty array arr. 
Alice and Bob decided to play a game where in every round Alice and Bob will do one move. 
The rules of the game are as follows:

Every round, first Alice will remove the minimum element from nums, and then Bob does the same.
Now, first Bob will append the removed element in the array arr, and then Alice does the same.
The game continues until nums becomes empty.
Return the resulting array arr.

Example 1:

Input: nums = [5,4,2,3]
Output: [3,2,5,4]
Explanation: In round one, first Alice removes 2 and then Bob removes 3. 
Then in arr firstly Bob appends 3 and then Alice appends 2. So arr = [3,2].
At the begining of round two, nums = [5,4]. Now, first Alice removes 4 and then Bob removes 5. Then both append in arr which becomes [3,2,5,4].
*/
func numberGame(nums []int) []int {
	// sort and get the smallest in array
    // create an empty slices
	// get the second smallest in array and append in empty array first, then append the smallest

	sort.Ints(nums)
	result := make([]int,0,len(nums))

	for i:= 0; i < len(nums) ; i+=2{
		 // Append the second smallest first, then the smallest
		result = append(result,nums[i+1],nums[i])
	}

	return result
}



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

	for _, ch := range ls{
		// Thêm ký tự đó vào map với giá trị là true.
		result[ch] = true
	}

	// Trả về true nếu map có đủ 26 phần tử 
	return len(result) == 26
}



/*
There are n employees in a company, numbered from 0 to n - 1. Each employee i has worked for hours[i] hours in the company.
The company requires each employee to work for at least target hours.
You are given a 0-indexed array of non-negative integers hours of length n and a non-negative integer target.
Return the integer denoting the number of employees who worked at least target hours.
 

Example 1:
Input: hours = [0,1,2,3,4], target = 2
Output: 3
Explanation: The company wants each employee to work for at least 2 hours.
- Employee 0 worked for 0 hours and didn't meet the target.
- Employee 1 worked for 1 hours and didn't meet the target.
- Employee 2 worked for 2 hours and met the target.
- Employee 3 worked for 3 hours and met the target.
- Employee 4 worked for 4 hours and met the target.
There are 3 employees who met the target.


*/

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	count := 0
    for i := 0; i < len(hours); i++{
		if(hours[i] < 0){
			return 0
		}
		if(hours[i] >= target){
			count++
		}
	}
	return  count
}


// func firstPalindrome(words []string) string {
    
// }