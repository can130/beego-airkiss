package controllers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

type jsApiConfig struct {
	ticket string
	token string
}

var onlyJsapi jsApiConfig

const (
	appid = "wxcxxxxx"
	secret = "94cb55xxxx"
)

func init() {
	go func() {
		for {
			onlyJsapi.token,_ = getJsToken(appid, secret)
			onlyJsapi.ticket,_ = getJsTicket(onlyJsapi.token)
			time.Sleep(time.Second*7000)
		}
	}()
}

//获取token
func getJsToken(appid, secret string) (string, error) {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+appid+"&secret="+secret)
	if err != nil {
		fmt.Println("GetToken from url fail", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GetToken read body fail", err)
		return "", err
	}

	var tokenStruct = struct {
		Token string `json:"access_token"`
	}{}

	if err = json.Unmarshal(body, &tokenStruct); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return "", err
	}
	return tokenStruct.Token,nil
}

//api_ticket会在生成JsAPIConfig对象signature（签名）属性时用到。
//api_ticket获取地址：https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token=ACCESS_TOKEN
func getJsTicket(token string) (string, error) {
	resp,err := http.Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="+token)
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	var ticket = struct {
		Ticket string
	}{}
	if err = json.Unmarshal(body, &ticket); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return "", err
	}
	return ticket.Ticket,nil
}