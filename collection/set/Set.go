package set

type Set map[interface{}] interface{}


func (s *Set) Add(key interface{}){
	(*s)[key]=nil
}

func (s *Set) Contains(key interface{}) bool{
	_, ok := (*s)[key]
	return ok
}

func (s *Set) Remove(key interface{}){
   delete(*s, key)
}

func (s *Set) Size() int {	
	return len(*s)
}


func (s *Set) ToSlice() []interface{} {
	keys := make([]interface{}, 0, len(*s))
	for k := range *s {
        keys = append(keys, k)
    }
	return keys
}












