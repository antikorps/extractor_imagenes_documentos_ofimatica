# Extractor de imágenes para documentos de ofimática
Utilidad para extraer automáticamente las imágenes incorporadas en archivos de ofimática con extensión .dotx o .xlsx

Puede usarse para un archivo concretos o para buscar todos los archivos coincidentes en un directorio (incluyendo subdirectorios)

# Uso
No es necesario ningún tipo de instalación. Únicamente debe descargarse el binario correspondiente al sistema operativo en el que se vaya a utilizar de la carpeta bin.

Puede ejecutarse por línea de comandos o mediante una interfaz guiada por diálogos.

### Línea de comandos
Se espera que se pase un origen y un destino. El origen puede ser la ruta completa de un archivo o de un directorio. El destino es la ruta completa del directorio en el que se guardarán las imágenes.
```bash
./extractor_imagenes_documentos_ofimatica -origen /home/usuario/Documentos/archivosOfimatica/ -destino /home/usuario/Documentos/archivosOfimatica/imagenes
```

### Interfaz guiada por diálogos
La interfaz guiada por diálogos se inicia automáticamente si no se facilita un origen y un destino. Primeramente aparecerá un diálogo informativo:

![Diálogo de presentación](https://i.imgur.com/HhE12iF.png)

Posteriormente deberá seleccionarse el tipo de extracción deseada. La extracción individual permite seleccionar un archivo .docx o .xlsx, mientras que con la extracción múltiple aparecerá un selector de archivos o carpetas.

![Tipo de extracción](https://i.imgur.com/FxOwTeX.png) 

Tras este primer selector de archivos, aparecerá un segundo en el que deberá escogerse el directorio en el que se guardarán las imágenes extraidas.

Tras estos pasos deberá esperarse al diálogo final en el que se recogerá el resumen de la ejecución, mostrandose cualquier error en el caso de que se haya producido.

![Resultado final](https://i.imgur.com/Sr1RbBw.png)