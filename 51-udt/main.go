package main

import "fmt"

type customMap map[string]any

// type StrSlice string
type CustomSlice []string

func (c CustomSlice) Length() int {
	if c == nil {
		return -1 // just to say that it is nil,
		// and since no other error return type is used, assume -1 means somethign went wrong with the slice
	}
	return len(c)
}

func (cm customMap) GetKeys() []string {
	keys := make([]string, len(cm))
	index := 0
	for ks := range cm {
		keys[index] = ks
		index++
	}
	return keys
}

func (cm customMap) GetValues() []any {
	values := make([]any, 0)
	for _, val := range cm {
		values = append(values, val) // when you know the len better create slice with len and iterate thru the index
	}
	return values
}

func main() {

	//var map1 map[string]interface{}
	var cmap1 customMap = make(customMap)
	cmap1["522316"] = "Near Guntur"
	cmap1["560086"] = "Yeshvantpur, Bangaluru"
	cmap1["560096"] = "Rajajinagar, Bangaluru"

	keys := cmap1.GetKeys()
	values := cmap1.GetValues()

	fmt.Println("cmap1 Keys:", keys)
	fmt.Println("cmap1 Values:", values)
	fmt.Println("cmap1 Length:", CustomSlice(keys).Length())

	//var map1 map[string]interface{}
	map1 := make(map[string]any)
	map1["91"] = "India"
	map1["90"] = "Pakisthan"
	map1["66"] = "Singapore"
	map1["97"] = "UAE"

	keys = customMap(map1).GetKeys()
	values = customMap(map1).GetValues()
	fmt.Println("map1 Keys:", keys)
	fmt.Println("map1 Values:", values)

}

// want to create a map that has more methods
