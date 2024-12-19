package block

import (
	"log"
	"time"
)

type Block struct{}

func (b *Block) Name() string {
	return "block"
}

func (b *Block) Run() {
	log.Println(b.Name(), "run")
	<-time.After(time.Second) // 阻塞一秒
}
