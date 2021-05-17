package main

import (
	"fmt"	
	"github.com/partha-sen/ostd/collection/set"
	"github.com/partha-sen/ostd/collection/list"
)


func main(){

	set:=set.SortedSet{}

	set.Add(Student{Name:"Ram", Class:1, Roll:1})
	set.Add(Student{Name:"Sham", Class:1, Roll:2})
	set.Add(Student{Name:"Jadu", Class:1, Roll:3})
	
	fmt.Println(set)

	set.Remove(Student{Name:"aaaa", Class:1, Roll:3})

	fmt.Println(set)

	list:=list.List{}
	list.Add(Student{Name:"Ram", Class:1, Roll:1})
	list.Add(Student{Name:"Sham", Class:1, Roll:1})
	list.Add(Student{Name:"Ram", Class:1, Roll:1})

	fmt.Println(list)

	list.Delete(1)
    
	fmt.Println(list)

}