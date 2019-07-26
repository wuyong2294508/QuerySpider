package main

import (
	"github.com/robfig/cron"
	"fmt"
	"net/http"
	"net/smtp"
	"io/ioutil"
	"strings"
	//"time"
)

func main() {
    //fmt.Printf("hello, wuyong!\n");
	c := cron.New();
	query();
	c.AddFunc("@hourly", query);
	c.Start();
	select{};
}

func query() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	url := "http://query.ruankao.org.cn/certificate";
	resp, err := http.Get(url);
	if err != nil {
		fmt.Printf("http.Get Error!\n");
		return;
	}
	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		fmt.Printf("ioutil.ReadAll(resp.Body) Error!\n");
		return;
	}
	//fmt.Printf(string(body));
	if strings.Contains(string(body), "2019年上半年") {
		//fmt.Printf("2019年上半年证书查询已开放，请查询！");
		sendEmail("2019年上半年证书查询已开放，请查询！\nhttp://query.ruankao.org.cn/certificate")
	}
}

func sendEmail(body string) {
    auth := smtp.PlainAuth("", "*********@qq.com", "*********", "smtp.qq.com")
    to := []string{"********@qq.com"}
    nickname := "证书查询"
    user := "*********@qq.com"
    subject := "2019年上半年证书查询提醒"
    content_type := "Content-Type: text/plain; charset=UTF-8"
    //body := "This is the email body."
    msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
        "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
    if err != nil {
        fmt.Printf("send mail error: %v", err)
    }
}