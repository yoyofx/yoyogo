package mq

/*
	bs := NewObserver()
	chBroadcast := bs.Run()
	chA := bs.Listener()
	chB := bs.Listener()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range chA {
			fmt.Println("A", v)
		}
		wg.Done()
	}()
	go func() {
		for v := range chB {
			fmt.Println("B", v)
		}
		wg.Done()
	}()
	for i := 0; i < 3; i++ {
		chBroadcast <- i
	}
	bs.RemoveListener(chA)
	for i := 3; i < 6; i++ {
		chBroadcast <- i
	}
	close(chBroadcast)
	wg.Wait()
*/

type Observer struct {
	// 这是消费者监听的通道
	chBroadcast chan int
	// 转发给这些channel
	chListeners []chan int
	// 添加消费者
	chNewRequests chan (chan int)
	// 移除消费者
	chRemoveRequests chan (chan int)
}

// NewObserver 创建一个广播服务
func NewObserver() *Observer {
	return &Observer{
		chBroadcast:      make(chan int),
		chListeners:      make([]chan int, 3),
		chNewRequests:    make(chan (chan int)),
		chRemoveRequests: make(chan (chan int)),
	}
}

// Listener 这会创建一个新消费者并返回一个监听通道
func (bs *Observer) Listener() chan int {
	ch := make(chan int)
	bs.chNewRequests <- ch
	return ch
}

// RemoveListener 移除一个消费者
func (bs *Observer) RemoveListener(ch chan int) {
	bs.chRemoveRequests <- ch
}
func (bs *Observer) addListener(ch chan int) {
	for i, v := range bs.chListeners {
		if v == nil {
			bs.chListeners[i] = ch
			return
		}
	}
	bs.chListeners = append(bs.chListeners, ch)
}

func (bs *Observer) removeListener(ch chan int) {
	for i, v := range bs.chListeners {
		if v == ch {
			bs.chListeners[i] = nil
			// 一定要关闭! 否则监听它的groutine将会一直block
			close(ch)
			return
		}
	}
}

func (bs *Observer) Run() chan int {
	go func() {
		for {
			// 处理新建消费者或者移除消费者
			select {
			case newCh := <-bs.chNewRequests:
				bs.addListener(newCh)
			case removeCh := <-bs.chRemoveRequests:
				bs.removeListener(removeCh)
			case v, ok := <-bs.chBroadcast:
				// 如果广播通道关闭，则关闭掉所有的消费者通道
				if !ok {
					goto terminate
				}
				// 将值转发到所有的消费者channel
				for _, dstCh := range bs.chListeners {
					if dstCh == nil {
						continue
					}
					dstCh <- v
				}
			}
		}
	terminate:
		//关闭所有的消费通道
		for _, dstCh := range bs.chListeners {
			if dstCh == nil {
				continue
			}
			close(dstCh)

		}
	}()
	return bs.chBroadcast
}
