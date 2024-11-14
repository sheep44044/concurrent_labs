package semaphore

type Semaphore struct {
	//you need to write here
}

func NewSemaphore(v int) *Semaphore {
	return &Semaphore{
		//you need to write here
	}

}

// sem_wait(consumer) : -1
func (sem *Semaphore) P() {
	panic("you need to write from here , and delete this line")
}

// sem_post(producter) : +1
func (sem *Semaphore) V() {
	panic("you need to write from here , and delete this line")
}
