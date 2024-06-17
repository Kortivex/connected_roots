package logger

import (
	"context"
	"fmt"
	"github.com/thejerf/suture/v4"
	"sync"

	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"

	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/service"
)

type Service struct {
	service.Service
	Logger *logger.Logger
	conf   *config.Config
}

// NewService This function creates a Service using an existing config.Config and returns it to the caller.
func NewService(name string, conf *config.Config) *Service {
	srv := &Service{
		Service: service.Service{
			Name:     name,
			Started:  make(chan bool),
			Status:   make(chan int),
			Release:  make(chan bool),
			Stop:     make(chan bool),
			Existing: 0,
			M:        sync.Mutex{},
			Running:  false,
		},
		conf: conf,
	}

	return srv
}

// provide return the logger.Logger instance with the config.Config that is loaded.
func (s *Service) provide() *logger.Logger {
	log := logger.NewLogger(s.conf.App.LogLevel)
	log.WithTag(commons.TagServiceConnectedRoots)
	log.Debug(fmt.Sprintf("%s loaded", s.Name))
	return &log
}

// Serve method from interface suture.Service to handle Service cycle life.
func (s *Service) Serve(ctx context.Context) error {
	s.M.Lock()
	if s.Existing != 0 {
		(&sync.Mutex{}).Unlock()
	}
	s.Existing++
	s.Running = true
	s.M.Unlock()

	defer func() {
		s.M.Lock()
		s.Running = false
		s.M.Unlock()
	}()

	releaseExistence := func() {
		s.M.Lock()
		s.Existing--
		s.M.Unlock()
	}

	// Start Service
	s.Logger = s.provide()
	s.Started <- true

	useStopChan := false

	for {
		select {
		case val := <-s.Status:
			switch val {
			case service.Run:
			case service.Heartbeat:
			case service.Fail:
				releaseExistence()
				if useStopChan {
					s.Stop <- true
				}
				return nil
			case service.Panic:
				releaseExistence()
				panic(service.ErrPanicService)
			case service.Hang:
				<-s.Release
			case service.UseStopChan:
				useStopChan = true
			case service.TerminateTree:
				return suture.ErrTerminateSupervisorTree
			case service.DoNotRestart:
				return suture.ErrDoNotRestart
			}
		case <-ctx.Done():
			releaseExistence()
			if useStopChan {
				s.Stop <- true
			}
			return fmt.Errorf(service.ErrFailureServiceEnding+": %w", ctx.Err())
		}
	}
}
