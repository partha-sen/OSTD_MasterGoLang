package main

import(
	"fmt"
)

func pangrams(s string) string {

	for i := byte(65); i <= 90; i++ { 
      var found bool
	 

		for j :=0; j < len(s); j++ {
           if(s[j]==i || s[j]==(i+32)){
				found=true;
				break;
		   } 
		}		  

		if(!found){
			return "not pangram"
		}
	}
	
  
	return "pangram"
}


func main(){	
	fmt.Printf("String is %v ", pangrams("abCdefGhijklmnOpqrStuvWxyz"))
}

