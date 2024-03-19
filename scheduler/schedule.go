package scheduler

type Config struct {
	DataSize  string        // Represents the size of the data
	Mode    string        // Represents whether the server should execute
	// sequentially or in parallel
	// If Mode == "s"  then run the sequential version
	// If Mode == "p"  then run the regular parallel version
	// If Mode == "w"  then run the parallel work-stealing queue version
	// These are the only values for Version
	ThreadCount int // Represents the number of threads to spawn
}

func Schedule(config Config) {
	if config.Mode == "s" {
		RunSequential(config)
	} else if config.Mode == "p" {
		RunParallel(config)
	} else if config.Mode == "w" {
		RunWorkStealing(config)
	} else {
		panic("Invalid mode given. Please use s, p, or w.\n")
	}
}
