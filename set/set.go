package set

import (
	"fmt"
)

type Set struct {
	Elements map[interface{}]struct{}
}

func NewSet() *Set {
	set := Set{
		Elements: make(map[interface{}]struct{}),
	}
	return &set
}

//Add adds an element to struct
func (s *Set) Add(elem interface{}) {
	s.Elements[elem] = struct{}{}
}

func (s *Set) Delete(elem interface{}) error {
	_, exist := s.Elements[elem]
	if !exist {
		return fmt.Errorf("element %v does not exist in the list", elem)
	}
	delete(s.Elements, elem)
	return nil
}

func (s *Set) Contains(elem interface{}) bool {
	_, exist := s.Elements[elem]
	return exist
}

func (s *Set) List() []interface{} {
	list := make([]interface{}, 0)
	for k, _ := range s.Elements {
		list = append(list, k)
	}
	return list
}

func (s *Set) Print() {

	for k := range s.Elements {
		fmt.Println(k)
	}

}
