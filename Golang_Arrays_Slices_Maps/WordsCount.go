package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"strings"
)

func countWord(strArr []string) map[string]string{
	var wc = make(map[string]string)
	for i, str := range strArr {
		words:=strings.Split(str, " ")
		for _, word := range words {
			if val, ok := wc[word]; ok {
			    wc[word] =val + " " + strconv.Itoa(i+1)
			}else{
				wc[word] = strconv.Itoa(i+1)
			}
		}
	}
  return wc    
}


func main(){	
	
	 scanner := bufio.NewScanner(os.Stdin)
     fmt.Print("No of input:")
     scanner.Scan()
	 text:= scanner.Text()
	 count, _ := strconv.Atoi(text)
     
	 inputStrArr:=make([]string, count)
     
    for i := 0; i < count; i++ {
		scanner.Scan()
	    text:= scanner.Text()
		inputStrArr[i]=text
	}

	 fmt.Printf("%v \n", countWord(inputStrArr))
	

}
