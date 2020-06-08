package lib

var pcbLib = []byte(`
# mk_track returns a PCB track with the specified parameters.
def mk_track(start = XY(), end = XY(), width = 0.25, layer=layers.front.copper, net=0):
    return Track(
        start = start,
        end = end,
        width = width,
        layer = layer,
        net_index = net,
    )

# mk_via returns a PCB via with the specified parameters.
def mk_via(at=XY(), size = 0.8, drill = 0.4, layers=[layers.front.copper,layers.back.copper], net=0, type=ViaThrough):
    return Via(
        at = at,
        size = size,
        drill = drill,
        layers = layers,
        net_index = net,
        type=type
    )

# mk_line returns a PCB line with the specified parameters.
def mk_line(start=XY(), end=XY(), layer=layers.edge, width=defaults.width):
    return Line(
        layer = layer,
        width = width,
        start = start,
        end = end,
    )

# mk_arc returns a PCB arc with the specified parameters.
def mk_arc(center=XY(), end=XY(), angle=90.0, width=defaults.width, layer=layers.front.silkscreen):
    return Arc(
        width = width,
        layer = layer,
        start = center,
        end = end,
        angle = angle,
    )

# mk_text returns a PCB text element with the specified parameters.
def mk_text(pos=XYZ(), layer=layers.front.silkscreen, size=XY(1,1), thickness=defaults.thickness, hidden=False, content=""):
    return Text(
        layer = layer,
        text = content,
        at = pos,
        hidden = hidden,
        effects = TextEffects(font_size = size, thickness = thickness),
    )

pcb = struct(
    line  = mk_line,
    arc   = mk_arc,
    text  = mk_text,
    via   = mk_via,
    track = mk_track,
)
`)
