package search2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if i+1 > len(matrix) && matrix[i+1][0] >= target {
			i++
			continue
		}
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
