package mem

import (
	"log"
	"pprof/constans"
)

type Mem struct {
	buffer [][constans.Mi]byte // 每个切片1M
}

func (m *Mem) Name() string {
	return "mem"
}

func (m *Mem) Run() {
	log.Println(m.Name(), "run")
	for len(m.buffer)*constans.Mi < constans.Gi {
		m.buffer = append(m.buffer, [constans.Mi]byte{})
	}
}
