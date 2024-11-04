package main

import (
	"fmt"
)

func main() {
	//slcie1 := []int{1, 2, 3}
	//short hand declaration of a map
	map1 := map[int]string{522001: "Guntur1", 522002: "Guntur-2", 560086: "Yeshvantpur,Bengaluru"}
	for k, v := range map1 {
		println("Key:", k, "Value:", v)
	}
	map1[560096] = "Rajajinagar,Bengaluru"
	for k, v := range map1 {
		println("Key:", k, "Value:", v)
	}

	delete(map1, 522002)
	println("after delete")
	for k, v := range map1 {
		println("Key:", k, "Value:", v)
	}
	// delete an element from a map
	var map2 map[int]string

	delete(map2, 522316)

	err := Delete(map2, 522316)
	if err != nil {
		println(err.Error())
	} else {
		println("deleted successfully from then map")
	}

	// to check whether key already exists or not

	v1 := map1[522316]
	println(v1)

	map1[645911] = ""
	v2 := map1[645911]
	println(v2)

	v3, ok3 := map1[522316]

	if !ok3 {
		println("no value on map1 for the key:", 522316)
	} else {
		fmt.Printf("value:%v for the key 522316\n", v3)
	}

	v3, ok3 = map1[522001]

	if !ok3 {
		println("no value on map1 for the key:", 522001)
	} else {
		fmt.Printf("value:%v for the key: 522001\n", v3)
	}

	// var any1 any
	// v4, ok1 := any1.(int)

}

func Delete(m map[int]string, key int) error {
	if m == nil {
		//return errors.New("nil map")
		return fmt.Errorf("nil map")
	}
	delete(m, key)
	return nil
}
