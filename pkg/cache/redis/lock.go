package redis

import "time"

type Lock struct {
	ops Ops
}
const (
	LockPrefix = "dlock:"
	LockValue = "DLOCK"
)

/**
获取一个锁，如果锁的key已经存在根据传入的第二个参数的时间进行等待，直到超时或者拿到锁为止。
切记业务完成后一定要及时释放锁。
*/
func (lock *Lock) GetDLock(key string, waitSecond int) (error, bool) {
	successGetLock, err := lock.ops.SetNX(LockPrefix+key, LockValue)
	if err != nil {
		return err, successGetLock
	}
	if !successGetLock {
		beginTime := time.Now()
		for {
			successGetLock, err = lock.ops.SetNX(LockPrefix+key, LockValue)
			if successGetLock {
				break
			}
			if err != nil {
				break
			}
			currentTime := time.Now()
			diffTime := currentTime.Sub(beginTime)
			if diffTime.Seconds() >= float64(waitSecond) {
				break
			}
		}
	}
	return err, successGetLock
}

/**
释放锁
*/
func (lock *Lock) DisposeLock(key string) (error, bool) {
	res, err := lock.ops.DeleteKey(LockPrefix + key)
	return err, res > 0
}