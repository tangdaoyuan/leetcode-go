package _1812

func squareIsWhite(coordinates string) bool {
	var row = coordinates[0] - 'a' + 1
	var col = coordinates[1]

	return (row+col-1)%2 == 0
}
