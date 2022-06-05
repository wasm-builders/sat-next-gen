package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/suborbital/sat/sat"
)

func startSat(satFunction *sat.Sat, satConfig *sat.Config) {
	fmt.Println("🌍:", satConfig.Port)
	satFunction.Start()
}

func main() {
	
	wasmHelloModuleConfig, _ := sat.ConfigFromRunnableArg("hello/hello.wasm")
	
	wasmHeyModuleConfig, _ := sat.ConfigFromRunnableArg("hey/hey.wasm")
	
	helloFunction, _ := sat.New(wasmHelloModuleConfig, nil)
	heyFunction, _ := sat.New(wasmHeyModuleConfig, nil)

	

	helloResult, err := helloFunction.Exec([]byte("Bob"))
	if err != nil {
		log.Fatal(err)
	}
	heyResult, err := heyFunction.Exec([]byte("Bob"))
	if err != nil {
		log.Fatal(err)
	}

	go startSat(helloFunction, wasmHelloModuleConfig)
	go startSat(heyFunction, wasmHeyModuleConfig)	

	fmt.Println(string(helloResult.Output))
	fmt.Println(string(heyResult.Output))

	// make sure that the go program won't exit
	<-make(chan bool)


}