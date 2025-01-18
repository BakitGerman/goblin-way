package landplot

func findSmallestSquareArea(width int, height int) (int, int) {
	if width%height == 0 {
		if width > height {
			return width / 2, height
		}
		return width, height / 2
	}
	if width > height {
		return findSmallestSquareArea(width%height, height)
	} else {
		return findSmallestSquareArea(width, height%width)
	}
}
