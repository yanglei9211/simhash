package simhash

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"
	"go/token"
)

const f = 64
const win_size = 4
const reg = "[\u4e00-\u9fcca-zA-Z0-9_]+"

func hashfunc(x string) string{
	h := md5.New()
	h.Write([]byte(x))
	r := h.Sum(nil)
	return fmt.Sprintf("%x", r)
}

type Simhash struct {
	data 		string
	features	[]string
	value 		uint64
	f, win_size int
}

func (s *Simhash) Init(data string) {
	r := regexp.MustCompile(reg)
	res := r.FindAllString(data, -1)
	for _, rs := range res {
		s.data += rs
	}
	s.win_size = win_size
	s.f = f
}


func (s *Simhash) Tokenize() {
	rs := strings.ToLower(s.data)
	res := make([]string, 0, len(rs)-s.win_size+1)
	for st := 0; st + s.win_size <= len(rs); st++ {
		p := rs[st:st+s.win_size]
		res = append(res, p)
	}
	s.dataToken = res
}

func (s *Simhash) buildByText() {

}

func (s *Simhash) buildByFeatures() {
	hashs := make([]string, 0, len(s.features))
	for w := range(s.features){
		hashs = append(hashs, hashfunc(w))
	}
	v := make([]int, s.f)

}
