package api

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"exiliumgf/core"
	"log"
	"net/http"
	"time"
)

type MemberInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"Message"`
	Data    *struct {
		User struct {
			Score int `json:"score"`
		} `json:"user"`
	}
}

func MemberInfo() {
	// 尝试获取信息
	s := "{}"
	r := bytes.NewReader([]byte(s))

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, core.MemberInfoUrl, r)
	if err != nil {
		log.Printf("创建请求失败: req:%s err:%v \n", core.MemberInfoUrl, err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MyBot/1.0)")
	req.Header.Set("Authorization", core.Authorization)
	req.Header.Set("Host", "gf2-bbs-api.exiliumgf.com")
	req.Header.Set("Origin", "https://gf2-bbs.exiliumgf.com")
	req.Header.Set("Referer", "https://gf2-bbs.exiliumgf.com/")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求 %s 执行失败: %v \n", core.MemberInfoUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("http status : %d, 登录失败 \n", resp.StatusCode)
		core.Authorization = ""
		core.Score = 0
	} else {
		if resp.Body != nil {
			var memberInfoResp MemberInfoResp
			err := json.NewDecoder(resp.Body).Decode(&memberInfoResp)
			if err != nil && memberInfoResp.Code == 0 {
				log.Printf("解析响应体失败: %v \n", err)
				return
			}
			if memberInfoResp.Code != 0 {
				log.Printf("获取用户信息失败: %s \n", memberInfoResp.Message)
				return
			}
			log.Printf("获取用户信息成功: %v \n", memberInfoResp)
			core.Score = memberInfoResp.Data.User.Score
		}
	}
}

type LoginResp struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    *struct {
		Account struct {
			Token      string `json:"token"`
			PlatformId int    `json:"platform_id"`
			ChannelId  int    `json:"channel_id"`
			Uid        int    `json:"uid"`
		} `json:"account"`
	} `json:"data"`
}

type LoginReq struct {
	AccountName string `json:"account_name"`
	Passwd      string `json:"passwd"`
	Source      string `json:"source"`
}

func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Login(account string, password string) {
	MemberInfo()
	if core.Authorization != "" {
		log.Println("已登录")
		return
	}

	log.Println("尝试登录")
	reqBody := LoginReq{
		AccountName: account,
		Passwd:      MD5Hash(password),
		Source:      "phone",
	}
	marshal, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	r := bytes.NewReader(marshal)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, core.LoginUrl, r)
	if err != nil {
		log.Printf("创建请求失败: req:%s err:%v \n", core.LoginUrl, err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MyBot/1.0)")
	req.Header.Set("Host", "gf2-bbs-api.exiliumgf.com")
	req.Header.Set("Origin", "https://gf2-bbs.exiliumgf.com")
	req.Header.Set("Referer", "https://gf2-bbs.exiliumgf.com/")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求 %s 执行失败: %v \n", core.LoginUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("非预期状态码: %d \n", resp.StatusCode)
	}

	if resp.Body != nil {
		var loginResp LoginResp
		err := json.NewDecoder(resp.Body).Decode(&loginResp)
		if err != nil && loginResp.Code == 0 {
			log.Printf("解析响应体失败: %v \n", err)
			return
		}
		if loginResp.Code != 0 {
			log.Printf("登录失败: %s \n", loginResp.Message)
			return
		}
		core.Authorization = loginResp.Data.Account.Token
		log.Printf("登录成功: %v \n", loginResp)
	}
}

func SignIn() {
	log.Println("尝试签到")
	s := "{}"
	r := bytes.NewReader([]byte(s))
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, core.SignIn, r)
	if err != nil {
		log.Printf("创建请求失败: req:%s err:%v \n", core.SignIn, err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MyBot/1.0)")
	req.Header.Set("Authorization", core.Authorization)
	req.Header.Set("Host", "gf2-bbs-api.exiliumgf.com")
	req.Header.Set("Origin", "https://gf2-bbs.exiliumgf.com")
	req.Header.Set("Referer", "https://gf2-bbs.exiliumgf.com/")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求 %s 执行失败: %v \n", core.SignIn, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("非预期状态码: %d \n", resp.StatusCode)
	}
	log.Println("签到结束")
}

type SignInStatusResp struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    *struct {
		HasSignIn bool `json:"has_sign_in"`
	} `json:"data"`
}

func SignInStatus() bool {
	s := "{}"
	r := bytes.NewReader([]byte(s))
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, core.SignInStatus, r)
	if err != nil {
		log.Printf("创建请求失败: req:%s err:%v \n", core.SignIn, err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MyBot/1.0)")
	req.Header.Set("Authorization", core.Authorization)
	req.Header.Set("Host", "gf2-bbs-api.exiliumgf.com")
	req.Header.Set("Origin", "https://gf2-bbs.exiliumgf.com")
	req.Header.Set("Referer", "https://gf2-bbs.exiliumgf.com/")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求 %s 执行失败: %v \n", core.SignIn, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("非预期状态码: %d \n", resp.StatusCode)
	}
	if resp.Body != nil {
		var signInStatusResp SignInStatusResp
		err := json.NewDecoder(resp.Body).Decode(&signInStatusResp)
		if err != nil && signInStatusResp.Code == 0 {
			log.Printf("解析响应体失败: %v \n", err)
			return signInStatusResp.Data.HasSignIn
		}
	}
	return false
}
