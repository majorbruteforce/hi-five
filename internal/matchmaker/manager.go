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
	queue   []*list.List
	config  Config
	Ingress chan Candidate
	Egress  chan [][2]string
}

func NewManager(config Config) *MatchManager {
	return &MatchManager{
		queue:  make([]*list.List, 0),
		config: config,
	}
}

func (m *MatchManager) NewMatchMaker() error {
	// Accepts incoming Candidates through Ingress, creates matches
	// and facilitates session creation.
	//
	// When target for the strategy is met, it sends the
	// the list to MatchCandidatesBatch and pipes the result
	// to broadcast.CreateBatchSessions
	return nil
}
