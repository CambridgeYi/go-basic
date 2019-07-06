package main

import (
	"fmt"
	"time"
)

func main(){
	go func(){
		fmt.Println("test concurrent")
	}()
	time.Sleep(time.Microsecond)
	// chanbase0 {} dead lock
}
