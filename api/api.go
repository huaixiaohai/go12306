package api

import (
	"fmt"
	"lzx/go12306/http"
)

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

type LoginConfResp struct {
}
type LoginConfReq struct {
}

// LoginConf 查看用户信息
func LoginConf() {
	const url = `https://kyfw.12306.cn/otn/login/conf`

	resp := &LoginConfResp{}
	req := &LoginConfReq{}
	http.Post(url, req, resp)
}

type QueryTicketResp struct {
}

// QueryTicket 查询车次信息
func QueryTicket(date string) (*QueryTicketResp, error) {
	//const url = `https://kyfw.12306.cn/otn/leftTicket/queryZ?leftTicketDTO.train_date=2023-01-23&leftTicketDTO.from_station=BJP&leftTicketDTO.to_station=YNV&purpose_codes=ADULT`
	var url = fmt.Sprintf(`https://kyfw.12306.cn/otn/leftTicket/queryZ?leftTicketDTO.train_date=%s&leftTicketDTO.from_station=%s&leftTicketDTO.to_station=%s&purpose_codes=%s`, "2023-01-23", "BJP", "YNV", "ADULT")

	resp := &QueryTicketResp{}
	err := http.Get(url, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

type SubmitOrderResp struct {
}

// SubmitOrder 提交订单请求
func SubmitOrder() (*SubmitOrderResp, error) {
	const url = `https://kyfw.12306.cn/otn/leftTicket/submitOrderRequest`

	reqForm := map[string][]string{
		"secretStr":               {},
		"train_date":              {"2-23-01-22"},
		"back_tran_date":          {"2023-01-13"},
		"tour_flag":               {"dc"},
		"purpose_codes":           {"ADULT"},
		"query_from_station_name": {"北京"},
		"query_to_station_name":   {"运城北"},
	}

	resp := &SubmitOrderResp{}
	err := http.Post(url, reqForm, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
