package main

import (
    "fmt"
    "regexp"
    "crypto/md5"
    "strconv"
)

const f = 64
const win_size = 4
const reg = `[\w\u4e00-\u9fcc]+`

type md5_value struct {
    h,l uint16
}

func hashfun(x string) string{
    h := md5.New()
    h.Write([]byte(x))
    r := h.Sum(nil)
    return fmt.Sprintf("%x", r)
}

func distance_int64(u1, u2 uint64) int {
    cnt := 0
    x := (u1^u2)&((1<<f)-1)
    for x > 0 {
        cnt ++;
        x &= (x-1)
    }
    return cnt
}

func distance(s1, s2 string) int {
    u1l, _ := strconv.ParseUint(s1[:16], 16, 0)
    u2l, _ := strconv.ParseUint(s2[:16], 16, 0)
    u1r, _ := strconv.ParseUint(s1[16:], 16, 0)
    u2r, _ := strconv.ParseUint(s2[16:], 16, 0)
    res := distance_int64(u1l, u2l) + distance_int64(u1r, u2r)
    return res
}


type Simhash struct {
    data, value string
    feature
    win_size int
}

//目前只支持传入string,后续追加以特征,simhash值来创建
func (s *Simhash) init(data string) {
    r := regexp.MustCompile("[\u4e00-\u9fcca-zA-Z0-9_]+")
    res := r.FindAllString(data, -1)
    for _, rs := range res{
        s.data += rs
    }
    s.win_size = win_size
}


func (s *Simhash)slide() []string {
    for st:=0; st + s.win_size < len(s.data); st++ {
        p := s.data[st:st+s.win_size]
        fmt.Println(p)
    }
    return []string{}
}


func main(){
    s1 := "阿#!道1夫@A生a活空@d间D__adsf"
    //s2 := "ffff1231231231230"
    ts := Simhash{}
    ts.init(s1)
    fmt.Println(ts.data)
    ts.slide()
}