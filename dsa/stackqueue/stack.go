// Package stack implements a stack.
package stackqueue

import (
	"fmt"
	"strings"
)

// S is the internal representation of the data structure.
type Stack struct {
	storage []interface{}
	i       int
}

// Init initializes the stack data structure.
// A stack must be initialized before it can be used.
// O(1)
func (s *Stack) Init(size int) {
	// the stack needs to have at least capacity for 1 element
	// defaulting here instead of panicing to allow for benchmarking
	if size == 0 {
		size = 1
	}
	s.storage = make([]interface{}, size)
	s.i = -1
}

// Push adds a new element to the top of the stack.
// O(1)
func (s *Stack) Push(v interface{}) {
	// dynamically increase the size of storage as needed
	if s.i+1 == cap(s.storage) {
		ns := make([]interface{}, cap(s.storage)*2)
		copy(ns, s.storage)
		s.storage = ns
	}

	s.i++
	s.storage[s.i] = v
}

// Pop removes the top element from the stack.
// O(1).
func (s *Stack) Pop() interface{} {
	if s.i < 0 {
		return nil
	}

	v := s.storage[s.i]
	s.storage[s.i] = nil
	s.i--

	return v
}

// Peek returns the top element from the stack without removing it.
// O(1)
func (s *Stack) Peek() interface{} {
	if s.i < 0 {
		return nil
	}
	return s.storage[s.i]
}

func (s *Stack) PeekStr() string {
	if s.i < 0 {
		return ""
	}
	return fmt.Sprint(s.storage[s.i])
}

// IsEmpty returns true if the stack has no elements.
// O(1)
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

// Len returns the number of elements in the stack.
// O(1)
func (s *Stack) Len() int {
	return s.i + 1
}

/*   Stack problems starting  */

var preceMap map[string]int

func createPREMap() {
	preceMap = map[string]int{
		"+": 1, "-": 1,
		"%": 2, "/": 2, "*": 2,
		"^": 3,
	}
}

// a+b*c+c -> abc*+d+
// a+b*(c^d-e)^(f+g*h)-i -> abcd^e-fgh*+^*+i-
func ItoP(expr string) {
	createPREMap()
	var sb strings.Builder
	stack := new(Stack)
	stack.Init(10)
	for _, c := range expr {
		s := string(c)
		//fmt.Print(s)
		if isOperand(s) {
			sb.WriteString(s)
		} else if s == "(" {
			stack.Push(s)
		} else if s == ")" {
			for !stack.IsEmpty() {
				t := peek(stack)
				if t == "(" {
					stack.Pop()
					break
				}
				sb.WriteString(pop(stack))

			}

		} else {
			for (precedence(s) <= precedence(peek(stack))) && !stack.IsEmpty() {
				sb.WriteString(pop(stack))
			}
			stack.Push(s)
		}
	}
	for !stack.IsEmpty() {
		//fmt.Println(stack.Peek())
		sb.WriteString(pop(stack))
		stack.Pop()
	}
	fmt.Println(sb.String())
}

func pop(s *Stack) string {
	value, _ := s.Pop().(string)
	return value
}

func peek(s *Stack) string {
	value, _ := s.Pop().(string)
	return value
}

func precedence(s string) int {
	value, exist := preceMap[s]
	if exist {
		return value
	} else {
		return -1
	}

}

func isOperand(s string) bool {
	if s == "(" || s == ")" {
		return false
	}
	_, exist := preceMap[s]
	return !exist
}

// [()]{}{[()()]()}
//
//	([]){()}
func IsBalanced(expr string) bool {
	stack := new(Stack)
	stack.Init(50)
	charMap := map[string]string{"}": "{", ")": "(", "]": "["}
	for _, c := range expr {
		s := string(c)
		val, exist := charMap[s]
		if exist && !stack.IsEmpty() {
			if pop(stack) == val {

			} else {
				return false
			}
		} else {
			stack.Push(s)
		}

	}
	return stack.IsEmpty()
}

// {"":"","":b,"c":[1,2,3]}
var charMap map[string]string

func IsWellFormedJson(expr string) bool {
	stack := new(Stack)
	stack.Init(50)
	charMap = map[string]string{"}": "{", "]": "[", "\"": "\"", ",": ",", ":": ":"}
	notJson := false
	for _, c := range expr {
		s := string(c)
		if IsSymbol(s) {
			switch s {
			case "}":
				if !(pop(stack) == "}") {
					notJson = true
				}
				break
			case "]":
				if !(pop(stack) == "[") {
					notJson = true
				}
				break
			case "\"":

				break
			case ":":
				break
			case ",":
				break
			default:

			}

		} else {
			if !(IsSymbol(peek(stack)) && IsSymbol(peek(stack))) {
				stack.Push(s)
			}
		}
		if notJson {
			return notJson
		}
	}
	return stack.IsEmpty()
}

func IsSymbol(s string) bool {
	_, exist := charMap[s]
	return exist
}
