package pkgs

type StrSet struct {
	m map[string]bool
}

func NewStrSet(strings ...string) StrSet {
	m := make(map[string]bool)
	for _, s := range strings {
		m[s] = true
	}
	return StrSet{m}
}

func (s1 StrSet) Has(s string) bool {
	_, ok := s1.m[s]
	return ok
}

func (s1 StrSet) List() []string {
	list := make([]string, len(s1.m))
	i := 0
	for k := range s1.m {
		list[i] = k
		i++
	}
	return list
}

func (s1 StrSet) Difference(s2 StrSet) StrSet {
	var diff []string
	for k := range s1.m {
		_, ok := s2.m[k]
		if !ok {
			diff = append(diff, k)
		}
	}
	return NewStrSet(diff...)
}

func (s1 StrSet) Size() int {
	return len(s1.m)
}
