// ex9.2 popcount function with sync.Once
package popcount

import "sync"

// pc[i] is the population count of i.
var pc [256]byte
var loadTableOnce sync.Once

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func loadTable() [256]byte {
	loadTableOnce.Do(func() {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
	})
	return pc
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	pc := loadTable()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
