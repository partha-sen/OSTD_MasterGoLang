package stack

type Stack interface{
	Push(a any)
	Clear()
	Pop() (any, error)
	Size() int
}
