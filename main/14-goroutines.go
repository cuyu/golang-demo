package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"strconv"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		time.Sleep(100 * time.Millisecond)
	}
	c <- sum // send sum to c
}

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // close the channel
}

func fibSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			// channel c may get more results if it has buffer (has buffer means it is ready and can be selected)
			fmt.Printf("put result %s to channel\n", strconv.Itoa(x))
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func ClockBoom() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Exercise!
func _WalkTree(t *tree.Tree, ch chan int, isRoot bool) {
	// Walks the tree t sending all values from the tree to the channel ch.
	if t.Right != nil {
		_WalkTree(t.Right, ch, false)
	}
	ch <- t.Value
	if t.Left != nil {
		_WalkTree(t.Left, ch, false)
	}
	if isRoot {
		close(ch)
	}
}

func WalkTree(t *tree.Tree, ch chan int) {
	_WalkTree(t, ch, true)
}

func SameTree(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go WalkTree(t1, c1)
	go WalkTree(t2, c2)
	for {
		i1, ok1 := <-c1
		i2, ok2 := <-c2
		if i1 != i2 || ok1 != ok2{
			return false
		} else if !ok1 {
			return true
		}
	}
}

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime. Format is `go func(x,y,z)`
	// Note the evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of `f` happens in the new goroutine.
	// Goroutines run in the same address space, so access to shared memory must be synchronized (like use `channel`).
	go say("world")
	say("hello")

	// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
	// (The data flows in the direction of the arrow.)
	// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
	// (channel有点像是Python中的Queue，只不过Go中从底层就支持并优化了，注意下面的x和y的值并不是确定的，要看哪个goroutine先执行完成)
	s := []int{7, 2, 8, -9, 4, 0}
	// Channels can be buffered. Provide the buffer length as the second argument to `make`.
	// (类似于Python中Queue的max size，buffer满了会block要进入channel的goroutine)
	c := make(chan int, 2)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	// A sender can close a channel to indicate that no more values will be sent.
	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	// (Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.)
	d := make(chan int, 10)
	go fib(cap(d), d)
	for i := range d {
		fmt.Println(i)
	}

	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
	e := make(chan int, 5)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-e)
		}
		quit <- 0
	}()
	fibSelect(e, quit)

	// The default case in a select is run if no other case is ready.
	ClockBoom()

	// Exercise!
	tree1 := tree.New(1)
	tree2 := tree.New(2)
	fmt.Println(SameTree(tree1, tree2))
}
