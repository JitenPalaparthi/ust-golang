package main

func main() {
	// slice1 := []any{10, 10, 20, 30, 40, 50, 60, 70}
	// arr1 := [...]int{10, 10, 20, 30, 40, 50, 60, 70}
	// slice2 := make([]any, 1000)
	// for i := range slice2 {
	// 	slice2[i] = rand.Intn(10000)
	// }

	// println(slice1)
	// println(slice2)
	// for _, v := range arr1 {
	// 	print(v)
	// }
	// println()

	ptr1 := getSquare1(100)
	println("ptr1:", *ptr1)

	ptr2 := new(int)
	getSquare2(100, ptr2)
	println("ptr2:", *ptr2)

	ptr4 := new(int)
	getSquare3(100, ptr4)
	println("ptr4:", *ptr4)

	var ptr3 *int
	ptr3 = new(int) // uncomment to know why there is runtime error
	getSquare3(100, ptr3)
	println("ptr3:", *ptr3)

	var str1 string = "Hello world"
	println(str1)

	str2 := ".Hey Minds"
	str1 = "Hello UST World" + str2 + `"Trying to workout how does this work"var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    ... Use conn to send and receive messages.
}`

	println(str1)

	//json.Unmarshal()

}

func getSquare1(num int) *int {
	out := new(int)
	*out = num * num
	return out
}

func getSquare2(num int, out *int) {
	*out = num * num
}

func getSquare3(num int, out *int) {
	if out == nil {
		println("yes it is nil")
		out = new(int)
	}
	*out = num * num
}
