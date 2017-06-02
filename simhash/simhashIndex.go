package simhash

import "fmt"

type IndexNode struct {
	simhash		Simhash
	obj_id 		string
}

type SimhashIndex struct {
	bucket map[string]string
	f, k int
	offsets []int
}

func (s *SimhashIndex) Init(nodes []IndexNode){
	s.f = f
	s.k = k
	s.offsets = make([]int, 0, s.k+1)
	for i := 0; i < s.k+1; i++ {
		s.offsets = append(s.offsets, s.f / (s.k+1) * i)
	}
	fmt.Println(s.offsets)
	for _, r := range nodes {
		s.Add(string(r.simhash.value), r.obj_id)
	}
}

func (s *SimhashIndex) Add(key, value string) {
	if !s.Has(key) {
		s.bucket[key] = value
	}
}

func (s *SimhashIndex) Del(key string) {
	delete(s.bucket, key)
}

func (s *SimhashIndex) Has(key string)bool {
	_, has := s.bucket[key]
	return has
}

func (s *SimhashIndex) Size() int {
	return len(s.bucket)
}

func (s *SimhashIndex) getKeys(value uint64) []string {
	res := make([]string, 0, len(s.offsets))
	for i, offset := range(s.offsets) {
		var m uint64
		if (i == len(s.offsets)-1) {
			m = 1 << uint(s.f - offset) -1
		} else {
			m = 1 << uint(s.offsets[i+1] - offset) - 1
		}
		c := value >> uint(offset) & m
		res = append(res, fmt.Sprintf("%x:%x", c, i))
	}
	return res
}
