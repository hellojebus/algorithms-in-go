package main

import "fmt"

func main(){
	arr := []int{2,5,4,1,3,5,7}
	hasDupe, val:= checkForDupe(arr)
	fmt.Println(hasDupe, val)
}

func checkForDupe(arr []int) (bool, int) {
	for index, value := range arr {
		if value != index + 1 {
			if arr[index] == arr[ arr[index] - 1 ]{
				return true, value
			}
		}
		arr[index], arr[ arr[index] - 1 ] = arr[ arr[index] - 1 ], arr[index]
		fmt.Println(index, arr)
	}

	for index, value := range arr {
		fmt.Println(arr[index], value)
		if(index + 1 != value){
			return true, value
		}
	}

	fmt.Println("final", arr)
	return false, 0
}