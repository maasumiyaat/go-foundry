package strings

import (
	"strings"
	"unicode"
)

func TrimString() string {
	str := "\t Hello, World\n "
	// fmt.Printf("Before Trim Length: %d String:%v\n", len(str), str)
	trim := strings.TrimSpace(str)
	// fmt.Printf("After Trim Length: %d String:%v\n", len(trim), trim)

	return trim
}

func IsOnlyLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func AllUpperCase(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func IfContains(w, l string) bool {
	return strings.Contains(w, l)
}

func DefangIPaddr(addr string) string {
	return strings.Replace(addr, ".", "[.]", -1)
}

func FinalValueAfterOperations(operations []string) int {
	x := 0

	for _, i := range operations {
		if i == "++X" || i == "X++" {
			x++
		} else if i == "--X" || i == "X--" {
			x--
		}
	}
	return x
}

func NumJewelsInStones(jewels string, stones string) int {

	count := 0

	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(jewels); j++ {
			if stones[i] == jewels[j] {
				count++
			}
		}
	}

	return count
}

func MostWordsFound(sentences []string) int {
	wordCounter := 0

	for _, r := range sentences {
		splitted := strings.Split(r, " ")
		if len(splitted) > wordCounter {
			wordCounter = len(splitted)
		}
	}
	return wordCounter
}

func Interpret(command string) string {
	command = strings.Replace(command, "()", "o", -1)
	command = strings.Replace(command, "(al)", "al", -1)

	return command
}

func InterpretLoop(command string) string {
	ans := ""

	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			ans = ans + "G"
		} else if command[i] == '(' {
			if command[i+1] == ')' {
				ans = ans + "o"
			} else if command[i+1] == 'a' {
				ans = ans + "al"
			}
		}
	}

	return ans
}

func stackLength(STACK []string) int {
	// fmt.Println("STACK: ", STACK, "len: ", len(STACK))
	return len(STACK)
}

func stackTop(STACK []string) string {
	if stackLength(STACK) == 0 {
		return "0"
	} else {
		return STACK[len(STACK)-1]
	}
}

// func stackPush(STACK []string, val string) {
// 	STACK = append(STACK, val)
// 	return
// }

// func stackPop(STACK []string) {
// 	STACK = STACK[:len(STACK)-1]
// }

func IsValidParentheses(s string) bool {

	var STACK []string

	for _, i := range s {

		if i == '(' || i == '{' || i == '[' {
			// stackPush(STACK, string(i))
			STACK = append(STACK, string(i))
		} else {
			top := stackTop(STACK)
			if i == ')' && top == "(" ||
				i == '}' && top == "{" ||
				i == ']' && top == "[" {
				STACK = STACK[:len(STACK)-1]
				// stackPop(STACK)
			}
		}
	}

	if stackLength(STACK) == 0 {
		return true
	} else {
		return false
	}
}

// return the index of the first occurrence of needle in haystack
func LC28(haystack string, needle string) int {
	// return strings.Index(haystack, needle)
	if needle == "" {
		return 0
	}

	hLen := len(haystack)
	nLen := len(needle)

	for i := 0; i <= hLen-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}

	return -1
}

func LongestCommonPrefix(strs []string) string {
	var result string
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(strs[i]); j++ {
			if i == 0 {
				result = result + string(strs[i][j])
			} else {
				if j >= len(result) || strs[i][j] != result[j] {
					result = result[:j]
					break
				}
			}
		}
	}
	return result
}

func LongestCommonPrefixOptimized(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
