package mergesort

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	resultChan := make(chan []int)
	defer close(resultChan)
	go mergeSort(input, resultChan)
	return <-resultChan
}

func mergeSort(input []int, output chan<- []int) {
	inputLength := len(input)

	if inputLength <= 1 {
		output <- input
		return
	}

	half := inputLength / 2

	left := input[0:half]
	right := input[half:]

	leftChan := make(chan []int)
	defer close(leftChan)
	rightChan := make(chan []int)
	defer close(rightChan)

	go mergeSort(left, leftChan)
	go mergeSort(right, rightChan)

	leftSorted := <-leftChan
	rightSorted := <-rightChan

	var resultArray []int

	li := 0
	ri := 0

	var (
		isLeftOver  bool
		isRightOver bool
	)

	for {

		if isLeftOver && isRightOver {
			break
		}

		if (leftSorted[li] < rightSorted[ri] || isRightOver) && !isLeftOver {
			resultArray = append(resultArray, leftSorted[li])

			if li+1 == len(leftSorted) {
				isLeftOver = true
				continue
			}

			li++
			continue
		}

		resultArray = append(resultArray, rightSorted[ri])

		if ri+1 == len(rightSorted) {
			isRightOver = true
			continue
		}

		ri++
	}

	output <- resultArray
}
