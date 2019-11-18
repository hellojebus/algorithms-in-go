package main

import "fmt"

func main(){
	arr := []int{2,1,4,1,3}
	hasDupe := checkForDupe(arr)
	fmt.Println(hasDupe)
}

func checkForDupe(arr []int) bool {
	for index, value := range arr {
		if value != index + 1 {
			//fmt.Println(index, "processed")
			if arr[index] == arr[ arr[index] - 1 ]{
				return true
			}
		}
		//fmt.Println(arr[index], arr[ arr[index] - 1 ])
		arr[index], arr[ arr[index] - 1 ] = arr[ arr[index] - 1 ], arr[index]
		fmt.Println(arr)
	}

	return false
}