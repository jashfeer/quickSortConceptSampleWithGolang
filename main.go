package main

import (
	"fmt"
	"math/rand"
)

func quikSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	left, right := 0, len(ar)-1
	pivot := rand.Int() % len(ar)
	ar[pivot], ar[right] = ar[right], ar[pivot]
	for i,_:=range ar{
		if ar[i]<ar[right]{
			ar[left],ar[i]=ar[i],ar[left]
			left++
		}
	}
	ar[left],ar[right]=ar[right],ar[left]

	quikSort(ar[:left])
	quikSort(ar[left+1:])
	return ar
}

func main() {
	slice:=[]int{10,4}
	fmt.Println("unSorted : ",slice)
	slice=quikSort(slice)
	fmt.Println("Sorted : ",slice)
}
