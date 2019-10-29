package kcgen

// Layer represents the layer on which a graphical element resides.
type Layer int

// Layers.
const (
	LayerFrontFab Layer = iota
	LayerBackFab
	LayerFrontSilkscreen
	LayerBackSilkscreen
	LayerEdgeCuts
	LayerAllCopper
	LayerAllMask
	LayerFrontCopper
	LayerBackCopper
	LayerFrontPaste
	LayerBackPaste
	LayerFrontMask
	LayerBackMask
	LayerFrontCourtyard
	LayerBackCourtyard
)

// Strictname returns the string representing the layer used in the Kicad 4 (and probably later) formats.
func (l Layer) Strictname() string {
	switch l {
	case LayerFrontFab:
		return "F.Fab"
	case LayerBackFab:
		return "B.Fab"
	case LayerFrontSilkscreen:
		return "F.SilkS"
	case LayerBackSilkscreen:
		return "B.SilkS"
	case LayerEdgeCuts:
		return "Edge.Cuts"
	case LayerAllCopper:
		return "*.Cu"
	case LayerAllMask:
		return "*.Mask"
	case LayerFrontCopper:
		return "F.Cu"
	case LayerBackCopper:
		return "B.Cu"
	case LayerFrontPaste:
		return "F.Paste"
	case LayerBackPaste:
		return "B.Paste"
	case LayerFrontMask:
		return "F.Mask"
	case LayerBackMask:
		return "B.Mask"
	case LayerFrontCourtyard:
		return "F.CrtYd"
	case LayerBackCourtyard:
		return "B.CrtYd"
	}
	panic("invalid layer")
}

// String returns a human representation of the layer
func (l Layer) String() string {
	switch l {
	case LayerFrontFab:
		return "Front Fabrication"
	case LayerBackFab:
		return "Back Fabrication"
	case LayerFrontSilkscreen:
		return "Front Silkscreen"
	case LayerBackSilkscreen:
		return "Back Silkscreen"
	case LayerEdgeCuts:
		return "Board Outline"
	case LayerAllCopper:
		return "All Copper"
	case LayerAllMask:
		return "All Mask"
	case LayerFrontCopper:
		return "Front Copper"
	case LayerBackCopper:
		return "Back Copper"
	case LayerFrontPaste:
		return "Front Paste"
	case LayerBackPaste:
		return "Back Paste"
	case LayerFrontMask:
		return "Front Mask"
	case LayerBackMask:
		return "Back Mask"
	case LayerFrontCourtyard:
		return "Front Courtyard"
	case LayerBackCourtyard:
		return "Back Courtyard"
	}
	return "?"
}
