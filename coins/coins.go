package coins

func Piles(n int) int {
	return piles(n, n)
}

var m = make(map[int]map[int]int)

func piles(n, k int) int {

	if n == 0 && k == 0 {
		return 1
	}

	if k == 0 {
		return 0
	}

	if _, nok := m[n]; !nok {
		m[n] = make(map[int]int)
	}

	if v := m[n][k]; v != 0 {
		return v
	}

	if k > n {
		m[n][k] = piles(n, n)
		return m[n][k]
	}

	m[n][k] = piles(n, k-1) + piles(n-k, k)
	return m[n][k]
}
