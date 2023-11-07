package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	assertErrorToNilf("Could not load .env: %s", godotenv.Load(".env"))
	assertErrorToNilf("could not download and install drivers for playwright: %v", playwright.Install())
	pw, err := playwright.Run()
	assertErrorToNilf("could not start playwright: %v", err)
	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)})
	assertErrorToNilf("could not launch browser: %v", err)
	page, err := browser.NewPage()
	assertErrorToNilf("could not create page: %v", err)
	_, err = page.Goto("https://login.secure.ninetyone.com")
	assertErrorToNilf("could not goto: %v", err)
	assertErrorToNilf("could not fill in username: %v", page.GetByLabel("Email").Fill(os.Getenv("EMAIL")))
	assertErrorToNilf("could not fill in password: %v", page.GetByLabel("Password").Fill(os.Getenv("PASSWORD")))
	assertErrorToNilf("could not log in: %v", page.GetByRole("button", playwright.PageGetByRoleOptions{Name: "LOG IN"}).Click())
	assertErrorToNilf("could not fill in 2FA code: %v", page.GetByPlaceholder("Enter the 6-digit code").Fill(os.Getenv("2FA_CODE")))
	assertErrorToNilf("could not log in after 2FA: %v", page.GetByRole("button", playwright.PageGetByRoleOptions{Name: ""}).Click())
	assertErrorToNilf("could not close browser: %w", browser.Close())
	assertErrorToNilf("could not stop Playwright: %w", pw.Stop())
}
