<a href="https://gowebly.org" target="_blank" title="Go to the Gowebly's Complete user guide"><img width="256px" alt="gowebly logo" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-logo.svg"></a>

# The Gowebly CLI – CLI-инструмент нового поколения для создания удивительных веб-приложений на языке Go с использованием htmx и hyperscript

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

[English][repo_url] | **Русский** | [简体中文][repo_readme_cn_url] | [Español][repo_readme_es_url]

С помощью этого CLI-инструмента можно легко создавать удивительные веб-приложения с **Go** на бэкенде, с использованием [**htmx**][htmx_url] и [**hyperscript**][hyperscript_url] и наиболее популярных атомарных/утилитарных **CSS-фреймворков** на фронтенде.

Особенности:

- 100% **свободный** и **открытый исходный код** под лицензией [Apache 2.0][repo_license_url];
- Для **любого** уровня знаний и технической экспертизы разработчика;
- [**Хорошо документирован**][docs_url], содержит большое количество советов и подсказок от авторов;
- Кроссплатформенность и мультиархитектурность позволяет **успешно запускаться** на GNU/Linux, MS Windows (включая WSL) и Apple macOS;
- Умный CLI, который **делает большую часть** рутинной настройки и подготовки к работе;
- Помогает быстрее освоить стек технологий **Go** + **htmx** + **hyperscript**;
- Поддержка наиболее популярных веб-фреймворков Go из коробки, таких как **Fiber**, **Gin**, **Echo**, **Chi** и **HttpRouter**;
- Поддержка способа разработки веб-приложений с использованием шаблонизатора **Templ** с функцией hot-reloading;
- Возможность простого добавления в проект готового к использованию и полностью настроенного атомарного/утилитарного **CSS-фреймворка**;
- Готовность к установке в качестве **PWA** (Progressive Web App) в браузере или на мобильном устройстве;
- Поддерживает режим **hot-reloading** для ваших CSS-стилей;
- Имеет библиотеку **удобных** помощников для вашего Go-кода;
- Содержит исчерпывающий **пример** использования из коробки.

> [!NOTE]
> Чтобы дать вам полное представление о проекте, мы записали короткое 📺 [видео][gowebly_youtube_video_url] и подготовили вводную 📝 [статью][gowebly_devto_article_url], демонстрирующую основные возможности **Gowebly** CLI.

## ⚡️ Быстрый старт

Сначала [скачайте][go_download_url] и установите **Go**. Требуется версия `1.21` (или выше).

Теперь вы можете использовать **Gowebly** CLI без установки. Просто запустите его с помощью [`go run`][go_run_url] для создания нового проекта с конфигурацией [по умолчанию][repo_default_config]:

```console
go run github.com/gowebly/gowebly@latest create
```

Вот и все! 🔥 В текущей папке создано замечательное веб-приложение.

Оно будет использовать пакет **net/http** (в качестве бэкенда Go) и пакет **html/template** (в качестве шаблонизатора). Библиотеки **htmx** и **hyperscript** уже доступны в ваших HTML-шаблонах.

### 🐳 Docker-путь к быстрому старту

