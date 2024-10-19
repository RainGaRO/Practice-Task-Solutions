package set

type Set struct {
	Collection []interface{}
}

func (s *Set) Size() int {
	return len(s.Collection)
}

func (s *Set) Contains(value interface{}) (int, bool) {
	for k, v := range s.Collection {
		if v == value {
			return k, true
		}
	}
	return 0, false
}

func (s *Set) Add(value interface{}) bool {
	if _, exist := s.Contains(value); exist == false {
		s.Collection = append(s.Collection, value)
		return true
	}
	return false
}

func (s *Set) Remove(value interface{}) bool {
	if i, exist := s.Contains(value); exist == true {
		length := s.Size()

		s.Collection[i] = s.Collection[length-1]
		s.Collection = s.Collection[:length-1]

		return true
	}

	return false
}

func (s *Set) Union(set *Set) *Set {
	union := &Set{}

	for _, v := range s.Collection {
		union.Add(v)
	}

	for _, v := range set.Collection {
		union.Add(v)
	}

	return union
}

func (s *Set) Difference(set *Set) *Set {
	difference := &Set{}

	for _, v := range s.Collection {
		if _, exist := set.Contains(v); exist == false {
			difference.Add(v)
		}
	}

	return difference
}

func (s *Set) Intersection(set *Set) *Set {
	intersection := &Set{}

	for _, v := range s.Collection {
		if _, exist := set.Contains(v); exist == true {
			intersection.Add(v)
		}
	}
	return intersection
}
