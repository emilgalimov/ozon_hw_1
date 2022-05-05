package warriors

type Coordinates struct {
	x int
	y int
	v bool
}

func (c Coordinates) isSame(coord Coordinates) bool {
	if abs(c.x-coord.x) <= 1 && abs(c.y-coord.y) <= 1 {
		return true
	}
	return false
}

func Count(image string) int {

	var coords []Coordinates

	var x int
	var y int

	for _, digit := range image {
		switch digit {
		case '1':
			coords = append(coords, Coordinates{x, y, true})
		case '0':
			coords = append(coords, Coordinates{x, y, false})
		case '\n':
			y++
			x = 0
			continue
		}
		x++
	}

	var usedCoords []Coordinates
	var result int

	for _, c := range coords {

		if !c.v {
			continue
		}

		isNew := true
		for _, u := range usedCoords {
			if u.isSame(c) {
				isNew = false
				continue
			}
		}

		if isNew {
			result++
		}
		usedCoords = append(usedCoords, c)
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
