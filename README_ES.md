<img width="256px" alt="gowebly logo" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-logo.svg">

# gowebly – Una herramienta CLI de última generación para crear increíbles aplicaciones web en Go utilizando htmx e hyperscript

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

[English][repo_url] | [Русский][repo_readme_ru_url] | 
[简体中文][repo_readme_cn_url] | **Español**

Esta herramienta CLI puede construir fácilmente increíbles aplicaciones web 
con **Go** en el backend, utilizando [**htmx**][htmx_url] & 
[**hyperscript**][hyperscript_url] y los más populares atomic/utility-first 
**frameworks CSS** en el frontend.

Características:

- 100% **gratis** y **de código abierto** bajo la licencia
  [Apache 2.0][repo_license_url] licencia;
- Para **cualquier** nivel de conocimiento y experiencia técnica del 
  desarrollador;
- **Bien documentado**, con muchos consejos y ayudas de los autores;
- Multiplataforma y multiarquitectura permite **ejecutar con éxito** en 
  GNU/Linux, MS Windows (incluido WSL) y Apple macOS;
- CLI inteligente que **hace la mayor parte** de la rutina de configuración 
  y preparación para la producción;
- Ayuda a entrar más rápidamente en la pila tecnológica **Go** + **htmx** + 
  **hyperscript**;
- La posibilidad de añadir simplemente a tu proyecto un **framework CSS** 
  atómico/utility-first listo para usar y completamente personalizado;
- Admite el modo **live-reloading** para sus estilos CSS.
- Dispone de una biblioteca de **ayudantes de fácil uso** para su código Go;
- Contiene un completo **ejemplo** de cómo utilizarlo fuera de la caja.

> 💬 De los autores: Para que conozcas a fondo el proyecto, hemos grabado un 
> breve [📺 vídeo][gowebly_youtube_video_url] y preparado un 
> [📝 artículo][gowebly_devto_article_url] de introducción en el que se 
> muestran las principales características de la CLI `gowebly`.

## ⚡️ Inicio rápido

Primero, [descarga][go_download_url] e instala **Go**. Se requiere la 
versión `1.21` (o superior).

