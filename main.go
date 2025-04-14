package main

import "fmt"

var arr1 = [3]string{"I", "Love", "You"}

func main() {
	var arr [2]int

	arr[1] = 6

	fmt.Println(arr)
	fmt.Println(arr1)

}
