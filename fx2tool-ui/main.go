package main

import (
	"log"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/Jockey66666/fx2tool/cmd"
)

var (
	lineEditor = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	findPresetBtn = new(widget.Clickable)
	listPresetBtn = new(widget.Clickable)
	logOutBtn     = new(widget.Clickable)
	switchEnvBtn  = new(widget.Clickable)
	isStage       = true
	list          = &layout.List{
		Axis: layout.Vertical,
	}
	progress            = 0
	progressIncrementer chan int
	topLabel            = "FX2 Tool"
	rootPath            = "~/Documents/PositiveGrid/BIAS_FX2/GlobalPresets"
)

type (
	// D : alias of layout.Dimensions
	D = layout.Dimensions

	// C : alias of layout.Context
	C = layout.Context
)

func main() {

	progressIncrementer = make(chan int)
	gofont.Register()

	go func() {
		for {
			time.Sleep(30 * time.Millisecond)
			progressIncrementer <- 1
		}
	}()

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(640), unit.Dp(400)), app.Title(topLabel))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme()

	var ops op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.ClipboardEvent:
				lineEditor.SetText(e.Text)
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e.Queue, e.Config, e.Size)
				showWidget(gtx, th)
				e.Frame(gtx.Ops)
			}
		case p := <-progressIncrementer:
			progress += p
			if progress > 100 {
				progress = 0
			}
			w.Invalidate()
		}
	}
}

func showWidget(gtx layout.Context, th *material.Theme) layout.Dimensions {
	for _, e := range lineEditor.Events() {
		if e, ok := e.(widget.SubmitEvent); ok {
			topLabel = e.Text
			lineEditor.SetText("")
		}
	}
	widgets := []layout.Widget{
		material.H3(th, topLabel).Layout,
		func(gtx C) D {
			e := material.Editor(th, lineEditor, "Preset name")
			return e.Layout(gtx)
		},
		func(gtx C) D {
			in := layout.UniformInset(unit.Dp(8))
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for findPresetBtn.Clicked() {
							cmd.FindPreset(rootPath, lineEditor.Text())
						}
						return material.Button(th, findPresetBtn, "Find preset").Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for listPresetBtn.Clicked() {
							cmd.ListAllPresets(rootPath)
						}
						return material.Button(th, listPresetBtn, "Show all presets").Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {

						for switchEnvBtn.Clicked() {
							if isStage {
								cmd.SwitchToProdction()
							} else {
								cmd.SwitchToStaging()
							}
							isStage = !isStage
						}

						displayText := "Switch to staging"
						if isStage {
							displayText = "Switch to production"
						}

						return material.Button(th, switchEnvBtn, displayText).Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for logOutBtn.Clicked() {
							cmd.Logout()
						}
						return material.Button(th, logOutBtn, "Log out").Layout(gtx)
					})
				}),
			)
		},
		func(gtx C) D {
			return material.ProgressBar(th, progress).Layout(gtx)
		},
	}

	return list.Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})
}
