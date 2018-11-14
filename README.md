# 开发 web 服务程序

## 1、概述
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

### 任务目标

 1. 熟悉 go 服务器工作原理 
 2. 基于现有 web 库，编写一个简单 web 应用类似 cloudgo。 
 3. 使用 curl 工具访问 web
 4. 程序 对 web 执行压力测试

相关知识
课件：http://blog.csdn.net/pmlpml/article/details/78404838

## 2、任务要求
### 基本要求

 1. 编程 web 服务程序 类似 cloudgo 应用。 
   * 要求有详细的注释 
   * 是否使用框架、选哪个框架自己决定 请在 README.md说明你决策的依据
 4. 使用 curl 测试，将测试结果写入 README.md 
 5. 使用 ab 测试，将测试结果写入README.md。并解释重要参数。

### 扩展要求

选择以下一个或多个任务，以博客的形式提交。

 1. 选择 net/http 源码，通过源码分析、解释一些关键功能实现
 2. 选择简单的库，如 mux 等，通过源码分析、解释它是如何实现扩展的原理，包括一些 golang 程序设计技巧。 
 3. 在 docker hub 申请账号，从 github 构建 cloudgo 的 docker 镜像，最后在 Amazon 云容器服务中部署。 
  * 实现 Github - Travis CI - Docker hub - Amazon “不落地”云软件开发流水线 其他 web 开发话题

## 3. 实现说明

### （1) Martini框架
要开发一个Web应用与服务，有许多的开发框架可以选择，根据程序的实际应用情况可以按照以下分类作为选择参考：

 - 简单应用：应选择自带库 net/http 
 - 一般 web 应用与服务开发：建议选择轻量组件 gorilla/mux + codegangsta/negroni + …
 - web 开发: beego、Martini、revel ……

这次实验使用了Martini框架，感受一下成熟的web开发框架魅力。Martini框架是使用Go语言作为开发语言的一个强力快速构建模块化web应用与服务的开发框架。Martini是一个专门用来处理Web相关内容的框架，它并无自带有关ORM或详细的分层内容。  

**安装：**  
首先在gowork/src/github.com目录下新建文件夹go-martini，并进入到该目录下，下载Martini包：`git clone https://github.com/go-martini/martini.git`  
然后在gowork/src/github.com目录下新建文件夹codegangsta，同样进入到该目录下，下载inject包（安装go-martini要用到）:`git clone https://github.com/codegangsta/inject.git`  
最后，返回到gowork/src/github.com/go-martini/martini目录下，输入`go install`即可。  

### （2）开发 cloudgo 应用
程序很简单，借助Martini框架很快就能实现一个web服务程序。

1. 创建一个martini实例；
2. 使用该实例接收对\的GET方法请求；
3. 运行server

具体可看代码。  

## 4. 实现结果  
1. 运行web服务器  
写好代码之后，在CloudGo目录下，输入`go install`。然后在gowork/bin目录下输入'CloudGo'，启动服务器，终端显示在8080端口监听http链接请求。这时，打开浏览器，进入localhost:8080（默认8080端口启动服务），可以看到如下结果。  
浏览器进入localhost:8080后，终端显示接收到一个GET方法，并返回HTTP状态码200，表示OK，请求成功。  
![run_server](https://github.com/sysuxwh/CloudGo/blob/master/pic/runServer.png)  
手动输入参数，将服务端口改为9999，再次启动服务器，运行结果如下：  
![run_server2](https://github.com/sysuxwh/CloudGo/blob/master/pic/runServer_2.png)  
除了状态码200外，还收到了一个404状态码，表示找不到favicon.ico(网页图标/url图标)

2. curl工具测试  
curl是一个控制台程序，可以精确控制 HTTP 请求的每一个细节。配合 shell 程序可以简单、重复给服务器发送不同的请求序列，调试程序或分析输出。另外，curl 是 linux 系统自带的命令行工具。  
另外打开一个终端机，输入命令`curl -v localhost:8080`（服务器在8080端口监听http链接请求），会出现以下信息，显示成功访问localhost:8080。  
![curl](https://github.com/sysuxwh/CloudGo/blob/master/pic/curl.png)

3. ab工具压力测试  
首先安装 Apache web 压力测试程序，在终端中输入`yum -y install httpd-tools`即可。  
确保服务器在8080端口监听http链接请求，然后输入命令`ab -n 1000 -c 100 localhost:8080/`（不要漏了最后的/），测试结果如下图——所有请求均接收到了。  
![test](https://github.com/sysuxwh/CloudGo/blob/master/pic/ab.png)
