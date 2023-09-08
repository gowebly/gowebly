# gowebly – A next-generation CLI tool for building amazing web apps in Go using htmx & hyperscript

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][repo_url]
[![License][repo_license_img]][repo_license_url]

**English** | [Русский][repo_readme_ru_url] | [中文][repo_readme_cn_url] | 
[Español][repo_readme_es_url]

This CLI tool can easily build amazing web applications with **Go** 
on the backend using [**htmx**][htmx_url] & [**hyperscript**][hyperscript_url] 
and the most popular atomic/utility-first **CSS frameworks** on the frontend.

Features:

- 100% **free** and **open source** under the [Apache 2.0][repo_license_url] 
  license.
- For **any** level of developer's knowledge and technical expertise.
- Helps to get into the **Go** + **htmx** + **hyperscript** technology 
  stack faster.
- Smart CLI that **does most** of the routine setup and preparation for 
  production.
- Contains a comprehensive **example** of how to use it out of the box.
- **Well-documented**, with a lot of tips and assists from the authors.

## ⚡️ Quick start

First, [download][go_download_url] and install **Go**. Version `1.21` (or 
higher) is required.

Now, you can use `gowebly` without installation. Just [`go run`][go_run_url] it 
with options to create a new project:

```console
go run github.com/gowebly/gowebly@latest create built-in unocss
```

That's it! 🔥 Your wonderful web application (in this example using the 
built-in `net/http` package) with **htmx** & **hyperscript** and **CSS 
framework** features (in this example UnoCSS) is available in Go HTML templates.

Let's make useful web project and deploy them to the Internet 🚀

### 🔹 A full Go-way to quick start

If you still want to install `gowebly` CLI to your system by Golang, use the 
[`go install`][go_install_url] command:

```console
go install github.com/gowebly/gowebly@latest
```

### 🍺 Homebrew-way to quick start

GNU/Linux and Apple macOS users available way to install `gowebly` via 
[Homebrew][brew_url].

Tap a new formula:

```console
brew tap gowebly/tap
```

Install `gowebly` CLI:

```console
brew install gowebly/tap/gowebly
```

### 🐳 Docker-way to quick start

Feel free to using `gowebly` CLI from our 
[official Docker image][docker_image_url] and run it in the isolated container:

```console
docker run \
    --rm -it -v ${PWD}:${PWD} -w ${PWD} \
    gowebly/gowebly:latest [COMMAND] [OPTIONS]
```

### 📦 Other way to quick start

Download a ready-made Windows `exe` files, `deb`, `rpm`, `apk`, or Arch Linux 
packages on the [Releases][repo_releases_url] page.

## 📖 Complete user guide

To get a complete guide to use and understand the basic principles of the
`gowebly` CLI, we have prepared a comprehensive explanation of each command at
once in this README file.

Don't forget to go to the translation page for your language (current is 
**English**): [Русский][repo_readme_ru_url], [中文][repo_readme_cn_url],
[Español][repo_readme_es_url].

> ⚡️ From the authors: We always treasure your time and want you to start 
> building really great web products on this awesome technology stack as 
> soon as possible!

We hope you find answers to all your questions! 👌 But, if you do not find 
needed information, feel free to create an [issue][repo_issues_url] or send a 
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

- **htmx**: latest non-production version from CDN;
- **hyperscript**: latest non-production version from CDN;
- (_optionally_) **CSS framework**: latest non-production version from CDN;

In development mode, only official supported CDNs from developers 
will be used: 

- [unpkg.com][unpkg_url] for **htmx** and **hyperscript**;
- [tailwindcss.com][tailwindcss_cdn_url] for **Tailwind CSS**;
- [jsDelivr][jsdelivr_url] for **UnoCSS**.

Every time you make `run` command for your project:

1. CLI embed CDN versions of **htmx** and **hyperscript** from the
   official (and trusted) resources to your Go HTML templates in a regular 
   `<script>` tag into the block called `gowebly-body-scripts` (usually, 
   placed on the bottom of the `<body>` tag);
2. (_optionally_) CLI embed a CDN version of the chosen **CSS framework** 
   from the official (and trusted) resource to your Go HTML templates in a 
   regular `<link>` tag into the block called `gowebly-head-styles` (usually, 
   placed on the bottom of the `<head>` tag);
3. CLI start a project's backend (on the port `5000`) via simple `go run` 
   command.

### `build`

Command to build your project for **production** and prepare containers for 
deploy.

```console
gowebly build
```

> 💡 Note: the `gowebly` CLI search for YAML config file (`.gowebly.yml`) for
> the project in the current folder.

The following library versions will be supplied in Go HTML templates:

- **htmx**: minified production version, selected in the config file;
- **hyperscript**: minified production version, selected in the config file;
- (_optionally_) **CSS framework**: latest production tree-shaking & 
  minified version;

Every time you make `build` command for your project:

1. CLI scan and validate the YAML config file (`.gowebly.yml`), apply all 
   settings to the current project;
2. CLI download minified versions of **htmx** and **hyperscript** from the 
   official (and trusted) resources;
   - Embed them into your Go HTML templates (inline-style) to the block 
     called `gowebly-body-scripts` (usually, placed on the bottom of the 
     `<body>` tag);
3. (_optionally_) CLI prepare a minified (and tree-shaking) version of the 
   chosen **CSS framework** via [Vite][vite_url] tool;
   - Embed them into your Go HTML templates (inline-style) to the block 
     called `gowebly-head-styles` (usually, placed on the bottom of the 
     `<head>` tag);
4. CLI generate a clear and well-documented `docker-compose.yml` file in the 
   root of the project folder for deploy it in isolated Docker containers 
   via [Portainer][portainer_url] (_recommended_) or manually to your remote 
   server.

## ✨ Solving case

...

## 🎯 Motivation

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

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- Readme links -->

[repo_readme_en_url]: https://github.com/gowebly/gowebly/blob/main/README.md
[repo_readme_ru_url]: https://github.com/gowebly/gowebly/blob/main/README_RU.md
[repo_readme_cn_url]: https://github.com/gowebly/gowebly/blob/main/README_CN.md
[repo_readme_es_url]: https://github.com/gowebly/gowebly/blob/main/README_ES.md
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
