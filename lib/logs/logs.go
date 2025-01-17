package assembler

import (
	"log"
	"runtime"
	"time"
)

func LogMemoryUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	log.Printf("Memory Usage: Alloc=%dKB, TotalAlloc=%dKB, Sys=%dKB, NumGC=%d",
		memStats.Alloc/1024, memStats.TotalAlloc/1024, memStats.Sys/1024, memStats.NumGC)
}

func LogExecutionTime(start time.Time, taskName string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", taskName, elapsed)
}
