package list

import (
	"fmt"
	"reflect"
	"sync"
)

type List struct {
	len       int64
	firstNode *node
	mutex     sync.Mutex
}

func NewList() *List {
	l := &List{
		len:       0,
		firstNode: nil,
	}
	return l
}

type node struct {
	data     any
	nextNode *node
}

func (l *List) Get(index int64) (val any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	node := l.firstNode
	for i := int64(0); i < l.len; i++ {
		if i == index {
			return node.data
		}
		node = node.nextNode
	}

	return nil
}

func (l *List) String() string {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var str_l []string

	node := l.firstNode

	for i := int64(0); i < l.len; i++ {

		str := fmt.Sprintf("%+v", node.data)
		str_l = append(str_l, str)
		node = node.nextNode
	}

	return fmt.Sprintf("%v", str_l)
}

/*
func (l *List) String() string {
	var str string

	node := l.firstNode

	for i := int64(0); i < l.len; i++ {
		str_val := fmt.Sprintf("%v", node.data)
		str += str_val
		node = node.nextNode
	}

	return str
}

*/

// Add Добавляет элемент в список. Возвращает индекс элемента. Индекс -1 означает неудачу при добавлении элемента.
func (l *List) Add(data any) (index int64) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.firstNode == nil {
		n := &node{}
		n.data = data
		l.firstNode = n
		l.len++
		return l.len - 1
	}
	// Первая нода уже существует и тип данных списка "задан"
	aType := reflect.TypeOf(l.firstNode.data)
	bType := reflect.TypeOf(data)
	if aType != bType {
		fmt.Println("\tWrong type:", bType, "\n\tNeed  type:", aType)
		return -1
	}
	nn := l.firstNode
	for {
		if nn.nextNode == nil {
			break
		}
		nn = nn.nextNode
	}
	n := &node{}
	n.data = data
	nn.nextNode = n
	l.len++
	return l.len - 1
}

func (l *List) Delete(index int64) (ok bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.firstNode == nil {
		return false
	}
	if index >= l.len {
		return false
	}
	nn := l.firstNode
	if index == 0 {
		l.firstNode = nn.nextNode
		l.len--
		return true
	}
	for i := int64(0); i < index-1; i++ {
		nn = nn.nextNode
	}
	delNode := nn.nextNode
	nn.nextNode = delNode.nextNode
	delNode.nextNode = nil
	l.len--
	return true
}

func (l *List) Print() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	fmt.Println("Length", l.len)
	if l.firstNode == nil {
		return
	}
	nn := l.firstNode

	for {
		if nn.nextNode == nil {
			fmt.Println(nn.data)
			break
		}
		fmt.Println(nn.data)
		nn = nn.nextNode
	}
}

//func aGrB(a any, b any) bool {
//	reqBodyBytes := new(bytes.Buffer)
//	json.NewEncoder(reqBodyBytes).Encode(a)
//	aBytes := reqBodyBytes.Bytes() // this is the []byte
//
//	reqBodyBytes = new(bytes.Buffer)
//	json.NewEncoder(reqBodyBytes).Encode(b)
//	bBytes := reqBodyBytes.Bytes() // this is the []byte
//	if len(aBytes) > len(bBytes) {
//		return true
//	}
//	if len(aBytes) < len(bBytes) {
//		return false
//	}
//	for i := 0; i < len(bBytes); i++ {
//		if aBytes[i] > bBytes[i] {
//			return true
//		}
//	}
//	return false
//}

// SortIncrease принимает функцию more которая принимает 2 значения i и j. Если i > j, то more возвращает true. Иначе false.
func (l *List) SortIncrease(more func(i, j any) bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.firstNode == nil {
		fmt.Println("Empty list")
		return
	}

	for i := int64(0); i < l.len-1; i++ {
		currentNode := l.firstNode
		for {
			if currentNode.nextNode == nil {
				break
			}
			if more(currentNode.data, currentNode.nextNode.data) {
				currentNode.nextNode.data, currentNode.data = currentNode.data, currentNode.nextNode.data
			}
			currentNode = currentNode.nextNode
		}
	}
}

// SortIncrease принимает функцию more которая принимает 2 значения i и j. Если i > j, то more возвращает true. Иначе false.
func (l *List) SortDecrease(more func(i, j any) bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.firstNode == nil {
		fmt.Println("Empty list")
		return
	}

	for i := int64(0); i < l.len-1; i++ {
		currentNode := l.firstNode
		for {
			if currentNode.nextNode == nil {
				break
			}
			if !more(currentNode.data, currentNode.nextNode.data) {
				currentNode.nextNode.data, currentNode.data = currentNode.data, currentNode.nextNode.data
			}
			currentNode = currentNode.nextNode
		}
	}
}

func (l *List) SortIncreaseNode(more func(i, j any) bool) {
	/*
		switch l.len {

		case 1:
			return
		case 2:
			var (
				a *node
				b *node
			)
			a = l.firstNode
			b = a.nextNode
			if a.data > b.data {
				l.firstNode = b
				a.nextNode = nil
				b.nextNode = a
				return
			}
		default:
			links := make([]*node, 0)
			nn := l.firstNode
			for {
				if nn.nextNode == nil {
					break
				}
				links = append(links, nn)
				nn = nn.nextNode
			}
			links = append(links, nn)

			for j := int64(0); j < l.len; j++ {
				for i := int64(1); i < l.len; i++ {
					//code
					a := links[i-1]
					b := links[i]
					if a.data > b.data {
						if i == 1 {
							l.firstNode = links[i]
							a.nextNode = links[i+1]
							b.nextNode = links[i-1]

						} else {

							a.nextNode = b.nextNode
							b.nextNode = a
							c := links[i-2]
							c.nextNode = b

						}
					}
				}
			}

		}

	*/
}
