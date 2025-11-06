package main

import (
	"fmt"
)
/* 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效  */
const (
	LEFT_BRACE = "{"
	RIGHT_BRACE = "}"
	LEFT_BRACKET = "["
	RIGHT_BRACKET = "]"
	LEFT_PAREN = "("
	RIGHT_PAREN = ")"
)

type Stack struct {
	data []string
}
// 压入元素
func (s *Stack) Push(val string) {
	s.data = append(s.data, val)
}
// 弹出栈顶元素
func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	str := s.data[len(s.data) - 1] // 栈顶元素
	s.data = s.data[:len(s.data)-1] // 弹出栈顶元素, 切片
	return str
}
// 查看栈顶元素
func (s *Stack) Peek() string {
	if len(s.data) == 0 {
		return ""
	}
	return s.data[len(s.data)-1] // 栈顶元素
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func isValid(str string) bool {
	stack := &Stack{}
	for v := range str {
		// fmt.Println("str : ", str, " v: ", v, " str[v]: ", str[v], " stack: ", stack)
		fmt.Println("str[v]: ", str[v], " 栈顶元素，",stack.Peek())
		if str[v] == '{' || str[v] == '[' || str[v] == '(' {
			stack.Push(string(str[v]))
		} else if str[v] == '}' {
			if stack.IsEmpty() {
				fmt.Println(str, " } error")
				return false
			} else if stack.Peek() == LEFT_BRACE {
				fmt.Println(str, " : 栈顶取出", stack.Peek())
				stack.Pop()
			} else {
				fmt.Println(str, " } error")
				return false
			}
		} else if str[v] == ']' {
			if stack.IsEmpty() {
				fmt.Println(str, " ] error")
				return false
			} else if stack.Peek() == LEFT_BRACKET {
				fmt.Println(str, " : 栈顶取出", stack.Peek())
				stack.Pop()
			} else {
				fmt.Println(str, " ] error")
				return false
			}
		} else if str[v] == ')' {
			if stack.IsEmpty() {
				fmt.Println(str, " ) error")
				return false
			} else if stack.Peek() == LEFT_PAREN {
				fmt.Println(str, " : 栈顶取出", stack.Peek())
				stack.Pop()
			} else {
				fmt.Println(str, " ) error")
				return false
			}
		}
	}
	if stack.IsEmpty() {
		fmt.Println(str, " : is Valid", )
		return true
	} else {
		fmt.Println(str, ": is not Valid")
		return false
	}
}

func main() {
	var str string = "[{}()]"
	var str2 string = "([)]"
	isValid(str)
	isValid(str2)
}