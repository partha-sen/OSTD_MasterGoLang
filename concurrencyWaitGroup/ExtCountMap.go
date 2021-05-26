package main

type extCountMap map[string]int

func (sumMap *extCountMap) add(itemMap extCountMap){
    for key, value := range itemMap {
		(*sumMap)[key] += value       
	}
}