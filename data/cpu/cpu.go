package cpu

import "log"

type Cpu struct{}

func (c *Cpu) Name() string {
	return "cpu"
}

func (c *Cpu) Run() {
	log.Println(c.Name(), "run")
	for i := 0; i < 10000000000; i++ {

	}
}
