package api

import (
	"context"
	"encoding/json"
	"exiliumgf/core"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type TopicListReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Sort     int    `json:"sort"`
	Type     int    `json:"type"`
	Keyword  string `json:"keyword"`
}

type TopicListResp struct {
	Code    int    `json:"code"`
	Message string `json:"Message"`
	Data    *struct {
		List []struct {
			TopicId int `json:"topic_id"`
		} `json:"list"`
	} `json:"data"`
}

func TopicList() {
	log.Printf("开始查看，点赞，访问话题 \n")
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, core.ListUrl, nil)
	if err != nil {
		log.Printf("创建请求失败: req:%s err:%v \n", core.ListUrl, err)
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
		log.Printf("请求 %s 执行失败: %v \n", core.ListUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("http status : %d, 获取话题列表失败 \n", resp.StatusCode)
	}

	if resp.Body != nil {
		var topicListResp TopicListResp
		err := json.NewDecoder(resp.Body).Decode(&topicListResp)
		if err != nil {
			log.Printf("解析响应体失败: %v \n", err)
			return
		}
		log.Printf("获取话题列表成功: %v \n", topicListResp)

		// 准备点赞，浏览
		for _, v := range topicListResp.Data.List {
			Info(v.TopicId)
			Like(v.TopicId)
			Share(v.TopicId)
		}
	}
	log.Printf("结束查看，点赞，访问话题 \n")
}

type ShareReq struct {
	TopicId int `json:"topic_id"`
}

func Share(id int) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf(core.Share, id, id), nil)
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
		log.Printf("http status : %d, 分享话题失败 \n", resp.StatusCode)
	}

	if resp.StatusCode == http.StatusOK {
		log.Printf(" 分享话题成功: \n")
	}
}

type InfoReq struct {
	TopicId int `json:"topic_id"`
}

func Info(id int) {

	infoReq := &InfoReq{
		TopicId: id,
	}
	marshal, err := json.Marshal(infoReq)
	if err != nil {
		log.Printf("marshal err:%v \n", err)
		return
	}
	r := strings.NewReader(string(marshal))
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf(core.Info, id), r)
	if err != nil {
		log.Printf("创建请求失败: req:%s err:%v \n", core.Info, err)
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
		log.Printf("请求 %s 执行失败: %v \n", core.Info, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("http status : %d, 查看话题失败 \n", resp.StatusCode)
	}
	if resp.StatusCode == http.StatusOK {
		log.Printf(" 查看话题成功: \n")
	}
}

type LikeReq struct {
	TopicId int `json:"topic_id"`
}

func Like(id int) {

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf(core.Like, id, id), nil)
	if err != nil {
		log.Printf("创建请求失败: req:%s err:%v \n", core.Like, err)
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
		log.Printf("请求 %s 执行失败: %v \n", core.Like, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("http status : %d, 分享失败 \n", resp.StatusCode)
	}
	if resp.StatusCode == http.StatusOK {
		log.Printf("点赞成功: \n")
	}
}
