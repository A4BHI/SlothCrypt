package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// // Modal confirmation dialog
	// modal := tview.NewModal().
	// 	SetText("Do you want to quit??").
	// 	AddButtons([]string{"Yes", "No"}).
	// 	SetDoneFunc(func(buttonIndex int, buttonLabel string) {
	// 		if buttonLabel == "Yes" {
	// 			app.Stop()
	// 		} else {
	// 			app.SetRoot(grid, true)
	// 		}
	// 	})

	// // Quit button
	// quitButton := tview.NewButton("Byee").
	// 	SetSelectedFunc(func() {
	// 		app.SetRoot(modal, true)
	// 	})

	// // Login and Register buttons
	// loginButton := tview.NewButton("Login").SetSelectedFunc(func() {
	// 	app.SetRoot(nil, true) // Replace with actual login screen
	// })

	// registerButton := tview.NewButton("Register").SetSelectedFunc(func() {
	// 	app.SetRoot(nil, true) // Replace with actual register screen
	// })

	// Side-by-side button grid
	// buttonRow := tview.NewGrid().
	// 	SetRows(1).                   // Single row
	// 	SetColumns(-5, 1, 1, 12, -1). // Two equal columns
	// 	AddItem(loginButton, 0, 0, 1, 1, 0, 0, true).
	// 	AddItem(registerButton, 0, 1, 1, 1, 0, 0, false)

	// TextView for welcome message
	homepage := tview.NewTextView().
		SetText("Welcome to SlothScrypt.\nWe ensure a high level of protection for your files using AES-256.").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	homepage.SetBorder(true).SetBackgroundColor(tcell.ColorGreen)

	flex1 := tview.NewFlex().AddItem(homepage, 4, 1, false).SetDirection(tview.FlexRow)
	flex1.SetBorder(true)
	// Main layout grid
	flex := tview.NewFlex().AddItem(flex1, 10, 1, false).SetDirection(tview.FlexRow)
	flex.SetBorder(true)

	if err := app.SetRoot(flex, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
