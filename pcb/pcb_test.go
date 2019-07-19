package pcb

import (
	"reflect"
	"testing"
)

func TestPCB(t *testing.T) {
	p, err := DecodeFile("testdata/t1.kicad_pcb")
	if err != nil {
		t.Fatalf("DecodeFile() failed: %v", err)
	}

	if got, want := p.FormatVersion, 4; got != want {
		t.Errorf("p.FormatVersion = %v, want %v", got, want)
	}

	if got, want := len(p.LayersByName), 20; got != want {
		t.Errorf("len(p.LayersByName) = %v, want %v", got, want)
		t.Logf("p.LayersByName = %+v", p.LayersByName)
	}
	if got, want := p.LayersByName["F.Mask"].Type, "user"; got != want {
		t.Errorf("p.LayersByName[\"F.Mask\"].Type = %v, want %v", got, want)
		t.Logf("p.LayersByName[\"F.Mask\"] = %+v", p.LayersByName["F.Mask"])
	}

	if got, want := len(p.LayersByName), 20; got != want {
		t.Errorf("len(p.LayersByName) = %v, want %v", got, want)
		t.Logf("p.LayersByName = %+v", p.LayersByName)
	}
	if got, want := p.LayersByName["F.Mask"].Type, "user"; got != want {
		t.Errorf("p.LayersByName[\"F.Mask\"].Type = %v, want %v", got, want)
		t.Logf("p.LayersByName[\"F.Mask\"] = %+v", p.LayersByName["F.Mask"])
	}

	if got, want := len(p.Nets), 7; got != want {
		t.Errorf("len(p.Nets) = %v, want %v", got, want)
		t.Logf("p.Nets = %+v", p.Nets)
	}
	if got, want := p.Nets[1].Name, "GND"; got != want {
		t.Errorf("p.Nets[1].Name = %v, want %v", got, want)
		t.Logf("p.Nets[1] = %+v", p.Nets[1].Name)
	}

	if got, want := len(p.Zones), 1; got != want {
		t.Errorf("len(p.Zones) = %v, want %v", got, want)
		t.Logf("p.Zones = %+v", p.Zones)
	}
	if got, want := p.Zones[0].NetName, "GND"; got != want {
		t.Errorf("p.Zones[0].NetName = %v, want %v", got, want)
	}
	if got, want := p.Zones[0].Layer, "B.Cu"; got != want {
		t.Errorf("p.Zones[0].Layer = %v, want %v", got, want)
	}

	if got, want := len(p.Tracks), 44; got != want {
		t.Errorf("len(p.Tracks) = %v, want %v", got, want)
		t.Logf("p.Tracks = %+v", p.Tracks)
	}
	if got, want := p.Tracks[11].NetIndex, 2; got != want {
		t.Errorf("p.Tracks[11].NetIndex = %v, want %v", got, want)
	}

	if got, want := len(p.Vias), 1; got != want {
		t.Errorf("len(p.Vias) = %v, want %v", got, want)
		t.Logf("p.Vias = %+v", p.Vias)
	}
	if got, want := p.Vias[0].NetIndex, 1; got != want {
		t.Errorf("p.Vias[0].NetIndex = %v, want %v", got, want)
	}
	if got, want := p.Vias[0].X, 88.1; got != want {
		t.Errorf("p.Vias[0].X = %v, want %v", got, want)
	}
	if got, want := p.Vias[0].Drill, 0.4; got != want {
		t.Errorf("p.Vias[0].Drill = %v, want %v", got, want)
	}
	if got, want := p.Vias[0].Layers, []string{"F.Cu", "B.Cu"}; !reflect.DeepEqual(got, want) {
		t.Errorf("p.Vias[0].Layers = %v, want %v", got, want)
	}

	if got, want := len(p.NetClasses), 1; got != want {
		t.Errorf("len(p.NetClasses) = %v, want %v", got, want)
		t.Logf("p.NetClasses = %+v", p.NetClasses)
	}
	if got, want := p.NetClasses[0].Name, "Default"; got != want {
		t.Errorf("p.NetClasses[0].Name = %v, want %v", got, want)
	}
	if got, want := p.NetClasses[0].TraceWidth, 0.25; got != want {
		t.Errorf("p.NetClasses[0].TraceWidth = %v, want %v", got, want)
	}
	if got, want := p.NetClasses[0].Nets[0], "/BUS_A"; got != want {
		t.Errorf("p.NetClasses[0].Nets[0] = %v, want %v", got, want)
	}
}
