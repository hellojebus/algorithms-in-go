package main

import "fmt"

func main(){
	arr := []int{7,1,4,2,5,3,6,8}
	hasDupe := checkForDupe(arr)
	fmt.Println(hasDupe)
}

func checkForDupe(arr []int) bool {
	for index, value := range arr {
		if value != index + 1 {
			if arr[index] == arr[ arr[index] - 1 ]{
				return true
			}
		} else {
			//fmt.Println("Skipping", index)
		}
		fmt.Println("Swapping", arr[index], "with", arr[ arr[index] - 1 ])
		arr[index], arr[ arr[index] - 1 ] = arr[ arr[index] - 1 ], arr[index]
		fmt.Println(index, arr)
	}

	fmt.Println("final", arr)
	return false
}