package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("START")

	//boolean型のchannel作成
	ch := make(chan bool)

	//goroutine作成
	go func() {
		time.Sleep(2 * time.Second)
		ch <- true
	}()

	isFin := <-ch
	close(ch)

	fmt.Println(isFin)
	fmt.Println("FINISH")
}
