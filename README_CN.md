<a href="https://gowebly.org" target="_blank" title="Go to the Gowebly's Complete user guide"><img width="256px" alt="gowebly logo" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-logo.svg"></a>

# The Gowebly CLI – 新一代 CLI 工具，可使用 htmx 和 hyperscript 在 Go 中构建令人惊叹的网络应用程序

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

[English][repo_url] | [Русский][repo_readme_ru_url] | **简体中文** | [Español][repo_readme_es_url]

这款 CLI 工具可在后端使用 **Go**，在前端使用 [**htmx**][htmx_url] 和 [**hyperscript**][hyperscript_url] 以及最流行的原子/实用工具优先 **CSS 框架**，轻松构建令人惊叹的网络应用程序。

特点：

- 根据 [Apache 2.0][repo_license_url] 许可，100% 免费开源；
- 适合任何知识水平和技术专长的开发人员；
- [文档齐全][docs_url]，作者提供了大量提示和帮助；
- 跨平台和多体系结构，可在 GNU/Linux、MS Windows（包括 WSL）和 Apple macOS 上成功运行；
- 智能 CLI 可完成大部分常规设置和生产准备工作；
- 有助于更快进入 Go + htmx + hyperscript 技术栈；
- 开箱即支持最流行的 Go 网络框架，如 Fiber、Gin、Echo、Chi 和 HttpRouter；
- 支持使用带有热加载功能的 Templ 模板引擎开发网络应用程序的方式；
- 可为您的项目添加即用型和完全定制的原子/实用工具优先 CSS 框架；
- 可在浏览器或移动设备中安装为 PWA（渐进式网络应用程序）；
- 支持 CSS 样式的热加载模式；
- 为你的 Go 代码提供用户友好地帮助程序库；
- 包含如何使用该框架的综合示例。

> [!NOTE]
> 为了让您全面了解该项目，我们录制了一段简短的 📺 [视频][gowebly_youtube_vide_url]，并准备了一篇介绍 📝 [文章][gowebly_devto_article_url]，演示 Gowebly CLI 的主要功能。

## ⚡️ 快速启动

首先，[下载][go_download_url] 并安装 Go。需要安装 `1.21`（或更高版本）。

现在，无需安装即可使用 Gowebly CLI。只需使用 [`go run`][go_run_url] 运行它，即可创建一个新项目，并使用 [default][repo_default_config] 配置：

```console
go run github.com/gowebly/gowebly@latest create
```

就是这样！一个精彩的网络应用程序已经在当前文件夹中创建。

它将使用 net/http 包（作为 Go 后端）和 html/template 包（作为模板引擎）。HTML 模板中已经包含了 htmx 和 hyperscript 库。

### 🐳 通过 Docker 快速启动

