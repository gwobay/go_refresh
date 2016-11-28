//sample codes
package main

//import "time"
import "fmt"
import "math/rand"
import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t==nil {
	return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}
func Insert(t *tree.Tree, val int) *tree.Tree {
	switch {
	case t==nil:
			new1 := &tree.Tree{nil, val, nil}
			return new1
	case val <= t.Value:
			if t.Left==nil {
				t.Left = &tree.Tree{nil, val, nil}
				return t.Left
			}
			Insert(t.Left, val)
	case val > t.Value:
			if t.Right==nil {
				t.Right = &tree.Tree{nil, val, nil}
				return t.Right
			}
			Insert(t.Right, val)	
	default:
	}
	return t
}
func CreateNew(initV int, n int) *tree.Tree {
	newT := &tree.Tree{nil, initV, nil}
	for i:=1; i<n; i++ {
		Insert(newT, int(rand.Int31()))
	}
	return newT
}
//logic is
//1.create execution job function, i.e., Walk
//2.create a thread to exectue the job with
//  preset parameters
//3. use the parameter returned for comparison
func Dispatcher(t1 *tree.Tree, ch0 chan int) {
	go func() {
		Walk(t1, ch0)
		close(ch0)
	}()
}
//----------
// CheckIfSameSeq determines whether the trees
// t1 and t2 contain the same values.
func CheckIfSameSeq(t1, t2 *tree.Tree) bool{
	//use returned value for comparison
	//they are sequences so to use 2 channels
	//to compare 1 by 1
	//those 2 channels are also carried by the
	//thread
	ch1 := make(chan int)
	ch2 := make(chan int)
	Dispatcher(t1, ch1)
	Dispatcher(t2, ch2)
	for {
		v1, failed1 := <-ch1
		v2, failed2 := <-ch2
		fmt.Println(v1, v2)
		if !failed1 || !failed2 {
			//here I need someone to
			//stop the channel so as to
			//notify the walk process finishs???
			//so I need another object to do that
			//this is how the Dispatcher comes in
			return (failed1 == failed2)
		}
		if v1 != v2 {
			//not the same sequence stop
			break
		}		
	}
	//if they are the same 
	//returned already
	return false
}

func LocateSame(t1, t2 *tree.Tree, n int) (bool, []int){
	//use returned value for comparison
	//they are sequences so to use 2 channels
	//to compare 1 by 1
	//those 2 channels are also carried by the
	//thread
	ch1 := make(chan int)
	ch2 := make(chan int)
 	var chOkBuf []int
	//ch2Buf := make(chan int, n)
	Dispatcher(t1, ch1)
	Dispatcher(t2, ch2)
	var ch1Buf, ch2Buf []int
	for i:=0; i<=n; i++{
		v1, failed1 := <-ch1
		v2, failed2 := <-ch2
		//fmt.Println(v1, v2)
		if !failed1 || !failed2 {
			break
		}
		ch1Buf=append(ch1Buf, v1)
		ch2Buf=append(ch2Buf, v2)
	}
	j := 0
	for i:=0; i<n; i++ {
		fmt.Println(ch1Buf[i])
		for ; j<n; j++ {
			fmt.Println("===>", ch2Buf[j])
			if ch2Buf[j] < ch1Buf[i]{
				continue
			}
			if ch2Buf[j]==ch1Buf[i] {
			fmt.Println("Got it!!", ch1Buf[i], "at ",i,"=", ch2Buf[j], "at", j)
				chOkBuf=append(chOkBuf, ch1Buf[i])
			}
			break
		}
	}
		
}
func main() {
	new1, new2 := CreateNew(1, 20), CreateNew(2, 20)
	fmt.Println(LocateSame(new1, new2, 20))
	new1, new2 = tree.New(6), tree.New(4)
	fmt.Println(LocateSame(new1, new2, 10))
	fmt.Println(CheckIfSameSeq(tree.New(4), tree.New(4)))
}