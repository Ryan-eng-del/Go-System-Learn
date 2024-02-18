package main

// package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	// wg := &sync.WaitGroup{}
	// wg.Add(1)

	// go func(wg *sync.WaitGroup) {
	// 	fmt.Println("wait")
	// 	wg.Wait()
	// }(wg)

	// time.Sleep(400 * time.Millisecond)

	// go func(wg *sync.WaitGroup) {
	// 	time.Sleep(3 * time.Second)
	// 	wg.Done()
	// }(wg)

	s := make(chan int64, 2)
	go func(){
		time.Sleep(3 * time.Second)
		s <- 2
	}()

	fmt.Println("s", <- s)

}

// func main() {
// 	WalkDir("/Users/max/Documents/coding/Backend")
// }
func WalkDir(dirs ...string) {

	if len(dirs) == 0 {
		dirs = []string{"./"}
	}

	fileSizeTotal := make(chan int64, 1)

	wg := &sync.WaitGroup{}
	
	for _ ,dir := range dirs {
		wg.Add(1)
		go walkDir(fileSizeTotal, dir, wg)
	}

	go func (wg *sync.WaitGroup) {
		wg.Wait()
		close(fileSizeTotal)
	}(wg)

	fileSize := make(chan int64, 1)
	fileFiles := make(chan int64, 1)

	go func (sizes chan <- int64, files chan <- int64){
		var fileSize, fileFiles int64
		for size := range fileSizeTotal {
			fileFiles++
			fileSize += size
		}
		sizes <- fileSize
		files <- fileFiles
	}(fileSize, fileFiles)

	fmt.Printf("size: %.2fMB, files: %d\n", float64(<-fileSize)/1024/1024, <-fileFiles)
}



func walkDir(fileSizeCh chan <- int64, dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, fileItem := range filesInfo(dir) {
		if fileItem.IsDir() {
			subDir :=  filepath.Join(dir, fileItem.Name())
			wg.Add(1)
			go walkDir(fileSizeCh, subDir, wg)
		} else {
			fileSizeCh <- fileItem.Size()
		}
	}
}


func filesInfo(dir string) []fs.FileInfo{
	entires, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
		return []fs.FileInfo{}
	}
	files := make([]fs.FileInfo, 0, len(entires))
	for _, entry := range entires {
		if info, err := entry.Info(); err == nil {
			files = append(files, info)
		}
	}
	return files
}