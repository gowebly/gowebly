<a href="https://gowebly.org" target="_blank" title="Go to the Gowebly's Complete user guide"><img width="256px" alt="gowebly logo" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-logo.svg"></a>

# The Gowebly CLI – Una herramienta CLI de última generación para crear increíbles aplicaciones web en Go utilizando htmx e hyperscript.

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

[English][repo_url] | [Русский][repo_readme_ru_url] | [简体中文][repo_readme_cn_url] | **Español**

Esta herramienta CLI puede construir fácilmente increíbles aplicaciones web con **Go** en el backend, utilizando [**htmx**][htmx_url] & [**hyperscript**][hyperscript_url] y los más populares atomic/utility-first **CSS frameworks** en el frontend.

Características:

- 100% **gratuito** y **de código abierto** bajo la licencia [Apache 2.0][repo_license_url];
- Para **cualquier** nivel de conocimientos y experiencia técnica del desarrollador;
- [**Bien documentado**][docs_url], con muchos consejos y ayudas de los autores;
- Multiplataforma y multiarquitectura que permite **ejecutar con éxito** en GNU/Linux, MS Windows (incluyendo WSL) y Apple macOS;
- CLI inteligente que **hace la mayor parte** de la rutina de configuración y preparación para la producción;
- Ayuda a entrar más rápido en la pila tecnológica **Go** + **htmx** + **hyperscript**;
- Soporta los frameworks web Go más populares, como **Fiber**, **Gin**, **Echo**, **Chi** y **HttpRouter**;
- Soporta la forma en que se desarrollan las aplicaciones web utilizando el motor de plantillas **Templ** con recarga en caliente;
- La posibilidad de añadir simplemente a tu proyecto un **CSS framework** atómico/utility-first listo para usar y completamente personalizado;
- Listo para instalar como **PWA** (Progressive Web App) en tu navegador o dispositivo móvil;
- Soporta **live-reloading mode** para tus estilos CSS;
- Tiene una biblioteca de **ayudantes de fácil uso** para tu código Go;
- Contiene un completo **ejemplo** de cómo utilizarlo fuera de la caja.

> [!NOTE]
> Para que conozcas a fondo el proyecto, hemos grabado un breve 📺[vídeo][gowebly_youtube_video_url] y preparado una introducción 📝[artículo][gowebly_devto_article_url] demostrando las principales características de la CLI de **Gowebly**.

## ⚡️ Inicio rápido

Primero, [descarga][go_download_url] e instala **Go**. Se requiere la versión `1.21` (o superior).

Ahora, puedes usar la CLI **Gowebly** sin instalación. Simplemente, ejecútalo con [`go run`][go_run_url] para crear un nuevo proyecto con una configuración [default][repo_default_config]:

```console
go run github.com/gowebly/gowebly@latest create
```

¡Ya está! 🔥 Se ha creado una maravillosa aplicación web en la carpeta actual.

Utilizará el paquete **net/http** (como backend Go) y el paquete **html/template** (como motor de plantillas). Las librerías **htmx** y **hyperscript** ya están disponibles en sus plantillas HTML.

### 🐳 Docker-manera de empezar rápido

