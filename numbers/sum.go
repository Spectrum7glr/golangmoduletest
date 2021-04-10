package numbers

import "errors"

func Add2terms(i int, j int) int {
	return i + j
}

func Add(numbers ...int) (error, int) {

	sum := 0
	if len(numbers) < 2 {
		return errors.New("provide at least 2 numbers"), sum
	} else {
		for _, v := range numbers {
			sum += v
		}

		return nil, sum
	}

}
