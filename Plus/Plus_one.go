package main

import "fmt"

func main() {

	fmt.Println(plusOne([]int{1, 2, 3}))
}

func plusOne(digits []int) []int {
	i := len(digits) - 1
	carry := 1
	for i >= 0 || carry > 0 {
		if i < 0 {
			digits = append([]int{carry}, digits...)
			break
		}
		sum := digits[i] + carry
		carry = sum / 10
		digits[i] = sum % 10
		i--
	}
	return digits
}

git config --global user.name "sathish9717"
       git config --global user.email " chukkasathish9717@gmail.com"