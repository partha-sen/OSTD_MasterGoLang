package stack

import(
	"errors"
)

type any interface{}
type anySlice [] any

func New() Stack {
	a:=make(anySlice, 0, 16)
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

	length:=len(*s)

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




