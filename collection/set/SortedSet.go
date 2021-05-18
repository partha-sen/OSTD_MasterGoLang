package set

type SortedSet []Comparable

const serchLimit = 10


func (s *SortedSet) insert(start int, end int, c Comparable){

    length:=len(*s)
    i := start

    for ; i <= end; i++ {
        if (*s)[i].CompareTo(c)>0 {
            break;
        }   
    }

    if i>0 && (*s)[i-1].CompareTo(c) == 0 {
       // if element exic    
       return
    }

    if i==length {      
       //adding at the end          
       *s=append(*s, c) 
    }else if i<length {
       //adding in the middle or beginning
       *s=append(*s, c)
	   copy((*s)[i+1:], (*s)[i:])
	   (*s)[i]=c
    }

}

func (s *SortedSet) findRange(start int, end int, val Comparable) (int, int){
    if (end-start) <= serchLimit {
       return start, end
    }else{
		mid:=(end-start)/2 + start
	    if val.CompareTo((*s)[mid]) < 0 {
			end=mid
		}else{
			start=mid
		}
    	 (*s).findRange(start, end, val)
	  }
    return start, end
}



func (s *SortedSet) Add(c Comparable){
    start, end := (*s).findRange(0, len(*s)-1, c)
    (*s).insert(start, end, c)
}

func (s *SortedSet) Remove(c Comparable) bool{
    start, end := (*s).findRange(0, len(*s)-1, c)
    index:=(*s).getIndex(start, end, c)
    return (*s).delete(index)
}

func (s *SortedSet) delete(i int) bool{
	if i>=0 && i<len(*s) {
		*s=append((*s)[:i], (*s)[i+1:len(*s)]...)
		return true
	}
	return false
}


func (s *SortedSet) getIndex(start int, end int, c Comparable) int {

    for i := start; i <= end; i++ {
        if (*s)[i].CompareTo(c) == 0 {
            return i
        }
    }

    return -1
}


func (s *SortedSet) Contains(c Comparable) bool{
    start, end := (*s).findRange(0, len(*s)-1, c)
    index:=(*s).getIndex(start, end, c)
    if index >= 0 {
        return true
    }else{
        return false
    }
}

func (s *SortedSet) Size() int {	
	return len(*s)
}

