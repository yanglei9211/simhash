package simhash

type StrSet map[string]struct{}

func NewStrSet(strs []string) StrSet {
	set := make(StrSet, len(strs))
	return set
}

func (set StrSet) Add(s string) StrSet{
	set[s] = struct {}{}
	return set
}

func (set StrSet) Del(s string) StrSet{
	delete(set, s)
	return set
}

func (set StrSet) Has(s string) bool{
	_, has := set[s]
	return has
}

func (set StrSet) AddList(ss []string) StrSet{
	for _, s := range ss {
		set.Add(s)
	}
	return set
}

func (set StrSet) ToList() []string{
	ss := make([]string, 0, len(set))
	for k := range set {
		ss = append(ss, k)
	}
	return ss
}