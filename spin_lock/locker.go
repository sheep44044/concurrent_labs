package spinlock

type Spinlock struct {
	value bool
}

func (sl *Spinlock) Lock() {
	for {
		if !sl.value {
			sl.value = true
			return
		}
	}
}

func (sl *Spinlock) Unlock() {
	sl.value = false
}
