<div align="center">
<h1>eryajfctl</h1>

[![Auth](https://img.shields.io/badge/Auth-eryajf-ff69b4)](https://github.com/eryajf)
[![GitHub Pull Requests](https://img.shields.io/github/stars/eryajf/eryajfctl)](https://github.com/eryajf/eryajfctl/stargazers)
[![HitCount](https://views.whatilearened.today/views/github/eryajf/eryajfctl.svg)](https://github.com/eryajf/eryajfctl)
[![GitHub license](https://img.shields.io/github/license/eryajf/eryajfctl)](https://github.com/eryajf/eryajfctl/blob/main/LICENSE)
[![](https://img.shields.io/badge/Awesome-MyStarList-c780fa?logo=Awesome-Lists)](https://github.com/eryajf/awesome-stars-eryajf#readme)

<p> 🌉 基于Cobra库快速开发类似kubectl一样的命令行工具框架 🌉</p>

<img src="https://camo.githubusercontent.com/82291b0fe831bfc6781e07fc5090cbd0a8b912bb8b8d4fec0696c881834f81ac/68747470733a2f2f70726f626f742e6d656469612f394575424971676170492e676966" width="800"  height="3">

</div>

运维也可以如此优雅！快用这个框架打造一个专属于你的工具箱吧！

通过这个框架，你可以快速上手，直接构建你想要的运维工具，而不必再考虑配置，框架设计等内容。

## 如何使用

先拷贝配置文件:

```sh
cp config.example.yml config.yml
```

执行如下指令，运行示例参数：

```sh
$ go run main.go ex getconfig -w "hello, This is eryajfctl"
通过配置文件获取到的用户名: eryajf
通过配置文件获取到的密码: 123456
通过命令行获取到的内容是: hello, This is eryajfctl
```

> 其中ex为一级参数，getconfig为二级参数，大多数场景下，分两个层级就够用了，你可以把一级参数当做归类，比如 jenkins, gitlab，二级参数当做功能参数，再往后的参数则是该二级参数所需要的运行时参数。

也可以编译成二进制，然后通过如下方式查看帮助信息：

```
# 编译
$ make build

#运行测试
$ ./eryajfctl ex getconfig -h
通过命令行获取配置信息

Usage:
  eryajfctl ex getconfig [flags]

Flags:
  -h, --help          help for getconfig
  -w, --word string   测试参数 (default "你好，这是测试")
```

## 开始开发

你可以直接参考ex参数的流程，开发新的参数，从而满足实际使用需求。

如果你的配置文件中有敏感数据，可以考虑结合go-bindata来使用，通过执行：

```
go-bindata -o=./public/bindata_config.go -pkg=public config.yml
```

然后更改 [public/config.go](https://github.com/eryajf/eryajfctl/blob/4cd30714062e5b65746bdb5f100f19bfe38ed52e/public/config.go#L28) 中的配置信息读取方式，接着进入开发即可。

如果后续本地的config.yml配置有更新，则再次执行上边的命令，将配置文件注入到 `bindata_config.go` 即可。

这样做的一个好处是，你的二进制放到服务器等地运行的时候，既不需要添加config.yml文件，也能防止文件中的敏感信息暴漏出去。

## 感谢开源

此框架建立在如下几个优秀的开源项目之上：

- [gopkg.in/yaml.v3 v3.0.1](https://github.com/go-yaml/yaml)
- [github.com/spf13/cobra v1.2.1](https://github.com/spf13/cobra)

## 其他参考

如果你想熟悉了解此框架的详细用法，还可以参考我的如下两篇文章：

- [使用go-bindata将文件编译进二进制](https://wiki.eryajf.net/pages/2bf6c3/)
- [利用cobra库快速开发类似kubectl一样的命令行工具](https://wiki.eryajf.net/pages/5c4163/)