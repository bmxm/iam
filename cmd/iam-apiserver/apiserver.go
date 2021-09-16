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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiserver.NewApp("iam-apiserver").Run()
}

// 推荐使用cURL工具来执行HTTP请求


// Insomnia 是一款带UI界面的工具