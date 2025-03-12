package main

// import "fmt"

// func isValid(s string) bool {
// 	stack := []rune{}

// 	for _, char := range s {
// 		if char == '(' {
// 			stack = append(stack, char)
// 		} else if char == ')' {
// 			if len(stack) == 0 || stack[len(stack)-1] != '(' {
// 				return false
// 			}
// 			stack = stack[:len(stack)-1]
// 		}
// 	}

// 	return len(stack) == 0
// }

// func main() {
// 	s := "()"
// 	fmt.Println(isValid(s))
// }
