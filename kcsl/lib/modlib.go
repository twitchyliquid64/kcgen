package lib

var modLib = []byte(`
# mk_line is a convenience function for generating a line on the
# front silkscreen.
def mk_line(start=XY(), end=XY(), layer=layers.front.silkscreen, width=defaults.width):
    return ModGraphic("fp_line", ModLine(
        layer = layer,
        width = width,
        start = start,
        end = end,
    ))

# mk_text is a convenience function for generating text on
# the front silkscreen.
def mk_text(pos=XYZ(), layer=layers.front.silkscreen, size=XY(1,1), thickness=defaults.thickness, content=""):
    return ModGraphic("fp_text", ModText(
        kind = text.user,
        layer = layer,
        text = content,
        at = pos,
        effects = TextEffects(font_size = size, thickness = thickness),
    ))

def mk_poly(points, pos = XYZ(), layer=layers.front.silkscreen, width = defaults.width):
    return ModGraphic("fp_poly", ModPolygon(
        points = points,
        layer = layer,
        width = width,
    ))

def mk_circle(center=XY(), end=XY(), layer=layers.front.silkscreen, width=defaults.width):
    return ModGraphic("fp_circle", ModCircle(
        center = center,
        end = end,
        layer = layer,
        width = width,
    ))


def mk_smd_pad(ident="", pos=XY(), size=XY(1.4, 1.8), layers=layers.front.smd, shape=shape.rect, round_ratio=0.25):
  return Pad(ident,
        at = pos,
        size = size,
        layers = layers,
        surface = pad.smd,
        shape = shape,
        round_rect_r_ratio = round_ratio)

def mk_th_pad(ident="", pos=XY(), size=XY(1.7, 1.7), drill=XY(1,1), layers=layers.th, shape=shape.oval):
  return Pad(ident,
        at = pos,
        size = size,
        drill_size = drill,
        layers = layers,
        surface = pad.through_hole,
        shape = shape)

graphics = struct(
    line = mk_line,
    text = mk_text,
    poly = mk_poly,
    circle = mk_circle,
)

pads = struct(
  smd = mk_smd_pad,
  th = mk_th_pad,
)

`)
