package simhash

import (
	"fmt"
	"strings"
)

type IndexNode struct {
	Sim   Simhash
	ObjId string
}

func (s *IndexNode) Init(sh Simhash, obj_id string) {
	s.Sim = sh
	s.ObjId = obj_id
}

type SimhashIndex struct {
	bucket  map[string]StrSet
	f, k    int
	offsets []int
}

func (s *SimhashIndex) Init(nodes []IndexNode) {
	s.f = f
	s.k = k
	s.bucket = make(map[string]StrSet)
	s.offsets = make([]int, 0, s.k+1)
	for i := 0; i < s.k+1; i++ {
		s.offsets = append(s.offsets, s.f/(s.k+1)*i)
	}
	for _, node := range nodes {
		s.Add(node)
	}
}

func (s *SimhashIndex) Add(node IndexNode) {
	keys := s.getKeys(node.Sim.Value())
	for _, key := range keys {
		if _, found := s.bucket[key]; !found {
			s.bucket[key] = StrSet{}
		}
		v := fmt.Sprintf("%x,%s", node.Sim.Value(), node.ObjId)
		s.bucket[key].Add(v)
	}
}

func (s *SimhashIndex) Del(node IndexNode) {
	keys := s.getKeys(node.Sim.Value())
	for _, key := range keys {
		v := fmt.Sprintf("%x,%s", node.Sim.Value(), node.ObjId)
		if s.bucket[key].Has(v) {
			s.bucket[key].Del(v)
		}
	}
}

func (s *SimhashIndex) Has(key string) bool {
	_, has := s.bucket[key]
	return has
}

func (s *SimhashIndex) Size() int {
	return len(s.bucket)
}

func (s *SimhashIndex) getKeys(value uint64) []string {
	res := make([]string, 0, len(s.offsets))
	for i, offset := range s.offsets {
		var m uint64
		if i == len(s.offsets)-1 {
			m = 1<<uint(s.f-offset) - 1
		} else {
			m = 1<<uint(s.offsets[i+1]-offset) - 1
		}
		c := value >> uint(offset) & m
		res = append(res, fmt.Sprintf("%x:%x", c, i))
	}
	return res
}

func (s *SimhashIndex) GetNearDups(sim Simhash) []string {
	ans := make(StrSet)
	keys := s.getKeys(sim.Value())
	var sim2, objId string
	c := make([]string, 2)
	for _, key := range keys {
		if dups, found := s.bucket[key]; found {
			if len(dups) > 200 {
				//TODO LOG.WARNING
				info := fmt.Sprintf("Big bucket found. key:%s, len:%d", key, len(dups))
				fmt.Println(info)
			}
			for dup := range dups {
				c = strings.Split(dup, ",")
				if len(c) != 2 {
					panic("inter error")
				}
				sim2 = c[0]
				objId = c[1]
				tpSim := Simhash{}
				tpSim.InitByHex(sim2)
				dis := sim.distance(tpSim)
				if dis <= s.k {
					ans.Add(objId)
				}
			}
		}
	}
	return ans.ToList()
}
