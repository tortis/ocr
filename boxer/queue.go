package boxer

import "fmt"

type point struct {
	x, y int
}

type node struct {
	p    point
	next *node
}

type queue struct {
	front  *node
	back   *node
	length int
}

func newQueue() *queue {
	return &queue{}
}

func (this *queue) q(p point) {
	if this.length == 0 {
		this.front = &node{
			p: p,
		}
		this.back = this.front
	} else {
		this.back.next = &node{
			p: p,
		}
		this.back = this.back.next
	}
	this.length += 1
}

func (this *queue) dq() *point {
	if this.length == 0 {
		return nil
	}
	this.length -= 1
	r := &this.front.p
	this.front = this.front.next
	if this.length == 0 {
		this.back = nil
	}
	return r
}

func (this *queue) peek() *point {
	if this.front == nil {
		return nil
	}
	return &this.front.p
}

func (this *queue) String() string {
	s := ""
	c := this.front
	for c != nil {
		s += fmt.Sprintf("%v -> ", c.p)
		c = c.next
	}
	s += "nil\n"
	return s
}
