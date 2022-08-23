package main

import (
	"fmt"
	"time"
)

func main() {

	goRoutine1()

	fmt.Println("")

	goRoutine2()

}

func process(name string) {
	time.Sleep(2 * time.Second)
	fmt.Println(name)
}

func goRoutine1() {
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

func goRoutine2() {
	isFin1 := make(chan bool)
	isFin2 := make(chan bool)
	isFin3 := make(chan bool)

	fmt.Println("START2")
	go func() {
		process("A")
		isFin1 <- true
	}()
	go func() {
		process("B")
		isFin2 <- true
	}()
	go func() {
		process("C")
		isFin3 <- true
	}()

	<-isFin1
	<-isFin2
	<-isFin3
	fmt.Println("FINISH2")
}
