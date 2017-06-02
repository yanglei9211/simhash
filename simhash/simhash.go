package simhash

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"
)

const f = 64
const win_size = 4
//const reg = "[\u4e00-\u9fcca-zA-Z0-9_]+"
const reg = `[\w\p{Han}]+`

func hashfunc(x string) uint64{
	h := md5.New()
	h.Write([]byte(x))
	r := h.Sum(nil)
	var res uint64
	rs := fmt.Sprintf("%x", r[8:])
	fmt.Sscanf(rs, "%x", &res)
	return res
}

type Simhash struct {
	data 		string
	features	[]string
	value 		uint64
	f, win_size int
}

func (s *Simhash) Init(data string) {
	fmt.Println(reg)
	r := regexp.MustCompile(reg)
	fmt.Println(r)
	res := r.FindAllString(data, -1)
	for _, rs := range res {
		s.data += rs
	}
	s.win_size = win_size
	s.f = f
	s.buildByText()
}

func (s *Simhash) Value() uint64 {
	return s.value
}

func (s *Simhash) Tokenize() {
	rs := strings.ToLower(s.data)
	res := make([]string, 0, len(rs)-s.win_size+1)
	for st := 0; st + s.win_size <= len(rs); st++ {
		p := rs[st:st+s.win_size]
		res = append(res, p)
	}
	s.features = res
}

func (s *Simhash) buildByText() {
	s.Tokenize()
	s.buildByFeatures()
}

func (s *Simhash) buildByFeatures() {
	hashs := make([]uint64, 0, len(s.features))
	for _, w := range(s.features){
		hashs = append(hashs, hashfunc(w))
	}
	v := make([]int, s.f)
	masks := make([]uint64, 0, s.f)
	for i := 0; i < s.f; i++ {
		masks = append(masks, 1<<uint(i))
	}
	for _, h := range hashs {
		for i := 0; i < s.f; i++ {
			if (h & masks[i]) > 0 {
				v[i]++
			} else {
				v[i]--
			}
		}
	}
	fmt.Println(v)
	var ans uint64
	for i := 0; i < s.f; i++ {
		if v[i] >= 0 {
			ans |= masks[i]
		}
	}
	s.value = ans
}
