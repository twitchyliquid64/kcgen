package lib

var drawLib = []byte(`
load("mod.lib", m="graphics")
load("pcb.lib", "pcb")

# draw_pcb_outline returns a set of lines that connects the given
# points.
def draw_pcb_outline(points=[],
  layer=layers.front.silkscreen,
  width=defaults.width,
  inclusive=True):
  return draw_outline(
    line = pcb.line,
    points = points,
    layer = layer,
    width = width,
    inclusive = inclusive,
  )


# draw_mod_outline returns a set of lines that connects the given
# points.
def draw_mod_outline(points=[],
  layer=layers.front.silkscreen,
  width=defaults.width,
  inclusive=True):
  return draw_outline(
    line = m.line,
    points = points,
    layer = layer,
    width = width,
    inclusive = inclusive,
  )

def draw_outline(line, points, layer, width, inclusive):
  out = []
  last = None
  for pt in points:
    if last:
      out.append(line(
        layer = layer,
        width = width,
        start = last,
        end   = pt,
      ))
    last = pt

  if inclusive and len(points) > 0:
    end = None
    if type(out[-1]) == "Line":
      end = out[-1].end
    elif type(out[-1]) == "ModGraphic":
      end = out[-1].renderable.end
    out.append(line(
      layer = layer,
      width = width,
      start = end,
      end   = points[0],
    ))
  return out

draw = struct(
  mod = struct(
    outline = draw_mod_outline,
  ),
  pcb = struct(
    outline = draw_pcb_outline,
  )
)
`)
