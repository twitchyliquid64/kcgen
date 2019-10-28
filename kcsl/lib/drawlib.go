package lib

var drawLib = []byte(`
load("mod.lib", m="graphics")

# draw_mod_outline returns a set of lines that connects the given
# points.
def draw_outline(points=[],
  layer=layers.front.silkscreen,
  width=defaults.width,
  inclusive=True):
  out = []
  last = None
  for pt in points:
    if last:
      out.append(m.line(
        layer = layer,
        width = width,
        start = last,
        end   = pt,
      ))
    last = pt

  if inclusive and len(points) > 0:
    out.append(m.line(
      layer = layer,
      width = width,
      start = out[-1].renderable.end,
      end   = points[0],
    ))
  return out

draw = struct(
  mod = struct(
    outline = draw_outline,
  )
)
`)
