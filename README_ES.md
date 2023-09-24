# gowebly – Una herramienta CLI de nueva generación para crear increíbles aplicaciones web en Go utilizando htmx & hyperscript

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][repo_url]
[![License][repo_license_img]][repo_license_url]

[English][repo_url] | [Русский][repo_readme_ru_url] | 
[中文][repo_readme_cn_url] | **Español**

Esta herramienta CLI puede construir fácilmente increíbles aplicaciones web 
con **Go** en el backend usando [**htmx**][htmx_url] & 
[**hyperscript**][hyperscript_url] y los más populares **frameworks CSS** 
atomic/utility-first en el frontend.

Funciones:

- 100% **libre** y de **código abierto** bajo la licencia
  [Apache 2.0][repo_license_url].
- Para **cualquier** nivel de conocimiento y experiencia técnica del 
  desarrollador.
- Ayuda a introducirse más rápidamente en la pila tecnológica **Go** + 
  **htmx** + **hyperscript**.
- CLI inteligente que **hace la mayor parte** de la rutina de configuración 
  y preparación para la producción.
- Posibilidad de añadir simplemente un **framework CSS** 
  atomic/utility-first a un proyecto;
- Contiene un **ejemplo** completo de cómo usarlo fuera de la caja.
- **Bien documentado**, con muchos consejos y ayudas de los autores.

## ⚡️ Inicio rápido

En primer lugar, [descargue][go_download_url] e instale **Go**. Se requiere 
la versión `1.21` (o superior).

Ahora, puedes usar `gowebly` sin instalación. Simplemente, 
[`go run`][go_run_url] para crear un nuevo proyecto con un archivo 
[config][repo_default_config] por defecto:

```console
go run github.com/gowebly/gowebly@latest create
```

¡Ya está! 🔥 Una maravillosa aplicación web, utilizando el paquete 
incorporado **net/http** (como backend de Go), **htmx** & **hyperscript** 
está disponible en sus plantillas Go HTML.

### 🔹 Un completo Go-way de inicio rápido

Si todavía quieres instalar `gowebly` CLI en tu sistema por Golang, usa el 
comando [`go install`][go_install_url]:

```console
go install github.com/gowebly/gowebly@latest
```

### 🍺 Un Homebrew-way de inicio rápido

GNU/Linux y Apple macOS usuarios disponibles manera de instalar `gowebly` CLI a 
través de [Homebrew][brew_url].

Toque una nueva fórmula:

```console
brew tap gowebly/tap
```

Instalar `gowebly`:

```console
brew install gowebly/tap/gowebly
```

### 🐳 Un Docker-way de inicio rápido

Siéntete libre de usar `gowebly` CLI desde nuestra 
[imagen Docker oficial][docker_image_url] y ejecutarlo en el contenedor aislado:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} gowebly/gowebly:latest create
```

### 📦 Otra forma de inicio rápido

Descargue archivos `exe` Windows, `deb`, `rpm`, `apk` o paquetes Arch Linux 
listos para usar en la página [Releases][repo_releases_url].

## 📖 Guía completa del usuario

Para obtener una guía completa para utilizar y comprender los principios 
básicos de la CLI `gowebly`, hemos preparado una explicación exhaustiva de 
cada comando a la vez en este archivo README.

> ⚡️ De los autores: Siempre valoramos tu tiempo y queremos que empieces a 
> crear productos web realmente geniales en esta impresionante pila 
> tecnológica lo antes posible.

Esperamos que encuentre respuesta a todas sus preguntas! 👌 Pero, si no 
encuentra la información necesaria, no dude en crear una 
[issue][repo_issues_url] o enviar un [PR][repo_pull_request_url] a este 
repositorio.

### `init`

Comando para crear un archivo de configuración **default** 
([`.gowebly.yml`][repo_default_config]) en la carpeta actual.

```console
gowebly init
```

> 💡 Nota: Por supuesto, puedes saltarte este paso si te sientes cómodo con 
> la siguiente configuración por defecto para tu nuevo proyecto:
>
> - Ir al backend con el paquete **net/http**;
> - El puerto del servidor es `5000`, tiempo de espera (en segundos): `5` para 
> lectura, `10` para escritura;
> - Últimas versiones de **htmx** & **hyperscript**;
> - Sin ningún framework CSS (solo estilos por defecto para el código de 
> ejemplo).

Normalmente, el archivo de configuración creado contiene las siguientes 
opciones:

```yaml
backend:
  name: default
  port: 5000
  timeout:
    read: 5
    write: 10
  
