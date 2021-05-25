package main

type extCountMap map[string]int

func (sumMap *extCountMap) add(itemMap extCountMap){
    for key, value := range itemMap {
       if oldVal, ok:=(*sumMap)[key]; ok {
		  (*sumMap)[key]=oldVal+value
	   }else {
		(*sumMap)[key]=value
	   }
	}
}