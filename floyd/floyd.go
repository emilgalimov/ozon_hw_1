package floyd

// Triangle makes a Floyd's triangle matrix with rows count.
func Triangle(rows int) [][]int {
	var result [][]int
	number := 1

	for lineI := 1; lineI <= rows; lineI++ {
		var line []int
		for i := 0; i < lineI; i++ {
			line = append(line, number)
			number++
		}
		result = append(result, line)
	}

	return result
}