Siéntete libre de usar la CLI **Gowebly** de nuestra [imagen Docker oficial][docker_image_url] y ejecutarla en el contenedor aislado:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} gowebly/gowebly:latest create
```

> [!IMPORTANT]
> Esta imagen Docker funciona **solo** en los sistemas GNU/Linux (`amd64` o `arm64`, incluido WSL).

### 🍺 Homebrew-manera de inicio rápido

GNU/Linux y Apple macOS, usuarios disponibles manera de instalar **Gowebly** CLI a través de [Homebrew][brew_url].

Toque una nueva fórmula:

```console
brew tap gowebly/tap
```

Instale:

```console
brew install gowebly/tap/gowebly
```

### 📦 Otra forma de empezar rápidamente

Descargue archivos `exe` ya creados para paquetes Windows, `deb`, `rpm`, `apk` o Arch Linux desde la página [Releases][repo_releases_url].

## 📖 Guía de usuario completa

¡Siempre valoramos tu tiempo y queremos que empieces a construir productos web realmente geniales sobre esta impresionante pila tecnológica lo antes posible! Por lo tanto, para obtener una guía completa para usar y entender los principios básicos de la CLI **Gowebly**, hemos preparado una explicación completa del proyecto en este 📖 [**Guía de usuario completa**][docs_url].

<a href="https://gowebly.org" target="_blank" title="Go to the Gowebly's Complete user guide"><img width="360px" alt="gowebly docs banner" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-docs-banner.svg"></a>

Nos hemos preocupado de que te resulte **lo más cómodo** aprender esta maravillosa herramienta, por lo que cada comando CLI tiene una descripción textual suficiente, así como un diagrama visual de cómo funciona.

### El camino del aprendizaje

Se recomienda encarecidamente empezar a explorar con breves artículos introductorios "[**¿Qué es Gowebly CLI?**][docs_what_is_gowebly_cli_url]" y "[**¿Cómo funciona?**][docs_how_it_works_url]" para entender el principio básico y los principales componentes incorporados en el **Gowebly** CLI.

Los siguientes pasos son:

1. [Instala la CLI en tu sistema](https://gowebly.org/complete-user-guide/installation)
2. [Configura tu proyecto](https://gowebly.org/complete-user-guide/configuration)
3. [Empezar a crear un nuevo proyecto](https://gowebly.org/complete-user-guide/create-new-project)
4. [Ejecutar su proyecto localmente](https://gowebly.org/complete-user-guide/run-your-project)
5. [Construir su proyecto para la producción](https://gowebly.org/complete-user-guide/build-project)

Espero que encuentres respuesta a todas tus preguntas! 😉

## 🎯 Motivación para crear

Cuéntanos, ¿cuántas veces has tenido que empezar un nuevo proyecto desde cero y has tenido que hacer dolorosas configuraciones manuales? 🤔 Especialmente, cuando recién te estás familiarizando con una nueva tecnología o stack, donde todo es nuevo para ti.

Para muchos desarrolladores, _incluidos nosotros_, este proceso es lo más tedioso e incluso deprimente posible, y no conlleva ninguna carga de trabajo útil. Es un proceso **muy** frustrante que puede alejar mucho a cualquier desarrollador de la tecnología.

¿Por qué no dar todo ese horrible trabajo manual a las máquinas? Dejemos que hagan todo el trabajo duro por nosotros, y nos limitaremos a crear increíbles productos web sin tener que pensar en construir y desplegar.

Es por eso que hemos creado la CLI **Gowebly** y su biblioteca de ayudantes, que le ayuda a iniciar una increíble aplicación web en **Go** utilizando **htmx**, **hyperscript** y populares atómica / utilidad-primero **CSS frameworks**.

Estamos aquí para salvarte (_y salvarnos_) de este dolor rutinario! ✨

> [!NOTE]
> Anteriormente, ya hemos salvado el mundo una vez, fue [Create Go App][cgapp_url] (sip, ese también es nuestro proyecto). Las estadísticas de [GitHub stars][cgapp_stars_url] de este proyecto no pueden mentir: más de **2.2k** desarrolladores de cualquier nivel y diferentes países comienzan un nuevo proyecto a través de esta herramienta CLI.

## 🏆 Una cooperación en la que todos ganan

Si te ha gustado **Gowebly** CLI y te ha resultado útil para tus tareas, por favor, haz clic en el botón 👁️ **Watch** para no perderte las notificaciones sobre nuevas versiones, ¡y dale una 🌟 **Estrella GitHub**!

Realmente **nos motiva** a hacer este producto **aún** mejor.

<img width="100%" alt="gowebly star and watch" src="https://github.com/gowebly/gowebly/assets/11155743/6f92ec26-1fe3-44c6-9a13-3abd3ffa58eb">

Y ahora, ¡te invito a participar en este proyecto! Trabajemos **juntos** para crear y popularizar la herramienta **más útil** para desarrolladores en la web hoy en día.

- [Issues][repo_issues_url]: haz preguntas y envía tus funcionalidades.
- [Pull requests][repo_pull_request_url]: envía tus mejoras a la corriente.
- [Discussions][repo_disscussions_url]: discute y comparte tus ideas.
- Di unas palabras sobre el proyecto en tus redes sociales y blogs (Dev.to, Medium, Хабр, etc.).

¡Tus PR, problemas y cualquier palabra son bienvenidos! Gracias 😘

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
[docs_what_is_gowebly_cli_url]: https://gowebly.org/getting-started
[docs_how_it_works_url]: https://gowebly.org/getting-started/how-does-it-work

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