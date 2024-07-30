package matchmaker

import (
	"container/list"
	"time"
)

type Candidate struct {
	ID       string
	Keywords []string
}

const (
	// Initiate batch matching after config.targetLength
	StrategyBatchSize int = 0
	// Initiate batch matching after config.targetTime
	StrategyBatchTime int = 1
	// Initiate a continous match and dispatch strategy
	StrategyRapid int = 2
)

type Config struct {
	targetLength int
	targetTime   time.Duration
	strategy     int
}

type MatchManager struct {
	queue    []*list.List
	config   Config
	Pipeline chan Candidate
}
