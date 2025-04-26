import "sync/atomic"

type spinlock struct {
    value int32 
}

func (sl *spinlock) Lock() {
    for {
        if atomic.CompareAndSwapInt32(&sl.value, 0, 1) { 
            return
        }
    }
}

func (sl *spinlock) Unlock() {
    atomic.StoreInt32(&sl.value, 0) 
}
