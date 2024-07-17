package stackqueue

import (
	"fmt"
	"testing"
)

func TestSPushPop(t *testing.T) {
	fmt.Println("Running push operation")
	s := new(Stack)
	s.Init(iterations / 2)

	for i := 0; i < iterations; i++ {
		s.Push(i)
	}

	for i := iterations - 1; i >= 0; i-- {
		testPopS(t, s, i)
	}
}

func TestInitPushSmallestStack(t *testing.T) {
	// Arrange.
	s := new(Stack)

	// Act.
	s.Init(1)

	for i := 0; i < 4; i++ {
		s.Push(i)
	}

	// Assert.
	for i := 3; i >= 0; i-- {
		testPopS(t, s, i)
	}
}

func TestPeek(t *testing.T) {
	s := new(Stack)
	s.Init(10)

	s.Push("a")
	testPeek(t, s, "a")

	s.Push("b")
	testPeek(t, s, "b")

	s.Pop()
	testPeek(t, s, "a")

	s.Pop()
	testPeek(t, s, nil)
}

func TestSLen(t *testing.T) {
	s := new(Stack)
	s.Init(iterations / 4)

	for i := 0; i < iterations; i++ {
		s.Push(i)
	}

	if l := s.Len(); l != iterations {
		t.Errorf("Stack length was expected to be %v, but is %v", iterations, l)
	}

	s.Pop()
	if l := s.Len(); l != iterations-1 {
		t.Errorf("Stack length was expected to be %v, but is %v", iterations-1, l)
	}
}

func TestSIsEmpty(t *testing.T) {
	s := new(Stack)
	s.Init(2)

	if s.IsEmpty() != true {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)

	if s.IsEmpty() != false {
		t.Errorf("Stack should not be empty")
	}
}

func testPopS(t *testing.T, s *Stack, e interface{}) {
	if v := s.Pop(); v != e {
		t.Errorf("Popping expected %v, got %v", e, v)
	}
}

func testPeek(t *testing.T, s *Stack, e interface{}) {
	if v := s.Peek(); v != e {
		t.Errorf("Peeking expected %v, got %v", e, v)
	}
}

func BenchmarkPushNoResize(b *testing.B) {
	s := new(Stack)
	s.Init(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPushResize(b *testing.B) {
	s := new(Stack)
	s.Init(b.N / 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPopS(b *testing.B) {
	s := new(Stack)
	s.Init(b.N)
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

/* stack problems test */

func TestItoP(t *testing.T) {
	//ItoP("a+b*c+d")
	output := InfixToPostfix("a+b*c+d")
	fmt.Printf("output =  %v  ", output)
	// InfixToPostfix("a+b*(c^d-e)^(f+g*h)-i")
}

func TestBalanced(t *testing.T) {
	balanced := isBalanced("[()]{}{[()()]()}")
	if !balanced {
		t.Fail()
	}
}

func TestCalculator(t *testing.T) {

	output := calculate("3+2*2")
	expected := 7
	if output != expected {
		t.Errorf("Expected = %v and Output = %v ", expected, output)
		t.Fail()
	}
}
