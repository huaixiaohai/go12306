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
	request.Header.Set("Cookie", `_passport_session=50397f4ff6db4174b964e02e9b7e74ad9928; BIGipServerotn=2246574346.24610.0000; guidesStatus=off; highContrastMode=defaltMode; cursorStatus=off; BIGipServerpassport=904397066.50215.0000; RAIL_EXPIRATION=1673723415731; RAIL_DEVICEID=TBLLXKpFwagAYQFRCwef0ldXuQS5MBD69vexXtyWzNRY-IRNJYgbps1zBCT1tFTGAo1CXoDjEIFT4u-NUHKwXLh0Igfp6v8rgwv6qTSc4MAPbDU-PvxIlJUxFuq_0-Jlh0mY2kC6wydfkRjs0GO9iYvUkbdqDV78; route=495c805987d0f5c8c84b14f60212447d; _jc_save_fromStation=%u5317%u4EAC%2CBJP; _jc_save_toStation=%u8FD0%u57CE%u5317%2CABV; _jc_save_toDate=2023-01-11; _jc_save_wfdc_flag=dc; current_captcha_type=Z; _jc_save_fromDate=2023-01-24; fo=fwufmbb5z70v9qssZJmvbAAJWbU6oLqVztClDN8aNOCHnL6n3zu89fxTPZl20LMKYjv-IPtFyy1gtQq_-xkCmZftXkcHXWWOTylrIr8UXlfpvQptUEqeMYd5LOAyQDBd4I0lJEhH48jqpVMgoIHdw-5_nepcwpZfhh8ep16QdzW6mYkzWvQ_da5OFpc`)
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
