package main

import (
	"fmt"	
	"github.com/partha-sen/ostd/collection/set"
)


func main(){
    /*
	studentSet:=set.SortedSet{}
	
	studentSet.Add(Student{Name:"Name2", Class:1, Roll:2})	
	studentSet.Add(Student{Name:"Name5", Class:1, Roll:5})
	studentSet.Add(Student{Name:"Name6", Class:1, Roll:6})	
	studentSet.Add(Student{Name:"Name8", Class:1, Roll:8})
	studentSet.Add(Student{Name:"Name9", Class:1, Roll:9})

	s10:=Student{Name:"Name10", Class:1, Roll:10}
	studentSet.Add(s10)
  
	studentSet.Add(Student{Name:"Name3", Class:1, Roll:3})
	studentSet.Add(Student{Name:"Name4", Class:1, Roll:4})
	studentSet.Add(Student{Name:"Name7", Class:1, Roll:7})

	s12:=Student{Name:"Name12", Class:1, Roll:12}
	studentSet.Add(s12)

	s11:=Student{Name:"Name11", Class:1, Roll:11}
	studentSet.Add(s11)

	s1:=Student{Name:"Name1", Class:1, Roll:1}
	studentSet.Add(s1)
	
	fmt.Println(studentSet)*/


	 set:=set.Set{}
	 set.Add("AA")
	 set.Add("B")
	 set.Add("AA")
	 set.Add("C")

	 fmt.Println(set.Contains("C"))

	 all:=set.ToSlice()

	 for _, v := range all {
		 fmt.Println(v)
	 }















}