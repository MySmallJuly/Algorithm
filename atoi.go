package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type back struct {
	res bool
	methed func(s string) float64
}
func main() {
	var count string = "-123456"
	fmt.Printf("转换前 ： %v   ->  %T\n",count,count)
	result := atoi(count)
	fmt.Printf("转换后 ： %v   ->  %T\n",result,result)
}
func atoi(s string) float64 {
	var deback float64
		if s[0] == '+' || s[0] == '-' {
			back := decide(s[1:])
			if back.res {
			deback = -back.methed(s[1:])
		}
			return deback
	}
			back := decide(s)

			if back.res {
				deback = back.methed(s)
			}
			return deback

}

func finddpoint(s string) float64{
	if count_point(s)>1 {
		log.Printf("字符串格式错误: %s",s)
		return 0
	}
	count := strings.Split(s,"")
	conf := count[:point(count)]
	var ab float64
	for k, v := range conf {
		i := exchange(v)
		k := math.Pow(10,float64(len(conf)-k-1))
		s := i * k
		ab += s
	}
	conl := count[point(count)+1:]

	var al float64
	for k1, v := range conl {
		i := exchange(v)
		k := math.Pow(10,-float64(k1+1))
		s := i * k
		al += s
	}
	all := ab + al
	return all
}

func finde(s string) float64 {
	if s[0] == 'e' || s[0] == 'E' || count_e(s) > 1 || s[len(s)-1] == '.'{
		log.Printf("字符串格式错误: %s",s)
		return 0
	}
	count := strings.Split(s,"")
	var ab float64
	var al float64
	conf := count[:eee(count)]
	for k, v := range conf {
		i := exchange(v)
		k := math.Pow(10,float64(len(conf)-k-1))
		s := i * k
		ab += s
	}


	conl := count[eee(count)+1:]
	for k, v := range conl {
		i := exchange(v)
		k := math.Pow(10,float64(len(conf)-k-1))
		s := i * k
		al += s
	}
	all := ab * math.Pow(math.E,al)
	return all
}

func findint(s string)  float64{
	var all float64
	count := strings.Split(s,"")
	for k, v := range count {
		i := exchange(v)
		k := math.Pow(10,float64(len(s)-k-1))
		s := i * k
		all += s
	}
	return all
}

func point(s []string) int{
	var index int
	for m,n := range s {
		if n == "." {
			index = m
		}
	}
	return index
}


func eee(s []string) int{
	var index int
	for m,n := range s {
		if n == "e" || n == "E"{
			index = m
		}
	}
	return index
}

func exchange(s string) float64 {
	r := []rune(s)
	re := r[0] - '0'
	return float64(re)
}
func decide(s string) back {
	for _, v := range []rune(s) {
		if (v > 57 && v < 69) || (v > 69 && v < 101) || (v > 101) || (v > 46 && v < 48) || (v < 46){
			log.Printf("字符串格式错误: %s" ,s)
			return back{res:false,methed:nil}
		}
	}
	for _, v := range []rune(s) {
	if v == 46 {
		return back{res: true, methed: finddpoint}
	}
	for _, v := range []rune(s) {
		if v == 69 || v == 101 {
				//fmt.Println("e")
				return back{res:true,methed:finde}
			}
		}
	}
	return back{res:true,methed:findint}
}
func count_e(s string) int {
	var count int
	for i := 0;i<len(s) ;i++{
		if s[i] == 'e' || s[i] == 'E' {
			count++
		}
	}
	return count
}

func count_point(s string) int {
	var count int
	for i := 0;i<len(s) ;i++{
		if s[i] == '.' {
			count++
		}
	}
	return count
}