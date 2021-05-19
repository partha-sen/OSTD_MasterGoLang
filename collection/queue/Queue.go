package queue

import(
	"errors"
)

type any interface{}
type Queue [] any


func (q *Queue) Add(a any){
    *q=append(*q, a)
}

func (s *Queue) Clear(){
    *s=(*s)[:0]
}


func (q *Queue) Remove() (any, error) {

	var (
		rtn any
		err error
	) 

	length:=len(*q)

	if length > 0 {
		rtn = (*q)[0]
		*q=(*q)[1:length]
	}else{
        err=errors.New("Record not found")
	}  

	return rtn, err
}

func (q *Queue) Size() int {
	return len(*q)
}




