package floyd

func Triangle(rowsCount int) [][]int {
	triangle := make([][]int, 0, rowsCount)
	value := 0

	for iRow := 1; iRow <= rowsCount; iRow++ {
		row := make([]int, 0, iRow)
		for col := 1; col <= iRow; col++ {
			value++
			row = append(row, value)
		}
		triangle = append(triangle, row)
	}
	return triangle
}
