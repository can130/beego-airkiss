package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"io"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	link := "http://xxxxx.chuquanl.com/"
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	nonce := uuid.String()
	timestampstr := fmt.Sprintf("%d", time.Now().Unix())
	src := "jsapi_ticket=" + onlyJsapi.ticket + "&noncestr=" + nonce + "&timestamp=" + timestampstr + "&url=" + link
	h := sha1.New()
	io.WriteString(h, src)
	signature := fmt.Sprintf("%x", h.Sum(nil))

	c.Data["Appid"] = appid
	c.Data["Noncestr"] = nonce
	c.Data["Timestamp"] = timestampstr
	c.Data["Signature"] = signature
	c.TplName = "index.html"
}