请随意使用我们 [官方 Docker 镜像][docker_image_url] 中的 Gowebly CLI，并在隔离的容器中运行它：

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} gowebly/gowebly:latest create
```

> [!IMPORTANT]
> 该 Docker 镜像仅适用于 GNU/Linux 系统（`amd64` 或 `arm64`，包括 WSL）。

### 🍺 通过 Homebrew 快速启动

GNU/Linux 和苹果 macOS 用户可通过 [Homebrew][brew_url] 安装 Gowebly CLI。

点击新公式：

```console
brew tap gowebly/tap
```

安装：

```console
brew install gowebly/tap/gowebly
```

### 📦 其他快速入门方法

从 [Releases][repo_releases_url] 页面为 Windows、`deb`、`rpm`、`apk` 或 Arch Linux 软件包下载现成的 `exe` 文件。

## 📖 完整的用户指南

我们一直非常珍惜您的时间，并希望您能尽快开始在这个超棒的技术栈上构建真正出色的网络产品！因此，为了获得完整的使用指南并了解 Gowebly CLI 的基本原理，我们在此 📖 [完整用户指南][docs_url] 中准备了一份全面的项目说明。

<a href="https://gowebly.org" target="_blank" title="Go to the Gowebly's Complete user guide"><img width="360px" alt="gowebly docs banner" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-docs-banner.svg"></a>

为了让您尽可能轻松地学习这一出色的工具，我们对每条 CLI 命令都进行了充分的文字说明，并提供了工作原理的直观图。

> [!IMPORTANT]
> 别忘了将文档切换为你的语言，这样学习新知识会更轻松！支持的语言： [English][docs_url], [Русский][docs_ru_url], [简体中文][docs_zh_hk_url], [Español][docs_es_url].

### 学习路径

我们强烈建议您从介绍性短文 "[什么是 Gowebly CLI][docs_what_is_gowebly_cli_url]" 和 "[它是如何工作的？][docs_how_it_works_url]" 开始，了解 Gowebly CLI 的基本原理和主要组件。

接下来的步骤是

1. [将 CLI 安装到系统中](https://gowebly.org/zh_HK/complete-user-guide/installation)
2. [配置您的项目](https://gowebly.org/zh_HK/complete-user-guide/configuration)
3. [开始创建新项目](https://gowebly.org/zh_HK/complete-user-guide/create-new-project)
4. [本地运行项目](https://gowebly.org/zh_HK/complete-user-guide/run-your-project)
5. [为生产构建项目](https://gowebly.org/zh_HK/complete-user-guide/build-project)

希望你能找到所有问题的答案！ 😉

## 🎯 创建动机

请告诉我们，您有多少次不得不从零开始启动一个新项目，并进行痛苦的手动配置？🤔 尤其是当你刚刚熟悉一种新技术或堆栈时，一切对你来说都是新的。

对于包括我们在内的许多开发人员来说，这个过程尽可能乏味甚至令人沮丧，而且没有任何有用的工作量。这是一个非常令人沮丧的过程，会让任何开发人员远离技术。

为什么不把这些可怕的手工工作交给机器呢？让它们为我们完成所有艰巨的工作，我们只需创建出色的网络产品，而不必考虑构建和部署。

这就是我们创建 Gowebly CLI 及其助手库的原因，它可以帮助你使用 htmx、hyperscript 和流行的原子/实用工具优先 CSS 框架，用 Go 语言创建出色的网络应用程序。

我们就是要将你（和我们自己）从这种日常的痛苦中解救出来！ ✨

> [!NOTE]
> 此前，我们已经拯救过世界一次，那就是 [Create Go App][cgapp_url]（没错，那也是我们的项目）。这个项目的 [GitHub stars][cgapp_stars_url] 统计数字不会说谎：超过 2.2 千名不同级别、不同国家的开发者通过这个 CLI 工具启动了新项目。

## 🏆 双赢合作

如果您喜欢 Gowebly CLI 并发现它对您的工作非常有用，请点击 👁️ Watch 按钮以避免错过新版本通知，并给它一个 🌟 GitHub Star！

这将激励我们把产品做得更好。

<img width="100%" alt="gowebly star and watch" src="https://github.com/gowebly/gowebly/assets/11155743/6f92ec26-1fe3-44c6-9a13-3abd3ffa58eb">

现在，我邀请您参与这个项目！让我们共同努力，创建并推广当今网络上对开发人员最有用的工具。

- [Issues][repo_issues_url]：提出问题并提交您的功能。
- [Pull requests][repo_pull_request_url]：提交您对当前版本的改进。
- [Discussions][repo_disscussions_url]: 讨论并分享您的想法。
- 在你的社交网络和博客（Dev.to、Medium、Хабр 等）上说几句关于项目的话。

欢迎您提出意见、问题和建议！谢谢 😘

### 🌟 Stargazers

[![gowebly stargazers][repo_badge_stargazers_img]][repo_stargazers_url]

## ⚠️ License

[`The Gowebly CLI`][repo_url] is free and open-source software licensed under the [Apache 2.0 License][repo_license_url], created and supported by [Vic Shóstak][author_url] and the [True Web Artisans][true_web_artisans_url] team with 🩵 for people and robots. Official logo distributed under the [Creative Commons License][repo_cc_license_url] (CC BY-SA 4.0 International).

<!-- Go links -->

[go_download_url]: https://golang.org/dl/
[go_run_url]: https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program
[go_install_url]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_report_url]: https://goreportcard.com/report/github.com/gowebly/gowebly
[go_dev_url]: https://pkg.go.dev/github.com/gowebly/gowebly
[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[go_code_coverage_url]: https://codecov.io/gh/gowebly/gowebly
[go_code_coverage_img]: https://img.shields.io/codecov/c/gh/gowebly/gowebly.svg?logo=codecov&style=for-the-badge
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none

<!-- Repository links -->

[repo_url]: https://github.com/gowebly/gowebly
[repo_issues_url]: https://github.com/gowebly/gowebly/issues
[repo_pull_request_url]: https://github.com/gowebly/gowebly/pulls
[repo_releases_url]: https://github.com/gowebly/gowebly/releases
[repo_disscussions_url]: https://github.com/gowebly/gowebly/discussions
[repo_license_url]: https://github.com/gowebly/gowebly/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_license_url]: https://creativecommons.org/licenses/by-sa/4.0/
[repo_readme_ru_url]: https://github.com/gowebly/gowebly/blob/main/README_RU.md
[repo_readme_cn_url]: https://github.com/gowebly/gowebly/blob/main/README_CN.md
[repo_readme_es_url]: https://github.com/gowebly/gowebly/blob/main/README_ES.md
[repo_stargazers_url]: https://github.com/gowebly/gowebly/stargazers
[repo_badge_stargazers_img]: https://user-images.githubusercontent.com/11155743/275514241-8ecdf4bd-c35e-4e28-a937-b0a63aa1dbaf.png
[repo_default_config_url]: https://github.com/koddr/gowebly/blob/main/internal/attachments/configs/default.yml

<!-- Docs links -->

[docs_url]: https://gowebly.org
[docs_ru_url]: https://gowebly.org/ru/
[docs_zh_hk_url]: https://gowebly.org/zh_HK/
[docs_es_url]: https://gowebly.org/es/
[docs_what_is_gowebly_cli_url]: https://gowebly.org/zh_HK/getting-started
[docs_how_it_works_url]: https://gowebly.org/zh_HK/getting-started/how-does-it-work

<!-- Author links -->

[author_url]: https://github.com/koddr
[true_web_artisans_url]: https://github.com/truewebartisans

<!-- Readme links -->

[gowebly_helpers_url]: https://github.com/gowebly/helpers
[gowebly_youtube_video_url]: https://www.youtube.com/watch?v=qazYscnLku4
[gowebly_devto_article_url]: https://dev.to/koddr/a-next-generation-cli-tool-for-building-amazing-web-apps-in-go-using-htmx-hyperscript-336d
[cgapp_url]: https://github.com/create-go-app/cli
[cgapp_stars_url]: https://github.com/create-go-app/cli/stargazers
[htmx_url]: https://htmx.org
[hyperscript_url]: https://hyperscript.org
[brew_url]: https://brew.sh
[docker_image_url]: https://hub.docker.com/repository/docker/gowebly/gowebly