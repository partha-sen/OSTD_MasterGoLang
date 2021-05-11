package main

import(
	"fmt"
)

func camelcase(s string) int32 {
	var length int32 
	for _, v:= range s {
		if(v>=65 && v<=90){
			length=length+1
		}
    }
	return length+1
}

func main(){	
	fmt.Printf("No of words %d ", camelcase("abcOne"))
}