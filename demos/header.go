package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

/*
	TODO: move this into a separate component
*/

type HeaderIcon struct {
	vecty.Core
	IsTopPage bool
}

func (hi *HeaderIcon) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class(
				"mdc-icon-button", "material-icons", "mdc-top-app-bar__navigation-icon",
			),
			vecty.Attribute("title", "home"),
		),
		func() vecty.ComponentOrHTML {
			if hi.IsTopPage {
				return elem.Image(
					vecty.Markup(
						prop.Src("/assets/images/ic_component_24px_white.svg"),
						prop.Alt("Material logo"),
					),
				)
			}

			return elem.Italic(
				vecty.Markup(
					prop.Alt("Menu button"),
					vecty.Class("material-icons"),
				),
				vecty.Text("menu"),
			)
		}(),
	)
}

type HeaderBar struct {
	vecty.Core
	IsTopPage bool
}

func (hb *HeaderBar) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/HeaderBar.css")

	return elem.Header(
		vecty.Markup(
			vecty.Class("mdc-top-app-bar", "catalog-top-app-bar"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-top-app-bar__row"),
			),
			elem.Section(
				vecty.Markup(
					vecty.Class(
						"mdc-top-app-bar__section", "mdc-top-app-bar__section--align-start",
					),
				),
				&HeaderIcon{IsTopPage: hb.IsTopPage},
				elem.Span(
					vecty.Markup(
						vecty.Class(
							"mdc-top-app-bar__title", "catalog-top-app-bar__title",
						),
					),

					elem.Span(
						vecty.Markup(
							vecty.Class(
								"catalog-top-app-bar__title--small-screen",
							),
						),
						vecty.Text("MDC Web"),
					),
					elem.Span(
						vecty.Markup(
							vecty.Class(
								"catalog-top-app-bar__title--large-screen",
							),
						),
						vecty.Text("Material Components for the Web"),
					),
				),
			),
		),
	)
}
