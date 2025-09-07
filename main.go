package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// create new window
		w := new(app.Window)
		w.Option(app.Title("Port scanner"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

func testConnection(ip string, ports []string) bool {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, port), timeout)
		if err != nil {
			fmt.Println("Connecting error:", err)
			return false
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Opened", net.JoinHostPort(ip, port))
		}
	}
	return true
}

func pickPorts(s string) []string {
	switch s {
	case "RDP":
		var ptrs []string
		ptrs = append(ptrs, "3389")
		return ptrs

	case "LDAP":
		var ptrs []string
		ptrs = append(ptrs, "389")
		return ptrs

	default:
		var ptrs []string
		ptrs = append(ptrs, "21")
		return ptrs
	}
}

func draw(window *app.Window) error {
	theme := material.NewTheme()
	var resultEditor widget.Editor
	var ipEditor widget.Editor
	var ops op.Ops
	var button widget.Clickable
	var connectionIsValid bool
	var resultText string = "Waiting for the test.."
	var ipText string = "IP" //192.168.1.248
	var protocolRadioButton widget.Enum
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			if button.Clicked(gtx) {
				rbtnEnumValue, _ := protocolRadioButton.Focused()
				fmt.Println(rbtnEnumValue)
				connectionIsValid = testConnection(ipEditor.Text(), pickPorts(rbtnEnumValue))
				if connectionIsValid {
					resultText = "Success"
				} else {
					resultText = "test failed"
				}

			}

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEvenly,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						editorText := material.Editor(theme, &ipEditor, ipText)
						ipEditor.SingleLine = true
						ipEditor.Alignment = text.Middle
						return editorText.Layout(gtx)
					},
				),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "RDP", "RDP")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "LDAP", "ldap")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						editorText := material.Editor(theme, &resultEditor, resultText)
						resultEditor.SingleLine = true
						resultEditor.Alignment = text.Middle
						return editorText.Layout(gtx)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &button, "Test!")
						btn.CornerRadius = 20
						btn.TextSize = 15
						return btn.Layout(gtx)
					},
				),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(15)}.Layout,
				),
			)
			e.Frame(gtx.Ops)
		}
	}
}
