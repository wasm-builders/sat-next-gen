package main

import (
	"fmt"
	"time"
	//"log"
	"sync"

	//"os"

	"github.com/suborbital/sat/sat"
)

var wg sync.WaitGroup // instanciation de notre structure WaitGroup


func startSat(satFunction *sat.Sat, satConfig *sat.Config, stopCh <-chan struct{}) {
	defer wg.Done()
	fmt.Println("🚀 started goroutine 🌍:", satConfig.Port)

	select {

		case <-stopCh:
			fmt.Println("⏹️ stopped goroutine 🌍:", satConfig.Port)

		default:
			satFunction.Start()
	}



	
}

func main() {
	
	wasmHelloModuleConfig, _ := sat.ConfigFromRunnableArg("hello/hello.wasm")
	
	wasmHeyModuleConfig, _ := sat.ConfigFromRunnableArg("hey/hey.wasm")
	
	helloFunction, _ := sat.New(wasmHelloModuleConfig, nil)
	heyFunction, _ := sat.New(wasmHeyModuleConfig, nil)

	
	/*
	helloResult, err := helloFunction.Exec([]byte("Bob"))
	if err != nil {
		log.Fatal(err)
	}
	heyResult, err := heyFunction.Exec([]byte("Bob"))
	if err != nil {
		log.Fatal(err)
	}
	*/

	wg.Add(1)
	stopHelloCh := make(chan struct{})
	go startSat(helloFunction, wasmHelloModuleConfig, stopHelloCh)

	wg.Add(1)
	stopHeyCh := make(chan struct{})
	go startSat(heyFunction, wasmHeyModuleConfig, stopHeyCh)	




	/*
	fmt.Println(string(helloResult.Output))
	fmt.Println(string(heyResult.Output))
	*/
	time.Sleep(time.Second * 5)
	fmt.Println("👋")
  close(stopHeyCh)
	fmt.Println("😉")
	wg.Wait()

	fmt.Println("🎉")
	// make sure that the go program won't exit
	//<-make(chan bool)


}