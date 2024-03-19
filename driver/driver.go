package main

import (
	"proj3-redesigned/scheduler"
	"fmt"
	"os"
	"strconv"
	"time"
)

const usage = "Usage: ./driver/driver.go <data_size> <mode> <number of threads>\n" +
	"data_size = (big) functions with big sized domains, (medium) functions with medium sized domains, (small) functions with small sized doamins\n" +
	"mode = (s) run sequentially, (p) run in parallel normally, (w) run in parallel with work stealing\n" +
	"number of threads = number of threads to spawn in p or w mode (not required for mode s)"

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}
	config := scheduler.Config{DataSize: os.Args[1], Mode: os.Args[2]}
	if len(os.Args) == 4 { 
		num, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(usage)
			return
		}
		config.ThreadCount = num	
	}

	start := time.Now()
	scheduler.Schedule(config)
	end := time.Since(start).Seconds()
	fmt.Printf("%.2f\n", end)
}
