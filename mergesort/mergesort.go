package mergesort

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	return mergeSort(input)
}

func mergeSort(input []int) []int {
	inputLength := len(input)

	if inputLength <= 1 {
		return input
	}

	half := inputLength / 2

	left := input[0:half]
	right := input[half:]

	leftSorted := mergeSort(left)
	rightSorted := mergeSort(right)

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

	return resultArray
}
