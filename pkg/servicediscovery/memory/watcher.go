package memory

import (
	"errors"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
)

type Watcher struct {
	id   string
	wo   servicediscovery.WatchOptions
	Res  chan *servicediscovery.Result
	exit chan bool
}

func NewWatcher() *Watcher {
	return &Watcher{
		id:   "test",
		Res:  make(chan *servicediscovery.Result),
		exit: make(chan bool),
	}
}

func (m *Watcher) Next() (*servicediscovery.Result, error) {
	for {
		select {
		case r := <-m.Res:
			if len(m.wo.Service) > 0 && m.wo.Service != r.Service.Name {
				continue
			}
			return r, nil
		case <-m.exit:
			return nil, errors.New("watcher stopped")
		}
	}
}

func (m *Watcher) Stop() {
	select {
	case <-m.exit:
		return
	default:
		close(m.exit)
	}
}
