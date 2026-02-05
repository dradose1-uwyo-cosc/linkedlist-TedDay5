package ds

import "errors"

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

// Insert at the tail
func (l *LinkedList) Insert(value string) {
	newNode := &Node{data: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}

	l.Size++
}

// Insert at a specific position
func (l *LinkedList) InsertAt(position int, value string) error {
	if position < 0 || position > l.Size {
		return errors.New("position out of range")
	}

	if position == 0 {
		newNode := &Node{data: value, Next: l.Head}
		l.Head = newNode
		if l.Size == 0 {
			l.Tail = newNode
		}
		l.Size++
		return nil
	}

	curr := l.Head
	for i := 0; i < position-1; i++ {
		curr = curr.Next
	}

	newNode := &Node{data: value, Next: curr.Next}
	curr.Next = newNode

	if newNode.Next == nil {
		l.Tail = newNode
	}

	l.Size++
	return nil
}

// Remove first occurrence of value
func (l *LinkedList) Remove(value string) error {
	if l.Head == nil {
		return errors.New("list empty")
	}

	if l.Head.data == value {
		l.Head = l.Head.Next
		l.Size--
		if l.Head == nil {
			l.Tail = nil
		}
		return nil
	}

	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.data == value {
			prev.Next = curr.Next
			if curr == l.Tail {
				l.Tail = prev
			}
			l.Size--
			return nil
		}
		prev = curr
		curr = curr.Next
	}

	return errors.New("value not found")
}

// Remove all occurrences of value
func (l *LinkedList) RemoveAll(value string) error {
	removed := false

	for l.Head != nil && l.Head.data == value {
		l.Head = l.Head.Next
		l.Size--
		removed = true
	}

	if l.Head == nil {
		l.Tail = nil
		if removed {
			return nil
		}
		return errors.New("value not found")
	}

	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.data == value {
			prev.Next = curr.Next
			l.Size--
			removed = true
			if curr == l.Tail {
				l.Tail = prev
			}
			curr = prev.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	if !removed {
		return errors.New("value not found")
	}
	return nil
}

// Remove node at position
func (l *LinkedList) RemoveAt(pos int) error {
	if pos < 0 || pos >= l.Size {
		return errors.New("position out of range")
	}

	if pos == 0 {
		l.Head = l.Head.Next
		l.Size--
		if l.Size == 0 {
			l.Tail = nil
		}
		return nil
	}

	curr := l.Head
	for i := 0; i < pos-1; i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next
	l.Size--

	if curr.Next == nil {
		l.Tail = curr
	}

	return nil
}

func (l *LinkedList) IsEmpty() bool {
	return l.Size == 0
}

func (l *LinkedList) GetSize() int {
	return l.Size
}

func (l *LinkedList) Reverse() {
	var prev *Node
	curr := l.Head
	l.Tail = l.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	l.Head = prev
}

func (l *LinkedList) PrintList() {
	curr := l.Head
	for curr != nil {
		print(curr.data, " ")
		curr = curr.Next
	}
	println()
}
