package DependencyInjection

import "github.com/maxzhang1985/inject"

type DefaultServiceProvider struct {
	container *inject.Container
}

func (d DefaultServiceProvider) GetService(refObject interface{}) (err error) {
	err = d.container.Extract(refObject)
	if err != nil {
		panic(err)
	}
	return err
}

func (d DefaultServiceProvider) GetServiceByName(refObject interface{}, name string) (err error) {
	err = d.container.Extract(refObject, inject.Name(name))
	if err != nil {
		panic(err)
	}
	return err
}
