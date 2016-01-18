package main

type Ship struct {
	name   string
	length int
	hits   int
}

func (s *Ship) Dead() bool {
	if s.hits >= s.length {
		return true
	}
	return false
}

func (s *Ship) Hit() {
	s.hits += 1
}
