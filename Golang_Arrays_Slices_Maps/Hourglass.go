package main

import(
	"fmt"
	"math"
)

func hourglassSum(arr [][]int32) int32 {

	rtn:=int32(math.MinInt32)	

	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr[i])-1; j++ {
			sum:=(arr[i-1][j-1]+arr[i-1][j]+arr[i-1][j+1] +arr[i][j]+  arr[i+1][j-1]+arr[i+1][j]+arr[i+1][j+1] )
			if sum>rtn {
				rtn=sum
			}
		}	
	}

   return rtn
}



func main(){	

	a:= [][]int32{{1, 1, 1, 0, 0, 0},
				  {0, 1, 0, 0, 0, 0},
				  {1, 1, 1, 0, 0, 0},
				  {0, 0, 2, 4, 4, 0},
				  {0, 0, 0, 2, 0, 0},
				  {0, 0, 1, 2, 4, 0}}


	fmt.Printf("highest hourglass sum is %v \n", hourglassSum(a))				  

	b:= [][]int32{{-9, -9, -9,  1, 1, 1}, 
				  {0, -9,  0,  4, 3, 2},
				  {-9, -9, -9,  1, 2, 3},
				  {0,  0,  8,  6, 6, 0},
				  {0,  0,  0, -2, 0, 0},
				  {0,  0, 1,  2, 4, 0}}
				  
				 
		fmt.Printf("highest hourglass sum is %v ", hourglassSum(b))

}