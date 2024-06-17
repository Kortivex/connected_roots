package service

import "sync"

const (
	Run = iota
	Fail
	Panic
	Hang
	Heartbeat
	UseStopChan
	TerminateTree
	DoNotRestart

	ErrFailureServiceEnding = "failure while service end"
	ErrPanicService         = "service panic"

	PingOK = "PING(OK)"
	PingKO = "PING(KO)"
)

type Service struct {
	Name     string
	Started  chan bool
	Status   chan int
	Release  chan bool
	Stop     chan bool
	Existing int

	M       sync.Mutex
	Running bool
}