Не стесняйтесь использовать **Gowebly** CLI из нашего [официального Docker-образа][docker_image_url] и запускать его в изолированном контейнере:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} gowebly/gowebly:latest create
```

> [!IMPORTANT]
> Данный Docker-образ работает **только** на системах GNU/Linux (`amd64` или `arm64`, включая WSL).

### 🍺 Homebrew-путь к быстрому старту

Пользователям GNU/Linux и Apple macOS доступен способ установки **Gowebly** CLI через [Homebrew][brew_url].

Тапните новую формулу:

```console
brew tap gowebly/tap
```

Установите:

```console
brew install gowebly/tap/gowebly
```

### 📦 Другой способ быстрого старта

Скачайте готовые `exe` файлы для Windows, а также пакеты `deb`, `rpm`, `apk` или Arch Linux со страницы [Releases][repo_releases_url].

## 📖 Полное руководство пользователя

Мы всегда ценим ваше время и хотим, чтобы вы как можно скорее начали создавать действительно отличные веб-продукты на этом потрясающем технологическом стеке! Поэтому, чтобы получить полное руководство по использованию и понять основные принципы работы **Gowebly** CLI, мы подготовили исчерпывающее объяснение проекта в этом 📖 [**Полном руководстве пользователя**][docs_url].

<a href="https://gowebly.org" target="_blank" title="Go to the Gowebly's Complete user guide"><img width="360px" alt="gowebly docs banner" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-docs-banner.svg"></a>

Мы постарались сделать так, чтобы вам было **как можно удобнее** осваивать этот замечательный инструмент, поэтому каждая команда CLI имеет достаточное текстовое описание, а также визуальную схему её работы.

### Путь обучения

Настоятельно рекомендуется начать изучение с краткой вводной статьи «[**How does it work?**][docs_how_it_works_url]», чтобы понять основной принцип работы и основные компоненты, встроенные в **Gowebly** CLI.

Следующие шаги:

1. [Установите CLI на свою систему][docs_installation_url].
2. [Сконфигурируйте проект][docs_configuring_url]
3. [Начните создание нового проекта][docs_create_new_project_url]
4. [Запустите проект локально][docs_run_project_url]
5. [Сборка проекта для продакшена][docs_build_project_url]

Надеемся, что вы найдете ответы на все свои вопросы! 😉

## 🎯 Мотивация к созданию

Скажите, как часто вам приходилось начинать новый проект с нуля и делать мучительные ручные настройки? 🤔 Особенно, когда вы только начинаете знакомиться с новой технологией или стеком, где все для вас в новинку.

Для многих разработчиков, _в том числе и для нас_, этот процесс максимально утомителен и даже уныл, и не несет никакой полезной нагрузки. Это **очень** удручающий процесс, который может сильно оттолкнуть любого разработчика от технологии.

Почему бы просто не отдать всю эту ужасную ручную работу машинам? Пусть они сделают всю тяжелую работу за нас, а мы будем просто создавать потрясающие веб-продукты и не думать о сборке и развертывании.

Именно поэтому мы создали **Gowebly** CLI и библиотеку хелперов, которая поможет вам создавать удивительные веб-приложения на **Go** с использованием **htmx**, **hyperscript** и популярных атомарных/утилитарных **CSS-фреймворков**.

Мы здесь, чтобы избавить вас (_и себя_) от этой рутинной боли! ✨

> [!NOTE]
> Ранее мы уже однажды спасли мир — это был [Create Go App][cgapp_url] (да, это тоже наш проект). Статистика [GitHub stars][cgapp_stars_url] этого проекта не может врать: более **2,2 тыс.** разработчиков любого уровня и разных стран начинают новый проект с помощью этого CLI-инструмента.

## 🏆 Беспроигрышное сотрудничество

Если вам понравился **Gowebly** CLI и вы нашли его полезным для решения своих задач, пожалуйста, нажмите на кнопку 👁️ **Watch**, чтобы не пропустить уведомления о выходе новых версий, и поставьте ему 🌟 **GitHub Star**!

Это очень **мотивирует** нас делать этот продукт **еще** лучше.

<img width="100%" alt="gowebly star and watch" src="https://github.com/gowebly/gowebly/assets/11155743/6f92ec26-1fe3-44c6-9a13-3abd3ffa58eb">

А теперь я приглашаю вас принять участие в этом проекте! Давайте работать **вместе** над созданием и популяризацией **самого полезного** инструмента для разработчиков в Интернете на сегодня.

- [Issues][repo_issues_url]: задавайте вопросы и предлагайте свои улучшения.
- [Pull requests][repo_pull_request_url]: присылайте свои улучшения.
- [Discussions][repo_disscussions_url]: обсудить и поделиться своими идеями.
- Скажите несколько слов о проекте в своих социальных сетях и блогах (Dev.to, Medium, Хабр и т.д.).

Ваши PR, проблемы и любые слова приветствуются! Спасибо 😘

### 🌟 Stargazers

[![gowebly stargazers][repo_badge_stargazers_img]][repo_stargazers_url]

## ⚠️ License

[`The Gowebly CLI`][repo_url] is free and open-source software licensed under the [Apache 2.0 License][repo_license_url], created and supported by [Vic Shóstak][author_url] with 🩵 for people and robots. Official logo distributed under the [Creative Commons License][repo_cc_license_url] (CC BY-SA 4.0 International).

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
[repo_disscussions_url]: https://github.com/gowebly/gowebly/discussions
[repo_releases_url]: https://github.com/gowebly/gowebly/releases
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
[docs_how_it_works_url]: https://gowebly.org/complete-user-guide/how-does-it-work/index.html
[docs_installation_url]: https://gowebly.org/complete-user-guide/installation/index.html
[docs_configuring_url]: https://gowebly.org/complete-user-guide/configuration/index.html
[docs_create_new_project_url]: https://gowebly.org/complete-user-guide/create-new-project/index.html
[docs_run_project_url]: https://gowebly.org/complete-user-guide/run-your-project/index.html
[docs_build_project_url]: https://gowebly.org/complete-user-guide/build-your-project/index.html

<!-- Author links -->

[author_url]: https://github.com/koddr

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