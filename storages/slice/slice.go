package slice

import (
	"fmt"
	"reflect"
	"sync"
)

type Slice struct {
	sl    []any
	mutex sync.Mutex
}

func NewSlice() *Slice {
	l := &Slice{}
	return l
}

// Add Добавляет элемент в слайс. Возвращает индекс элемента. Индекс -1 означает неудачу при добавлении элемента.
func (s *Slice) Add(data any) (index int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.sl) != 0 {
		// слайс НЕ пустой
		aType := reflect.TypeOf(s.sl[0])
		bType := reflect.TypeOf(data)
		if aType != bType {
			fmt.Println("\tWrong type:", bType, "\n\tNeed  type:", aType)
			return -1
		}
	}
	v := append(s.sl, data)
	s.sl = v
	return int64(len(s.sl)) - 1
}

func (s *Slice) Delete(index int64) (ok bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if index >= int64(len(s.sl)) {
		return false
	}
	s.sl = append(s.sl[:index], s.sl[index+1:]...)
	return true
}

func (s *Slice) Get(index int64) (val any) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if index > int64(len(s.sl)) {
		return nil
	} else {
		return s.sl[index]
	}
}

func (s *Slice) Print() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	fmt.Println("Length", len(s.sl))
	//fmt.Println(s.sl)
	for i := 0; i < len(s.sl); i++ {
		fmt.Println(s.sl[i])
	}

}

func (s *Slice) SortIncrease(more func(i, j any) bool) {

	s.Sort(true, more)
}

func (s *Slice) SortDecrease(more func(i, j any) bool) {
	s.Sort(false, more)
}

func (s *Slice) String() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return fmt.Sprintf("%+v", s.sl)
}

func (s *Slice) Len() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return int64(len(s.sl))
}

/*
func (s *Slice) String() string {
	var str string

	for i := 0; i < len(s.sl); i++ {
		str += fmt.Sprintf("%v", s.sl[i])
	}

	return str
}*/

func (s *Slice) Sort(increase bool, more func(i, j any) bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	switch increase {
	case true:
		for j := 1; j < len(s.sl); j++ {
			for i := 1; i < len(s.sl); i++ {
				if more(s.sl[i-1], s.sl[i]) { //if s.sl[i-1] > s.sl[i] {
					a := s.sl[i-1]
					s.sl[i-1] = s.sl[i]
					s.sl[i] = a
				}
			}
		}
	case false:
		for j := 1; j < len(s.sl); j++ {
			for i := 1; i < len(s.sl); i++ {
				if !more(s.sl[i-1], s.sl[i]) { //if s.sl[i-1] < s.sl[i] {
					a := s.sl[i-1]
					s.sl[i-1] = s.sl[i]
					s.sl[i] = a
				}
			}
		}
	}

}
