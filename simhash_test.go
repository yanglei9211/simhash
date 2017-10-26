package simhash

import (
	"testing"
	"fmt"
)

func TestSimhash(t *testing.T) {
	//

	ss := []string{"放大是否打算","发生的旅客合法的顺口溜合法的索科洛夫哈萨克了会发生", "11223344", "aabbcc",
		"AABBCC", "AaBbcC", "放大和克赖斯基大师傅但撒风", "是电饭煲adfsgadsg阿嘎色噶多少 12341234",
		"", "", "1", "23", "1122334455667788",
	}
	vs := []uint64{5044479088192766471,
		12744145622408002668,
		8705364617710974090,
		17823150184088196196,
		17823150184088196196,
		17823150184088196196,
		10626750315874733277,
		3800745475638554855,
		16825458760271544958,
		16825458760271544958,
		994258241967195291,
		6739558535857285837,
		17914821209901149642,
	}
	for idx, s := range ss {
		sim := Simhash{}
		sim.Init(s)
		if sim.Value() != vs[idx] {
			fmt.Println(ss[idx], "calc: ", sim.Value(), "ans: ", vs[idx])
			t.Error("error")
		} else {
			fmt.Println(ss[idx], "calc: ", sim.Value(), "ans: ", vs[idx])
			fmt.Println("check ok")
		}
	}
	t.Log("ok")
}

// d4  1d  8c  d9  8f 00 b2 04 e9 80  09 98 ec  f8  42 7e
//[212 29 140 217 143 0 178 4 233 128 9 152 236 248 66 126]
//
	//h := md5.New()
	//h.Write([]byte{})
	//rr := h.Sum(nil)
	//fmt.Println(rr)
	//fmt.Println("------------")
	////rrs := fmt.Sprintf("%x", rr)
	//rrrs := fmt.Sprintf("%x", rr[8:])
	//rrss := fmt.Sprintf("%x", rr[len(rr)-8:])
	//fmt.Println(rrrs)
	//fmt.Println(rrss)
	//var res uint64
	//fmt.Sscanf(rrrs, "%x", &res)
	//fmt.Println(res)