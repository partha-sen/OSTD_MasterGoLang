package main

import(
	"fmt"
	"sort"
)

func findMedian(arr []int32) int32 {
	
	intArr:=make([]int, len(arr))
	for i, v := range arr {
		intArr[i]=int(v)
	}

    sort.Ints(intArr)
	length:=len(arr)
	return int32(intArr[length/2])
}




func main(){	
	//r:=[]int32{1,2,3,4,5}
	r:=[]int32{5,4,3,2,1}
	fmt.Printf("Median is %v ", findMedian(r))
}

