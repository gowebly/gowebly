# gowebly 新一代 CLI 工具，可使用 htmx 和 hyperscript 在 Go 中构建令人惊叹的网络应用程序

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][repo_url]
[![License][repo_license_img]][repo_license_url]

[English][repo_url] | [Русский][repo_readme_ru_url] | **中文** | 
[Español][repo_readme_es_url]

这款 CLI 工具可在后端使用 **Go**，在前端使用 [**htmx**][htmx_url] 和 
[**hyperscript**][hyperscript_url] 以及最流行的原子/实用工具优先 **CSS 
frameworks**，轻松构建令人惊叹的网络应用程序。

功能：

- 根据[Apache 2.0][repo_license_url]许可，100%免费开源。
- 适用于任何知识水平和技术专长的开发人员。
- 有助于更快地进入 Go + htmx + hyperscript 技术栈。
- 智能 CLI 可完成大部分常规设置和生产准备工作。
- 为项目添加原子/实用工具优先 CSS 框架的可能性；
- 包含开箱即用的综合示例。
- 文档详实，作者提供了大量提示和帮助。

## ⚡️ 快速启动

首先，[下载][go_download_url] 并安装 Go。需要安装 `1.21`（或更高版本）。

现在，你可以在不安装的情况下使用 `gowebly`。只需 [`go run`][go_run_url] 它，并在选项中创建一个新项目：

```console
go run github.com/gowebly/gowebly@latest create built-in
```

就是这样！🔥 您的精彩网络应用程序（本例中使用内置的 `net/http` 软件包）与 **htmx** 和 **hyperscript** 可在 Go 
HTML 模板中使用。

### 🔹 快速启动的完整 Go-way

如果您仍想通过 Golang 在系统中安装 `gowebly` CLI，请使用 [`go install`][go_install_url]命令：

```console
go install github.com/gowebly/gowebly@latest
```

### 🍺 Homebrew 快速启动方式

GNU/Linux 和苹果 macOS 用户可通过 [Homebrew][brew_url]安装 `gowebly` CLI。

点击新公式：

```console
brew tap gowebly/tap
```

安装 `gowebly`：

```console
brew install gowebly/tap/gowebly
```

### 🐳 Docker 快速启动方式

请随意使用我们[官方 Docker 镜像][docker_image_url]中的 `gowebly` CLI，并在隔离的容器中运行它：

```console
docker run \
    --rm -it -v ${PWD}:${PWD} -w ${PWD} \
    gowebly/gowebly:latest [COMMAND] [OPTIONS]
```

### 📦 其他快速启动方法§

在 [Releases][repo_releases_url] 页面下载现成的 Windows `exe` 文件、`deb`、`rpm`、`apk` 或 
Arch Linux 软件包。

## 📖 完整的用户指南

为了获得完整的使用指南并理解 `gowebly` CLI 的基本原理，我们在 README 文件中对每条命令都做了全面的解释。

> ⚡️ 作者的话 我们一直非常珍惜您的时间，并希望您能尽快开始在这个超棒的技术栈上构建真正出色的网络产品！

我们希望您能找到所有问题的答案！👌 但是，如果您没有找到所需的信息，请随时创建一个 [issue][repo_issues_url] 或发送一个 
[PR][repo_pull_request_url] 到此版本库。

### `create`

命令使用给定的 Go 后端、htmx 和 hyperscript 以及（可选）最流行的原子/实用工具优先 CSS 框架创建新项目。

```console
gowebly create [BACKEND]
```

您可以为项目选择 Go 后台：

| 后台         | 说明                                           |
|------------|----------------------------------------------|
| `built-in` | 创建一个新项目，Go 后台内置 [net/http][net_http_url] 软件包 |
| `fiber`    | 使用[Fiber][fiber_url] 网络框架创建一个带有 Go 后端的新项目    |
| `echo`     | 使用[Echo][echo_url] 网络框架创建一个带有 Go 后端的新项目      |
| `chi`      | 使用 Go 后端和 [chi][chi_url] 可组合路由器创建一个新项目       |

### `add`

命令将最流行的原子/实用工具优先 CSS 框架之一添加到你的项目中。这是可选项，不是必需项。

```console
gowebly add [CSS_FRAMEWORK]
```

> 💡 注意： `gowebly` CLI 会搜索当前文件夹中项目的 YAML 配置文件（`.gowebly.yml`）。

您可以选择 CSS 框架：

| CSS 框架        | 说明                                       |
|---------------|------------------------------------------|
| `tailwindcss` | 在项目前端添加 [Tailwind CSS][tailwindcss_url]。 |
| `unocss`      | 在项目前端添加 [UnoCSS][unocss_url]。            |

### `run`

命令在开发（非生产）模式下运行项目。

```console
gowebly run
```

> 💡 注意： `gowebly` CLI 会搜索当前文件夹中项目的 YAML 配置文件（`.gowebly.yml`）。

Go HTML 模板中将提供以下库版本：

