package boot

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// app启动地址
var appAddress = ""
var appDaemon = false

func SwaggerInit() {
	cmd := exec.Command("swag", "init", "--output", "app/swagger")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}

func InitService(container *framework.NiceContainer) error {
	// 从服务容器中获取kernel的服务实例
	kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
	// 从kernel服务实例中获取引擎
	core := kernelService.HttpEngine()
	flag.BoolVar(&appDaemon, "d", false, "start app daemon")
	flag.StringVar(&appAddress, "address", "", "设置app启动的地址，默认为:8888")
	if appAddress == "" {
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		if envService.Get("ADDRESS") != "" {
			appAddress = envService.Get("ADDRESS")
		} else {
			configService := container.MustMake(contract.ConfigKey).(contract.Config)
			if configService.IsExist("app.address") {
				appAddress = configService.GetString("app.address")
			} else {
				appAddress = ":8888"
			}
		}
	}
	// 创建一个Server服务
	server := &http.Server{
		Handler: core,
		Addr:    appAddress,
	}
	fmt.Println("app serve url:", appAddress)
	if err := startAppServe(server, container); err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

// 启动AppServer, 这个函数会将当前goroutine阻塞
func startAppServe(server *http.Server, c framework.Container) error {
	// 这个goroutine是启动服务的goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	closeWait := 3
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	if configService.IsExist("app.close_wait") {
		closeWait = configService.GetInt("app.close_wait")
	}
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(closeWait)*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		return err
	}
	return nil
}
