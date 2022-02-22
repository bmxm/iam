// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// apiserver is the api server for iam-apiserver service.
// it is responsible for serving the platform RESTful resource management.
package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/marmotedu/iam/internal/apiserver"
)

// 通过一个名为 iam-apiserver 的进程，对外提供RESTful API接口
// 完成用户、秘钥、策略三种REST资源的增删改查。

// 代码实现可以分为四部分：配置处理、启动流程、请求处理流程、代码架构
/*
配置处理 ： Options配置、应用配置、HTTP/GRPC服务配置
	- Options配置：用来构建命令行参数，它的值来源于命令行选项或者配置文件（也可能是二者Merge后的配置）。
		Options可以用来构建应用架构，Options配置也是应用配置的输入。

	- 应用配置： iam-apiserver 组件中需要的一切配置。如：启动HTTP/GRPC需要配置监听地址和端口，初始化数据库地址、用户名、密码等。

	- HTTP/GRPC服务配置： 启动HTTP/GRPC服务需要的配置
*/

/*
1 iam-apiserver 作为 iam 系统的核心组件，需要第一个安装
	source scripts/install/environment.sh
	make build BINS=iam-apiserver
	sudo cp _output/platforms/linux/amd64/iam-apiserver ${IAM_INSTALL_DIR}/bin

2 生成并安装 iam-apiserver 的配置文件（iam-apiserver.yaml）
	./scripts/genconfig.sh scripts/install/environment.sh configs/iam-apiserver.yaml > iam-apiserver.yaml
	sudo mv iam-apiserver.yaml ${IAM_CONFIG_DIR}

3 创建并安装 iam-apiserver systemd unit 文件
	./scripts/genconfig.sh scripts/install/environment.sh init/iam-apiserver.service > iam-apiserver.service
	sudo mv iam-apiserver.service /etc/systemd/system/

4 启动 iam-apiserver 服务：
	sudo systemctl daemon-reload
	sudo systemctl enable iam-apiserver
	sudo systemctl restart iam-apiserver
	systemctl status iam-apiserver # 查看 iam-apiserver 运行状态，如果输出中包含 active (running)字样说明 iam-apiserver 成功启动
*/

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiserver.NewApp("iam-apiserver").Run()
}

// 推荐使用cURL工具来执行HTTP请求

// Insomnia 是一款带UI界面的工具
