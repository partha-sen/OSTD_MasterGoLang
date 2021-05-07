package main

import(
	"fmt"
)

type SortedSet []string

func (set *SortedSet) isPrevDuplicate(startIndex int, currIndex int, val string) bool{
    if currIndex > startIndex {
		if (*set)[currIndex-1]==val{
			return true
		}
	}	
	return false
}


func (set *SortedSet) insert(start int, end int, val string) {
    if (end-start)<=10 {	  
      for i := start; i <= end; i++ {	
		if (*set)[i]>val {		
			if (*set).isPrevDuplicate(start, i, val) {			
				return
			}	
			*set=append((*set),val)
			copy((*set)[i+1:], (*set)[i:])
			(*set)[i]=val
			break;
		}
	  }
	}else{
		mid:=(end-start)/2 + start
	    if val<(*set)[mid] {
				end=mid
		}else{
			start=mid
		}
    	 (*set).insert(start, end, val)
	  }
}


func (set *SortedSet) add(s string){
	if (len(*set)>0){
		if (*set)[len(*set)-1] < s {
           *set=append(*set, s)
		}else{
			(*set).insert(0, len(*set)-1, s)
		}		
	}else{
	   *set=append(*set, s)
	}
	
}

func (set *SortedSet) remove(s string){
	index:=(*set).getIndex(s)
    (*set).delete(index)
}

func (set *SortedSet) delete(i int) bool{
	if i>=0 && i<len(*set) {
		first:=(*set)[:i]
		last:=(*set)[i+1:len(*set)]
		*set=append(first, last...)
		return true
	}
	return false
}


func (set *SortedSet) findIndex(start int, end int, val string) int{
    if (end-start)<=10 {
      for i := start; i <= end; i++ {
		if (*set)[i]==val {
			return i
		}
	  }
	}else{
		mid:=(end-start)/2+start
	    if val<(*set)[mid] {
				end=mid
		}else{
			start=mid
		}
    	return (*set).findIndex(start, end, val)
	  }

	return -1
}


func (set *SortedSet) getIndex(val string) int{
   return (*set).findIndex(0, len(*set)-1, val)
}  


func main(){
    set:=SortedSet{}
	fmt.Printf("C=%v len=%d cap=%d \n", set,  len(set), cap(set))	
	var str byte
	for i := 90; i >64; i-- {
		if i!=80 {
			str=byte(i)	
			fmt.Printf("%v",string(str))
			set.add(string(str))
		}	
	 
	}	
	fmt.Println("")
	set.add("P")

	fmt.Printf("set=%v \n", set)	

}





