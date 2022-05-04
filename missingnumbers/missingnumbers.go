package missingnumbers

func Missing(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{1, 2}
	}

	miss := make([]int, 0, 2)
	i := 1
	for {
		if !contains(numbers, i) {
			miss = append(miss, i)
			if len(miss) == 2 {
				break
			}
		}
		i++
	}
	return miss
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
