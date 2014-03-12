package sleepsort

import "time"

var SleepTime time.Duration = time.Microsecond * 700
func SleepSort(arr []int) {
	sync := make(chan int)
	for _,v := range arr {
		go func(i int) {
			time.Sleep(SleepTime * time.Duration(i))
			sync <- i
		}(v)
	}
	for i,_ := range arr {
		arr[i] = <-sync
	}
}
