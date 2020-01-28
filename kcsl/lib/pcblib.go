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

pcb = struct(
    line = mk_line,
)
`)