Ahora, puedes utilizar `gowebly` sin instalación. Simplemente, 
[go run`][go_run_url] para crear un nuevo proyecto con una 
[configuración][repo_default_config] por defecto:

```console
go run github.com/gowebly/gowebly@latest create
```

¡Ya está! 🔥 Una maravillosa aplicación web, usando el paquete integrado 
**net/http** (como backend de Go), **htmx** & **hyperscript** está 
disponible en tus plantillas HTML de Go.

### 🔹 Un completo Go-way de inicio rápido

Si todavía quieres instalar `gowebly` CLI en tu sistema por Golang, usa el 
comando [`go install`][go_install_url]:

```console
go install github.com/gowebly/gowebly@latest
```

### 🍺 Homebrew-manera de inicio rápido

GNU/Linux y Apple macOS usuarios disponibles manera de instalar `gowebly` 
CLI a través de [Homebrew][brew_url].

Toque una nueva fórmula:

```console
brew tap gowebly/tap
```

Instale `gowebly`:

```console
brew install gowebly/tap/gowebly
```

### 🐳 Docker-manera de inicio rápido

Siéntete libre de usar `gowebly` CLI desde nuestra 
[imagen Docker oficial][docker_image_url] y ejecutarlo en el contenedor aislado:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} gowebly/gowebly:latest create
```

### 📦 Otra forma de empezar rápidamente

Descargue archivos `exe` para Windows, `deb`, `rpm`, `apk` o Arch Linux 
desde la página [Releases][repo_releases_url].

## 📖 Guía completa del usuario

Para obtener una guía completa de uso y comprender los principios básicos de 
la CLI `gowebly`, hemos preparado una explicación exhaustiva de cada comando 
a la vez en este archivo README.

> 💬 De los autores: Siempre valoramos tu tiempo y queremos que empieces a 
> crear productos web realmente geniales en esta impresionante pila 
> tecnológica lo antes posible.

Esperamos que encuentres respuesta a todas tus preguntas 👌 Pero, si no 
encuentras la información que necesitas, no dudes en crear una 
[issue][repo_issues_url] o enviar un [PR][repo_pull_request_url] a este 
repositorio.

### `init`

Comando para crear un archivo de configuración por **defecto** 
([`.gowebly.yml`][repo_default_config]) en la carpeta actual.

```console
gowebly init
```

> 💡 Nota: Por supuesto, puedes saltarte este paso si te sientes cómodo con 
> la siguiente configuración por defecto para tu nuevo proyecto:
>
> - Los nombres de los módulos Go (`go.mod`) y `package.json` se establecen 
> en **project**;
> - Sin ningún framework Go para la parte backend (sólo paquete integrado 
> **net/http**);
> - Sin ningún framework CSS para la parte frontend (solo estilos por defecto 
> para el código de ejemplo);
> - El entorno de ejecución JavaScript para la parte frontend utilizará 
> **Node.js**;
> - El puerto del servidor es `5000`, tiempo de espera (en segundos): `5` 
> para lectura, `10` para escritura;
> - Últimas versiones de **htmx** & **hyperscript**.

<img width="720" alt="gowebly init" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_init.png">

Normalmente, un archivo de configuración creado contiene las siguientes 
opciones:

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
   runtime_environment: default # (string) option can be one of the values: 'bun', or 'default'
   htmx: latest # (string) option can be any existing version
   hyperscript: latest # (string) option can be any existing version
```

Pero, puedes elegir cualquier **Go framework** para el backend de tu proyecto:

| Go framework | Descripción                                                                      |
|--------------|----------------------------------------------------------------------------------|
| `default`    | No uses ningún framework Go (solo el paquete integrado [net/http][net_http_url]) |
| `fiber`      | Utilizar un backend Go con el framework web [Fiber][fiber_url]                   |
| `echo`       | Utilizar un backend Go con el framework web [Echo][echo_url]                     |
| `chi`        | Utilizar un backend Go con el enrutador componible [chi][chi_url]                |

Además, puedes elegir versiones de **htmx**, **hyperscript**, y uno de los 
más populares atomic/utility-first **CSS framework** para tu proyecto:

| CSS framework | Descripción                                                                          |
|---------------|--------------------------------------------------------------------------------------|
| `default`     | No utilice ningún framework CSS (solo estilos por defecto para el ejemplo de código) |
| `tailwindcss` | Utilizar [Tailwind CSS][tailwindcss_url] como framework CSS                          |
| `unocss`      | Utilizar [UnoCSS][unocss_url] como framework CSS                                     |

Además, puede establecer uno de los entornos de ejecución de JavaScript para 
su parte frontend:

| JavaScript runtime | Descripción                                                             |
|--------------------|-------------------------------------------------------------------------|
| `default`          | Utilizar [Node.js][nodejs_url] como en torno de ejecución de JavaScript |
| `bun`              | Utilizar [Bun][bun_url] como en torno de ejecución de JavaScript        |

### `create`

Comando para crear un nuevo proyecto con el backend **Go**, **htmx** & 
**hyperscript**, y (_opcionalmente_) atomic/utility-first **CSS framework**.

```console
gowebly create
```

> 💡 Nota: Si no ejecutas el comando `init` para crear un archivo de 
> configuración (`.gowebly.yml`), el CLI de `gowebly` crea un nuevo proyecto 
> con una configuración [por defecto][repo_default_config].

<img width="720" alt="gowebly create" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_create.png">

Cada vez que haga `crear` comando para su proyecto:

1. CLI válida la configuración y aplica todos los ajustes al proyecto actual;
2. CLI prepara la parte backend de tu proyecto (genera la estructura del 
   proyecto y los archivos de utilidades necesarios, ejecuta `go mod tidy`);
3. CLI prepara la parte frontend de tu proyecto (genera los archivos de 
   utilidades necesarios, ejecuta `npm|bun install` y `npm|bun run build:dev` 
   por primera vez);
4. CLI descarga versiones minimizadas de **htmx** y **hyperscript** (desde 
   el CDN oficial y de confianza [unpkg.com][unpkg_url]) a la carpeta
   `./static` y las coloca como etiquetas separadas `<script>` (al final de la 
   etiqueta `<body>`) en la plantilla Go HTML
   [`templates/main.html`][repo_main_layout]. 

Típicamente, un proyecto creado contiene los siguientes archivos y carpetas:

```console
.
├── assets
│   └── styles.css
├── static
│   ├── favicons
│   │   ├── apple-touch-icon.png
│   │   ├── favicon.png
│   │   ├── favicon.svg
│   │   └── favicon.ico
│   ├── images
│   │   └── logo.svg
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
├── package.json
└── server.go
```

### `run`

Comando para ejecutar su proyecto en modo **desarrollo** (no producción).

```console
gowebly run
```

> 💡 Nota: Si no ejecutas el comando `init` para crear un fichero de 
> configuración (`.gowebly.yml`), el CLI de `gowebly` ejecuta tu proyecto 
> con una configuración [por defecto][repo_default_config].

<img width="720" alt="gowebly run" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_run.png">

Cada vez que haga `run` comando para su proyecto:

1. CLI válida la configuración y aplica todos los ajustes al proyecto actual;
2. CLI prepara la parte frontend de tu proyecto (ejecuta
   `npm|bun run watch`);
3. CLI prepara una versión de desarrollo (no de producción) del **framework 
   CSS** seleccionado en la carpeta `./static` y lo coloca como una 
   etiqueta `<link>` (al final de la etiqueta `<head>`) en la plantilla Go 
   HTML [`templates/main.html`][repo_main_layout];
4. CLI inicia el backend de un proyecto con los ajustes de la configuración 
   por defecto (o del archivo de configuración `.gowebly.yml`) mediante un 
   simple comando `go run`.

### `build`

Comando para construir su proyecto para **producción** y preparar los 
archivos Docker para el despliegue.

```console
gowebly build [OPTION]
```

> 💡 Nota: Si no ejecutas el comando `init` para crear un archivo de 
> configuración (`.gowebly.yml`), el CLI de `gowebly` construye tu proyecto 
> con una configuración [por defecto][repo_default_config].

<img width="720" alt="gowebly build" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_build.png">

Puedes añadir las siguientes opciones:

| Opción          | Descripción                                                                               | Requerido? |
|-----------------|-------------------------------------------------------------------------------------------|------------|
| `--skip-docker` | Omitir el proceso de generación de los archivos Docker (es útil si usted tiene su propio) | no         |

Cada vez que haga `build` comando para su proyecto:

1. CLI válida la configuración y aplica todos los ajustes al proyecto actual;
2. CLI descarga versiones minimizadas de **htmx** y **hyperscript** (desde 
   el CDN oficial y de confianza [unpkg.com][unpkg_url]) a la carpeta
   `./static` y las coloca como etiquetas separadas `<script>` (al final de la 
   etiqueta `<body>`) en la plantilla Go HTML
   [`templates/main.html`][repo_main_layout];
3. CLI prepara una versión de producción del **CSS framework** seleccionado 
   y lo coloca como una etiqueta `<link>` (al final de la etiqueta `<head>`) 
   en la plantilla Go HTML [`templates/main.html`][repo_main_layout];
4. Si la opción `--skip-docker` no está activada, CLI genera unos archivos 
   Docker claros y bien documentados (`.dockerignore`, `Dockerfile`, 
   `docker-compose.yml`) en la raíz de la carpeta del proyecto para 
   desplegarlo en contenedores aislados vía [Portainer][portainer_url] 
   (_recomendado_), o manualmente, a tu servidor remoto.

### `doctor`

Comando para mostrar **información** útil sobre tu sistema.

```console
gowebly doctor
```

> 💡 Nota: Esto es muy útil para el proceso de autodepuración, o para crear 
> un nuevo [issue][repo_issues_url] con un informe de error en este 
> repositorio de GitHub.

<img width="720" alt="gowebly doctor" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_doctor.png">

Cada vez que haga `doctor` comando para su sistema:

1. CLI comprueba las versiones de todas las herramientas necesarias para que 
   tu proyecto tenga éxito (como `gowebly` CLI, Go, Node.js, Docker, Docker 
   Compose, etc.);
2. CLI produce un informe con la versión instalada para cada herramienta.

## 🙋 Ayudas fáciles de usar

La CLI `gowebly` tiene una librería de [helpers][gowebly_helpers_url] 
amigables para tu código. Esto te ayudará a empezar a construir hermosas 
aplicaciones web en Go aún más rápido.

```console
go get -u github.com/gowebly/helpers
```

> 💡 Nota: La librería `gowebly helpers` está **ya** incluida en el backend 
> Go que se crea con el comando `create`, pero puedes usar estos helpers en 
> otros proyectos también.

## 🎯 Motivación para crear

Cuéntanos, ¿cuántas veces has tenido que empezar un nuevo proyecto desde 
cero y has tenido que hacer dolorosas configuraciones manuales? 🤔 
Especialmente, cuando recién te estás familiarizando con una nueva 
tecnología o stack, donde todo es nuevo para ti.

Para muchos desarrolladores, _incluidos nosotros_, este proceso es lo más 
tedioso e incluso deprimente posible, y no conlleva ninguna carga de trabajo 
útil. Es un proceso **muy** frustrante que puede alejar mucho a cualquier 
desarrollador de la tecnología.

¿Por qué no dar todo ese horrible trabajo manual a las máquinas? Dejemos que 
hagan todo el trabajo duro por nosotros, y nos limitaremos a crear 
increíbles productos web sin tener que pensar en compilar y desplegar.

Es por eso que hemos generado la CLI `gowebly` y su biblioteca de ayudantes, 
que le ayuda a iniciar una increíble aplicación web en **Go** utilizando 
**htmx**, **hyperscript** y populares atómica/utilidad-primero 
**frameworks CSS**.

¡Estamos aquí para salvarte (_y salvarnos_) de este dolor rutinario! ✨

> 💬 De los autores: Anteriormente, ya hemos salvado el mundo una vez, fue 
> [Create Go App][cgapp_url] (sí, ese también es nuestro proyecto). 
> Las estadísticas de [GitHub stars][cgapp_stars_url] de este proyecto no 
> pueden mentir: más de **2.2k** desarrolladores de cualquier nivel y 
> diferentes países comienzan un nuevo proyecto a través de esta herramienta 
> CLI.

## 🏆 Una cooperación beneficiosa para todos

¡Si te ha gustado el proyecto `gowebly` y te ha resultado útil para tus 
tareas, por favor, dale un 🌟 **GitHub Star** y haz clic en 👁️ **Watch** para 
no perderte las notificaciones sobre nuevas versiones!

<img width="100%" alt="gowebly star and watch" src="https://github.com/gowebly/gowebly/assets/11155743/6f92ec26-1fe3-44c6-9a13-3abd3ffa58eb">

Y ahora, ¡te invito a participar en este proyecto! Trabajemos **juntos** 
para crear la herramienta **más útil** para desarrolladores en la web hoy en 
día.

- [Issues][repo_issues_url]: haz preguntas y envía tus funcionalidades.
- [Pull requests][repo_pull_request_url]: envía tus mejoras a la corriente.

¡Tus PR e issues son bienvenidos! Gracias 😘

### 🌟 Guardianes de la Galaxia

[![gowebly stargazers][repo_badge_reporoster_url]][repo_stargazers_url]

## ⚠️ Licencia

[`gowebly`][repo_url] es un software libre y de código abierto licenciado 
bajo la [Licencia Apache 2.0][repo_license_url], creado y soportado por 
[Vic Shóstak][author_url] con 🩵 para personas y robots. Logotipo oficial 
distribuido bajo [Licencia Creative Commons][repo_cc_license_url] (CC BY-SA
4.0 International).

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
[repo_license_url]: https://github.com/gowebly/gowebly/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_license_url]: https://creativecommons.org/licenses/by-sa/4.0/
[repo_readme_ru_url]: https://github.com/gowebly/gowebly/blob/main/README_RU.md
[repo_readme_cn_url]: https://github.com/gowebly/gowebly/blob/main/README_CN.md
[repo_readme_es_url]: https://github.com/gowebly/gowebly/blob/main/README_ES.md
[repo_default_config]: https://github.com/gowebly/gowebly/blob/main/internal/attachments/configs/default.yml
[repo_main_layout]: https://github.com/gowebly/gowebly/blob/main/internal/attachments/templates/frontend/main.html
[repo_stargazers_url]: https://github.com/gowebly/gowebly/stargazers
[repo_badge_reporoster_url]: https://reporoster.com/stars/notext/gowebly/gowebly

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- Readme links -->

[gowebly_youtube_video_url]: https://www.youtube.com/watch?v=qazYscnLku4
[gowebly_devto_article_url]: https://dev.to/koddr/a-next-generation-cli-tool-for-building-amazing-web-apps-in-go-using-htmx-hyperscript-336d
[gowebly_helpers_url]: https://github.com/gowebly/helpers
[cgapp_url]: https://github.com/create-go-app/cli
[cgapp_stars_url]: https://github.com/create-go-app/cli/stargazers
[nodejs_url]: https://nodejs.org
[bun_url]: https://bun.sh
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
