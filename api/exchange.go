package api

import (
	"bytes"
	"context"
	"encoding/json"
	"exiliumgf/core"
	"log"
	"net/http"
	"strings"
	"time"
)

type ExchangeListResp struct {
	Code int `json:"code"`
	Data *struct {
		List []struct {
			ExchangeId       int    `json:"exchange_id"`
			MaxExchangeCount int    `json:"max_exchange_count"`
			ExchangeCount    int    `json:"exchange_count"`
			UseScore         int    `json:"use_score"`
			ItemName         string `json:"item_name"`
		} `json:"list"`
	} `json:"data"`
}

func ExchangeList() {
	log.Println("尝试兑换")

	s := "{}"
	r := bytes.NewReader([]byte(s))
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, core.ExchangeList, r)
	if err != nil {
		log.Printf("创建请求失败: req:%s err:%v", core.MemberInfoUrl, err)
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
		log.Printf("请求执行失败: %v \n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("http status : %d, 获取话题列表失败 \n", resp.StatusCode)
	}

	if resp.Body != nil {
		var exchangeListResp ExchangeListResp
		err := json.NewDecoder(resp.Body).Decode(&exchangeListResp)
		if err != nil {
			log.Printf("解析响应体失败: %v \n", err)
			return
		}

		for _, v := range exchangeListResp.Data.List {

			ok := false
			for _, id := range core.ExchangeAllowed {
				if id == v.ExchangeId {
					ok = true
				}
			}
			if !ok {
				continue
			}

			log.Printf("道具：%s，可兑换：%d", v.ItemName, v.MaxExchangeCount-v.ExchangeCount)
			if v.ExchangeCount < v.MaxExchangeCount && core.Score >= v.UseScore {
				Exchange(v.ExchangeId)
			}
		}
	}
	log.Println("兑换结束")
}

type ExchangeReq struct {
	ExchangeId int `json:"exchange_id"`
}

func Exchange(id int) {
	exchangeReq := ExchangeReq{
		ExchangeId: id,
	}
	marshal, err := json.Marshal(exchangeReq)
	if err != nil {
		log.Printf("marshal err:%v", err)
		return
	}
	r := strings.NewReader(string(marshal))

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, core.Exchange, r)
	if err != nil {
		log.Printf("创建请求失败: req:%s err:%v", core.Exchange, err)
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
		log.Printf("请求执行失败: %v \n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("http status : %d, 兑换失败 \n", resp.StatusCode)
	}
	if resp.StatusCode == http.StatusOK {
		log.Printf("兑换成功 \n")
	}
}
