package main

import (
	"fmt"
	"testing"
)
var teststring []string = []string{t1,t2,t3,t4,t5,t6,t7,t8,t9}
var (
	t1 = "123.456"
	t2 = "-123.456"
	t3 = "m123.456"
	t4 = "123.456m"
	t5 = "123456"
	t6 = "asdfg"
	t7 = "asdfge"
	t8= "asdfge."
	t9= ".asdfge."
)
func TestAtoi(t *testing.T)  {
	for k, v := range teststring {
		if atoi(v) == 0 {
			fmt.Printf("第%v个  err ,  %s\n", k+1, v)
		}
	}
}