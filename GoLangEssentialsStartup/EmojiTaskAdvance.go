package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
	"math/rand"
	"time"
)




func main(){
	
	fmt.Println("Printing three random emoji")
	 
	allCode:=textCodes()
	var index int


	for i := 0; i < 3; i++ {
		index=randomNumber(len(allCode))
		emoji.Println(allCode[index])
		allCode=removeItem(allCode, index)
	}

}


func removeItem(a []string, i int) []string{
	a[i] = a[len(a)-1] // Copy last element to index i.
    a[len(a)-1] = ""   // Erase last element (write zero value).
    a = a[:len(a)-1]
	return a
}


func textCodes() []string{	

	mapCode:=emoji.CodeMap()
	allCode := make([]string, 0, len(mapCode))

    for k:= range mapCode {
		allCode = append(allCode, k)
    }

	return allCode
	
}


func randomNumber(len int) int{
    rand.Seed(time.Now().UnixNano())
	val:=rand.Intn(len)
	return val
}

