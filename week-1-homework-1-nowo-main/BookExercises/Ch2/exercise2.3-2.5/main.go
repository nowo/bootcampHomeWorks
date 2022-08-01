package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//2.3
func PopCountLoop(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

//2.4
func PopCountLoop64(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		count += int(pc[byte(x>>(i*64))])
	}
	return count
}

//2.5
func PopCountClearRightmost(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(PopCountLoop(10))
	fmt.Println(PopCountLoop64(10))
	fmt.Println(PopCountClearRightmost(10))
}
