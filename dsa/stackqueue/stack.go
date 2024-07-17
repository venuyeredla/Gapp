// Package stack implements a stack.
package stackqueue

import (
	"fmt"
	"strings"
)

type Stack struct {
	storage []any
	size    int
}

func (stack *Stack) Init(size int) {
	if size == 0 {
		size = 1
	}
	stack.storage = make([]any, size)
	stack.size = -1
}

func (stack *Stack) Push(v any) {
	if stack.size+1 == cap(stack.storage) {
		ns := make([]any, cap(stack.storage)*2)
		copy(ns, stack.storage)
		stack.storage = ns
	}
	stack.size++
	stack.storage[stack.size] = v
}

func (stack *Stack) Pop() any {
	if stack.size < 0 {
		return nil
	}
	v := stack.storage[stack.size]
	stack.storage[stack.size] = nil
	stack.size--
	return v
}

func (stack *Stack) Peek() any {
	if stack.size < 0 {
		return nil
	}
	return stack.storage[stack.size]
}

func (stack *Stack) IsEmpty() bool {
	return stack.Len() == 0
}

func (stack *Stack) Len() int {
	return stack.size + 1
}

/*   Stack problems starting  */

// a+b*c+c -> abc*+d+
// a+b*(c^d-e)^(f+g*h)-i -> abcd^e-fgh*+^*+i-
func InfixToPostfix(expr string) string {
	preceMap := map[string]int{
		"+": 1, "-": 1,
		"%": 2, "/": 2, "*": 2,
		"^": 3,
	}
	var sb strings.Builder
	stack := new(Stack)
	stack.Init(len(expr))
	var precedence = func(s string) int {
		value, exist := preceMap[s]
		if exist {
			return value
		} else {
			return -1
		}
	}

	var isOperand = func(s string) bool {
		if s == "(" || s == ")" {
			return false
		}
		_, exist := preceMap[s]
		return !exist
	}

	var peek = func() string { return stack.Peek().(string) }
	var pop = func() string { return stack.Pop().(string) }

	for _, c := range expr {
		s := string(c)
		//fmt.Print(s)
		if isOperand(s) {
			sb.WriteString(s)
		} else if s == "(" {
			stack.Push(s)
		} else if s == ")" {
			for !stack.IsEmpty() {
				t := peek()
				if t == "(" {
					stack.Pop()
					break
				}
				sb.WriteString(stack.Pop().(string))
			}
		} else {
			for !stack.IsEmpty() && (precedence(s) <= precedence(peek())) {
				sb.WriteString(pop())
			}
			stack.Push(s)
		}
	}
	for !stack.IsEmpty() {
		//fmt.Println(stack.Peek())
		sb.WriteString(pop())
		stack.Pop()
	}
	fmt.Println()

	return sb.String()
}

func isBalanced(expr string) bool {
	stack := new(Stack)
	stack.Init(len(expr))
	charMap := map[string]string{"}": "{", ")": "(", "]": "["}
	for _, c := range expr {
		s := string(c)
		val, exist := charMap[s]
		if exist && !stack.IsEmpty() {
			svalue, _ := stack.Pop().(string)
			if svalue != val {
				return false
			}
		} else {
			stack.Push(s)
		}
	}
	return stack.IsEmpty()
}

func IsWellFormedJson(expr string) bool {
	stack := new(Stack)
	stack.Init(len(expr))
	charMap := map[string]string{"}": "{", "]": "[", "\"": "\"", ",": ",", ":": ":"}
	notJson := false
	for _, c := range expr {
		s := string(c)
		if _, exist := charMap[s]; exist {
			switch s {
			case "}":
				svalue, _ := stack.Pop().(string)
				if !(svalue == "}") {
					notJson = true
				}
			case "]":
				svalue, _ := stack.Pop().(string)
				if !(svalue == "[") {
					notJson = true
				}
			case "\"":

				break
			case ":":
				break
			case ",":
				break
			default:
			}

		} else {
			peek_value := stack.Pop().(string)
			if _, exist := charMap[peek_value]; !exist {
				stack.Push(s)
			}
		}
		if notJson {
			return notJson
		}
	}
	return stack.IsEmpty()
}

func calculate(s string) (answer int) {
	/* stack := new(Stack)
	stack.Init(len(s))

	for _, c := range s {
	//	num_oper := string(c)

	} */
	answer = 7
	return answer
}
