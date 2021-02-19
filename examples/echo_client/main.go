package main

import (
	"flag"
	"origin/examples/echo_client/client"
	"origin/src/logger"
	"origin/src/utils"
	"sync"
)


var clientCount = flag.Int("n", 1000, "bench client count")
var messageCount = flag.Int64("m", 100, "bench client message count")
var useGoroutine = flag.Bool("go", true, "bench client use goroutine")
var serverPort = flag.Int("port", 9600, "bench server port")
var logLevel = flag.String("log", "info", "bench client count")
var messageSize = flag.Int64("s", 100, "bench client message size")
var host = flag.String("host", "127.0.0.1", "bench server host")
var useSync = flag.Bool("sync", true, "bench client use sync")

func run(isGoroutine bool) {
	logger.Infof("Run goroutine %v", isGoroutine)
	var wg sync.WaitGroup
	for i := 0; i < *clientCount; i++ {
		bc := client.NewBenchClient(int64(i+1), &wg, *messageCount, *messageSize)
		bc.Sync = *useSync

		if isGoroutine {
			wg.Add(1)
			go bc.Start(*host, *serverPort)
		} else {
			bc.WaitGroup = nil
			bc.Start(*host, *serverPort)
		}
	}
	wg.Wait()
}

func main() {
	flag.Parse()
	logger.InitLogger(*logLevel)

	logger.Infof("Start benchmark: client %v message %v size %v sync %v", *clientCount, *messageCount, *messageSize, *useSync)

	delta := utils.ProfileFunc(func() {
		run(*useGoroutine)
	})

	seconds := float64(delta) / 1000000.0
	packetCount := int64(float64(*messageCount) * float64(*clientCount) / seconds)
	logger.Infof("Second packet count %v data size %v KB", packetCount, packetCount*(*messageSize)/1024)
}

