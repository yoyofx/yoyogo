package Abstractions

import "sync"

type ApplicationEvent struct {
	Data  interface{}
	Topic string
}

// ApplicationChannel 是一个能接收 ApplicationEvent 的 channel
type ApplicationChannel chan ApplicationEvent

// DataChannelSlice 是一个包含 DataChannels 数据的切片
type DataChannelSlice []ApplicationChannel

// ApplicationEventPublisher 存储有关订阅者感兴趣的特定主题的信息
type ApplicationEventPublisher struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func NewEventPublisher() *ApplicationEventPublisher {
	var eb = &ApplicationEventPublisher{
		subscribers: map[string]DataChannelSlice{},
	}
	return eb
}

func (eb *ApplicationEventPublisher) NewEvent() chan ApplicationEvent {
	return make(chan ApplicationEvent)
}

func (eb *ApplicationEventPublisher) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		// 这样做是因为切片引用相同的数组，即使它们是按值传递的
		// 因此我们正在使用我们的元素创建一个新切片，从而正确地保持锁定
		channels := append(DataChannelSlice{}, chans...)
		go func(data ApplicationEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(ApplicationEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

func (eb *ApplicationEventPublisher) Subscribe(topic string, ch ApplicationChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]ApplicationChannel{}, ch)
	}
	eb.rm.Unlock()
}
