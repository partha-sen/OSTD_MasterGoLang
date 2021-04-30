package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)




func main(){
	
	fmt.Println("Printing three random emoji")	 
	allCode:=textCodes()

	emoji.Println(allCode[0])
	emoji.Println(allCode[1])
	emoji.Println(allCode[2])

}

func textCodes() []string{	

	mapCode:=emoji.CodeMap()
	allCode := make([]string, 0, len(mapCode))

    for k:= range mapCode {
		allCode = append(allCode, k)
    }

	return allCode
	
}

