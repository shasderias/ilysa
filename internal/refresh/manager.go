package refresh

import (
	"context"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/shasderias/ilysa"
)

type Manager struct {
	*Configuration
	ID         string
	Logger     *Logger
	Restart    chan bool
	cancelFunc context.CancelFunc
	context    context.Context
	gil        *sync.Once
}

func New(c *Configuration) *Manager {
	return NewWithContext(c, context.Background())
}

func NewWithContext(c *Configuration, ctx context.Context) *Manager {
	ctx, cancelFunc := context.WithCancel(ctx)
	m := &Manager{
		Configuration: c,
		Logger:        NewLogger(c),
		cancelFunc:    cancelFunc,
		context:       ctx,
		gil:           &sync.Once{},
	}
	return m
}

func (m *Manager) Start() error {
	w, err := NewWatcher(m)
	if err != nil {
		return err
	}
	w.Start()
	go m.build(fsnotify.Event{Name: ":start:"})
	go func() {
		for {
			select {
			case event := <-w.Events():
				if event.Op != fsnotify.Chmod {
					go m.build(event)
				}
				w.Remove(event.Name)
				w.Add(event.Name)
			case <-m.context.Done():
				m.Logger.Error("done")
				break
			}
		}
	}()

	if m.Debug {
		go func() {
			for {
				select {
				case err := <-w.Errors():
					m.Logger.Error(err)
				case <-m.context.Done():
					break
				}
			}
		}()
	}
	select {}
}

func (m *Manager) build(event fsnotify.Event) {
	m.gil.Do(func() {
		defer func() {
			m.gil = &sync.Once{}
		}()
		err := func() error {
			now := time.Now()
			m.Logger.Print("Rebuild on: %s", event.Name)

			if err := ilysa.Invoke(m.ProjectDir); err != nil {
				return err
			}

			tt := time.Since(now)
			m.Logger.Success("Building Completed (Time: %s)", tt)
			return nil
		}()
		if err != nil {
			m.Logger.Error(err)
		}
	})
}
