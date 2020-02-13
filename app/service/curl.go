package service

import (
	"fmt"
	"dragon-fruit/app/global"
	"dragon-fruit/app/global/errorcode"
	"dragon-fruit/app/global/helper"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// sendGet CURL GET
func sendGet(apiURL string, header map[string]string, param map[string]interface{}) (body []byte, apiErr errorcode.Error) {
	time := time.Duration(global.TimeOut * time.Second)
	client := &http.Client{
		Timeout: time,
	}
	// 建立一個請求
	reqest, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001022, err.Error())

		return nil, apiErr
	}
	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}
	//組參數
	q := reqest.URL.Query()
	for pk, pv := range param {
		paramV := reflect.ValueOf(pv)
		if paramV.Kind() == reflect.Slice {
			for i := 0; i < paramV.Len(); i++ {
				value := paramV.Index(i)
				q.Add(pk, fmt.Sprintf("%v", value))
			}
			continue
		}
		q.Add(pk, fmt.Sprintf("%v", paramV))
	}
	reqest.URL.RawQuery = q.Encode()

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001025, err.Error())

		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("API_STATUS_ERROR: Status: %d, ErrorMsg: %v ", resp.StatusCode, err)
		apiErr = helper.ErrorHandle(global.WarnLog, 1001031, errMsg)

		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001028, err2.Error())

		return nil, apiErr
	}

	return body, apiErr
}

// sendPost CURL POST
func sendPost(apiURL string, header map[string]string, param map[string]interface{}) (body []byte, apiErr errorcode.Error) {
	// 組參數
	form := url.Values{}
	for pk, pv := range param {
		paramV := reflect.ValueOf(pv)
		if paramV.Kind() == reflect.Slice {
			for i := 0; i < paramV.Len(); i++ {
				value := paramV.Index(i)
				form.Add(pk, fmt.Sprintf("%v", value))
			}
			continue
		}
		form.Add(pk, fmt.Sprintf("%v", paramV))
	}

	// 建立一個請求
	time := time.Duration(global.TimeOut * time.Second)
	client := &http.Client{
		Timeout: time,
	}
	reqest, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001023, err.Error())

		return nil, apiErr
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001026, err.Error())

		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("API_STATUS_ERROR: Status: %d, ErrorMsg: %v ", resp.StatusCode, err)
		apiErr = helper.ErrorHandle(global.WarnLog, 1001032, errMsg)

		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001029, err2.Error())

		return nil, apiErr
	}

	return body, apiErr
}

// sendPut CURL PUT
func sendPut(apiURL string, header map[string]string, param map[string]interface{}) (body []byte, apiErr errorcode.Error) {
	// 組參數
	form := url.Values{}
	for pk, pv := range param {
		paramV := reflect.ValueOf(pv)
		if paramV.Kind() == reflect.Slice {
			for i := 0; i < paramV.Len(); i++ {
				value := paramV.Index(i)
				form.Add(pk, fmt.Sprintf("%v", value))
			}
			continue
		}
		form.Add(pk, fmt.Sprintf("%v", paramV))
	}

	// 建立一個請求
	time := time.Duration(global.TimeOut * time.Second)
	client := &http.Client{
		Timeout: time,
	}
	reqest, err := http.NewRequest(http.MethodPut, apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001024, err.Error())

		return nil, apiErr
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001027, err.Error())

		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("API_STATUS_ERROR: Status: %d, ErrorMsg: %v ", resp.StatusCode, err)
		apiErr = helper.ErrorHandle(global.WarnLog, 1001033, errMsg)

		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001030, err2.Error())

		return nil, apiErr
	}

	return body, apiErr
}
