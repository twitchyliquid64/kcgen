package lib

var shapeLib = []byte(`
load("math.lib", math="m")

def mk_box(
    width=0.0,
    height=0.0,
    center=XY(),
    rounded_radius=0.0,
    points_per_degree=0.25):
  top_left     = XY(center.x - width/2 + rounded_radius, center.y - height/2 + rounded_radius)
  top_right    = XY(center.x + width/2 - rounded_radius, center.y - height/2 + rounded_radius)
  bottom_left  = XY(center.x - width/2 + rounded_radius, center.y + height/2 - rounded_radius)
  bottom_right = XY(center.x + width/2 - rounded_radius, center.y + height/2 - rounded_radius)

  if rounded_radius < 0.000001:
      return [top_left, top_right, bottom_right, bottom_left]

  out  = math.point_arc(top_left,     radius=rounded_radius, start_angle=180, end_angle=270)
  out += math.point_arc(top_right,    radius=rounded_radius, start_angle=270, end_angle=360)
  out += math.point_arc(bottom_right, radius=rounded_radius, start_angle=0,   end_angle=90)
  out += math.point_arc(bottom_left, radius=rounded_radius,  start_angle=90,  end_angle=180)
  return out

shapes = struct(
  box = mk_box,
)

`)