frontend:
  htmx: latest
  hyperscript: latest
  css_framework: default
```

Pero, puede elegir cualquier **Go** backend con opciones de servidor para su 
proyecto (_esto es obligatorio_):

| Nombre del backend  | Descripción                                                              |
|---------------------|--------------------------------------------------------------------------|
| `default`           | Utilizar un backend Go con el paquete [net/http][net_http_url] integrado |
| `fiber`             | Utilizar un backend Go con el framework web [Fiber][fiber_url]           |
| `echo`              | Utilizar un backend Go con el framework web [Echo][echo_url]             |
| `chi`               | Utilizar un backend Go con el enrutador componible [chi][chi_url]        |

Adicionalmente, puedes elegir versiones de **htmx**, **hyperscript**, y uno 
de los más populares atomic/utility-first **framework CSS** para tu proyecto 
(_esto es opcional, no obligatorio_):

| Framework CSS | Descripción                                                                           |
|---------------|---------------------------------------------------------------------------------------|
| `default`     | No utilices ningún framework CSS (solo estilos por defecto para el código de ejemplo) |
| `tailwindcss` | Utilizar [Tailwind CSS][tailwindcss_url] como framework CSS                           |
| `unocss`      | Utilizar [UnoCSS][unocss_url] como framework CSS                                      |

### `create`

Comando para crear un nuevo proyecto con el backend **Go**, **htmx** & 
**hyperscript**, y (_opcionalmente_) atomic/utility-first **CSS framework**.

```console
gowebly create
```

> 💡 Nota: Si no ejecutas el comando `init` para crear un archivo de 
> configuración (`.gowebly.yml`), por defecto el CLI de `gowebly` crea un 
> nuevo proyecto con una configuración [por defecto][repo_default_config].

Normalmente, el proyecto creado contiene los siguientes archivos y carpetas:

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
├── go.mod
├── go.sum
├── handlers.go
├── main.go
├── package-lock.json
├── package.json
└── server.go
```

### `run`

Comando para ejecutar su proyecto en modo de **desarrollo** (no producción).

```console
gowebly run
```

> 💡 Nota: Si no ejecutas el comando `init` para crear un archivo de 
> configuración (`.gowebly.yml`), por defecto el CLI de `gowebly` ejecuta tu 
> proyecto con una configuración [por defecto][repo_default_config].

Las siguientes versiones de la biblioteca se suministrarán en las plantillas 
Go HTML:

- **htmx**: última versión no de producción desde CDN;
- **hyperscript**: última versión no de producción desde CDN;
- (_opcionalmente_) **framework CSS**: última versión no de producción desde 
  CDN;

En el modo de desarrollo, solo se utilizarán las CDN oficiales compatibles 
de los desarrolladores se utilizarán: 

- [unpkg.com][unpkg_url] para **htmx** & **hyperscript**;
- [tailwindcss.com][tailwindcss_cdn_url] para **Tailwind CSS**;
- [jsDelivr][jsdelivr_url] para **UnoCSS**.

Cada vez que haga `run` comando para su proyecto:

1. CLI válida la configuración y aplica todos los ajustes al proyecto actual;
2. CLI descarga versiones minimizadas de **htmx** y **hyperscript** (de CDN 
   oficiales y de confianza) a la carpeta `./static` y las coloca como 
   etiquetas separadas `<script>` (al final de la etiqueta `<body>`) en la 
   plantilla Go HTML [`templates/main.html`][repo_main_layout];
3. (_opcionalmente_) CLI descarga una versión no de producción del **framework 
   CSS** seleccionado (desde un CDN oficial y de confianza) a la carpeta
   `./static` y lo coloca como una etiqueta `<link>` (al final de la etiqueta 
   `<head>`) en la plantilla Go HTML [`templates/main.html`][repo_main_layout];
