package lib

var mathLib = []byte(`
def point_on_circle(center=XY(), radius=0.0, angle=0.0):
  x = center.x + (radius * math.cos(angle))
  y = center.y + (radius * math.sin(angle))
  return XY(x=x, y=y)

def point_on_circle_degrees(center=XY(), radius=0.0, angle=0.0):
  return point_on_circle(center=center, radius=radius, angle=(math.pi * angle)/180.0)

# point_arc returns the set of points in an arc from start_angle to end_angle.
def point_arc(center=XY(),
    radius=0.0,
    start_angle=0.0,
    end_angle=90.0,
    points_per_degree=0.5):
  diff = end_angle - start_angle
  c = int((diff) * points_per_degree)
  return [
    point_on_circle_degrees(
      center=center,
      radius=radius,
      angle=start_angle + diff * (i/c))
    for i in range(c)
  ] + [point_on_circle_degrees(center=center, radius=radius, angle=end_angle)]

m = struct(
  point_on_circle         = point_on_circle,
  point_on_circle_degrees = point_on_circle_degrees,
  point_arc               = point_arc,
)

`)
