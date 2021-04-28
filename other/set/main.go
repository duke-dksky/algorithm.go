package main

import "fmt"

var Exists = struct{}{}

type set struct {
	m map[interface{}]struct{}
}

func newSet(items ...interface{}) *set {
	s := &set{
		m: make(map[interface{}]struct{}),
	}
	s.insert(items...)
	return s
}

func (s *set) insert(items ...interface{}) {
	for _, item := range items {
		s.m[item] = Exists
	}
}

func (s *set) erase(item interface{}) {
	delete(s.m, item)
}

func (s *set) contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *set) size() int {
	return len(s.m)
}

func (s *set) equal(other *set) bool {
	if s.size() != other.size() {
		return false
	}
	for k := range s.m {
		if !other.contains(k) {
			return false
		}
	}
	return true
}

func (s *set) isSubset(other *set) bool {
	if s.size() > other.size() {
		return false
	}
	for k := range s.m {
		if !other.contains(k) {
			return false
		}
	}
	return true
}

func main() {
	s := newSet()
	s.insert(1)
	s.insert(2)
	s.insert(3)
	s.erase(3)
	fmt.Println(s.contains(3))
	fmt.Println(s.contains(2))
	fmt.Println(s.contains(1))
}
