package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)




func main(){
	
	fmt.Println("Printing three random emoji")	 
	allCode:=textCodes()

	for i := 0; i < 3; i++ {
		emoji.Println(allCode[i])
	}

}

func textCodes() []string{	

	mapCode:=emoji.CodeMap()
	allCode := make([]string, 0, len(mapCode))

    for k:= range mapCode {
		allCode = append(allCode, k)
    }

	return allCode
	
}

