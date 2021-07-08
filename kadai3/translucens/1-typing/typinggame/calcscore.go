package typinggame

// CalcScore caliculates player score based on player input
func CalcScore(correct, playerinput string) int {
	editlen := EditLength(correct, playerinput)
	correctlen := len(correct)

	if editlen > correctlen {
		return 0
	}

	return correctlen - editlen
}

// EditLength caliculates edit distance between two strings
func EditLength(str1, str2 string) int {
	len1 := len(str1)
	len2 := len(str2)

	str1 = " " + str1
	str2 = " " + str2

	table := make([][]int, len1+1)
	for i1 := range table {
		table[i1] = make([]int, len2+1)
		table[i1][0] = i1
	}
	for i2 := range table[0] {
		table[0][i2] = i2
	}

	for i1 := range table {
		if i1 == 0 {
			continue
		}

		for i2 := range table[i1] {
			if i2 == 0 {
				continue
			}

			cost := 0
			if str1[i1] != str2[i2] {
				cost = 1
			}
			table[i1][i2] = min(min(table[i1-1][i2]+1, table[i1][i2-1]+1), table[i1-1][i2-1]+cost)
		}
	}

	return table[len1][len2]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
