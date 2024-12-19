package main

import (
	"log"
	"os"
	"pprof/data"
	"pprof/data/block"
	"pprof/data/cpu"
	"pprof/data/goroutine"
	"pprof/data/mem"
	"pprof/data/mutex"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

// 硬编码采集
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

	cpufile, err := os.OpenFile("code_collection/out/cpu.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(cpufile) // 开始采集
	defer pprof.StopCPUProfile()         // 停止采集
	defer cpufile.Close()

	memfile, err := os.OpenFile("code_collection/out/mem.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.WriteHeapProfile(memfile) // 开始采集，这个不用close
	defer memfile.Close()

	tracefile, err := os.OpenFile("code_collection/out/trace.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = trace.Start(tracefile)
	defer trace.Stop()
	defer tracefile.Close()

	for i := 0; i < 10; i++ {
		for _, v := range cmd {
			v.Run()
		}
		time.Sleep(time.Second)
	}
}
