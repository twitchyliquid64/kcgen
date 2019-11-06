package lib

var flattenlib = []byte(`
load("mod.lib", m="graphics")
load("math.lib", pocd="point_on_circle_degrees")
# load("flatten.lib", "flatten")

def flatten_graphics(graphics):
  out = []
  for g in graphics:
    if getattr(g.renderable, "layer", None):
      g.renderable.layer = layers.front.fab
    out.append(g)
  return out

def flatten_pads(pads):
  out = []
  for p in pads:
    if str(p.shape) == str(shape.rect):
      out.append(m.line( # top line
        start =XY(p.at.x-p.size.x/2, p.at.y-p.size.y/2),
        end   =XY(p.at.x+p.size.x/2, p.at.y-p.size.y/2),
        layer = layers.front.fab)
      )
      out.append(m.line( # bottom line
        start = XY(p.at.x-p.size.x/2, p.at.y+p.size.y/2),
        end   = XY(p.at.x+p.size.x/2, p.at.y+p.size.y/2),
        layer = layers.front.fab)
      )
      out.append(m.line( # left line
        start = XY(p.at.x-p.size.x/2, p.at.y-p.size.y/2),
        end   = XY(p.at.x-p.size.x/2, p.at.y+p.size.y/2),
        layer = layers.front.fab)
      )
      out.append(m.line( # right line
        start = XY(p.at.x+p.size.x/2, p.at.y-p.size.y/2),
        end   = XY(p.at.x+p.size.x/2, p.at.y+p.size.y/2),
        layer = layers.front.fab)
      )
    elif str(p.shape) in [str(shape.circle), str(shape.oval)]:
      out.append(m.circle(
        center = XY(p.at.x, p.at.y),
        end    = XY(p.at.x+p.size.x/2, p.at.y+p.size.y/2),
        layer  = layers.front.fab)
      )

    if p.drill_size.x > 0:
      center = XY(p.at.x+p.drill_offset.x, p.at.y+p.drill_offset.y)
      out.append(m.circle(
        center = center,
        end    = XY(p.at.x+p.drill_offset.x+p.drill_size.x/2, p.at.y+p.drill_offset.y+p.drill_size.y/2),
        layer  = layers.front.fab)
      )
      br = pocd(center=center, radius=p.drill_size.x/2, angle=45.0)
      tr = pocd(center=center, radius=p.drill_size.x/2, angle=315.0)
      tl = pocd(center=center, radius=p.drill_size.x/2, angle=225.0)
      bl = pocd(center=center, radius=p.drill_size.x/2, angle=135.0)
      out.append(m.line( # right-to-left cross
        start = br,
        end   = tl,
        layer = layers.front.fab)
      )
      out.append(m.line( # left-to-right cross
        start = bl,
        end   = tr,
        layer = layers.front.fab)
      )
  return out

flatten = struct(
  pads = flatten_pads,
  graphics = flatten_graphics,
)
`)
