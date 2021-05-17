package list

type List []interface{}

func (l *List) Add(a interface{}){
	*l=append(*l, a)
}

func (l *List) Delete(i int) bool{
	if i>=0 && i<len(*l) {
		first:=(*l)[:i]
		last:=(*l)[i+1:len(*l)]
		*l=append(first, last...)
		return true
	}
	return false
}

func (l *List) Contains(a interface{}) bool{
	for _, v := range *l {
		if v == a {
			return true
		}
	}
   return false
}

