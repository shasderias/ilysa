package refresh

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/shasderias/ilysa/internal/filenotify"
)

type Watcher struct {
	filenotify.FileWatcher
	*Manager
	context context.Context
}

func NewWatcher(m *Manager) (*Watcher, error) {
	//var watcher filenotify.FileWatcher

	watcher := filenotify.NewPollingWatcher()

	return &Watcher{
		FileWatcher: watcher,
		Manager:     m,
		context:     m.context,
	}, nil
}

func (w *Watcher) Start() {
	go func() {
		for {
			err := filepath.Walk(w.ProjectDir, func(path string, info os.FileInfo, err error) error {
				if info == nil {
					w.cancelFunc()
					return errors.New("nil directory")
				}
				if info.IsDir() {
					if strings.HasPrefix(filepath.Base(path), "_") {
						return filepath.SkipDir
					}
					if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
						return filepath.SkipDir
					}
				}
				if w.isWatchedFile(path) {
					w.Add(path)
				}
				return nil
			})

			if err != nil {
				w.Logger.Error(err)
				w.context.Done()
				break
			}
			// sweep for new files every 1 second
			time.Sleep(1 * time.Second)
		}
	}()
}

var watchedExtensions = []string{".go"}

func (w Watcher) isWatchedFile(path string) bool {
	ext := filepath.Ext(path)

	for _, e := range watchedExtensions {
		if strings.TrimSpace(e) == ext {
			return true
		}
	}

	return false
}
