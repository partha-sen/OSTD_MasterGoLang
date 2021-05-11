package main

import(
	"fmt"
)

func contains(set []byte, s byte) bool {

	for _, val := range set {
		if val==s {
			return true
		}
	}
   return false
}


func stringConstruction(s string) int32 {

	set := make([]byte, 0, len(s))

	for i, _ := range s {
		if !contains(set, s[i]){
			set=append(set, s[i])
		}
	}
   
   return int32(len(set))
}


func main(){	
	fmt.Printf("Cost is %v ", stringConstruction("abcdabf"))
}


