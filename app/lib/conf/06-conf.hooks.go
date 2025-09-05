package conf

import ()

// In your ConfLib initialization:
func (c *ConfLib) RegisterHooks() {
	c.Hooks.AddPost("ConfPostView", c.ViewPagesSetup)
}
