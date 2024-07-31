package matchmaker

import (
	"fmt"
	"time"
)

type Candidate struct {
	ID       string
	Keywords []string
}

const (
	// Initiate batch matching after config.targetLength
	StrategySize int = 0
	// Initiate batch matching after config.targetTime
	StrategyTime int = 1
	// Initiate a continous match and dispatch strategy
	StrategyRapid int = 2
)

type Config struct {
	TargetBufferSize int
	TargetWaitTime   time.Duration
	Strategy         int
}

type MatchManager struct {
	buffer  []Candidate
	config  Config
	Ingress chan Candidate
	Egress  chan [][2]string
	stop    chan struct{}
}

func NewManager(config Config) *MatchManager {
	m := &MatchManager{
		buffer:  make([]Candidate, 0),
		config:  config,
		Ingress: make(chan Candidate),
		Egress:  make(chan [][2]string),
		stop:    make(chan struct{}),
	}
	m.NewMatchMaker()
	return m
}

func (m *MatchManager) NewMatchMaker() {

	switch m.config.Strategy {
	case StrategySize:
		go m.runStrategySize()
	case StrategyTime:
		go m.runStrategyTime()
	case StrategyRapid:
		go m.runStrategyRapid()
	default:
		fmt.Println("Unknown strategy")
	}
	fmt.Println("New matchmaker running")
}

func (m *MatchManager) runStrategySize() {
	for {
		select {
		case arrival, ok := <-m.Ingress:
			if !ok {
				fmt.Println("matchmaker ingress channel is closed and drained")
				return
			}
			m.buffer = append(m.buffer, arrival)
			if len(m.buffer) >= m.config.TargetBufferSize {
				matches := MatchCandidatesBatch(m.buffer)
				m.Egress <- matches
				m.buffer = m.buffer[:0]
			}
		case <-m.stop:
			return
		}
	}
}

func (m *MatchManager) runStrategyTime() {
	ticker := time.NewTicker(m.config.TargetWaitTime)
	defer ticker.Stop()

	for {
		select {
		case arrival, ok := <-m.Ingress:
			if !ok {
				fmt.Println("matchmaker ingress channel is closed and drained")
				return
			}
			m.buffer = append(m.buffer, arrival)
		case <-ticker.C:
			if len(m.buffer) > 0 {
				matches := MatchCandidatesBatch(m.buffer)
				m.Egress <- matches
				m.buffer = m.buffer[:0]
			}
		case <-m.stop:
			return
		}
	}
}

func (m *MatchManager) runStrategyRapid() {
	// Implement
}
