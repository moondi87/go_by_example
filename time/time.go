package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format(time.ANSIC))
	fmt.Println(now.Format(time.RFC3339))

	then := time.Date(2013, 11, 16, 20, 30, 50, 66666, time.UTC)
	fmt.Println(then)
	fmt.Printf("%d-%d-%d %d:%d:%d %v\n", then.Year(), int(then.Month()), then.Day(), then.Hour(), then.Minute(), then.Second(), then.Location())
	fmt.Println(then.Local())
	diff := now.Sub(then)
	fmt.Println(diff)

}
