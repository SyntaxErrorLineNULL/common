package cmd

import "sync"

type Commander struct {
	mutex *sync.Mutex
	wg    *sync.WaitGroup
}

func NewCommander() *Commander {
	return &Commander{}
}
