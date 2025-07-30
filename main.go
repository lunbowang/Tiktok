package main

import (
	"Tiktok/router/routers"
	"Tiktok/setting"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化配置
	setting.Init()

	// 配置路由
	r := routers.NewRouter()
	server := &http.Server{
		Addr:           "127.0.0.1:8081",
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("server run success,run on 127.0.0.1:8081")

	// 用于接收服务启动错误
	errChan := make(chan error, 1)
	go func() {
		// 启动服务，若失败则将错误发送到errChan
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("server listen: %w", err)
		}
	}()

	// 等待中断信号（如Ctrl+C）
	quit := make(chan os.Signal, 1)
	// 监听 SIGINT（终端中断）和 SIGTERM（进程终止）信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit //阻塞等待信号
	fmt.Println("收到终止信号，开始优雅关机...")

	// 设置合理的超时时间，等待正在处理的请求完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关机服务：
	// 1. 停止接收新请求
	// 2. 等待已接收的请求处理完成（在超时时间内）
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("服务强制下线:%v\n", err)
		// 若优雅关闭失败，强制退出
		os.Exit(1)
	}

	fmt.Println("服务已优雅关闭")
}
