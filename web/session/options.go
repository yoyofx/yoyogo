package session

import (
	"github.com/yoyofx/yoyogo/web/session/identity"
	"github.com/yoyofx/yoyogo/web/session/store"
	"github.com/yoyofxteam/dependencyinjection"
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

func (op *Options) AddSessionStoreFactory(storeFactoryCtor interface{}) {
	op.di.AddSingletonByImplements(storeFactoryCtor, new(store.ISessionStore))
}

func (op *Options) AddSessionIdentityFactory(identityFactoryCtor interface{}) {
	//identity.NewCookie("")
	op.di.AddSingletonByImplements(identityFactoryCtor, new(identity.IProvider))
}

func (op *Options) AddSessionMemoryStore(storer store.ISessionStore) {
	//store.NewMemory(3600)
	op.di.AddSingletonByImplements(func() store.ISessionStore { return storer }, new(store.ISessionStore))
}

func (op *Options) AddSessionIdentity(provider identity.IProvider) {
	//identity.NewCookie("")
	op.di.AddSingletonByImplements(func() identity.IProvider { return provider }, new(identity.IProvider))
}
