package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	var quitButton *tview.Button
	var flex *tview.Flex
	var flex1 *tview.Flex
	var loginform *tview.Form
	var loginButton *tview.Button
	var registerButton *tview.Button
	var buttonrow *tview.Flex

	loginButton = tview.NewButton("Login").SetSelectedFunc(func() {
		flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).AddItem(nil, 0, 1, false).AddItem(loginform, 25, 1, false).AddItem(nil, 0, 1, false), 0, 1, false), 9, 1, false).AddItem(nil, 0, 1, false)
		buttonrow.RemoveItem(loginButton)
		buttonrow.RemoveItem(registerButton)
	})

	registerButton = tview.NewButton("Register").SetSelectedFunc(func() {
		app.SetRoot(nil, true) // Replace with actual register screen
	})

	homepage := tview.NewTextView().
		SetText("Welcome to SlothScrypt.\n !! Lock it once, trust it forever !!\nWe ensure a high level of protection for your files using AES-256.").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	homepage.SetBorder(true).SetBackgroundColor(tcell.ColorGreen)
	modal := tview.NewModal().
		SetText("Do you want to quit??").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				app.Stop()
			} else {
				app.SetRoot(flex, true)
			}
		})
	// Quit button
	quitButton = tview.NewButton("Quit").
		SetSelectedFunc(func() {
			app.SetRoot(modal, true)
		})

	loginform = tview.NewForm().AddInputField("Username:", "", 11, nil, nil).AddInputField("Password:", "", 11, nil, nil).AddButton("Login ", nil).
		AddButton("Cancel ", nil)
	loginform.SetTitle("Login Now")
	loginform.SetButtonStyle(tcell.Style.Background(tcell.Style{}, tcell.ColorDarkOrchid))
	loginform.SetBorder(true)
	buttonrow = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(loginButton, 20, 0, true).
		AddItem(nil, 1, 1, false).
		AddItem(registerButton, 20, 1, true).
		AddItem(nil, 1, 1, false).
		AddItem(quitButton, 20, 1, true).
		AddItem(nil, 0, 1, false)
	flex1 = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(homepage, 5, 1, false).
		AddItem(nil, 0, 2, false).
		AddItem(buttonrow, 1, 1, true).
		AddItem(nil, 0, 1, false)
	flex1.SetBorder(true)

	flex = tview.NewFlex().AddItem(flex1, 10, 1, false).SetDirection(tview.FlexRow)
	flex.SetBorder(true)
	flex.AddItem(nil, 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
