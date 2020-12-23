package session

import (
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/web/session/identity"
	"github.com/yoyofx/yoyogo/web/session/store"
)

type Options struct {
	MaxLifeTime int64
	di          *dependencyinjection.ServiceCollection
}

var option = &Options{MaxLifeTime: 3600}

func UseSession(sc *dependencyinjection.ServiceCollection, opFunc func(options *Options)) {
	option.di = sc
	opFunc(option)
}

func (op *Options) AddSessionMemoryStore(storer store.ISessionStore) {
	//store.NewMemory(3600)
	op.di.AddSingletonByImplements(func() store.ISessionStore { return storer }, new(store.ISessionStore))
}

func (op *Options) AddSessionIdentity(provider identity.IProvider) {
	//identity.NewCookie("")
	op.di.AddSingletonByImplements(func() identity.IProvider { return provider }, new(identity.IProvider))
}
