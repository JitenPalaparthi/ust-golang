package main

func main() {

	slice1 := []int{}
	if slice1 == nil {
		println("nil slice")
	}
	/*
		slice1
		-----
		ptr: 0x202010af
		len: 0
		cap: 0
	*/
	slice1 = append(slice1, 10, 20, 30, 40) // can pass any number of arguments since it is variadic parameter

	s1 := sumOfNums() // zero values are there
	println("s1:", s1)
	s2 := sumOfNums(10, 20, 30, 40, 50, 60) // zero values are there
	println("s2:", s2)
	s3 := sumOfNums(10, 20, 30, 40, 50, 60, 12, 3, 2, 3, 23, 2, 3, 23, 23) // zero values are there
	println("s3:", s3)

	s4 := sumOfNums(slice1...)
	println("s4:", s4)

	arr1 := [5]int{10, 20, 30, 40, 50}

	s5 := sumOfNums(arr1[:]...)
	println("s5:", s5)

	//fmt.Println()

}

// ... variadic parameter
// the variadic parameter must be the last paramter

func sumOfNums(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func sumOfNMultiplyOfNums(nums ...int, n int) int { // cant write as variadic must be the last param
func sumOfNMultiplyOfNums(n int, nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += (v * n)
	}
	return sum
}
