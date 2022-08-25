package concurrent_test

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Pipeline_Print_Num_Letter_Interview(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	numChan := make(chan int, 1)
	letterChan := make(chan int, 1)
	startChan := make(chan int, 1)

	go func() {
		<-startChan
		for i := 1; i <= 10; i += 2 {
			fmt.Printf("%d%d", i, i+1)
			letterChan <- 1
			<-numChan
		}

		close(letterChan)
		wg.Done()
	}()

	go func() {
		letterSlice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
		for i := 0; i < len(letterSlice); i += 2 {
			<-letterChan
			fmt.Printf("%v%v", letterSlice[i], letterSlice[i+1])
			numChan <- 1
		}

		close(numChan)
		wg.Done()
	}()

	startChan <- 1
	wg.Wait()
	close(startChan)

	fmt.Printf("\nprint num letter done !**************************************************%d", GetGoroutineID())
}
