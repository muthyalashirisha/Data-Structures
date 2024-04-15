package problems

func generatePascalTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}
func NthRowOfPascalT(n int) []int {
	if n == 1 {
		return []int{1}
	}
	if n == 2 {
		return []int{1, 1}
	}

	nthRow := make([]int, n)
	nthRow[0], nthRow[1] = 1, 1

	for i := 2; i < n; i++ {
		nthRow[i] = 1
		for j := i - 1; j > 0; j-- {
			nthRow[j] += nthRow[j-1]
		}
	}

	return nthRow
}
