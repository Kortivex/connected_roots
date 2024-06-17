package config

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/knadh/koanf/v2"

	"github.com/Kortivex/connected_roots/pkg/service"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/thejerf/suture/v4"
)

const (
	LocalFile = "local.yaml"

	ErrMsgWrongFilePath          = "wrong file or path in parsing event"
	ErrMsgWrongConfigEnvVarsLoad = "wrong config env. vars load"
	ErrMsgWrongConfigUnmarshall  = "wrong config unmarshall"
)

type Service struct {
	service.Service
	Conf *Config
}

// NewService This function creates a new Service struct with a provided name, ready to be started and tracked.
func NewService(name string) *Service {
	return &Service{
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
	}
}

// provide reads, loads and fills the final config structure Config.
func (s *Service) provide() error {
	path := os.Getenv("CONFIG_PATH")

	k := koanf.New(".")
	err := k.Load(file.Provider(path+LocalFile), yaml.Parser())
	if err != nil {
		return fmt.Errorf(ErrMsgWrongFilePath+": %w", err)
	}

	err = k.Load(env.Provider("", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, "")), "_", ".")
	}), nil)
	if err != nil {
		return fmt.Errorf(ErrMsgWrongConfigEnvVarsLoad+": %w", err)
	}

	s.Conf = &Config{}
	err = k.Unmarshal("", &s.Conf)
	if err != nil {
		return fmt.Errorf(ErrMsgWrongConfigUnmarshall+": %w", err)
	}

	return nil
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
	if err := s.provide(); err != nil {
		releaseExistence()
		return err
	}
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
