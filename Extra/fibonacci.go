package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
  n_1 := 1
  n_2 := 0
	return func(x int) int {
		if x > 1 {
			rtn:=n_1+n_2
			n_2=n_1
			n_1=rtn
			return rtn
	    }else{
			return x
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f(i))
	}
}