- htmx：来自 CDN 的最新非生产版本；
- hyperscript: CDN 提供的最新非生产版本；
- (可选）CSS 框架：CDN 提供的最新非生产版本；

在开发模式下，只有开发人员提供的官方支持 CDN
将被使用：

- [unpkg.com][unpkg_url] 用于 htmx 和 hyperscript；
- [tailwindcss.com][tailwindcss_cdn_url] 用于 Tailwind CSS；
- [jsDelivr][jsdelivr_url] 用于 UnoCSS。

每次为您的项目执行 `run` 命令时：

1. 通过 CLI 将 htmx 和 hyperscript 的 CDN 版本嵌入到 Go HTML 模板中的常规 `<script>` 
标记中，并将其放入名为 `gowebly-body-scripts` 的块中（通常放在 `<body>` 标记的底部）；
2. (可选）通过 CLI 将所选 CSS 框架的 CDN 版本嵌入到 Go HTML 模板中，在名为 `gowebly-head-styles` 的块中使用常规的 `<link>` 标记（通常位于 `<head>` 标记的底部）；
3. CLI 通过简单的 `go run` 命令启动项目的后端（端口为 `5000`）。

### `build`

命令来构建用于生产的项目，并为部署容器做好准备。

```console
gowebly build
```

> 💡 注意： `gowebly` CLI 会搜索当前文件夹中项目的 YAML 配置文件（`.gowebly.yml`）。

Go HTML 模板中将提供以下库版本：

- htmx：在配置文件中选择的最小化生产版本；
- hyperscript: 已精简的生产版本，在配置文件中选择；
- (可选）CSS 框架：最新生产的树状结构和 最小化版本；

每次为项目执行 `build` 命令时：

1. CLI 扫描并验证 YAML 配置文件 (`.gowebly.yml`)，将所有设置应用于当前项目；
2. CLI 从官方（可信的）资源下载经过精简的 htmx 和 hyperscript 版本；
   - 将它们嵌入 Go HTML 模板（内联样式）中名为 "gowebly-body-scripts "的块（通常位于"<body>"标记的底部）；
3. (可选）通过[Vite][vite_url]工具，用 CLI 编写所选 CSS 框架的精简（和树化）版本；
   - 将它们嵌入 Go HTML 模板（内联样式）中名为 "gowebly-head-styles "的块（通常位于"<head>"标记的底部）；
4. CLI 会在项目文件夹根目录下生成一个清晰且文档齐全的 `docker-compose.yml` 文件，以便通过
   [Portainer][portainer_url]（推荐）将其部署到隔离的 Docker 容器中，或手动部署到远程服务器上。 

## 🎯 创作动机

请告诉我们，您有多少次不得不从头开始一个新项目，并进行痛苦的手动配置？进行痛苦的手动配置？🤔 
特别是当您刚刚开始熟悉新技术或堆栈时。尤其是当你刚刚熟悉一种新技术或堆栈时，一切对你来说都是新的。

对于包括我们在内的许多开发人员来说，这个过程尽可能乏味甚至令人沮丧，而且没有任何有用的工作量。这是一个非常令人沮丧的过程，会让任何开发人员远离技术。

为什么不把这些可怕的手工工作交给机器呢？让它们为我们完成所有艰苦的工作，我们只需创建出色的网络产品，而不必考虑构建和部署。

这就是我们创建 `gowebly` CLI 的原因，它可以帮助你使用 htmx、hyperscript 和流行的原子/实用工具优先 CSS 框架在 Go 
中创建出色的网络应用程序。

我们就是要把你（和我们自己）从这种日常的痛苦中解救出来！✨

## 🏆 双赢合作

现在，我邀请您参与这个项目！让我们共同努力，为开发人员打造当今网络上最实用的工具。

- [Issues][repo_issues_url]: 提出问题并提交您的功能。
- [Pull requests][repo_pull_request_url]: 提交您对当前版本的改进。

欢迎您提交 PR 和问题！谢谢 😘

## ⚠️ 许可证

[`gowebly`][repo_url]是根据[Apache 2.0 License][repo_license_url]授权的免费开源软件, 由
[Vic Shóstak][author_url]创建和支持, 其中🩵代表人和机器人。

<!-- Go links -->

[go_download_url]: https://golang.org/dl/
[go_run_url]: https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program
[go_install_url]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_report_url]: https://goreportcard.com/report/github.com/gowebly/gowebly
[go_dev_url]: https://pkg.go.dev/github.com/gowebly/gowebly
[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[go_code_coverage_img]: https://img.shields.io/badge/code_coverage-0%25-success?style=for-the-badge&logo=none
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none

<!-- Repository links -->

[repo_url]: https://github.com/gowebly/gowebly
[repo_issues_url]: https://github.com/gowebly/gowebly/issues
[repo_pull_request_url]: https://github.com/gowebly/gowebly/pulls
[repo_releases_url]: https://github.com/gowebly/gowebly/releases
[repo_license_url]: https://github.com/gowebly/gowebly/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_readme_ru_url]: https://github.com/gowebly/gowebly/blob/main/README_RU.md
[repo_readme_cn_url]: https://github.com/gowebly/gowebly/blob/main/README_CN.md
[repo_readme_es_url]: https://github.com/gowebly/gowebly/blob/main/README_ES.md

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- Readme links -->

[docker_image_url]: https://hub.docker.com/repository/docker/gowebly/gowebly
[portainer_url]: https://docs.portainer.io
[brew_url]: https://brew.sh
[vite_url]: https://vitejs.dev
[htmx_url]: https://htmx.org
[hyperscript_url]: https://hyperscript.org
[tailwindcss_url]: https://tailwindcss.com
[tailwindcss_cdn_url]: https://tailwindcss.com/docs/installation/play-cdn
[unocss_url]: https://unocss.dev
[unpkg_url]: https://unpkg.com
[jsdelivr_url]: https://www.jsdelivr.com
[net_http_url]: https://pkg.go.dev/net/http
[fiber_url]: https://github.com/gofiber/fiber
[echo_url]: https://github.com/labstack/echo
[chi_url]: https://github.com/go-chi/chi
