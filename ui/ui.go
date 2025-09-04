package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Loadui() {
	app := tview.NewApplication()

	var quitButton *tview.Button
	var flex *tview.Flex
	var flex1 *tview.Flex
	var loginform *tview.Form
	var loginButton *tview.Button
	var registerButton *tview.Button
	var buttonrow *tview.Flex
	var registerform *tview.Form
	var loginContainer *tview.Flex
	// var buttonrowcheck *tview.Flex
	var nilspace *tview.Flex
	var nilspace2 *tview.Flex

	loginform = tview.NewForm().AddInputField("Username:", "", 11, nil, nil).AddInputField("Password:", "", 11, nil, nil).AddButton("Login ", nil).
		AddButton("Cancel ", func() {
			flex1.AddItem(buttonrow, 1, 1, true)
			flex1.AddItem(nil, 0, 1, false)

			flex.RemoveItem(loginContainer)
			flex.RemoveItem(nilspace)
			flex.RemoveItem(nilspace2)
			// buttonrowcheck = flex1.AddItem(nil, 0, 1, false)

		})
	loginform.SetTitle("Login Now")
	loginform.SetTitleColor(tcell.ColorAqua).
		SetBorderColor(tcell.ColorFuchsia).
		SetBackgroundColor(tcell.ColorBlack)
	loginform.SetButtonStyle(tcell.StyleDefault.
		Background(tcell.ColorMediumPurple).
		Foreground(tcell.ColorBlack).
		Bold(true))
	loginform.SetBorder(true)

	loginContainer = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).AddItem(loginform, 25, 1, false).
			AddItem(nil, 0, 1, false), 0, 1, false)

	//Login Button
	loginButton = tview.NewButton("Login").SetSelectedFunc(func() {
		flex1.RemoveItem(buttonrow)
		// nilspace = flex.AddItem(nil, 0, 1, false)
		flex.AddItem(loginContainer, 9, 1, false)
		nilspace2 = flex.AddItem(nil, 0, 1, false)

		// flex1.RemoveItem(buttonrowcheck)
	})
	loginButton.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorMediumPurple).
		Foreground(tcell.ColorBlack).
		Bold(true))

	//Register Button
	registerButton = tview.NewButton("Register").SetSelectedFunc(func() {
		flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(nil, 0, 1, false).AddItem(registerform, 31, 1, false).
				AddItem(nil, 0, 1, false), 0, 1, false), 11, 1, false).AddItem(nil, 0, 1, false)
		buttonrow.RemoveItem(loginButton)
		buttonrow.RemoveItem(registerButton)
	})
	registerButton.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorHotPink).
		Foreground(tcell.ColorBlack).
		Bold(true))

	//Homepage (title)
	homepage := tview.NewTextView().
		SetText("Welcome to SlothScrypt.\n !! Lock it once, trust it forever !!\nWe ensure a high level of protection for your files using AES-256.").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	homepage.SetBorder(true).
		SetBorderColor(tcell.ColorFuchsia).
		SetBackgroundColor(tcell.ColorBlack)

	//Modal for Quit Button
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

	modal.SetBackgroundColor(tcell.ColorBlack).
		SetButtonBackgroundColor(tcell.ColorMediumPurple)

	// Quit button
	quitButton = tview.NewButton("Quit").
		SetSelectedFunc(func() {
			app.SetRoot(modal, true)
		})
	quitButton.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorRed).
		Foreground(tcell.ColorWhite).
		Bold(true))

	//Login Form

	//Register Form
	registerform = tview.NewForm().AddInputField("Username:", "", 17, nil, nil).AddInputField("Email:", "", 17, nil, nil).AddInputField("Password:", "", 17, nil, nil).AddButton(" Register ", nil).
		AddButton(" Cancel ", nil)
	registerform.SetTitle("Register")
	registerform.SetTitleColor(tcell.ColorAqua).
		SetBorderColor(tcell.ColorFuchsia).
		SetBackgroundColor(tcell.ColorBlack)
	registerform.SetButtonStyle(tcell.StyleDefault.Background(tcell.ColorMediumPurple).Foreground(tcell.ColorBlack).Bold(true))
	registerform.SetBorder(true)

	//Button Row Of Login,Register,Quit
	buttonrow = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(loginButton, 20, 0, true).
		AddItem(nil, 1, 1, false).
		AddItem(registerButton, 20, 1, true).
		AddItem(nil, 1, 1, false).
		AddItem(quitButton, 20, 1, true).
		AddItem(nil, 0, 1, false)

		//flex to align homepage and buttonrow
	flex1 = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(homepage, 5, 1, false).
		AddItem(nil, 0, 1, false).
		AddItem(buttonrow, 1, 1, true).
		AddItem(nil, 0, 1, false)
	flex1.SetBorder(true)
	flex1.SetBorderColor(tcell.ColorFuchsia)

	//main Flex
	flex = tview.NewFlex().AddItem(flex1, 10, 1, false).SetDirection(tview.FlexRow)
	flex.SetBorder(true)
	flex.AddItem(nil, 0, 1, false)
	flex.SetBorderColor(tcell.ColorAqua)
	flex.SetBackgroundColor(tcell.ColorBlack)

	if err := app.SetRoot(flex, true).EnableMouse(true).SetFocus(flex).SetFocus(flex1).SetFocus(loginform).Run(); err != nil {
		panic(err)
	}
}
