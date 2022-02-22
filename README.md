# IAM - 身份识别与访问管理系统

IAM = Identity and Access Management

IAM 是一个基于 Go 语言开发的

身份识别与访问管理系统，

用于对资源访问进行授权。

**最新稳定版本为：v1.0.8，建议基于稳定版安装测试**。


IAM 同时也具有以下能力：

1. 配合极客时间专栏 讲解如何用 Go 做企业级应用的开发，是该项目的理论课程，包含了项目各个知识点和构建思路的讲解，也会包含我的一线研发经验和建议。



2. 作为一个开发脚手架，供开发者克隆后二次开发，快速构建自己的应用。

IAM 项目会长期维护、定期更新，**欢迎兄弟们 Star & Contributing**


## 软件架构

![IAM架构](./docs/images/IAM架构.png)


## 快速开始

### 依赖检查

**Minimum Requirements**

- Hardware
  - 2 GB of Memory
  - 50 GB of Disk Space
- 操作系统：CentOS Linux 8.2 (64-bit)
- 正常访问外网

 **需求检查 & 依赖安装** 

 请参考：[](docs/guide/zh-CN/installation/installation-requirement.md)

### 构建

1. 代码包下载

```
$ git clone https://github.com/marmotedu/iam
```

2. 编译

```bash
$ cd iam
$ make
```

### 运行

```bash
./scripts/install/install.sh iam::install::install_iam    
```

## 使用指南

[IAM Documentation](docs/guide/zh-CN)


## 社区

You are encouraged to communicate most things via [GitHub issues](https://github.com/marmotedu/iam/issues/new/choose) or pull requests.


## 谁在用

如果你有项目在使用iam系统模板，也欢迎联系作者，加入使用案例。



04 规范设计（