package tests

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/memory"
	"testing"
	"time"
)

func TestWatcher(t *testing.T) {
	w := memory.NewWatcher()

	go func() {
		time.Sleep(3 * time.Second)
		w.Res <- &servicediscovery.Result{}
	}()

	_, err := w.Next()
	if err != nil {
		t.Fatal("unexpected err", err)
	}

	w.Stop()

	if _, err := w.Next(); err == nil {
		t.Fatal("expected error on Next()")
	}
}
