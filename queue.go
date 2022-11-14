package exqueue

import (
	"fmt"
	"strings"
	"sync"
)

type item struct {
	next  *item
	value interface{}
}

type Queue struct {
	headFirst *item
	headLast  *item
	tailFirst *item
	tailLast  *item
	mutex     *sync.Mutex
}

func New() *Queue {
	return &Queue{
		mutex: &sync.Mutex{},
	}
}

func (q *Queue) Pop() interface{} {
	if q.headFirst == nil {
		return nil
	}
	defer func() {
		if q.headFirst == q.tailFirst {
			q.tailFirst = q.headFirst.next
		}
		if q.headFirst == q.headLast {
			q.headLast = q.headFirst.next
		}
		q.headFirst = q.headFirst.next
	}()
	return q.headFirst.value
}

func (q *Queue) Push(val interface{}, prior bool) {
	newItem := &item{
		value: val,
	}
	if q.headFirst == nil {
		q.headFirst = newItem
		q.headLast = newItem
		q.tailFirst = newItem
		q.tailLast = newItem
		return
	}
	noHead := q.headFirst == q.tailFirst
	if prior {
		if noHead {
			newItem.next = q.tailFirst
			q.headFirst = newItem
		} else {
			newItem.next = q.headLast.next
			q.headLast.next = newItem
		}

		q.headLast = newItem
	} else {
		q.tailLast.next = newItem
		q.tailLast = newItem
	}

}

func (q *Queue) ToString() string {
	sb := strings.Builder{}
	cur := q.headFirst
	for cur != nil {
		if cur == q.headFirst {
			sb.WriteString("H>|")
		} else {
			sb.WriteString("  |")
		}

		if cur == q.headLast {
			sb.WriteString("<H")
		} else {
			sb.WriteString("  ")
		}

		sb.WriteString(":")

		if cur == q.tailFirst {
			sb.WriteString("T>|")
		} else {
			sb.WriteString("  |")
		}

		if cur == q.tailLast {
			sb.WriteString("<T")
		} else {
			sb.WriteString("  ")
		}

		sb.WriteString(fmt.Sprintf(" %v\n", cur.value))
		cur = cur.next
	}
	return sb.String()
}
