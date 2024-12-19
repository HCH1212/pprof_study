package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"pprof/data"
	"pprof/data/block"
	"pprof/data/cpu"
	"pprof/data/goroutine"
	"pprof/data/mem"
	"pprof/data/mutex"
	"runtime"
	"time"
)

// web网页采集
var cmd = []data.Cmd{
	&cpu.Cpu{},
	&mem.Mem{},
	&block.Block{},
	&goroutine.Goroutine{},
	&mutex.Mutex{},
}

func main() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(os.Stdout)

	// 开启对锁调用的跟踪
	runtime.SetMutexProfileFraction(1)
	// 开启对阻塞操作的跟踪
	runtime.SetBlockProfileRate(1)

	go func() {
		_ = http.ListenAndServe(":6060", nil)
	}()

	for {
		for _, v := range cmd {
			v.Run()
		}
		time.Sleep(time.Second)
	}
}
