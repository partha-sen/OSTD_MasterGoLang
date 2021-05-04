package main

import(
	"fmt"
)

func rotateLeft(d int, arr []int32) []int32 {
	size:=len(arr);
    x:= d % size 
    rotate:=make([]int32, len(arr))

	for i := 0; i < len(arr); i++ {
		index:= (i+x)%size
		rotate[i]=arr[index]
	}

   return rotate
}



func main(){	
	r:=[]int32{1,2,3}

	fmt.Printf("String is %v ", rotateLeft(4, r))
}

