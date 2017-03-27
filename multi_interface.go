package main

import (
    "fmt"
//    "regexp"
    "crypto/md5"
    "strconv"
)

const f = 64

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
    value int
}


func main(){
    s1 := "ffff1231231231238"
    s2 := "ffff1231231231230"
    r := distance(s1, s2)
    fmt.Println(r)
}