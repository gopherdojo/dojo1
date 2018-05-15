package split


func CreateArray(length ,limit int) []int {
	limits := make([]int, limit)
	downloadSize := length / limit

	for i := 0; i < limit - 1 ; i++ {
		limits[i] = downloadSize * (i + 1)
	}
	limits[limit - 1] = length
	return limits
}