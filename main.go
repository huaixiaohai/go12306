package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium/chrome"

	"github.com/tebeka/selenium"
)

const (
	userName = "16601116704"
	password = "2015.ami"
)

var chromeArgs = []string{
	"--lang=en-US",
	//"--no-default-browser-check", "--no-first-run",
	//"--no-sandbox", "--test-type",
	////"--window-size=1920,1080",
	//"--start-maximized",
	//"--start-maximized",
	//"--enable-automation",
	//"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	//"--headless",
}

func main() {

	const (
		chromeDriverPath = "C:\\Users\\zhanxiong.liu\\Downloads\\chromedriver_win32\\chromedriver.exe"
		port             = 9515
	)

	// Start a WebDriver server instance
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	selenium.SetDebug(false)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{
		"browserName":      "chrome",
		"pageLoadStrategy": "normal",
	}
	detach := true
	caps.AddChrome(chrome.Capabilities{
		//Path:         c.chromePath,
		Args: chromeArgs,
		//DebuggerAddr: c.debuggerAddr,
		ExcludeSwitches: []string{"enable-automation"},
		Detach:          &detach,
	})
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://kyfw.12306.cn/otn/resources/login.html"); err != nil {
		panic(err)
	}
	script := `Object.defineProperty(navigator, "webdriver", {get: () => false,});`
	x, err := wd.ExecuteScript(script, nil)
	if err != nil {
		panic(err)
	}
	println(x)
	// Get a reference to the text box containing code.
	userNameTxt, err := wd.FindElement(selenium.ByID, "J-userName")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := userNameTxt.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = userNameTxt.SendKeys(userName)
	if err != nil {
		panic(err)
	}

	//
	pwdTxt, err := wd.FindElement(selenium.ByID, "J-password")
	if err != nil {
		panic(err)
	}
	if err := pwdTxt.Clear(); err != nil {
		panic(err)
	}
	err = pwdTxt.SendKeys(password)
	if err != nil {
		panic(err)
	}

	loginBtn, err := wd.FindElement(selenium.ByID, "J-login")
	if err != nil {
		panic(err)
	}

	err = loginBtn.Click()
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 3)

	sliderGroup, err := wd.FindElement(selenium.ByID, "J-slide-passcode")
	if err != nil {
		panic(err)
	}

	slider, err := sliderGroup.FindElement(selenium.ByID, "nc_1_n1z")
	if err != nil {
		panic(err)
	}
	sliderSize, err := slider.Size()
	if err != nil {
		panic("")
	}

	sliderLoc, err := slider.Location()
	if err != nil {
		panic("")
	}

	sliderBox, err := wd.FindElement(selenium.ByID, "nc_1__scale_text")
	if err != nil {
		panic(err)
	}
	sliderBoxSize, err := sliderBox.Size()
	if err != nil {
		panic("")
	}
	sliderBoxLoc, err := sliderBox.Location()
	if err != nil {
		panic("")
	}

	fmt.Println(sliderLoc.X, sliderLoc.Y)
	fmt.Println(sliderBoxLoc.X, sliderBoxLoc.Y)
	sliderBoxLoc.Y = sliderBoxLoc.Y + 1000
	sliderBoxLoc.X = sliderBoxLoc.X + 1000
	fmt.Println(sliderSize.Width, sliderSize.Height)
	fmt.Println(sliderBoxSize.Width, sliderBoxSize.Height)
	offset := selenium.Point{X: sliderLoc.X, Y: sliderLoc.Y}
	offset1 := selenium.Point{X: sliderLoc.X + sliderBoxSize.Width, Y: sliderLoc.Y}
	wd.StorePointerActions(
		"mouse1",
		selenium.MousePointer,
		//// using selenium.FromViewport as the move origin
		//// which calculates the offset from 0,0.
		//// the other valid option is selenium.FromPointer.
		selenium.PointerMoveAction(0, offset, selenium.FromViewport),
		selenium.PointerPauseAction(250),
		selenium.PointerDownAction(selenium.LeftButton),
		selenium.PointerMoveAction(500, offset1, selenium.FromViewport),
		selenium.PointerPauseAction(250),
		selenium.PointerUpAction(selenium.LeftButton),
	)
	err = wd.PerformActions()
	if err != nil {
		panic("")
	}

	err = wd.ReleaseActions()
	if err != nil {
		panic("")
	}

	time.Sleep(10 * time.Second)

}
