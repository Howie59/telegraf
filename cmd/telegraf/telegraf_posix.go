//go:build !windows
// +build !windows

package main

// 执行入口
func run(inputFilters, outputFilters []string) {
	stop = make(chan struct{})
	reloadLoop(
		inputFilters,
		outputFilters,
	)
}
