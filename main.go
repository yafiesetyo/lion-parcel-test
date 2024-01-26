package main

import "fmt"

func main() {
	isPalindrome := palindrome("abba")
	fmt.Println("1. Is palindrome", isPalindrome)
	res := max([]int{1, 10, 12, 20, 5})
	fmt.Println("2. Max: ", res)

	fmt.Print("3. Steps")
	step(5)
	f := factorial(8)
	fmt.Println("4. Fatorial", f)
}

// palindrome
func palindrome(input string) bool {
	median := len(input) / 2
	i, j := 0, len(input)-1
	result := true

	for i < median && j > median {
		if input[i] == input[j] {
			i++
			j--
		} else {
			result = false
			break
		}
	}

	return result
}

// sort (max)
func max(in []int) int {
	sorted := mergeSort(in)
	return sorted[0]
}

func mergeSort(in []int) []int {
	if len(in) <= 1 {
		return in
	}

	mid := len(in) / 2
	left := in[:mid]
	right := in[mid:]

	l := mergeSort(left)
	r := mergeSort(right)

	return merge(l, r)
}

func merge(left, right []int) []int {
	merged := []int{}
	leftSize := len(left)
	rightSize := len(right)

	var (
		l = 0
		r = 0
	)

	for l < leftSize && r < rightSize {
		if left[l] > right[r] {
			merged = append(merged, left[l])
			l++
		} else {
			merged = append(merged, right[r])
			r++
		}
	}

	for l < leftSize {
		merged = append(merged, left[l])
		l++
	}

	for r < rightSize {
		merged = append(merged, right[r])
		r++
	}

	return merged
}

// steps
func step(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// factorial
func factorial(in int) int {
	if in == 0 || in == 1 {
		return 1
	}

	return in * factorial(in-1)
}
