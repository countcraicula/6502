package main

import (
	"cpu/cpu"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	data, err := ioutil.ReadFile("6502_functional_test.bin")
	if err != nil {
		log.Fatal(err)
	}
	c := cpu.CPU{}
	m := cpu.Memory(data)
	c.Reset(m)
	clock := cpu.NewClock(-1, false)

	for {
		clock.Step()
		c.Execute(clock, m)
		fmt.Scanln()
	}
}
