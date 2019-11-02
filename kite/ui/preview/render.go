package preview

import (
	"github.com/twitchyliquid64/kcgen/pcb"
)

func (p *Preview) Render(mod *pcb.Module) {
	p.mod = mod
	p.canvas.QueueDraw()
}
