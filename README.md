# kcgen

A standalone scripting engine for KiCad footprints and PCBs.

 - [x] Implement MVP
 - [x] Generation of Kicad Modules (footprints)
 - [ ] Ability to specify custom parameters, so script behaviour can be customized
 - [ ] Loading of existing modules so they can be edited / combined
 - [ ] Implement generation of text using custom fonts
 - [ ] Generate / edit KiCad PCBs

## Building

Make sure you have Go 1.13+ installed.

```shell
go get github.com/twitchyliquid64/kcgen
go build -o kcgen github.com/twitchyliquid64/kcgen/kcgen
# You should now have the binary kcgen in your current directory.
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

![Rounded box image](https://raw.githubusercontent.com/twitchyliquid64/kcgen/master/rounded_box.png)

### A longer example - SOIC-8

This example demonstrates some more advanced concepts, like using variables,
expressions, and list comprehensions to compute equations and loop multiple
times.

Execute the script like this: `./kcgen -o soic.kicad_mod soic.kcsl`

```python
load("mod.lib", m="graphics", p="pads")
load("shapes.lib", "shapes")
load("draw.lib", "draw")

# Configurable parameters.
pad_size          = XY(x=1.95, y=0.6)
dist_between_rows = 4.95
pitch             = 1.27
pins              = 8
extra_clearance   = XY(0.5, 0.3)

width  = dist_between_rows + pad_size.x
height = pins * pitch / 2
first_pad_y = -(height/2 - pitch/2)

mod = Mod(
    name = "SOIC-8_" + str(width) + "x" + str(height) + "_" + str(pitch),
    layer = layers.front.copper,
    description = "An 8 pin SOIC footprint.",
    tags = ["soic", "smd"],
    attrs = ["smd"],
    graphics = [m.ref(XYZ(0, 0))] +
      draw.mod.outline(shapes.box(width + extra_clearance.x,
                                  height + extra_clearance.y),
                       layer=layers.front.courtyard),
    pads = [ # left row
        p.smd(str(x+1),
            center = XY(-1 * dist_between_rows / 2, first_pad_y + x*pitch),
            size   = pad_size,
        ) for x in range(int(pins/2))
    ] + [ # right row
        p.smd(str(x+1+int(pins/2)),
            center = XY(dist_between_rows / 2, first_pad_y + x*pitch),
            size   = pad_size,
        ) for x in range(int(pins/2))
    ],
)
```

You can find more scripts in [kcgen/example](https://github.com/twitchyliquid64/kcgen/tree/master/kcgen/example)

## Scripting API

### Common functions & types

| Function      | Description   | Example |
| ------------- | ------------- | ------- |
| `load`  | Imports identifiers from a helper library or another script relative to the current directory.  | `load("draw.lib", "draw")` - Imports the 'draw' symbol from the draw helper library. |
| `print`  | Writes the argument to standard output.  | `print("Hello world")` - Prints _Hello world_. |
| `range`  | Generates a list of numbers based on the arguments provided. Identical to python. | `range(3)` generates `[0, 1, 2]`. |
| `XY` | Specifies coordinates in 2D. | `XY(1,2)` - coordinates are `x=1` and `y=2`.<br> `XY(x=3, y=4)` - coordinates are `x=3` and `y=4`. |
| `XYZ` | Specifies coordinates in 3D. | `XY(1,2,3)` - coordinates are `x=1`, `y=2`, and `z=3`.<br> `XYZ(x=3)` - coordinates are `x=3`, `y=0`, and `z=0`. |
| `Mod` | Generates a KiCad Module with the specified parameters. | See above. |

For a full list of Starlark constructs and builtin functions, please refer to the Starlark [language spec](https://github.com/bazelbuild/starlark/blob/master/spec.md).

### Constants

| Constant   |         |       |
| ---------- | ------- | ----- |
| `layers`   | Gives easy access to all the layer names, and the set of layer names typically used for smd & th pads. | layers.front.copper<br>layers.front.fab<br>layers.front.silkscreen<br>layers.front.courtyard<br>layers.front.paste<br>layers.back...<br><br>layers.smd<br>layers.th |
| `shape`    | Different kinds of pad shapes. | TODO |
| `defaults` | Typical values used as a default by KiCad | TODO |
| `pad`      | Different kinds of pad. | TODO |
| `text`     | Different kinds of module text element. | TODO |
| `zone_connect` | TODO | TODO |

### Helper libraries

In addition to native functions available in your scripts, there are a number
of helper libraries to help accomplish common tasks and make simple tasks
shorter.

#### `mod.lib`

`mod.lib` has shorthands for generating pads & graphical elements for modules. Import either the `graphics` or `pads` identifier (or both), like this:

##### Graphics

```python
load("mod.lib", m="graphics")
# You can now call graphics functions like m.blah.
```

| Function      | Description   | Example |
| ------------- | ------------- | ------- |
| `graphics.ref()`  | Places a text label that describes the reference (designator) of the module.  | `graphics.ref()` - Place reference at `XY(0,0)`.<br>`graphics.ref(XYZ(2,2))` - Place reference at `XY(2,2)`.<br>`graphics.ref(XYZ(2,2, text.vertical))` - Place vertical reference at `XY(2,2)`. |
| `graphics.text()` | Places a text element. | `graphics.text(pos=XYZ(1,2), content='a')` - Add a text element at `XY(1,2)` that says `A`. |
| `graphics.line()` | Places a line.<br>You can also specify `layer` and `width` attributes. | `graphics.line(start=XY(), end=XY(2,2))` - Add a line from `XY(0,0)` to `XY(2,2)`. |
| `graphics.circle()` | TODO | |
| `graphics.poly()` | TODO | |


##### Pads

```python
load("mod.lib", p="pads")
# You can now call pad functions like p.blah.
```

| Function      | Description   | Example |
| ------------- | ------------- | ------- |
| `pads.th()`   | Generates a through-hole pad. | `pads.th("2", center = XY(4, 3))` - Generates a pad called _2_ at `XY(4,3)`. <br><br>The drill defaults to `XY(1,1)` and pad size to `XY(1.7,1.7)`. |
| `pads.smd()` | Generates a surface-mount pad. | `p.smd("1", center = XY(1,3))` - Creates a smd pad at `XY(1,3)`.<br><br>The size defaults to `XY(1.4, 1.8)` and the shape to a rectangle. |

#### `draw.lib`

TODO

#### `math.lib`

TODO
