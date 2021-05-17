package set

type SortedSet []Comparable


func (s *SortedSet) Add(c Comparable){

    length:=len(*s)
    i := 0
    for ; i < length; i++ {
        if (*s)[i].CompareTo(c)>0 {
            break;
        }   
    }

    if i>0 && (*s)[i-1].CompareTo(c)==0 {
       // if element exic    
       return
    }

    if i==length {      
       //adding at the end          
       *s=append(*s, c) 
    }else if i<length {
       //adding in the middle or cart
       *s=append(*s, c)
	   copy((*s)[i+1:], (*s)[i:])
	   (*s)[i]=c
    }
}

func (s *SortedSet) Remove(c Comparable) bool{
    index:=(*s).getIndex(c)
    return (*s).delete(index)
}

func (s *SortedSet) delete(i int) bool{
	if i>=0 && i<len(*s) {
		*s=append((*s)[:i], (*s)[i+1:len(*s)]...)
		return true
	}
	return false
}


func (s *SortedSet) getIndex(c Comparable) int {

    for i, v := range *s {
        if v.CompareTo(c) == 0 {
           return i
        }
    } 

    return -1
}


func (s *SortedSet) Contains(c Comparable) bool{
       for _, v := range (*s) {
            if  v.CompareTo(c) == 0 {
                   return true
            }
       }
       return false
}

