package fib

import "time"

func Fib(u uint) uint {
	if u <= 1 {
    time.Sleep(1 * time.Second)
		return 1
	}
	return Fib(u-2) + Fib(u-1)
}
