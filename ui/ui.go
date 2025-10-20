package ui

import (
	"context"
	"fmt"
	"math/rand"
	"sloth/db"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/crypto/bcrypt"
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
	var registerContainer *tview.Flex
	// var buttonrowcheck *tview.Flex
	var nilspace *tview.Flex
	var nilspace2 *tview.Flex
	var nilspace3 *tview.Flex
	var alert *tview.Modal
	var username *tview.InputField
	var password *tview.InputField

	username = tview.NewInputField().SetLabel("Username:")
	password = tview.NewInputField().SetLabel("Password:")

	alert = tview.NewModal().
		SetText("Account Created Successfully :) ").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {

				app.SetRoot(flex, true).SetFocus(loginform)

			}
		})

	alert.SetBackgroundColor(tcell.ColorBlack).
		SetButtonBackgroundColor(tcell.ColorMediumPurple)

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
	// Register Form
	registerform = tview.NewForm().
		AddFormItem(username).
		AddFormItem(password).
		AddButton("Register ", func() {

			userid := rand.Intn(100) + 100
			conn := db.Connect()
			var n string
			conn.QueryRow(context.TODO(), "Select username from users where username=$1", username.GetText()).Scan(&n)

			if n == "" {
				hashedpass, err := bcrypt.GenerateFromPassword([]byte(password.GetText()), 12)
				if err != nil {
					fmt.Println("Error generating hashedpass", err)
					return
				}
				_, err = conn.Exec(context.TODO(), "Insert into users values($1,$2,$3)", userid, username.GetText(), hashedpass)
				if err == nil {
					app.SetRoot(alert, true)

				}
			} else {
				return
			}

		}).
		AddButton("Cancel ", func() {
			flex1.AddItem(buttonrow, 1, 1, true)
			nilspace3 = flex1.AddItem(nil, 0, 1, false)
			flex.RemoveItem(nilspace)
			flex.RemoveItem(registerContainer)
			flex.RemoveItem(nilspace2)
		})

	registerform.SetTitle("Register Now")
	registerform.SetTitleColor(tcell.ColorAqua).
		SetBorderColor(tcell.ColorFuchsia).
		SetBackgroundColor(tcell.ColorBlack)
	registerform.SetButtonStyle(
		tcell.StyleDefault.
			Background(tcell.ColorMediumPurple).
			Foreground(tcell.ColorBlack).
			Bold(true),
	)
	registerform.SetBorder(true)

	registerContainer = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(nil, 0, 1, false).
				AddItem(registerform, 25, 1, false).
				AddItem(nil, 0, 1, false),
			0, 1, false,
		)

	registerButton = tview.NewButton("Register").SetSelectedFunc(func() {
		flex1.RemoveItem(buttonrow)
		flex1.RemoveItem(nilspace3)
		nilspace = flex.AddItem(nil, 0, 1, false)
		flex.AddItem(registerContainer, 9, 1, false)
		nilspace2 = flex.AddItem(nil, 0, 1, false)
	})

	registerButton.SetStyle(
		tcell.StyleDefault.
			Background(tcell.ColorMediumPurple).
			Foreground(tcell.ColorBlack).
			Bold(true),
	)

	loginform = tview.NewForm().
		AddFormItem(username).
		AddFormItem(password).
		AddButton("Login ", nil).
		AddButton("Cancel ", func() {
			flex1.AddItem(buttonrow, 1, 1, true)
			nilspace3 = flex1.AddItem(nil, 0, 1, false)
			flex.RemoveItem(nilspace)
			flex.RemoveItem(loginContainer)

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
		flex1.RemoveItem(nilspace3)
		nilspace = flex.AddItem(nil, 0, 1, false)
		flex.AddItem(loginContainer, 9, 1, false)
		nilspace2 = flex.AddItem(nil, 0, 1, false)
		// flex.AddItem(nil, 0, 1, false)

		// flex1.RemoveItem(buttonrowcheck)
	})
	loginButton.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorMediumPurple).
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
	// flex.AddItem(nil, 0, 1, false)
	flex.SetBorderColor(tcell.ColorAqua)
	flex.SetBackgroundColor(tcell.ColorBlack)

	if err := app.SetRoot(flex, true).EnableMouse(true).SetFocus(flex).SetFocus(flex1).SetFocus(loginform).Run(); err != nil {
		panic(err)
	}
}
