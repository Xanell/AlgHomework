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

// CanFulfill - «склад»: можно ли собрать заказ order из товаров stock.
// Каждый товар на складе можно отдать один раз. По желанию.
//
//	функция canFulfill(stock, order):
//	    left = пустой словарь
//	    для каждого item в stock:
//	        left[item] = left[item] + 1
//	    для каждого item в order:
//	        если left[item] == 0:
//	            вернуть нет
//	        left[item] = left[item] − 1
//	    вернуть да
func CanFulfill(stock, order []string) bool {
	// TODO: ваш код
	return false
}

// MinStack - стек, который умеет за O(1) отвечать, какой в нём сейчас минимум.
// Со звёздочкой. Поля структуры придумайте сами.
type MinStack struct {
	// TODO: ваши поля
}

// Push кладёт x на верх стека.
func (s *MinStack) Push(x int) {
	// TODO: ваш код
}

// Pop снимает верхний элемент и возвращает его. Стек не пустой - это гарантирует вызывающий.
func (s *MinStack) Pop() int {
	// TODO: ваш код
	return 0
}

// Minimum возвращает самый маленький элемент в стеке за O(1). Стек не пустой.
func (s *MinStack) Minimum() int {
	// TODO: ваш код
	return 0
}
