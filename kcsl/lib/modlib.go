package lib

var modLib = []byte(`
# mk_line returns a graphical line with the specified parameters.
def mk_line(start=XY(), end=XY(), layer=layers.front.silkscreen, width=defaults.width):
    return ModGraphic("fp_line", ModLine(
        layer = layer,
        width = width,
        start = start,
        end = end,
    ))

# mk_arc returns a graphical arc with the specified parameters.
def mk_arc(center=XY(), end=XY(), angle=90.0, width=defaults.width, layer=layers.front.silkscreen):
    return ModGraphic("fp_arc", ModArc(
        width = width,
        layer = layer,
        start = center,
        end = end,
        angle = angle,
    ))

# mk_text returns a graphical text element with the specified parameters.
def mk_text(pos=XYZ(), layer=layers.front.silkscreen, size=XY(1,1), thickness=defaults.thickness, content=""):
    return ModGraphic("fp_text", ModText(
        kind = text.user,
        layer = layer,
        text = content,
        at = pos,
        effects = TextEffects(font_size = size, thickness = thickness),
    ))

# mk_ref returns a graphical text element that will be populated with the
# reference of the symbol where it is used.
def mk_ref(pos=XYZ(), layer=layers.front.silkscreen, size=XY(1,1), thickness=defaults.thickness, content="REF**"):
    return ModGraphic("fp_text", ModText(
        kind = text.reference,
        layer = layer,
        text = content,
        at = pos,
        effects = TextEffects(font_size = size, thickness = thickness),
    ))

# mk_poly returns a graphical polygon with the specified parameters.
def mk_poly(points, pos = XYZ(), layer=layers.front.silkscreen, width = defaults.width):
    return ModGraphic("fp_poly", ModPolygon(
        points = points,
        layer = layer,
        width = width,
    ))

# mk_circle returns a graphical circle with the specified parameters.
def mk_circle(center=XY(), end=XY(), layer=layers.front.silkscreen, width=defaults.width):
    return ModGraphic("fp_circle", ModCircle(
        center = center,
        end = end,
        layer = layer,
        width = width,
    ))


# mk_smd_pad returns an smd pad with the specified parameters.
def mk_smd_pad(ident="", center=XY(), size=XY(1.4, 1.8), layers=layers.front.smd, shape=shape.rect, round_ratio=0.25):
  return Pad(ident,
        at = center,
        size = size,
        layers = layers,
        surface = pad.smd,
        shape = shape,
        round_rect_r_ratio = round_ratio)

# mk_th_pad returns a through-hole pad with the specified parameters.
def mk_th_pad(ident="", center=XY(), size=XY(1.7, 1.7), drill=XY(1,1), layers=layers.th, shape=shape.oval):
  return Pad(ident,
        at = center,
        size = size,
        drill_size = drill,
        layers = layers,
        surface = pad.through_hole,
        shape = shape)

# mk_mod_via returns a fake pad which is semantically equivalent to a
# normal-sized via, but possible in a module using pads.
def mk_mod_via(ident="1", center=XY(), layers=layers.th):
  return Pad(ident,
    at = center,
    size = XY(0.8, 0.8),
    drill_size = XY(0.5, 0.5),
    layers = layers,
    surface = pad.through_hole,
    shape = shape.circle)

def filter_graphics(graphics=[], filter="fp_text", text_type=""):
  out = []
  for graphic in graphics:
    if filter != graphic.ident:
      if filter == "fp_text" and graphic.ident == "fp_text": # need to look at the text_type
        if text_type == graphic.renderable.kind:
          continue
      out.append(graphic)
  return out

graphics = struct(
    line = mk_line,
    text = mk_text,
    ref  = mk_ref,
    poly = mk_poly,
    circle = mk_circle,
    arc  = mk_arc,
    filter = filter_graphics,
)

pads = struct(
  smd     = mk_smd_pad,
  th      = mk_th_pad,
  mod_via = mk_mod_via,
)

`)
