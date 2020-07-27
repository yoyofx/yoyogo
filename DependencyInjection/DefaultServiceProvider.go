package DependencyInjection

import (
	"github.com/defval/inject/v2"
	"github.com/defval/inject/v2/di"
)

type DefaultServiceProvider struct {
	container *inject.Container
}

func (d DefaultServiceProvider) GetService(refObject interface{}) (err error) {
	err = d.container.Extract(refObject)
	return err
}

func (d DefaultServiceProvider) GetServiceByName(refObject interface{}, name string) (err error) {
	err = d.container.Extract(refObject, inject.Name(name))

	return err
}

func (d DefaultServiceProvider) GetGraph() string {
	var graph *di.Graph
	if err := d.container.Extract(&graph); err != nil {
		// handle err
	}

	return graph.String() // use string representation
}
