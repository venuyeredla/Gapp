package list

import (
	trand "Gapp/dsa/rand"
	"Gapp/dsa/types"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	mrand "math/rand"
	"testing"
)

var rand *mrand.Rand

func init() {
	seed := make([]byte, 8)
	if _, err := crand.Read(seed); err == nil {
		rand = trand.ThreadSafeRand(int64(binary.BigEndian.Uint64(seed)))
	} else {
		panic(err)
	}
}

func randstr(length int) types.String {
	slice := make([]byte, length)
	if _, err := crand.Read(slice); err != nil {
		panic(err)
	}
	return types.String(slice)
}

func randhex(length int) types.String {
	slice := make([]byte, length/2)
	if _, err := crand.Read(slice); err != nil {
		panic(err)
	}
	return types.String(hex.EncodeToString(slice))
}

/*

func TestEquals(t *testing.T) {
	al := list.New(10)
	ll := New()
	for i := 0; i < 10; i++ {
		ll.Push(randhex(4))
		al.Push(ll.Last())
	}
	if !ll.Equals(al) {
		t.Fatalf("ll != al, %v != %v", ll, al)
	}
	t.Logf("ll %v", ll)
	t.Logf("al %v", ll)
}

func TestPushPopSize(t *testing.T) {
	al := list.New(10)
	ll := New()
	for i := 0; i < 10; i++ {
		ll.Push(randhex(4))
		al.Push(ll.Last())
	}
	if !ll.Equals(al) {
		t.Fatalf("ll != al, %v != %v", ll, al)
	}
	for i := 0; i < 10; i++ {
		llItem, err := ll.Pop()
		if err != nil {
			t.Fatal(err)
		}
		alItem, err := al.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if !alItem.Equals(llItem) {
			t.Fatalf("llItem != alItem, %v != %v", llItem, alItem)
		}
	}
	if !ll.Equals(al) {
		t.Fatalf("ll != al, %v != %v", ll, al)
	}
}

func TestFrontPushPopSize(t *testing.T) {
	al := list.New(10)
	ll := New()
	for i := 0; i < 10; i++ {
		ll.EnqueFront(randhex(4))
		al.Insert(0, ll.First())
	}
	if !ll.Equals(al) {
		t.Fatalf("ll != al, %v != %v", ll, al)
	}
	for i := 0; i < 10; i++ {
		llItem, err := ll.DequeFront()
		if err != nil {
			t.Fatal(err)
		}
		alItem, err := al.Get(0)
		if err != nil {
			t.Fatal(err)
		}
		err = al.Remove(0)
		if err != nil {
			t.Fatal(err)
		}
		if !alItem.Equals(llItem) {
			t.Fatalf("llItem != alItem, %v != %v", llItem, alItem)
		}
	}
	if !ll.Equals(al) {
		t.Fatalf("ll != al, %v != %v", ll, al)
	}
}
*/

/**  Problems TestCases  **/

func TestAddTwo(t *testing.T) {

	list1 := New()
	list1.EnqueBack(types.Int(2))
	list1.EnqueBack(types.Int(4))
	list1.EnqueBack(types.Int(3))

	list2 := New()
	list2.EnqueBack(types.Int(5))
	list2.EnqueBack(types.Int(6))
	list2.EnqueBack(types.Int(4))

	result := AddTwoNumbers(list1, list2)

	temp := result.Head
	for temp != nil {
		fmt.Print(temp.Data.(types.Int))
		temp = temp.Next
	}

}

func TestReverse(t *testing.T) {

	list1 := New()
	list1.EnqueBack(types.Int(2))
	list1.EnqueBack(types.Int(4))
	list1.EnqueBack(types.Int(3))

	result := ReverseSLL(list1)
	temp := result.Head
	for temp != nil {
		fmt.Print(temp.Data.(types.Int))
		temp = temp.Next
	}
}
