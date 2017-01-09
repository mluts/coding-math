package tri

import (
	"../vector"
	"testing"
)

func TestPoints(t *testing.T) {
	tr := NewTriangle(vec.NewVec2(40, 50), 20, 10, 0)
	t.Logf("tr.Center.X is %f", tr.Center.X)
	t.Logf("tr.Center.Y is %f", tr.Center.Y)

	points := tr.Points()
	top := points[0]

	wanted := float64(40)
	if top.X != wanted {
		t.Errorf("Wanted top.X to eq %f but have %f", wanted, top.X)
	}
}
