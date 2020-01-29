package lib

var pcbLib = []byte(`
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
    line = mk_line,
    arc  = mk_arc,
    text = mk_text,
)
`)
