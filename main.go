package main

import (
	"extractor_imagenes_documentos_ofimatica/extractor"
	"flag"
)

func main() {
	var origen, destino string
	flag.StringVar(&origen, "origen", "", "ruta completa al archivo o directorio con los documentos docx o xlsx de los que se busca extraer las imágenes")
	flag.StringVar(&destino, "destino", "", "ruta completa del directorio en el que se guardarán las imágenes extraídas")
	flag.Parse()

	e := extractor.Configurar(origen, destino)
	e.Extraer()
	e.Informar()
}
