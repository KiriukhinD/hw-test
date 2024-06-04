package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type List2 struct {
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() *List2 {
	return new(List2)
}

///////////////////////////

func (l *List2) Len() int {
	return l.len
}

func (l *List2) Front() *ListItem {
	return l.front
}

func (l *List2) Back() *ListItem {
	return l.back
}

func (l *List2) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: l.front}
	if l.front != nil {
		l.front.Prev = newItem
	}
	l.front = newItem
	if l.back == nil {
		l.back = newItem
	}
	l.len++
	return newItem
}

func (l *List2) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: l.back}
	if l.back != nil {
		l.back.Next = newItem
	}
	l.back = newItem
	if l.front == nil {
		l.front = newItem
	}
	l.len++
	return newItem
}

func (l *List2) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	l.len--
}

func (l *List2) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}

	l.Remove(i)
	i.Prev = nil
	i.Next = l.front
	l.front.Prev = i
	l.front = i
}
