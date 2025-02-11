package main

// import "fmt"

// func isPalindrome(x int) bool {
// 	// If the number is negative or ends with zero (but is not zero), it's not a palindrome
// 	if x < 0 || (x%10 == 0 && x != 0) {
// 		return false
// 	}
	
// 	reversedHalf := 0
// 	// Reverse the second half of the number
// 	for x > reversedHalf {
// 		digit := x % 10
// 		reversedHalf = reversedHalf*10 + digit
// 		x /= 10
// 	}

// 	// Check if the first half is equal to the reversed second half
// 	return x == reversedHalf || x == reversedHalf/10
// }

// func main() {
// 	// Test cases
// 	fmt.Println(isPalindrome(121))  // Output: true
// 	fmt.Println(isPalindrome(-121)) // Output: false
// 	fmt.Println(isPalindrome(10))   // Output: false
// 	fmt.Println(isPalindrome(0))
// }
