package main

import(
	"fmt"
)

type list []string


func NewList() list{
	return make(list,0, 16)
}


func (l *list) add(c string){
       *l=append(*l, c)
}

func (l *list) delete(i int) bool{
	if i>=0 && i<len(*l) {
		first:=(*l)[:i]
		last:=(*l)[i+1:len(*l)]
		*l=append(first, last...)
		return true
	}
	return false
}

func (l *list) contains(s string) bool{
	for _, str := range *l {
		if str==s {
			return true
		}
	}
   return false
}

func main(){
    lst:=NewList();
	fmt.Printf("C=%v len=%d cap=%d \n", lst,  len(lst), cap(lst))	
	lst.add("One")
	fmt.Printf("C=%v len=%d cap=%d \n", lst,  len(lst), cap(lst))	
	lst.add("Two")
	fmt.Printf("C=%v len=%d cap=%d \n", lst,  len(lst), cap(lst))
	lst.add("Three")
	fmt.Printf("C=%v len=%d cap=%d \n", lst,  len(lst), cap(lst))
	lst.delete(0)
	fmt.Printf("C=%v len=%d cap=%d \n", lst,  len(lst), cap(lst))
	fmt.Printf("contains=%v \n", lst.contains("One"))	
	fmt.Printf("contains=%v \n", lst.contains("Two"))	
}

