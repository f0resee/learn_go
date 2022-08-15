package lockfreequeue

import "sync"

type Node struct {
	data int
	next *Node
}
type LockQueue struct {
	head  *Node
	tail  *Node
	mutex sync.Mutex
}

func MakeLockQueue() *LockQueue {
	lkQueue := &LockQueue{
		mutex: sync.Mutex{},
	}
	lkQueue.head = &Node{}
	lkQueue.tail = lkQueue.head
	return lkQueue
}
func (l *LockQueue) Enqueue(d int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	newNode := &Node{
		data: d,
		next: l.tail,
	}
	l.head.next = newNode
	l.head = newNode
}

func (l *LockQueue) Dequeue() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	ret := l.tail.data
	l.tail = l.tail.next
	return ret
}
