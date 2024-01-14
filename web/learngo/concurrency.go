package learngo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func TestGoRoutine() {
	wg.Add(2)
	go fun1()
	go fun2()
	wg.Wait()
}
func fun1() {
	for i := 0; i < 10; i++ {
		fmt.Println("fun1,  ->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}
func fun2() {
	for i := 0; i < 10; i++ {
		fmt.Println("fun2,  ->", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}

var wait sync.WaitGroup
var count int

func increment(s string) {
	for i := 0; i < 10; i++ {
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)

	}
	wait.Done()

}
func RaceTest() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value ", count)
}

var mutex sync.Mutex

func Mutexincrement(s string) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)
		mutex.Unlock()

		//  Atmoic variable    atomic.AddInt64(&count,1)

	}
	wait.Done()

}
func MutexTest() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value ", count)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func ChannelTest() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
