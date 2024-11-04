package main

func main() {

	var map1 map[string]string // declaration of map

	if map1 == nil {
		map1 = make(map[string]string)
	}
	map1["522001"] = "Guntur-1"
	map1["522002"] = "Guntur-2"
	map1["560086"] = "Yeshvantpur,Bengaluru"
	map1["560096"] = "RajaniNagar,Bengaluru"

	map2 := make(map[bool]string)
	map2[true] = "yes"
	map2[false] = "no"

	// slice1 := []int{1, 2, 3}
	// slice2 := []int{4, 5, 6}
	// if slice1 == slice2 {

	// }

	// map3 := make(map[[]int]int) // this does not work bcz == operator can't be used on slices

	for key, value := range map1 {
		println("key:", key, "value:", value)
	}

}

// maps
// keys are mapped to values
// maps are very good performers becase O(1)
// buckets:

// 92279febea1e62fb5c67 0514d2fa8fd704d76286
// bucket number        element in the bucket

// what data type can be the key ?
// any type that can apply == operator can be a key in the map
// the order of elements is not guaranteed
// the maps are not thread safe
