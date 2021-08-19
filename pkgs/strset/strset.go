package strset

type strset struct {
	m map[string]bool
}

func New(strings ...string) strset {
	m := make(map[string]bool)
	for _, s := range strings {
		m[s] = true
	}
	return strset{m}
}

func (s1 strset) Has(s string) bool {
	_, ok := s1.m[s]
	return ok
}

func (s1 strset) List() []string {
	list := make([]string, len(s1.m))
	i := 0
	for k := range s1.m {
		list[i] = k
		i++
	}
	return list
}

func (s1 strset) Difference(s2 strset) strset {
	var diff []string
	for k, _ := range s1.m {
		_, ok := s2.m[k]
		if ok {
			diff = append(diff, k)
		}
	}
	return New(diff...)
}

func (s1 strset) Size() int {
	return len(s1.m)
}

