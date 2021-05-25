package main

import(
	"strings"
)

func ext(name string) string{
	var ext string
	arr:=strings.Split(name,".")
	length:=len(arr)
	if length>1 {
		ext=arr[length-1]
	}
	return ext
}