package stack

import(
	"errors"
)

type any interface{}
type Stack [] any


func (s *Stack) Push(a any){
    *s=append(*s, a)
}

func (s *Stack) Clear(){
    *s=(*s)[:0]
}


func (s *Stack) Pop() (any, error) {

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

func (s *Stack) Size() int {
	return len(*s)
}




