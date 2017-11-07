package simhash

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestSimhash(t *testing.T) {
	//

	ss := []string{"放大是否打算", "发生的旅客合法的顺口溜合法的索科洛夫哈萨克了会发生", "11223344", "aabbcc",
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

func TestSimhashIndex(t *testing.T) {
	fmt.Println("test 2")
	ss := []string{"放大是否打算", "发生的旅客合法的顺口溜合法的索科洛夫哈萨克了会发生", "11223344", "aabbcc",
		"AABBCC", "AaBbcC", "放大和克赖斯基大师傅但撒风", "是电饭煲adfsgadsg阿嘎色噶多少 12341234",
		"", "", "1", "23", "1122334455667788",
	}
	res := []Simhash{}
	for _, s := range ss {
		sim := Simhash{}
		sim.Init(s)
		res = append(res, sim)
	}

	test := []IndexNode{}
	s := SimhashIndex{}
	for idx := range res {
		tpNode := IndexNode{res[idx], fmt.Sprintf("%d", idx)}
		test = append(test, tpNode)
	}
	s.Init(test)

	toTestStr := []string{"发生的旅客合法的顺口溜合法的索科洛夫哈萨克了会发生", "1", "23", "是电饭煲adfsgadsg阿嘎色噶多少 12341234"}
	anss := []string{"1", "10", "11", "7"}
	for i := range toTestStr {
		ts := Simhash{}
		ts.Init(toTestStr[i])
		ans := s.GetNearDups(ts)
		fmt.Println(ans[0], anss[i])
		if len(ans) == 1 && ans[0] == anss[i] {
			fmt.Println("check ok")
		} else {
			t.Error("error")
		}

	}
}

func BenchmarkSimhash_Init(b *testing.B) {
	ss := []string{
		"过正方体上底面的对角线和下底面一顶点的平面截去一个三棱锥所得到的几何体如图所示它的俯视图为leftqquadright",
	}
	for i := 0; i < b.N; i++ {
		for _, rs := range ss {
			s := Simhash{}
			s.Init(rs)
		}
	}
}

func BenchmarkMd5(b *testing.B) {
	ss := []string{
		"过正方体上底面的对角线和下底面一顶点的平面截去一个三棱锥所得到的几何体如图所示它的俯视图为leftqquadright",
	}
	for i := 0; i < b.N; i++ {
		for _, rs := range ss {
			h := md5.New()
			h.Write([]byte(string(rs)))
			h.Sum(nil)
		}
	}
}

func TestMd(t *testing.T) {
	ss := "过正方体上底面的对角线和下底面一顶点的平面截去一个三棱锥所得到的几何体如图所示它的俯视图为leftqquadright"
	h := md5.New()
	h.Write([]byte(string(ss)))
	r := h.Sum(nil)
	fmt.Println(r)
	s := Simhash{}
	s.Init(ss)
	fmt.Println(s.value)
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
