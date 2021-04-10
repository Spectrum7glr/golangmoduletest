package numbers

func Add2terms(i int, j int) int {
	return i + j
}

func Add(numbers ...int) int {

	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum

}
