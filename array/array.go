package array

// Name is name
const Name string = "array"

// Main is a function
func Main() {
	// var arr [3]int
	// arr[0] = 1

	// arr1 := [3]int{1, 2, 3}
	// arr2 := [...]int{1, 2, 3, 4}

	// fmt.Println(len(arr2), cap(arr2))

	// // len 长度
	// // cap 联系存储空间

	// 循环

	// for i := 0; i < len(arr2); i++ {
	// 	fmt.Println(arr2[i])
	// }

	// for _, item := range arr2 {
	// 	fmt.Println(item)
	// }

	// 数据截取
	// [开始: 结束（不包含）] 不支持负数，不填表示全部
	// arrS := arr2[2:3]
	// fmt.Println(arrS)

	// slice 切片
	// var s0 []int

	// s0 = append(s0, 1)

	// s0 = append(s0, s0...)

	// fmt.Println(len(s0), cap(s0))

	// s1 := make([]int, 3, 5) // len cap

	// fmt.Println(len(s1), cap(s1))

	// year := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// s2 := year[3:5]

	// fmt.Println(len(s2), cap(s2))

	// array 可作比较 长度相同
	// slice 不可比较
}
