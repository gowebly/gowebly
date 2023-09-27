# gowebly – A next-generation CLI tool for building amazing web apps in Go using htmx & hyperscript

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][repo_url]
[![License][repo_license_img]][repo_license_url]

**English** | [Русский][repo_readme_ru_url] | [中文][repo_readme_cn_url] | 
[Español][repo_readme_es_url]

This CLI tool can easily build amazing web applications with **Go** 
on the backend, using [**htmx**][htmx_url] & [**hyperscript**][hyperscript_url] 
and the most popular atomic/utility-first **CSS frameworks** on the frontend.

Features:

- 100% **free** and **open source** under the [Apache 2.0][repo_license_url] 
  license.
- For **any** level of developer's knowledge and technical expertise.
- **Well-documented**, with a lot of tips and assists from the authors;
- Smart CLI that **does most** of the routine setup and preparation for 
  production.
- Helps to get into the **Go** + **htmx** + **hyperscript** technology 
  stack faster.
- The possibility of simply adding a ready-to-use and completely customized 
  atomic/utility-first **CSS framework** to your project.
- Has a library of **user-friendly** helpers for your Go code;
- Contains a comprehensive **example** of how to use it out of the box.

## ⚡️ Quick start

First, [download][go_download_url] and install **Go**. Version `1.21` (or 
higher) is required.

Now, you can use `gowebly` without installation. Just [`go run`][go_run_url] 
it to create a new project with a default [config][repo_default_config] file:

```console
go run github.com/gowebly/gowebly@latest create
```

That's it! 🔥 A wonderful web application, using the built-in **net/http** 
package (as a Go backend), **htmx** & **hyperscript** is available in your Go 
HTML templates.

### 🔹 A full Go-way to quick start

If you still want to install `gowebly` CLI to your system by Golang, use the 
[`go install`][go_install_url] command:

```console
go install github.com/gowebly/gowebly@latest
```

### 🍺 Homebrew-way to quick start

GNU/Linux and Apple macOS users available way to install `gowebly` CLI via 
[Homebrew][brew_url].

Tap a new formula:

```console
brew tap gowebly/tap
```

Install `gowebly`:

```console
brew install gowebly/tap/gowebly
```

### 🐳 Docker-way to quick start

Feel free to using `gowebly` CLI from our 
[official Docker image][docker_image_url] and run it in the isolated container:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} gowebly/gowebly:latest create
```

### 📦 Other way to quick start

Download a ready-made Windows `exe` files, `deb`, `rpm`, `apk`, or Arch Linux 
packages on the [Releases][repo_releases_url] page.

## 📖 Complete user guide

To get a complete guide to use and understand the basic principles of the
`gowebly` CLI, we have prepared a comprehensive explanation of each command at
once in this README file.

> ⚡️ From the authors: We always treasure your time and want you to start 
> building really great web products on this awesome technology stack as 
> soon as possible!

We hope you find answers to all of your questions! 👌 But, if you do not find 
needed information, feel free to create an [issue][repo_issues_url] or send a 
[PR][repo_pull_request_url] to this repository.

Don't forget to switch this page for your language (current is
**English**): [Русский][repo_readme_ru_url], [中文][repo_readme_cn_url],
[Español][repo_readme_es_url].

### `init`

Command to create a **default** config file 
([`.gowebly.yml`][repo_default_config]) in the current folder.

```console
gowebly init
```

> 💡 Note: Of course, you can skip this step if you're comfortable with the
> following default configuration for your new project:
>
> - Go module (`go.mod`) and `package.json` names are set to **project**;
> - Without any Go framework for the backend part (only built-in
> **net/http** package);
> - Without any CSS framework for the frontend part (only default styles for
> the code example);
> - Server port is `5000`, timeout (in seconds): `5` for read, `10` for write;
> - Latest versions of the **htmx** & **hyperscript**.

<img width="720" alt="gowebly init" src="https://github.com/gowebly/gowebly/assets/11155743/679dd0e1-ecd6-4cfb-b145-c9f551ab2d9c">

Typically, a created config file contains the following options:

```yaml
backend:
   module_name: project # (string) option can be any name of your Go module (for example, 'github.com/user/project')
   go_framework: default # (string) option can be one of the values: 'fiber', 'echo', 'chi', or 'default'
   port: 5000 # (int) option can be any port that is not taken up on your system
   timeout:
      read: 5 # (int) option can be any number of seconds, 5 is recommended
      write: 10 # (int) option can be any number of seconds, 10 is recommended

