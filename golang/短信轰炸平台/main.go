package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

// flag 参数
var name = flag.String("name", "89647828", "The greeting object.")
var password = flag.String("password", "896478218", "The greeting object.")

// 多个数字
var phones = flag.String("phones", "13241470518 13755776496", "The greeting object.")
var client = &http.Client{}
var cookie = ""

const (
	loginURL     = "https://aaccc.cc/index/index/login"
	orderAddAjax = "https://aaccc.cc/UserAjax/orderAddAjax"
)

func main() {
	flag.Parse()
	fmt.Println("Hello", *name)
	go func() {
		for {
			c := login()

			phonesNumbers := strings.Split(*phones, " ")
			for _, phone := range phonesNumbers {
				index(c, phone)
			}
			time.Sleep(time.Minute)
		}
	}()
	select {}
}

func login() string {
	url := "https://aaccc.cc/index/userajax/login"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", *name)
	_ = writer.WriteField("password", *password)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	c := res.Header.Get("Set-Cookie")
	cookie = strings.Split(c, ";")[0]
	return cookie
}

func index(cookie string, phone string) {
	fmt.Println("尝试添加号码：", phone)
	url := "https://aaccc.cc/index/UserAjax/orderAddAjax"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("phone", phone)
	_ = writer.WriteField("ts", "120")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", cookie+";thinkphp_show_page_trace=0|0")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
