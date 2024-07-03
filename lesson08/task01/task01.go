package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	files := []string{"./file1.txt", "./file2.txt", "./file3.txt"}

	sum := 0
	for _, file := range files {
		wg.Add(1)
		go readFile(file, ch, &wg)
		sum += <-ch
	}

	wg.Wait()
	fmt.Printf("合計文字数:%d\n", sum)
}

func readFile(path string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fp, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Can't open file,err:%v", err)
	}

	count := len(fp)

	ch <- count

}
