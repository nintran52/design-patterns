package main

import "github.com/nintran52/design-patterns/1-singleton/logger"

func main() {
	logger1 := logger.GetLogger()
	logger2 := logger.GetLogger()

	if logger1 == logger2 {
		println("Both logger instances are the same (singleton works).")
	} else {
		println("Logger instances are different (singleton failed).")
	}
}
