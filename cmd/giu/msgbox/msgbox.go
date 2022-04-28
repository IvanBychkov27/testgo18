package main

import (
	"fmt"
	"os"

	"github.com/AllenDang/giu"
)

func loop() {
	giu.Window("window").Layout(
		giu.PrepareMsgbox(),
		giu.Button("click me to see message box").OnClick(func() {
			giu.Msgbox("Info", "I'm a msgbox. press OK to close me")
		}),
		giu.Button("show yes-no dialog").OnClick(func() {
			giu.Msgbox("Question", "Exit? I'm yes-no dialog. Please take an action").
				Buttons(giu.MsgboxButtonsYesNo).
				ResultCallback(func(result giu.DialogResult) {
					switch result {
					case giu.DialogResultYes:
						fmt.Println("Yes clicked")
						os.Exit(0)
					case giu.DialogResultNo:
						fmt.Println("No clicked")
					}
				})
		}),
		giu.Button("show ok-cancel dialog").OnClick(func() {
			giu.Msgbox("ok-cancel", "I'm ok-cancel dialog").Buttons(giu.MsgboxButtonsOkCancel)
		}),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Msg box demo", 640, 480, 0)
	wnd.Run(loop)
}
