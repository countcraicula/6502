package main

import (
	"cpu/cpu"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
)

func main() {

	data, err := ioutil.ReadFile("6502_functional_test.bin")
	if err != nil {
		log.Fatal(err)
	}
	c := cpu.CPU{
		IRQ: make(chan bool),
	}
	m := cpu.Memory(data)
	c.Reset(m)
	clock := cpu.NewClock(-1, false)
	f, err := os.Create("profile.dat")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	c.Execute(clock, m)
	pprof.StopCPUProfile()
	fmt.Println(c.String())
}
