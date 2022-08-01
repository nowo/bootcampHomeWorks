package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	internalsBacking()
	internalSlice()
	observeLenCap()

}

//exercise 17
func internalsBacking() {
	nums := []int{56, 89, 15, 25, 30, 50}

	// ----------------------------------------
	// breaks the connection:
	// mine and nums now have different backing arrays

	// verbose solution:
	// var mine []int
	// mine = append(mine, nums[:3]...)

	// better solution (almost the same thing):
	mine := append([]int(nil), nums[:3]...)
	// ----------------------------------------

	mine[0], mine[1], mine[2] = -50, -100, -150
	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}

//exercise 18
func internalSlice() {

	debug.SetGCPercent(-1)

	report("initial memory usage")

	var array [size]int
	report("after declaring an array")

	array2 := array
	report("after copying the array")

	passArray(array)

	slice1 := array[:]
	slice2 := array[:1e3]
	slice3 := array[1e3:1e4]
	report("after slicings")

	passSlice(slice3)

	fmt.Println()
	fmt.Printf("Array's size : %d bytes.\n", unsafe.Sizeof(array))
	fmt.Printf("Array2's size: %d bytes.\n", unsafe.Sizeof(array2))
	fmt.Printf("Slice1's size: %d bytes.\n", unsafe.Sizeof(slice1))
	fmt.Printf("Slice2's size: %d bytes.\n", unsafe.Sizeof(slice2))
	fmt.Printf("Slice3's size: %d bytes.\n", unsafe.Sizeof(slice3))
}

func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

func passSlice(items []int) {
	report("inside passSlice")
}

func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}

//exercise 19
func observeLenCap() {
	// --- #1 ---
	// var games []string
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// games := []string{}
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// games = append(games, "pacman", "mario", "tetris", "doom")
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	games := []string{"pacman", "mario", "tetris", "doom"}
	fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// --- #2 ---
	fmt.Println()

	for i := 0; i <= len(games); i++ {
		s := games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(s), cap(s))
	}

	// --- #3 ---
	fmt.Println()

	zero := games[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	// --- #4 ---
	fmt.Println()

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}

	// --- #5 ---
	fmt.Println()

	zero = zero[:cap(zero)]
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}

	// --- #6 ---
	fmt.Println()

	zero[0] = "command & conquer"
	games[0] = "red alert"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// uncomment and see the error.
	// _ = games[:cap(games)+1]
	// or:
	// _ = games[:5]
}

//exercise 20
func obserCapGrowth() {
	var (
		nums   []int
		oldCap float64
	)

	// loop 10 million times
	for len(nums) < 1e7 {
		// get the capacity
		c := float64(cap(nums))

		// only print when the capacity changes
		if c == 0 || c != oldCap {
			// print also the growth ratio: c/oldCap
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(nums), c, c/oldCap)
		}

		// keep track of the previous capacity
		oldCap = c

		// append an arbitrary element to the slice
		nums = append(nums, 1)
	}
}