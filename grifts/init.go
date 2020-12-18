package grifts

import (
	"github.com/garretthaima/recipe/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
