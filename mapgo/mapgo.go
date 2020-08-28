package mapgo

// Main is a function
func Main() {
	// m := map[int]int{1: 1, 2: 2}
	// fmt.Println(m)
	// fmt.Printf("%d", len(m))

	m2 := map[int]int{}
	m2[2] = 16
	// fmt.Println(m2)
	// fmt.Printf("%d", len(m2))

	// m2 := make(map[string]int, 10 /*Initial Capacity*/) = map[string]int{}
	// m2["str"] = 2
	// fmt.Println(m2)

	// for k, v := range m2 {
	// 	fmt.Println(k, v)
	// 	delete(m2, k)
	// }

	// m2[2] = 0

	// // key不存在会返回0，不能用判断nil的方式判断是否存在

	// if v, ok := m2[2]; ok { // value boolean 值是否存在
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("not existing")
	// }

	// m := map[int]func(op int) int{}
	// m[1] = func(op int) int {
	// 	return op
	// }

	// fmt.Println(m[1](2))
	// m := map[int]bool{}

	// m[1] = true

	// n := 1
	// if m[n] {
	// 	fmt.Printf("%d is existing", n)
	// } else {
	// 	fmt.Printf("%d is not existing", n)
	// }
}
