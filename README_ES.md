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

Ahora, puedes usar `gowebly` sin instalación. Basta con [`go run`][go_run_url] 
con opciones para crear un nuevo proyecto:

```console
go run github.com/gowebly/gowebly@latest create
```

¡Ya está! 🔥 Tu maravillosa aplicación web (en este ejemplo, usando el 
paquete `net/http` incorporado) con **htmx** & **hyperscript** está 
disponible en plantillas Go HTML.

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
docker run \
    --rm -it -v ${PWD}:${PWD} -w ${PWD} \
    gowebly/gowebly:latest [COMMAND] [OPTIONS]
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

### `create`

Comando para crear un nuevo proyecto con el backend Go dado, **htmx** & 
**hyperscript**.

```console
gowebly create [BACKEND]
```

> 💡 Nota: Si no defines un backend Go, por defecto el CLI `gowebly` crea un 
> nuevo proyecto con el paquete incorporado [net/http][net_http_url].

Puedes elegir otro backend **Go** para tu proyecto:

| Backend    | Descripción                                                                                |
|------------|--------------------------------------------------------------------------------------------|
| `fiber`    | Crear un nuevo proyecto con Go backend con el framework web [Fiber][fiber_url]             |
| `echo`     | Crear un nuevo proyecto con Go backend con el framework web [Echo][echo_url]               |
| `chi`      | Crear un nuevo proyecto con Go backend con el [chi][chi_url] composable router             |

### `add`

Comando para añadir uno de los más populares atomic/utility-first 
**framework CSS** a tu proyecto. _Esto es opcional, no obligatorio._

```console
gowebly add [CSS_FRAMEWORK]
```

You can choose the **CSS framework**:

| CSS framework | Descripción                                                      |
|---------------|------------------------------------------------------------------|
| `tailwindcss` | Añade [Tailwind CSS][tailwindcss_url] al frontend de tu proyecto |
| `unocss`      | Añade [UnoCSS][unocss_url] al frontend de tu proyecto            |

### `run`

Comando para ejecutar su proyecto en modo de **desarrollo** (no producción).

```console
gowebly run
```

> 💡 Nota: `gowebly` CLI busca el archivo de configuración YAML
> (`.gowebly.yml`) para el proyecto en la carpeta actual.

Las siguientes versiones de la biblioteca se suministrarán en las plantillas 
Go HTML:

- **htmx**: última versión no de producción desde CDN;
- **hyperscript**: última versión no de producción desde CDN;
- (_optionally_) **framework CSS**: última versión no de producción desde CDN;

En el modo de desarrollo, solo se utilizarán las CDN oficiales compatibles 
de los desarrolladores se utilizarán: 

- [unpkg.com][unpkg_url] para **htmx** & **hyperscript**;
- [tailwindcss.com][tailwindcss_cdn_url] para **Tailwind CSS**;
- [jsDelivr][jsdelivr_url] para **UnoCSS**.

Cada vez que haga `run` comando para su proyecto:

1. CLI incrusta versiones CDN de **htmx** & **hyperscript** a tus plantillas 
   Go HTML en una etiqueta normal `<script>` dentro del bloque llamado 
   `gowebly-body-scripts` (normalmente, situado en la parte inferior de la 
   etiqueta `<body>`);
2. (_opcionalmente_) CLI incrustar una versión CDN del **framework CSS** 
   elegido a sus plantillas Go HTML en una etiqueta normal `<link>` en el 
   bloque llamado `gowebly-head-styles` (normalmente, situado en la parte 
   inferior de la etiqueta `<head>`);
3. CLI inicia el backend de un proyecto (en el puerto `5000`) mediante el 
   simple comando `go run` comando.

### `build`

Comando para construir su proyecto para **producción** y preparar 
contenedores para desplegar.

```console
gowebly build
```

> 💡 Nota: `gowebly` CLI busca el archivo de configuración YAML 
> (`.gowebly.yml`) para el proyecto en la carpeta actual.

Las siguientes versiones de la biblioteca se suministrarán en las plantillas
Go HTML:

- **htmx**: versión de producción minificada, seleccionada en el archivo de 
  configuración;
- **hyperscript**: versión de producción minificada, seleccionada en el 
  archivo de configuración;
- (_optionally_) **framework CSS**: última producción tree-shaking & versión 
  minificada;

Cada vez que haga `build` comando para su proyecto:

1. CLI escanear y validar el archivo de configuración YAML (`.gowebly.yml`), 
   aplicar toda la configuración al proyecto actual;
2. CLI descarga versiones minificadas de **htmx** & **hyperscript** de los 
   recursos oficiales (y de confianza);
   - Incrústelos en sus plantillas Go HTML (estilo inline) en el bloque 
     llamado `gowebly-body-scripts` (normalmente, situado en la parte 
     inferior de la etiqueta `<body>`);
3. (_opcionalmente_) CLI preparar una versión minificada (y arborescente) 
   del **framework CSS** elegido a través de la herramienta [Vite][vite_url];
   - Incrústalos en tus plantillas Go HTML (estilo inline) en el bloque 
     llamado `gowebly-head-styles` (normalmente, situado en la parte 
     inferior de la etiqueta `<head>`);
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
