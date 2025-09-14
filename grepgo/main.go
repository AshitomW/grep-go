package main

import (
	"fmt"
	"grep-go/worker"
	"grep-go/worklist"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexflint/go-arg"
)



func discoverDirectories(wl *worklist.Worklist,path string){
	entries, err := os.ReadDir(path)
	if err !=nil {
		fmt.Println("Read directory error : ", err);
		return 
	}
	
	for _,entry := range entries{
		if entry.IsDir() {
			nextPath := filepath.Join(path,entry.Name())
			discoverDirectories(wl, nextPath)
		}else{
			wl.Add(worklist.NewJob(filepath.Join(path,entry.Name())))
		}
	}

}


var args struct {
	SearchTerm string `arg:"positional,required"`
	SearchDirectory string `arg:"positional"`
}

func main(){
	arg.MustParse(&args)



	var workersWG sync.WaitGroup

	wl := worklist.New(100)
	results := make(chan worker.Result, 100)
	numWorkers := 10
	workersWG.Add(1)
	go func(){
		defer workersWG.Done()
		discoverDirectories(&wl,args.SearchDirectory)
		wl.Finalize(numWorkers)
	}()


	for range numWorkers{
		workersWG.Add(1)
		go func(){
			defer workersWG.Done()
			for{
				workEntry:= wl.Next()
				if workEntry.Path != ""{
					workerResult := worker.FindInFile(workEntry.Path,args.SearchTerm)
					if workerResult != nil {
						for _, r:= range workerResult.Inner{
							results <- r
						}
					}

				}else {
					 return 
				}
			}
		}()
	}

	blockWorkersWG := make(chan struct{})
	go func(){
		workersWG.Wait()
		close(blockWorkersWG)
	}()


	var displayWG sync.WaitGroup
	displayWG.Add(1)
	go func(){
		for{
			select {
				case r:= <- results:
					fmt.Printf("%v[%v]:%v\n",r.Path,r.LineNumber,r.Line)
			  case <- blockWorkersWG:
					if len(results) == 0{
					displayWG.Done()
					return
				}

			}
		}
	}()
	displayWG.Wait()
}


