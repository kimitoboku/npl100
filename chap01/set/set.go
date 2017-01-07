package set

type Set struct {
	m map[string]struct{}
}

func NetSet() *Set {
	set := &Set{}
	set.m = make(map[string]struct{})
	return set
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[string]struct{})
}

func (s *Set) Add(ss ...string) {
	if 0 < len(ss) {
		for _, i := range ss {
			s.m[i] = struct{}{}
		}
	}
}

func (s *Set) Remove(ss string) {
	delete(s.m, ss)
}

func (s *Set) RemoveAll(ss ...string) {
	if 0 < len(ss) {
		for _, i := range ss {
			s.Remove(i)
		}
	}
}

func (s *Set) Contains(ss string) bool {
	if _, ok := s.m[ss]; ok {
		return true
	} else {
		return false
	}
}

func (s *Set) Items() []string {
	list := make([]string, s.Size(), s.Size())
	i := 0
	for k, _ := range s.m {
		list[i] = k
		i++
	}
	return list
}

func (s *Set) Union(s2 *Set) *Set {
	set := NetSet()
	set.Add(s.Items()...)
	set.Add(s2.Items()...)
	return set
}

func (s *Set) Intersection(s2 *Set) *Set {
	set := NetSet()
	for k, _ := range s2.m {
		if s.Contains(k) {
			set.Add(k)
		}
	}
	return set
}
