package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func main() {

	goRoutine1()

	fmt.Println("")

	goRoutine2()

	result := testing.Benchmark(func(b *testing.B) { run("A", "B", "C", "D", "E") })
	fmt.Println(result.T)

}

func run(name ...string) {
	fmt.Println("START")
	wg := new(sync.WaitGroup)
	siFin := make(chan bool, len(name))

	for _, v := range name {
		//Max5で処理一つに対して、1増やす
		wg.Add(1)
		go process(v, siFin, wg)
	}

	wg.Wait()
	close(siFin)
	fmt.Println("Finish")
}

func process(name string, isFin chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println(name)
	isFin <- true
}

func process1(name string) {
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
		process1("A")
		isFin1 <- true
	}()
	go func() {
		process1("B")
		isFin2 <- true
	}()
	go func() {
		process1("C")
		isFin3 <- true
	}()

	<-isFin1
	<-isFin2
	<-isFin3
	fmt.Println("FINISH2")
}
