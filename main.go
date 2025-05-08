package main

import (
	"os"
	"runtime/pprof"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func main() {
	defer logger.Sync()

	f, _ := os.Create("heap.out")
	pprof.WriteHeapProfile(f)
	f.Close()

	logger.Info("Application started")
	processData()
	cleanupResources()
	logger.Info("Application finished")
}

func processData() {
	logger.Info("Starting data processing")
	time.Sleep(2 * time.Second)
	logger.Info("Data processing in progress")
	time.Sleep(2 * time.Second)
	logger.Info("Data processing finished")
	time.Sleep(2 * time.Second)

}

func cleanupResources() {
	logger.Info("Cleaning up resources")
	time.Sleep(2 * time.Second)
	logger.Info("Closing database connections")
	time.Sleep(2 * time.Second)
	logger.Info("Resources cleaned up")
	time.Sleep(2 * time.Second)
}
