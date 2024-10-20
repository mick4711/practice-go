package floyd

type (
	row      []int
	triangle []row
)

func Triangle(rowsCount int) triangle {
	result := make(triangle, 0, rowsCount)
	value := 0

	for iRow := 1; iRow <= rowsCount; iRow++ {
		line := make(row, 0, iRow)
		for col := 1; col <= iRow; col++ {
			value++
			line = append(line, value)
		}
		result = append(result, line)
	}
	return result
}
