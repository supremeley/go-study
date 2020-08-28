package panic

// Main is a function
func Main() {
	// fmt.Println("start")

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("recovered from", err)
	// 	}
	// 	// fmt.Println("Finally")
	// }()

	// panic(errors.New("wrong"))

}

// init is a function
func init() {
	// fmt.Println("panic init1")
}

func init() {
	// fmt.Println("panic init2")
}
