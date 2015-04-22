package main

func Bubble(input []int) {
	swap := true
	for swap {
		swap = false
		for i := range input[1:] {
			if input[i] > input[i+1] {
				Swap(input, i+1, i)
				swap = true
			}
		}
	}
}

func Insertion(input []int) {
	for i := range input[1:] {
		for j := i + 1; j > 0 && input[j-1] > input[j]; j-- {
			Swap(input, j, j-1)
		}
	}
}

func Selection(input []int) {
	min := 0
	for i := range input {
		min = i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}
		if min != i {
			Swap(input, i, min)
		}
	}
}

func Shell(input []int) {

}

func Merge(input []int) {

}

func Quick(input []int) {

}
