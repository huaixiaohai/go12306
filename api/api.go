package api

import "lzx/go12306/http"

//type CheckLoginVerifyReq struct {
//	UserName string `json:"username"`
//	AppID    string `json:"appid"`
//}

const OTN = "otn"

type CheckLoginVerifyResp struct {
	LoginCheckCode string `json:"login_check_code"`
	ResultCode     int64  `json:"result_code"`
	ResultMessage  string `json:"result_message"`
}

func CheckLoginVerify(userName string) (*CheckLoginVerifyResp, error) {
	const url = `https://kyfw.12306.cn/passport/web/checkLoginVerify`
	formReq := map[string][]string{
		"username": {userName},
		"appid":    {OTN},
	}
	resp := &CheckLoginVerifyResp{}
	err := http.PostForm(url, formReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//type SlidePasscodeReq struct {
//	SlideMode string
//	AppID     string
//	UserName  string
//}

type SlidePasscodeResp struct {
	IfCheckSlidePasscodeToken string `json:"if_check_slide_passcode_token"`
	ResultCode                string `json:"result_code"`
	ResultMessage             string `json:"result_message"`
}

func SlidePasscode(userName string) (*SlidePasscodeResp, error) {
	const url = `https://kyfw.12306.cn/passport/web/slide-passcode`
	formReq := map[string][]string{
		"slideMode": {"1"},
		"username":  {userName},
		"appid":     {OTN},
	}
	resp := &SlidePasscodeResp{}
	err := http.PostForm(url, formReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func AliSession() {
	const url = `https://ynuf.aliapp.org/service/um.json`
	formReq := map[string][]string{
		"data": {},
		"xa":   {},
		"xt":   {},
		"eft":  {},
	}
	resp := &SlidePasscodeResp{}
	err := http.PostForm(url, formReq, resp)
	if err != nil {
		println(err.Error())
	}
}

type LoginResp struct {
}

//func Login() (*LoginResp, error) {
//	const url = `https://kyfw.12306.cn/passport/web/login`
//	formReq := map[string][]string{
//		"slideMode": {"1"},
//		"username":  {userName},
//		"appid":     {OTN},
//	}
//	resp := &SlidePasscodeResp{}
//	err := http.PostForm(url, formReq, resp)
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
