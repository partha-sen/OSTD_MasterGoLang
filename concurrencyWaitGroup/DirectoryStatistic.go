package main

import(
	"fmt"
	"io/ioutil"
	"sync"
	"time"
	"path/filepath"
)


func calculateExtCount(dir string, chanMapCount chan map[string]int, wg *sync.WaitGroup){	
	defer wg.Done()
	
	var extMap =map[string]int{} 
	files, err := ioutil.ReadDir(dir)
	if err!=nil{		
		return 
	}

	for _, file := range files {			
		if file.IsDir() {
			wg.Add(1)
            go calculateExtCount(dir+"/"+file.Name(), chanMapCount, wg)            
		}else{
			ext:=filepath.Ext(file.Name())
			extMap[ext] += 1					
		}		
	}
	chanMapCount<-extMap
}


func getExtCount(dir string, chanMapCount chan map[string]int){
	var wg sync.WaitGroup
	wg.Add(1)
	go calculateExtCount(dir, chanMapCount, &wg)
	wg.Wait()
	close(chanMapCount)
}


func main(){

	start := time.Now()

	var extnMap =extCountMap{} 
	var rootDir = "C:/pss"
	var chanMapCount = make(chan map[string]int, 10)
	go getExtCount(rootDir, chanMapCount)

	for item := range chanMapCount {
		extnMap.add(item)
	}

	fmt.Println(extnMap)
	duration := time.Since(start)
	fmt.Println(duration)

}