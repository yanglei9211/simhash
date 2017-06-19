package simhash

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

const f = 64
const k = 2
const win_size = 4
//const reg = "[\u4e00-\u9fcca-zA-Z0-9_]+"
const reg = `[\w\p{Han}]+`

func hashfunc(x []rune) uint64{
	h := md5.New()
	h.Write([]byte(string(x)))
	r := h.Sum(nil)
	var res uint64
	rs := fmt.Sprintf("%x", r[8:])
	fmt.Sscanf(rs, "%x", &res)
	return res
}

type UtlString struct {
	s []rune
}

type Simhash struct {
	data 		[]rune
	features	[][]rune
	value 		uint64
	f, win_size int
}

func (s *Simhash) Init(data string) {
	r := regexp.MustCompile(reg)
	res := r.FindAllString(data, -1)
	var strData string
	for _, rs := range res {
		strData += rs
	}
	strData = strings.ToLower(strData)
	s.data = String2Utf8(strData)
	s.win_size = win_size
	s.f = f
	s.buildByText()
}

func (s *Simhash) InitByHex(h string, f int) {
	s.f = f
	fmt.Sscanf(h, "%x", &s.value)
}

func (s *Simhash) Value() uint64 {
	return s.value
}

func (s *Simhash) Tokenize() {
	res := make([][]rune, 0, len(s.data)-s.win_size+1)
	for st := 0; st + s.win_size <= len(s.data); st++ {
		p := s.data[st:st+s.win_size]
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
	var ans uint64
	for i := 0; i < s.f; i++ {
		if v[i] >= 0 {
			ans |= masks[i]
		}
	}
	s.value = ans
}

func (s Simhash) distance(another Simhash) int{
	if s.f != another.f {
		panic("inter error, can't compare")
	}
	return bitsCount(s.value ^ another.value)
}

func String2Utf8(word string) []rune {
	s := []byte(word)
	res := make([]rune, 0, len(s))
	for utf8.RuneCount(s) > 1 {
		r, size := utf8.DecodeRune(s)
		s = s[size:]
		res = append(res, r)
	}
	r, _ := utf8.DecodeRune(s)
	res = append(res, r)
	return res
}

func bitsCount(num uint64) int {
	num = num - ((num >> 1) & 0x5555555555555555)
	num = (num & 0x3333333333333333) + ((num >> 2) & 0x3333333333333333)
	return int((((num + (num >> 4)) & 0xF0F0F0F0F0F0F0F) * 0x101010101010101) >> 56)
}
