package main

import (
	"fmt"
	"image/color"
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
	"golang.org/x/exp/shiny/materialdesign/icons"
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

func pickPorts(s string, portNumber string) []string {
	switch s {
	case "RDP":
		var ptrs []string
		ptrs = append(ptrs, "3389")
		return ptrs

	case "LDAP":
		var ptrs []string
		ptrs = append(ptrs, "389")
		return ptrs

	case "DNS":
		var ptrs []string
		ptrs = append(ptrs, "53")
		return ptrs

	case "KRB1":
		var ptrs []string
		ptrs = append(ptrs, "88")
		return ptrs
	case "NTP":
		var ptrs []string
		ptrs = append(ptrs, "123")
		return ptrs
	case "RPCm":
		var ptrs []string
		ptrs = append(ptrs, "135")
		return ptrs
	case "LDAPS1":
		var ptrs []string
		ptrs = append(ptrs, "636")
		return ptrs
	case "SMB1":
		var ptrs []string
		ptrs = append(ptrs, "137", "138", "139")
		return ptrs
	case "SMB2":
		var ptrs []string
		ptrs = append(ptrs, "445")
		return ptrs
	case "KRB2":
		var ptrs []string
		ptrs = append(ptrs, "464")
		return ptrs
	case "LDAPgc":
		var ptrs []string
		ptrs = append(ptrs, "3268", "3269")
		return ptrs
	case "ADWS":
		var ptrs []string
		ptrs = append(ptrs, "9389")
		return ptrs
	case "OTHER":
		var ptrs []string
		ptrs = append(ptrs, portNumber)
		return ptrs
	default:
		var ptrs []string
		ptrs = append(ptrs, "21")
		return ptrs
	}
}

func draw(window *app.Window) error {
	theme := material.NewTheme()
	var (
		resultEditor        widget.Editor
		ipEditor            widget.Editor
		manualPort          widget.Editor
		ops                 op.Ops
		button              widget.Clickable
		coloredBut          widget.Clickable
		connectionIsValid   bool
		resultText          string = "Waiting for the test.."
		ipText              string = "IP" //192.168.1.248
		manualPortText      string = "1"
		protocolRadioButton widget.Enum
		resultButtonColor   bool
		col                 color.NRGBA
		colBack             color.NRGBA
	)

	toggleCheckBoxIcon, _ := widget.NewIcon(icons.ToggleCheckBox) //(icons.ToggleCheckBox) //widget.NewIcon(icons.ActionSearch)
	toggleCheckBoxBlankIcon, _ := widget.NewIcon(icons.ToggleCheckBoxOutlineBlank)
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			if button.Clicked(gtx) {
				rbtnEnumValue, _ := protocolRadioButton.Focused()
				fmt.Println(rbtnEnumValue)
				connectionIsValid = testConnection(ipEditor.Text(), pickPorts(rbtnEnumValue, manualPort.Text()))
				if connectionIsValid {
					resultText = "Success"
					resultButtonColor = true
				} else {
					resultText = "test failed"
					resultButtonColor = false
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
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "RDP", "RDP(3389)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "LDAP", "ldap(389)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "DNS", "DNS(53)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "KRB1", "Kerberos(88)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "KRB2", "KerberosPWD(464)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "NTP", "Time(123)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "RPCm", "RPC(135)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "LDAPS1", "ldaps(636)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "SMB1", "smb/netbios(137,138,139)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "SMB2", "smb(445)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "LDAPgc", "LDAP GC(3268,3269)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "ADWS", "AD WS(9389)")
					return buttonStyle.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEnd, Alignment: layout.Middle}.Layout(gtx, //Spacing: layout.SpaceEnd
						layout.Rigid(
							func(gtx layout.Context) layout.Dimensions {
								buttonStyle := material.RadioButton(theme, &protocolRadioButton, "OTHER", "Other:")
								return buttonStyle.Layout(gtx)
							},
						),
						layout.Rigid(
							func(gtx layout.Context) layout.Dimensions {
								manualText := material.Editor(theme, &manualPort, manualPortText)
								return manualText.Layout(gtx)
							},
						),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Spacing: layout.Spacing(layout.Middle), Alignment: layout.Middle}.Layout(gtx, //{Axis: layout.Horizontal, Spacing: layout.Spacing(layout.Middle)}
						layout.Rigid(
							func(gtx layout.Context) layout.Dimensions {
								editorText := material.Editor(theme, &resultEditor, resultText)
								resultEditor.SingleLine = true
								resultEditor.Alignment = text.Middle
								return editorText.Layout(gtx)
							},
						),
						layout.Rigid(
							layout.Spacer{Width: unit.Dp(25)}.Layout,
						),
						layout.Rigid(
							func(gtx layout.Context) layout.Dimensions {
								var clrBtn material.IconButtonStyle
								if resultButtonColor {
									col = color.NRGBA{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}
									colBack = color.NRGBA{R: 0x80, B: 0x00, G: 0xFF, A: 0xFF}
									clrBtn = material.IconButton(theme, &coloredBut, toggleCheckBoxIcon, "Test!")
								} else {
									col = color.NRGBA{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}
									colBack = color.NRGBA{R: 0xFF, B: 0x00, G: 0x80, A: 0xFF}
									clrBtn = material.IconButton(theme, &coloredBut, toggleCheckBoxBlankIcon, "Test!")
								}
								//clrBtn.Size = unit.Dp(15)
								clrBtn.Color = col
								clrBtn.Background = colBack
								return clrBtn.Layout(gtx)
							},
						))
				}),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &button, "Test!")
						btn.CornerRadius = 10
						btn.TextSize = 15
						return btn.Layout(gtx)
					},
				),
			)
			e.Frame(gtx.Ops)
		}
	}
}
