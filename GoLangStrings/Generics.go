package main

import(
	"fmt"
)

func reverse (s []byte) {
    first := 0
    last := len(s) - 1
    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}

func main(){	
	str:="abcdefgh"
	slice:=[]byte(str)
	reverse(slice)
	fmt.Printf("String is %v ", string(slice))
}

