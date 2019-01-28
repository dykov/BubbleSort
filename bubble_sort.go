package main

import "fmt"

func main()  {

	var array = []int{12,212,10,2,6,1,3,11,31,111,0}
	bubbleSort(array)
	fmt.Println(array)

}

func bubbleSort( array []int ) {

	var sorted = false
	for !sorted {
		sorted = true
		for i := 0 ; i < len(array)-1 ; i++ {
			if array[i] > array[i+1] {
				sorted = false
				array[i] , array[i+1] = array[i+1] , array[i]
			}
		}
		fmt.Println(array)
	}

}
