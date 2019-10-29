# kcgen

A standalone scripting engine for KiCad footprints and PCBs.

## Building

Make sure you have Go 1.13+ installed.

```shell
go get github.com/twitchyliquid64/kcgen
go build -o kcgen github.com/twitchyliquid64/kcgen/kcgen
```

## Quick start

Make a file `rounded_box.kcsl`

```python
load("mod.lib", m="graphics", p="pads")
load("shapes.lib", "shapes")
load("draw.lib", "draw")

mod = Mod(
    name = "test",
    layer = layers.front.copper,
    description = "This is a test module to demonstrate module generation.",
    attrs = ["virtual"],
    graphics = [m.ref()] +
      draw.mod.outline(shapes.box(10, 10, rounded_radius=1)),
)
```

Execute the script like this: `./kcgen -o rounded_box.kicad_mod rounded_box.kcsl`

## Scripting reference

TODO
