package main

import(
	"fmt"	
	"github.com/partha-sen/ostd/struct_interface/stack"
)


type parents struct{	 
	mother string	
	father string
}



func main(){	
	
	stack:=stack.New()
	stack.Push("Partha")
	stack.Push(1)
	stack.Push(parents{father:"Karamchand", mother:"Putlibai"})
	stack.Push(1.1)

	for x, err:=stack.Pop(); err == nil; x, err=stack.Pop() {
		switch v := x.(type) {
			case int:
				fmt.Println("int:", v)
			case float64:
				fmt.Println("float64:", v)
			case string:
				fmt.Println("string:", v)	
			case parents:
				fmt.Println("Father:", v.father, " Mother:", v.mother)	
			default:
				fmt.Println("unknown")
	    }
	}	

}