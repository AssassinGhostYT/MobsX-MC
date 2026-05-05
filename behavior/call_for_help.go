package behavior

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
)

// CallForHelpBehavior makes an entity alert others nearby when it is in danger.
type CallForHelpBehavior struct {
	// RangeX, RangeY, RangeZ define the area of the call for help.
	RangeX, RangeY, RangeZ int
	
	Alerted bool
}

func (c *CallForHelpBehavior) Priority() int { return 100 } // Very high priority

func (c *CallForHelpBehavior) CanRun(e api.Entity, w api.World) bool {
	return c.Alerted
}

func (c *CallForHelpBehavior) Run(e api.Entity, w api.World) {
	e.AlertOthers(c.RangeX, c.RangeY, c.RangeZ)
	c.Alerted = false
}
