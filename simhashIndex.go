package simhash

import (
	"fmt"
)

type IndexNode struct {
	simhash		Simhash
	obj_id 		string
}

func (s *IndexNode)Init (sh Simhash, obj_id string) {
	s.simhash = sh
	s.obj_id = obj_id
}

type SimhashIndex struct {
	bucket map[string]set.StrSet
	f, k int
	offsets []int
}

func (s *SimhashIndex) Init(nodes []IndexNode){
	s.f = f
	s.k = k
	s.bucket = make(map[string]set.StrSet)
	s.offsets = make([]int, 0, s.k+1)
	for i := 0; i < s.k+1; i++ {
		s.offsets = append(s.offsets, s.f / (s.k+1) * i)
	}
	fmt.Println(s.offsets)
	for _, node := range nodes {
		s.Add(node)
	}
	fmt.Print(s.bucket)
}

func (s *SimhashIndex) Add(node IndexNode) {
	keys := s.getKeys(node.simhash.Value())
	for _, key := range(keys) {
		if _, found := s.bucket[key]; !found {
			s.bucket[key] = set.StrSet{}
		}
		v := fmt.Sprintf("%x,%s", node.simhash.Value(), node.obj_id)
		s.bucket[key].Add(v)
	}
}

func (s *SimhashIndex) Del(node IndexNode) {
	keys := s.getKeys(node.simhash.Value())
	for _, key := range (keys) {
		v := fmt.Sprintf("%x,%s", node.simhash.Value(), node.obj_id)
		if s.bucket[key].Has(v) {
			s.bucket[key].Del(v)
		}
	}
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
