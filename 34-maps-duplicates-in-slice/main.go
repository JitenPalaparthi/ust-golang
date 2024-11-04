package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	slice1 := make([]int, 50)
	for i := range slice1 {
		slice1[i] = rand.Intn(30)
	}
	fmt.Println(slice1)

	// map1 := make(map[int]int)

	// for _, v := range slice1 {
	// 	v1, ok1 := map1[v]
	// 	if !ok1 {
	// 		map1[v] = 1
	// 	} else {
	// 		map1[v] = v1 + 1
	// 	}
	// }
	if map1, map2, err := GetDuplicatesAndNon(slice1); err != nil {
		println(err.Error())
	} else {
		println("duplicates---")
		println("Number of duplicates:", len(map1))
		for k, v := range map1 {
			println("Key:", k, "Value:", v)
		}
		println("non-duplicates---")
		println("Number of non-duplicates:", len(map2))
		for k, v := range map2 {
			println("Key:", k, "Value:", v)
		}
	}

}

func GetDuplicatesAndNon(slice []int) (duplicates map[int]int, nonduplicates map[int]int, err error) {
	if slice == nil {
		return nil, nil, errors.New("nil slice")
	}
	map1 := make(map[int]int)
	for _, v := range slice {
		v1, ok1 := map1[v]
		if !ok1 {
			map1[v] = 1
		} else {
			map1[v] = v1 + 1
		}
	}

	map2 := make(map[int]int)
	// delete if the value is 1, that means no duplicates
	for k, v := range map1 {
		if v == 1 {
			map2[k] = v
			delete(map1, k)
		}
	}

	return map1, map2, nil
}
