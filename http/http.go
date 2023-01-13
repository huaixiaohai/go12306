package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	urlPackage "net/url"
	"strings"
)

const (
	GET  = "GET"
	POST = "POST"
)

var cookie = ""

func SetCookie(str string) {
	cookie = str
}

func Get(url string, req, resp interface{}) error {
	return Do(GET, url, req, resp)
}

func Post(url string, req, resp interface{}) error {
	return Do(POST, url, req, resp)
}

func PostForm(url string, req map[string][]string, resp interface{}) error {
	//body, err := json.Marshal(req)
	//if err != nil {
	//	println(err.Error())
	//	return err
	//}
	var request *http.Request
	var err error
	request, err = http.NewRequest(POST, url, strings.NewReader(urlPackage.Values(req).Encode()))
	if err != nil {
		println(err.Error())
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	request.Header.Set("Cookie", cookie)
	var response *http.Response
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		println(err.Error())
		return err
	}
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		println(err.Error())
		return err
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func Do(method, url string, req interface{}, resp interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		println(err.Error())
		return err
	}
	var request *http.Request
	request, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		println(err.Error())
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	request.Header.Set("Cookie", `JSESSIONID=AEB1A12586B79817E84E65BBCC77C33C; tk=NOjmO7-WTLOBVF_kENFv_wc9rfzJcYCef_wkLky8TSrkAoU1TuvF6gubl1l0; guidesStatus=off; highContrastMode=defaltMode; cursorStatus=off; BIGipServerpassport=904397066.50215.0000; RAIL_EXPIRATION=1673723415731; RAIL_DEVICEID=TBLLXKpFwagAYQFRCwef0ldXuQS5MBD69vexXtyWzNRY-IRNJYgbps1zBCT1tFTGAo1CXoDjEIFT4u-NUHKwXLh0Igfp6v8rgwv6qTSc4MAPbDU-PvxIlJUxFuq_0-Jlh0mY2kC6wydfkRjs0GO9iYvUkbdqDV78; _jc_save_fromStation=%u5317%u4EAC%2CBJP; _jc_save_toStation=%u8FD0%u57CE%u5317%2CABV; _jc_save_toDate=2023-01-11; _jc_save_wfdc_flag=dc; current_captcha_type=Z; _jc_save_fromDate=2023-01-24; BIGipServerpool_passport=216269322.50215.0000; route=6f50b51faa11b987e576cdb301e545c4; BIGipServerotn=1105723658.64545.0000; uKey=e364ee66d4fccdc43331fd4b0f8deb425b428afe5d61fc87d341265700499ee8b6721ced9d73e6c175ee48e496bb92c3; fo=ftmf979gm4eq190kv63s7sAD0_LWPzcgIoQ6WOHiyMsNnWAK7cnZIn9p_Aoo9wSroEQa9VPeX-wxSmMxiA-ydTISeXsK0mt_9--ZBGxzTgkav495V46hor1xxI8CRxFnnHIgJkih5BdUeMRvwC5BdB0LuclRG6mmDCMWtN-nz_kq-Phh59BCZmyIvKo`)
	var response *http.Response
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		println(err.Error())
		return err
	}
	body, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		println(err.Error())
		return err
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
