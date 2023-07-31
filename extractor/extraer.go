package extractor

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func descomprimir(archivo Archivo, destino string, multiple bool, canal chan error) {
	defer wg.Done()
	if multiple {
		destino = filepath.Join(destino, archivo.Nombre)
		crearCarpetaError := os.Mkdir(destino, 0777)
		if crearCarpetaError != nil {
			mensajeError := fmt.Sprintf("error en el proceso de extracción de imágenes en el archivo %v, no se ha podido crear la carpeta en el directorio indicado: %v", archivo.Nombre, crearCarpetaError)
			canal <- errors.New(mensajeError)
			return
		}
	}

	zip, zipError := zip.OpenReader(archivo.Ruta)
	if zipError != nil {
		mensajeError := fmt.Sprintf("error en el proceso de extracción de imágenes en el archivo %v, no se ha podido leer el archivo zip: %v", archivo.Nombre, zipError)
		canal <- errors.New(mensajeError)
		return
	}
	defer zip.Close()

	for _, v := range zip.File {
		// Rutas en Linux xl/media/image1.png, quizá en Windows cambian
		rutaV, nombreV := filepath.Split(v.Name)
		partesV := strings.Split(rutaV, string(os.PathSeparator))
		if len(partesV) < 2 {
			continue
		}
		if partesV[1] != "media" {
			continue
		}

		imagen, imagenError := v.Open()
		if imagenError != nil {
			mensajeError := fmt.Sprintf("error en el proceso de extracción de imágenes en el archivo %v, no se ha podido leer la imagen %v: %v", archivo.Nombre, nombreV, imagenError)
			canal <- errors.New(mensajeError)
			return
		}
		defer imagen.Close()

		rutaImagenDestino := filepath.Join(destino, nombreV)
		imagenDestino, imagenDestinoError := os.Create(rutaImagenDestino)
		if imagenDestinoError != nil {
			mensajeError := fmt.Sprintf("error en el proceso de extracción de imágenes en el archivo %v, no se ha podido crear la imagen %v en el destino: %v", archivo.Nombre, nombreV, imagenDestinoError)
			canal <- errors.New(mensajeError)
			return
		}
		defer imagenDestino.Close()

		_, imagenCopiaError := io.Copy(imagenDestino, imagen)
		if imagenCopiaError != nil {
			mensajeError := fmt.Sprintf("error en el proceso de extracción de imágenes en el archivo %v, no se ha podido copiar la imagen %v en el destino: %v", archivo.Nombre, nombreV, imagenDestinoError)
			canal <- errors.New(mensajeError)
			return
		}
	}
}

func (e *Extractor) Extraer() {
	var procesados int
	procesar := true

	for procesar {
		canal := make(chan error)
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go descomprimir(e.Archivos[procesados], e.Destino, e.Multiple, canal)
			procesados++
			if procesados == len(e.Archivos) {
				procesar = false
				break
			}
		}
		go func() {
			wg.Wait()
			close(canal)
		}()
		for v := range canal {
			if v != nil {
				e.Errores = append(e.Errores, v.Error())
			}
		}
	}
}
