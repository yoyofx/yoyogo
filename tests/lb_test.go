package tests

import (
	"fmt"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/memory"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/strategy"
	"testing"
)

func TestLb(t *testing.T) {
	services := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5", "192.168.1.6"}

	sd := memory.NewServerDiscovery("demo", services)
	selector := strategy.NewRandom(sd, 10)
	for i := 0; i < 6; i++ {
		i1, _ := selector.Next("demo")
		fmt.Println(i1.GetHost())
	}
	fmt.Println("-------------------------------------")
	selector = strategy.NewRound(sd)
	for i := 0; i < 10; i++ {
		i1, _ := selector.Next("demo")
		fmt.Println(i1.GetHost())
	}

}
