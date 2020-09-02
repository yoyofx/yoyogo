package Test

import (
	"fmt"
	"github.com/yoyofx/yoyogo/Internal/ServiceDiscoveryProvider/LB"
	"github.com/yoyofx/yoyogo/Internal/ServiceDiscoveryProvider/Memory"
	"testing"
)

func TestLb(t *testing.T) {
	services := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5", "192.168.1.6"}

	sd := Memory.NewServerDiscovery("demo", services)
	selector := LB.NewRandom(sd, 10)
	for i := 0; i < 6; i++ {
		i1, _ := selector.Next("demo")
		fmt.Println(i1.GetHost())
	}
	fmt.Println("-------------------------------------")
	selector = LB.NewRound(sd)
	for i := 0; i < 10; i++ {
		i1, _ := selector.Next("demo")
		fmt.Println(i1.GetHost())
	}

}
