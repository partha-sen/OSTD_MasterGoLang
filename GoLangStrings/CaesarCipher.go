package main

import(
	"fmt"
)

func caesarCipher(s string, k int32) string {

	rtn:= make([]byte, len(s))

    for i := 0; i < len(s); i++ {
		rtn[i] = s[i] + byte(k)
		if(rtn[i]>'z'){
			rtn[i]=rtn[i]-'z'+'a'-1
		}
	}
	
	return string(rtn)
}


func main(){		
	fmt.Printf("Cipher is: %v ", caesarCipher("abcy", 2))
}