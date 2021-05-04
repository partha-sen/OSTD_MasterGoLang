package main

import(
	"fmt"
)

func reverseArray(a []int32) []int32 {
  
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		temp:=a[i]
        a[i]=a[j]
		a[j]=temp
    }
  return a
}




func main(){	
	r:=[]int32{1,2,3,4}
	fmt.Printf("Revers Array is %v ", reverseArray(r))
}