4. CLI inicia el backend de un proyecto en el puerto seleccionado mediante 
   el simple comando `go run`.

### `build`

Comando para construir su proyecto para **producción** y preparar 
contenedores para desplegar.

```console
gowebly build
```

> 💡 Nota: Si no ejecutas el comando `init` para crear un archivo de 
> configuración (`.gowebly.yml`), por defecto el CLI de `gowebly` construye 
> tu proyecto con una configuración [por defecto][repo_default_config].

Las siguientes versiones de la biblioteca se suministrarán en las plantillas
Go HTML:

- **htmx**: versión de producción minificada, seleccionada en el archivo de 
  configuración;
- **hyperscript**: versión de producción minificada, seleccionada en el 
  archivo de configuración;
- (_opcionalmente_) **framework CSS**: última versión de producción 
  minificada después de tree-shaking;

Cada vez que haga `build` comando para su proyecto:

1. CLI válida la configuración y aplica todos los ajustes al proyecto actual;
2. CLI descarga versiones minimizadas de **htmx** y **hyperscript** (de CDN
   oficiales y de confianza) a la carpeta `./static` y las coloca como
   etiquetas separadas `<script>` (al final de la etiqueta `<body>`) en la
   plantilla Go HTML [`templates/main.html`][repo_main_layout];
3. (_opcional_) CLI prepara una versión de producción del **CSS framework** 
   seleccionado (usando la herramienta [Vite][vite_url]) y lo coloca como 
   una etiqueta `<link>` (al final de la etiqueta `<head>`) en la plantilla 
   Go HTML [`templates/main.html`][repo_main_layout];
4. CLI generar un archivo claro y bien documentado `docker-compose.yml` en 
   la raíz de la carpeta del proyecto para desplegarlo en contenedores 
   Docker aislados a través de [Portainer][portainer_url] (_recomendado_) o 
   manualmente a su servidor remoto.

## 🎯 Motivación para crear

Cuéntanos, ¿cuántas veces has tenido que empezar un nuevo proyecto desde 
cero y has tenido que hacer dolorosas configuraciones manuales? 🤔 Sobre todo 
cuando recién te estás familiarizando con una nueva tecnología o stack, 
donde todo es nuevo para ti.

Para muchos desarrolladores, _incluidos nosotros_, este proceso es lo más 
tedioso e incluso deprimente posible, y no conlleva ninguna carga de trabajo 
útil. Es un proceso **muy** frustrante que puede alejar mucho a cualquier 
desarrollador de la tecnología.

¿Por qué no dar todo ese horrible trabajo manual a las máquinas? Dejemos que 
hagan todo el trabajo duro por nosotros, y nos limitaremos a crear 
increíbles productos web sin tener que pensar en compilar y desplegar.

Por eso hemos creado la CLI `gowebly`, que te ayuda a crear increíbles 
aplicaciones web en **Go** usando **htmx**, **hyperscript** y populares 
frameworks atómicos/utility-first **CSS**.

¡Estamos aquí para salvarte (_y a mí mismo_) de este dolor rutinario! ✨

## 🏆 Una cooperación beneficiosa para todos

Y ahora, ¡te invito a participar en este proyecto! Trabajemos **juntos** 
para crear la herramienta **más útil** para desarrolladores en la web hoy en 
día.

- [Issues][repo_issues_url]: pregunte y envíe sus características.
- [Pull requests][repo_pull_request_url]: envíe sus mejoras a la corriente.

Sus relaciones públicas y problemas son bienvenidos. Gracias a todos 😘

## ⚠️ Licencia

[`gowebly`][repo_url] es software libre y de código abierto licenciado
bajo la [Licencia Apache 2.0][repo_license_url], creado y soportado por
[Vic Shóstak][author_url] con 🩵 para personas y robots.

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
[repo_main_layout]: https://github.com/gowebly/gowebly/blob/main/internal/attachments/frontend/gowebly/main.html

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
