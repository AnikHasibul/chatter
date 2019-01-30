package components

import (
	"fmt"
	"time"

	"github.com/anikhasibul/chatter/app/model"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type ChatWindow struct {
	vecty.Core
	Session   string
	ActiveNow int
	EvFunc    func(string) *vecty.EventListener
	Chats     model.Chat
}

func (v *ChatWindow) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Header(
			vecty.Markup(
				vecty.Style("margin-bottom", "80px"),
				vecty.Class(
					"container",
					"top",
					"white",
					"padding",
					"card-2",
					"bar",
				),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class(
						"xlarge",
						"left",
						"text-blue",
					),
				),
				vecty.Text("ChaTter"),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class(
						"xlarge",
						"right",
						"tag",
						"round",
						"green",
					),
				),
				vecty.Text(fmt.Sprint(v.ActiveNow)),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class(
					"container",
				),
			),
			v.ChatBubbles(),
		),
		v.inputBox(),
	)
}

func (v *ChatWindow) ChatBubbles() *vecty.HTML {
	var bubbles vecty.List
	for _, m := range v.Chats {
		if m.From == "0" {
			bubbles = append(bubbles, elem.Div(
				elem.Heading6(
					vecty.Markup(
						vecty.Class(
							"small",
							"center",
							"text-red",
						),
					),
					vecty.Text(m.Message),
				),
			))
		} else {
			bubbles = append(bubbles, elem.Div(
				elem.Heading6(
					vecty.Markup(
						vecty.Class(
							"small",
							"text-grey",
						),
					),
					vecty.Text(m.From),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class(
							"light-grey",
							"round-large",
							"padding",
						),
						vecty.MarkupIf(m.From != "You",
							vecty.Class(
								"animate-zoom",
							),
						),
						vecty.MarkupIf(m.From == "You",
							vecty.Class(
								"animate-opacity",
							),
						),
					),
					vecty.Text(m.Message),
				),
				elem.Heading6(
					vecty.Markup(
						vecty.Class(
							"small",
							"text-grey",
						),
					),
					vecty.Text(time.Now().Format("2006/01/02 15:04:05")),
				),
			))
		}
	}
	return elem.Div(
		vecty.Markup(
			vecty.Style(
				"padding-bottom",
				"250px",
			),
		),
		bubbles,
	)
}

func (v *ChatWindow) inputBox() *vecty.HTML {
	idname := fmt.Sprint(time.Now().UnixNano())
	return elem.Form(
		vecty.Markup(
			vecty.Class(
				"bottom",
				"container",
			),
			vecty.MarkupIf(v.EvFunc != nil,
				v.EvFunc(idname),
			),
		),
		elem.Input(
			vecty.Markup(
				prop.ID(idname),
				prop.Placeholder("Your message here!"),
				vecty.Class(
					"input",
					"border-blue",
				),
				vecty.Style("margin-bottom", "20px"),
			),
		),
	)
}
