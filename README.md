<h1 align="center">Scraper-FLV</h1>

## Descripción

Scraper-flv es un script que realiza [web scraping](https://es.wikipedia.org/wiki/Web_scraping) en el sitio web de [AnimeFLV](https://www3.animeflv.net/) para extraer toda la información relevante de las series y películas de anime.

## Motivación

Dado que el sitio web no cuenta con una API pública oficial, he decidido crear este web scraper para extraer la información pública de los animes y luego exponerla a través de una API.

De esa manera, cualquier persona podrá crear su propia API, ya sea REST o GraphQL, o incluso desarrollar una aplicación móvil basada en el sitio web de AnimeFLV.

## Instalación

1. Clona el repositorio

```bash
git clone https://github.com/ZUR1C4T0/scraper-flv.git
```

2. Instala las dependencias

```bash
go mod download
```

3. Ejecuta el script

```bash
go run .
```


# Uso

Una vez ejecutado el script, este comenzará a extraer la información del sitio y la guardará en un archivo JSON llamado animes.json.

Al finalizar, el script mostrará por la terminal un breve resumen de la información que ha encontrado y almacenado.

<center>

![resultado del script](/script-result.png)

</center>


# Contribución

Si deseas colaborar, ya sea agregando características :sparkles:, corrigiendo errores :bug:, mejorando el código :lipstick:, o aumentando el rendimiento :zap:, por favor, envía una Pull Request para contribuir al crecimiento de este pequeño proyecto :rocket:.

# Licencia

Este proyecto está licenciado bajo la Licencia [MIT](https://en.wikipedia.org/wiki/MIT_License) - ver el archivo [LICENSE](./LICENSE) para más detalles.