frontend:
   package_name: project # (string) option can be any name of your package.json (for example, 'project')
   css_framework: default # (string) option can be one of the values: 'tailwindcss', 'unocss', or 'default'
   htmx: latest # (string) option can be any existing version
   hyperscript: latest # (string) option can be any existing version
```

But, you can choose any **Go framework** for your project's backend:

| Go framework | Description                                                                 |
|--------------|-----------------------------------------------------------------------------|
| `default`    | Don't use any Go framework (only built-in [net/http][net_http_url] package) |
| `fiber`      | Use a Go backend with the [Fiber][fiber_url] web framework                  |
| `echo`       | Use a Go backend with the [Echo][echo_url] web framework                    |
| `chi`        | Use a Go backend with the [chi][chi_url] composable router                  |

In additional, you can choose versions of the **htmx**, **hyperscript**, and 
one of the most popular atomic/utility-first **CSS framework** to your 
project:

| CSS framework | Description                                                            |
|---------------|------------------------------------------------------------------------|
| `default`     | Don't use any CSS framework (only default styles for the code example) |
| `tailwindcss` | Use the [Tailwind CSS][tailwindcss_url] as a CSS framework             |
| `unocss`      | Use the [UnoCSS][unocss_url] as a CSS framework                        |

### `create`

Command to create a new project with the **Go** backend, **htmx** & 
**hyperscript**, and (_optionally_) atomic/utility-first **CSS framework**.

```console
gowebly create
```

> 💡 Note: If you don't run `init` command to create a config file 
> (`.gowebly.yml`), the `gowebly` CLI creates a new project with a 
> [default][repo_default_config] configuration.

<img width="720" alt="gowebly create" src="https://github.com/gowebly/gowebly/assets/11155743/35b15677-4991-406d-b666-dfbc40beb1ce">

Every time you make `create` command for your project:

1. CLI validates the config and applies all settings to the current project;
2. CLI prepares the backend part of your project (generates the project 
   structure and needed utility files, runs `go mod tidy`);
3. CLI prepares the frontend part of your project (generates the needed utility 
   files, runs `npm install` and `npm run build:dev` for the first time);
4. CLI downloads minimized versions of **htmx** and **hyperscript** (from
   official and trusted [unpkg.com][unpkg_url] CDN) to the `./static` folder
   and places them as separated `<script>` tags (at the bottom of the
   `<body>` tag) in the Go HTML template 
   [`templates/main.html`][repo_main_layout].

Typically, a created project contains the following files and folders:

```console
.
├── assets
│   └── styles.css
├── static
│   ├── favicon.ico
│   ├── htmx.min.js
│   ├── hyperscript.min.js
│   └── styles.css
├── templates
│   ├── pages
│   │   └── index.html
│   └── main.html
├── .gitignore
├── go.mod
├── go.sum
├── handlers.go
├── main.go
├── package-lock.json
├── package.json
└── server.go
```

### `run`

Command to run your project in a **development** (non-production) mode.

```console
gowebly run
```

> 💡 Note: If you don't run `init` command to create a config file
> (`.gowebly.yml`), the `gowebly` CLI runs your project with a
> [default][repo_default_config] configuration.

<img width="720" alt="gowebly run" src="https://github.com/gowebly/gowebly/assets/11155743/51c05652-4601-4f8b-8722-20401d0099d1">

Every time you make `run` command for your project:

1. CLI validates the config and applies all settings to the current project;
2. CLI prepares the frontend part of your project (runs `npm run build:dev`);
3. CLI prepares a development (non-production) version of the selected **CSS 
   framework** to the `./static` folder and places it as a `<link>` tag (at 
   the bottom of the `<head>` tag) in the Go HTML template
   [`templates/main.html`][repo_main_layout];
4. CLI starts a project's backend with settings from the default 
   configuration (or from the `.gowebly.yml` config file) by a simple `go run` 
   command.

### `build`

Command to build your project for **production** and prepare Docker files for 
deploy.

```console
gowebly build [OPTION]
```

> 💡 Note: If you don't run `init` command to create a config file
> (`.gowebly.yml`), the `gowebly` CLI builds your project with a
> [default][repo_default_config] configuration.

<img width="720" alt="gowebly build" src="https://github.com/gowebly/gowebly/assets/11155743/ac35b01f-0596-4d33-832e-1618709497d3">

You might add the following options:

| Option name     | Description                                                                    | Required? |
|-----------------|--------------------------------------------------------------------------------|-----------|
| `--skip-docker` | Skip generation process for the Docker files (it's helpful if you've your own) | no        |

Every time you make `build` command for your project:

1. CLI validates the config and apply all settings to the current project;
2. CLI downloads minimized versions of **htmx** and **hyperscript** (from
   official and trusted [unpkg.com][unpkg_url] CDN) to the `./static` folder
   and places them as separated `<script>` tags (at the bottom of the
   `<body>` tag) in the Go HTML template
   [`templates/main.html`][repo_main_layout];
3. CLI prepares a production version of the selected **CSS framework** and 
   places it as a `<link>` tag (at the bottom of the `<head>` tag) in the Go 
   HTML template [`templates/main.html`][repo_main_layout];
4. If the `--skip-docker` option is not set, CLI generate a clear and 
   well-documented Docker files (`.dockerignore`, `Dockerfile`, 
   `docker-compose.yml`) in the root of the project folder to deploy it in 
   isolated Docker containers via [Portainer][portainer_url] (_recommended_),
   or manually, to your remote server.

## 🙋 User-friendly helpers

The `gowebly` CLI has a library of user-friendly [helpers][gowebly_helpers_url] 
for your code. This will help you start building beautiful web applications 
in Go even faster.

```console
go get -u github.com/gowebly/helpers
```

> 💡 Note: The `gowebly helpers` library is **already** included in the Go 
> backend that is created by the `create` command, but you can use these 
> helpers in other projects as well.

## 🎯 Motivation to create

Tell us, how often have you had to start a new project from scratch and had
to make painful manual configurations? 🤔 Especially, when you are just getting
acquainted with a new technology or stack, where everything is new to you.

For many developers, _including us_, this process is as tedious and even
depressing as possible, and doesn't carry any useful workload. It is a **very**
frustrating process that can push any developer away from technology a lot.

Why not just give all that awful manual work to machines? Let them do all
the hard work for us, and we will just create awesome web products and not
have to think about build and deploy.

That's why we created the `gowebly` CLI and its helpers library, which helps 
you start an amazing web applications in **Go** using **htmx**, 
**hyperscript** and popular atomic/utility-first **CSS frameworks**.

We are here to save you (_and ourselves_) from this routine pain! ✨

## 🏆 A win-win cooperation

And now, I invite you to participate in this project! Let's work **together** to
create the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.

Your PRs & issues are welcome! Thank you 😘

## ⚠️ License

[`gowebly`][repo_url] is free and open-source software licensed 
under the [Apache 2.0 License][repo_license_url], created and supported by 
[Vic Shóstak][author_url] with 🩵 for people and robots.

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
[repo_default_config]: https://github.com/gowebly/gowebly/blob/main/internal/attachments/configs/default.yml
[repo_main_layout]: https://github.com/gowebly/gowebly/blob/main/internal/attachments/templates/frontend/main.html

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- Readme links -->

[gowebly_helpers_url]: https://github.com/gowebly/helpers
[docker_image_url]: https://hub.docker.com/repository/docker/gowebly/gowebly
[portainer_url]: https://docs.portainer.io
[brew_url]: https://brew.sh
[htmx_url]: https://htmx.org
[hyperscript_url]: https://hyperscript.org
[tailwindcss_url]: https://tailwindcss.com
[unocss_url]: https://unocss.dev
[unpkg_url]: https://unpkg.com
[net_http_url]: https://pkg.go.dev/net/http
[fiber_url]: https://github.com/gofiber/fiber
[echo_url]: https://github.com/labstack/echo
[chi_url]: https://github.com/go-chi/chi
