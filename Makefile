## web网页采集
#.PHONY: pprof
#pprof:
#	@go tool pprof {路径} && web
#
# 基准测试采集
.PHONY: test
test:
	@go test -run ^$ -bench . ./data_test/ -blockprofile block.out -cpuprofile cpu.out -memprofile mem.out -mutexprofile mutex.out -trace trace.out -outputdir ./testout/
# test不能用makefile，复制命令后直接终端运行


# 最方便直观的一个
# go tool pprof -http :8080 {.out文件路径}