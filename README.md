<div align="center">
<h1>gitlabCodeSearch</h1>

[![Auth](https://img.shields.io/badge/Auth-eryajf-ff69b4)](https://github.com/eryajf)
[![GitHub Pull Requests](https://img.shields.io/github/stars/eryajf/gitlabCodeSearch)](https://github.com/eryajf/gitlabCodeSearch/stargazers)
[![HitCount](https://views.whatilearened.today/views/github/eryajf/gitlabCodeSearch.svg)](https://github.com/eryajf/gitlabCodeSearch)
[![GitHub license](https://img.shields.io/github/license/eryajf/gitlabCodeSearch)](https://github.com/eryajf/gitlabCodeSearch/blob/main/LICENSE)
[![](https://img.shields.io/badge/Awesome-MyStarList-c780fa?logo=Awesome-Lists)](https://github.com/eryajf/awesome-stars-eryajf#readme)

<p> 🌉 🔎 使用关键字搜索所有gitlab项目,并将结果输出为Excel 🌉</p>

<img src="https://camo.githubusercontent.com/82291b0fe831bfc6781e07fc5090cbd0a8b912bb8b8d4fec0696c881834f81ac/68747470733a2f2f70726f626f742e6d656469612f394575424971676170492e676966" width="800"  height="3">

</div>

日常工作中，我们在做一些变更的时候，经常会遇到对 gitlab 所有项目检索某个关键字的需求，这个工具，正是为解决这一需求而生的。

## 如何使用

先创建 gitlab 的 token，这一步就不介绍如何创建了。注意此 token 需要给：api 和 read_api 两个权限。

然后你可以通过配置文件，或者环境变量的方式将配置加载到程序当中。

通过环境变量能够快速运行项目，因此这里强烈推荐你使用此方式：

如果你熟悉 go 语言，则可以自己编译二进制，如果不熟悉，也可以直接在 releases 当中下载已经编译好的二进制。

```sh
gcs search -u 'https://gitlab.xxx.com' -t 'xxxxxxxxxxx' -w '测试'
```

效果如下：

![](https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20230902_000537.png)

运行成功以后，会自动在项目运行目录下将结果输出为 Excel。

你也可以直接通过 docker 来运行：

```sh
docker run -it registry.cn-hangzhou.aliyuncs.com/ali_eryajf/gcs gcs search -u 'https://gitlab.xxx.com' -t 'xxxxxxxxxxx' -w '测试'
```

如果你在使用过程中遇到问题，欢迎提交 issue。如果你还有其他需求，以及其他想法，欢迎提交 PR。

## 感谢开源

此框架建立在如下几个优秀的开源项目之上：

- [gopkg.in/yaml.v3 v3.0.1](https://github.com/go-yaml/yaml)
- [github.com/spf13/cobra v1.2.1](https://github.com/spf13/cobra)
