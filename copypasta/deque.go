package copypasta

// 另一种实现是直接 make 个一定大小的 slice，然后用两个下标 s t 模拟

// l-1,...1,0,0,1...,r-1
type deque struct {
	stackL, stackR []int
}

func (q *deque) len() int {
	return len(q.stackL) + len(q.stackR)
}

func (q *deque) pushL(v int) {
	q.stackL = append(q.stackL, v)
}

func (q *deque) pushR(v int) {
	q.stackR = append(q.stackR, v)
}

// panic if empty
func (q *deque) popL() (v int) {
	if len(q.stackL) > 0 {
		q.stackL, v = q.stackL[:len(q.stackL)-1], q.stackL[len(q.stackL)-1]
	} else {
		v, q.stackR = q.stackR[0], q.stackR[1:]
	}
	return
}

// panic if empty
func (q *deque) popR() (v int) {
	if len(q.stackR) > 0 {
		q.stackR, v = q.stackR[:len(q.stackR)-1], q.stackR[len(q.stackR)-1]
	} else {
		v, q.stackL = q.stackL[0], q.stackL[1:]
	}
	return
}
