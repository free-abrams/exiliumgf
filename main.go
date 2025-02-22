package main

import (
	"encoding/json"
	"exiliumgf/api"
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"math/rand"
	"os"
	"time"
)

var configFile = flag.String("f", "config.json", "the config file")

type Config struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func main() {
	// 设置标准库 log 的输出格式（自动包含文件名和行号）
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Lshortfile 显示短文件名和行号

	flag.Parse()
	var c Config
	file, err := os.ReadFile(*configFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(fmt.Errorf("解析配置文件失败: %v", err))
	}

	// 创建带中国时区的 Cron 调度器
	location, _ := time.LoadLocation("Asia/Shanghai")
	xcron := cron.New(cron.WithLocation(location))

	// 添加每日凌晨任务（每天 00:00 触发）
	_, err = xcron.AddFunc("0 0 * * *", func() {
		// 生成随机延时（5-6小时）
		delay := time.Duration(5*60+rand.Intn(60)) * time.Minute
		fmt.Printf("任务已触发，将在 %v 后执行\n", delay)

		// 设置延时执行
		time.AfterFunc(delay, func() {
			fmt.Printf("开始每日签到 @ %s\n", time.Now().Format(time.RFC3339))
			api.Login(c.Account, c.Password)
			api.MemberInfo()
			api.SignIn()
			api.TopicList()
			api.ExchangeList()
		})
	})

	if err != nil {
		panic(fmt.Sprintf("创建定时任务失败: %v", err))
	}

	// 启动定时任务
	xcron.Start()
	fmt.Println("定时服务已启动，等待任务触发...")

	// 保持程序运行
	select {}
}
