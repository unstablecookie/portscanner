package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
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
	case "RPC":
		var ptrs []string
		for i := 49152; i < 65536; i++ {
			str := strconv.Itoa(i)
			ptrs = append(ptrs, str)
		}
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
					buttonStyle := material.RadioButton(theme, &protocolRadioButton, "RPC", "RPC(49152-65535)")
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
