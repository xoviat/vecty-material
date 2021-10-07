package formfield

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/base/applyer"
	"github.com/vecty-components/material/components/formfield"
)

// FF is a vecty-material formfield component.
type FF struct {
	*base.MDC
	vecty.Core
	Root     vecty.MarkupOrChild   `vecty:"prop"`
	Input    vecty.ComponentOrHTML `vecty:"prop"`
	Label    string                `vecty:"prop"`
	AlignEnd bool                  `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (c *FF) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	inputID := applyer.FindID(c.Input)
	return elem.Div(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		c.Input,
		elem.Label(
			vecty.Markup(
				vecty.MarkupIf(inputID != "", prop.For(inputID)),
			),
			vecty.Text(c.Label),
		),
	)
}

func (c *FF) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = formfield.New()
	}
	vecty.Markup(
		vecty.Class("mdc-form-field"),
		vecty.MarkupIf(c.AlignEnd,
			vecty.Class("mdc-form-field--align-end"),
		),
	).Apply(h)
	c.MDC.RootElement = h
}
