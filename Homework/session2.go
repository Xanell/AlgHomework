package session2

import (
	"slices"
)

func HasDuplicate(a []int) bool {
	seen := make(map[int]bool)
	for _, value := range a {
		if _, ok := seen[value]; ok {
			return true
		}
		seen[value] = true
	}
	return false
}

func TwoItems(prices []int, k int) (int, int, bool) {
	seen := make(map[int]int)
	for i, p := range prices {
		need := k - p
		if j, ok := seen[need]; ok {
			return j, i, true
		}
		seen[p] = i
	}
	return 0, 0, false
}

func BracketsOK(s string) bool {
	pare := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != pare[char] {
				return false
			}
		}
	}
	return len(stack) == 0
}

func IsPalindrome(s string) bool {
	runes := []rune(s)
	left := 0
	right := len(runes) - 1
	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func IsAnagramSort(a, b string) bool {
	runesA := []rune(a)
	runesB := []rune(b)
	slices.Sort(runesA)
	slices.Sort(runesB)
	return slices.Equal(runesA, runesB)
}

func IsAnagramMap(a, b string) bool {
	count := map[rune]int{}
	for _, runes := range a {
		count[runes]++
	}
	for _, runes := range b {
		count[runes]--
	}
	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}

func CanFulfill(stock, order []string) bool {
	left := map[string]int{}
	for _, item := range stock {
		left[item]++
	}
	for _, item := range order {
		if left[item] == 0 {
			return false
		}
		left[item]--
	}
	return true
}

type MinStack struct {
	elements []int
	minimums []int
}

func (s *MinStack) Push(x int) {
	s.elements = append(s.elements, x)
	if len(s.minimums) == 0 || x < s.minimums[len(s.minimums)-1] {
		s.minimums = append(s.minimums, x)
	} else {
		s.minimums = append(s.minimums, s.minimums[len(s.minimums)-1])
	}
}

func (s *MinStack) Pop() int {
	popElement := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	s.minimums = s.minimums[:len(s.minimums)-1]
	return popElement
}

func (s *MinStack) Minimum() int {
	return s.minimums[len(s.minimums)-1]
}
