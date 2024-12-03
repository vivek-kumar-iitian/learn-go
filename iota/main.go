package main

import "fmt"

const (
	_ = iota
	kb = 1 << (10 * iota)
	mb = 1 << (10 * iota)
	gb = 1 << (10 * iota)
	tb = 1 << (10 * iota)
	pb = 1 << (10 * iota)
	eb = 1 << (10 * iota)
)
	

func main() {
	fmt.Println(kb,mb,gb,tb,pb,eb) 
}
