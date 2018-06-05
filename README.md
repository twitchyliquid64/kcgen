# kcgen

See the two example generators `kcbox/` and `kcdashcircle/`.

## Build instructions

Go 1.8+ is required to build, so go install that first.

```shell

mkdir /tmp/kcbuild
cd /tmp/kcbuild
export GOPATH=/tmp/kcbuild
go get github.com/twitchyliquid64/kcgen

# Build the things
go build -o kc-outline github.com/twitchyliquid64/kcgen/kcoutline
go build -o kc-magnet github.com/twitchyliquid64/kcgen/kcmagnet
go build -o kc-dash-circle github.com/twitchyliquid64/kcgen/kcdashcircle
go build -o kc-box github.com/twitchyliquid64/kcgen/kcbox

# You should now have /tmp/kcbuild/kc-outline etc
```
## kc-outline  usage

Make a 40mm by 20mm box with rounded corners and edge mounts:

![Box image](https://raw.githubusercontent.com/twitchyliquid64/kcgen/master/kcoutline%2040x20.png)

```shell
./kc-outline --make-mounts 40x20 40 20
```

Usage:

```
USAGE: kc-outline <module-name> <width> <height>
  -make-mounts
    	Generate mounting holes
  -o string
    	Where output is written (default "-")
  -radius float
    	Rounded edges radius (default 3)
  -refY float
    	Y-axis offset at which module designator is placed
  -resolution int
    	How many interpolations to make per degree
```

## kc-magnet usage

Make a PCB magnet with 10 windings, with a track thickness of 0.25mm and a clearance of 0.16mm.

![Magnet image](https://raw.githubusercontent.com/twitchyliquid64/kcgen/master/kcmagnet.png)

```shell
./kc-magnet 10x-magnet-module 0.4 0.2 10
```

Usage:

```
USAGE: ./kc-magnet <module-name> <trace-thickness> <trace-clearance> <windings>
  -o string
    	Where output is written (default "-")
  -resolution int
    	How many interpolations to make per degree (default 1)
  -skip-windings float
    	How many windings to skip on the inside (default 1)
```
