package stack

import(
	"fmt"	
	"errors"
)

type any interface{}
type anySlice [] any
const initCapacity=16

func New() Stack {
	a:=make(anySlice, 0, initCapacity)
	return &a
}


func (s *anySlice) Push(a any){
    *s=append(*s, a)
}

func (s *anySlice) Clear(){
    *s=(*s)[:0]
}


func (s *anySlice) Pop() (any, error) {

	var (
		rtn any
		err error
	) 

	length, capacity :=len(*s), cap(*s) 
	
	fmt.Println("len",len(*s),"cap",cap(*s))

	if( capacity>initCapacity && length <= capacity/2 ){
		//down size
		*s=append(anySlice{}, (*s)[:length]...)
	}

	if length > 0 {
		rtn = (*s)[length-1]
		*s=(*s)[:length-1]
	}else{
        err=errors.New("Record not found")
	}  

	return rtn, err
}

func (s *anySlice) Size() int {
	return len(*s)
}




