package calc

import "testing"

func TestCoils(t *testing.T) {
	circleCases := []struct {
		RequiredInductance float64
		OuterDia, InnerDia float64
		Turns              float64
	}{
		{
			RequiredInductance: 13225,
			OuterDia:           25,
			InnerDia:           2,
			Turns:              2,
		},
	}

	for i := range circleCases {
		c := circleCases[i]
		nH := RoundCoilInductance(c.Turns, c.OuterDia*1000, c.InnerDia*1000)
		t.Logf("Computed for circle testcase %d: nH = %v", i, nH)
		if nH != c.RequiredInductance {
			t.Errorf("Incorrect inductance for case %d", i)
		}
	}
}
