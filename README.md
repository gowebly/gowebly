# gowebly – A next-generation CLI tool for building amazing web applications in Go using htmx & hyperscript.

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][repo_url]
[![Wiki][repo_wiki_img]][repo_wiki_url]
[![License][repo_license_img]][repo_license_url]

**gowebly** description ...

Features:

- 100% **free** and **open source**.
- ...
- ...

## ⚡️ Quick start

First, [download][go_download_url] and install **Go**. Version `1.21` (or 
higher) is required.

Installation is done by using the [`go install`][go_install_url] command:

```console
go install github.com/gowebly/gowebly@latest
```

> 💡 Note: See the repository's [Release page][repo_releases_url], if you want
> to download ready-made MS Windows `exe` files, `deb`, `rpm`, `apk` or 
> `Arch Linux` packages.

Also, GNU/Linux and macOS users available way to install via 
[Homebrew][brew_url]:

```console
# Tap a new formula:
brew tap gowebly/tap

# Installation:
brew install gowebly/tap/gowebly
```

Next, run `gowebly` to create a new project:

```console
gowebly create built-in unocss
```

That's it! Your amazing Go web app with [htmx][htmx_url] & 
[hyperscript][hyperscript_url], and (_optionally_) **CSS framework** 
features are available in your Go HTML templates. Let's make useful web project 
and deploy them to the Internet 🚀

### 🐳 Docker-way to quick start

If you don't want to physically install `gowebly` to your system, you feel
free to using our [official Docker image][docker_image_url] and run it from
isolated container:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} gowebly/gowebly:latest [COMMANDS]
```

## 📖 Complete user guide

> ⚡️ From authors: We always treasure your time and want you to start building
> really great products on this great technology stack as soon as possible!

To get a complete guide to use and understand the basic principles of the 
`gowebly` CLI, we have prepared a comprehensive explanation of each command at 
once in this README file. 

We hope you find answers to all your questions! But, if you do not find needed 
information, feel free to create an [issue][repo_issues_url] or send a 
[PR][repo_pull_request_url] to this repository.

### `create`

Command to create a new project with the given Go backend, **htmx** & 
**hyperscript**, and (_optionally_) with the most popular 
atomic/utility-first **CSS framework**.

```console
gowebly create [BACKEND] [CSS_FRAMEWORK]
```

You can choose the **Go backend** for your project:

| Backend    | Description                                                 |
|------------|-------------------------------------------------------------|
| `built-in` | Go backend with a built-in [net/http][net_http_url] package |
| `fiber`    | Go backend with the [Fiber][fiber_url] web framework        |
| `echo`     | Go backend with the [Echo][echo_url] web framework          |
| `chi`      | Go backend with the [chi][chi_url]  composable router       |

_Optionally_, you can choose the **CSS framework**:

| CSS framework | Description                                         |
|---------------|-----------------------------------------------------|
| `tailwindcss` | Add [Tailwind CSS][tailwindcss_url] to your project |
| `unocss`      | Add [UnoCSS][unocss_url]  to your project           |

### `run`

Command to run your project in **development** mode.

```console
gowebly run
```

The following library versions will be supplied in Go HTML templates:

- **htmx**: latest version from CDN in a regular `<script>` tag;
- **hyperscript**: latest version from CDN in a regular `<script>` tag;
- (_optionally_) **CSS framework**: latest version from CDN in a regular 
  `<link>` tag;

In development mode, only official supported CDNs from developers 
will be used: 

- [unpkg.com][unpkg_url] for **htmx** and **hyperscript**;
- [tailwindcss.com][tailwindcss_cdn_url] for **Tailwind CSS**;
- [jsDelivr][jsdelivr_url] for **UnoCSS**.

### `build`

Command to build your project for **production** and prepare containers for 
deploy.

```console
gowebly build
```

The following library versions will be supplied in Go HTML templates:

- **htmx**: minified and embed to the `gowebly-body-scripts` block;
- **hyperscript**: minified and embed to the `gowebly-body-scripts` block 
  after **htmx** part;
- (_optionally_) **CSS framework**: tree-shaking & minified, and embed to 
  the `gowebly-head-styles` block;

> 💡 Note: the `gowebly` CLI search for YAML config file (`.gowebly.yml`) for 
> the project in the current folder.

Every time you run `build` command for your project:

1. CLI scan and validate the YAML config file (`.gowebly.yml`), apply all 
   settings to the current project;
2. Automatically download minified versions of **htmx** and **hyperscript** 
   from the official (and trusted) resources;
3. Embed them into your Go HTML templates (inline-style) to the block called 
   `gowebly-body-scripts` (usually, placed on the bottom of the `<body>` tag);
4. (_optionally_) CLI prepare a minified (and tree-shaking) version of the 
   chosen **CSS framework** via [Vite][vite_url] tool;
5. (_optionally_) Embed them into your Go HTML templates (inline-style) to 
   the block called `gowebly-head-styles` (usually, placed on the bottom of 
   the `<head>` tag);
6. CLI generate a clear and well-documented `docker-compose.yml` file in the 
   root of the project folder for deploy it in isolated Docker containers via 
   (_strongly recommended_) [Portainer][portainer_url] or manually to your 
   remote server.

## ✨ Solving case

...

## 💡 Motivation

...

## 🏆 A win-win cooperation

And now, I invite you to participate in this project! Let's work **together** to
create the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.

Your PRs & issues are welcome! Thank you 😘

## ⚠️ License

[`gowebly`][repo_url] is free and open-source software licensed 
under the [Apache 2.0 License][repo_license_url], created and supported with 🩵 
for people and robots by [Vic Shóstak][author_url].

<!-- Go links -->

[go_download_url]: https://golang.org/dl/
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
[repo_wiki_url]: https://github.com/gowebly/gowebly/wiki
[repo_license_url]: https://github.com/gowebly/gowebly/blob/main/LICENSE
[repo_wiki_img]: https://img.shields.io/badge/docs-wiki_page-blue?style=for-the-badge&logo=none
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- Readme links -->

[docker_image_url]: https://hub.docker.com/repository/docker/gowebly/gowebly
[portainer_url]: https://docs.portainer.io
[brew_url]: https://brew.sh
[unpkg_url]: https://unpkg.com
[vite_url]: https://vitejs.dev
[htmx_url]: https://htmx.org
[hyperscript_url]: https://hyperscript.org
[tailwindcss_url]: https://tailwindcss.com
[tailwindcss_cdn_url]: https://tailwindcss.com/docs/installation/play-cdn
[unocss_url]: https://unocss.dev
[jsdelivr_url]: https://www.jsdelivr.com
[net_http_url]: https://pkg.go.dev/net/http
[fiber_url]: https://gofiber.io
[echo_url]: https://echo.labstack.com
[chi_url]: https://go-chi.io
