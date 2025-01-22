package landplot

import "testing"

// вообще эту задачу по идее можно решить через нахождение наибольшего общего делителя
func TestLandPlot(t *testing.T) {
	gotWidth, gotHeight := findSmallestSquareArea(1680, 640)
	wantWidth, wantHeight := 80, 80
	if wantWidth != gotWidth || wantHeight != gotHeight {
		t.Errorf("findSmallestSquareArea(1680, 640), wantWidth %d, wantHeight %d, gotWidth %d, gotHeight %d",
			wantWidth, wantHeight, gotWidth, gotHeight)
	}
}
