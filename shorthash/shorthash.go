package shorthash

func length(v string) int {
	return len(v)
}

func GenerateShortHashes(dictionary string, len int) []string {

	if 0 == len {
		return []string{}
	}

	if 0 == length(dictionary) {
		return []string{}
	}

	unique := make(map[string]string)

	//первая "итерация"
	for _, v := range dictionary {
		unique[string(v)] = string(v)
	}

	i := 1

	//1+ итерация
	for {
		if i >= len {
			break
		}

		newUnique := make(map[string]string)
		for _, v := range unique {
			for _, v2 := range dictionary {
				newVal := v + string(v2)
				newUnique[newVal] = newVal
			}
		}

		for _, v := range newUnique {
			unique[v] = v
		}

		i++
	}

	var result []string

	for _, v := range unique {
		result = append(result, v)
	}

	return result
}
