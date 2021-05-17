package main


type Student struct {
	Name string
	Class int
	Roll int
} 


func (s1 Student) CompareTo(s2 interface{}) int{
	rtn:=s1.Class-s2.(Student).Class
	if rtn == 0 {
		rtn=s1.Roll-s2.(Student).Roll
	}	
	return rtn;
}







