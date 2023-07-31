package extractor

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ncruces/zenity"
)

func configurarMedianteDialogos(extensionesValidas []string) (string, string) {
	presentacionMensaje := fmt.Sprintf(`¡Hola!

Este extractor de imágenes sirve para archivos:
%v

Existen 2 posibilidades de extracción:
   - individual (un único archivo)
   - múltiple (todos los archivos de un directorio*)

Podrás escoger la opción en el próximo diálogo.
Luego tendrás que seleccionar el archivo o carpeta.
Finalmente deberás seleccionar un directorio de destino.

===============
* Implica búsqueda recursiva o en subdirectorios`, strings.Join(extensionesValidas, "\n"))

	presentacionError := zenity.Info(presentacionMensaje,
		zenity.Title("Presentación"),
		zenity.Width(600),
		zenity.InfoIcon)
	if presentacionError != nil {
		if presentacionError == zenity.ErrCanceled {
			os.Exit(0)
		}
		log.Fatalln("error crítico en el diálogo de presentación", presentacionError)
	}

	tipo, tipoError := zenity.List(
		"Selecciona el tipo de extracción:",
		[]string{"individual", "múltiple"},
		zenity.Title("Tipo de extracción"),
		zenity.DisallowEmpty(),
	)
	if tipoError != nil {
		log.Fatalln("error crítico en el diálogo de tipo")
	}

	var origen string
	var origenError error
	if tipo == "individual" {
		var extensionesSelector []string
		for _, ext := range extensionesValidas {
			extensionSelector := fmt.Sprintf("*%v", ext)
			extensionesSelector = append(extensionesSelector, extensionSelector)
		}
		origen, origenError = zenity.SelectFile(
			zenity.FileFilters{
				{Name: "Documentos de ofimática", Patterns: extensionesSelector, CaseFold: false},
			})
		if origenError != nil {
			if origenError == zenity.ErrCanceled {
				os.Exit(0)
			}
			log.Fatalln("error en el diálogo de selección individual", origenError)
		}
	}
	if tipo == "múltiple" {
		origen, origenError = zenity.SelectFile(
			zenity.Directory())
		if origenError != nil {
			if origenError == zenity.ErrCanceled {
				os.Exit(0)
			}
			log.Fatalln("error en el diálogo de selección múltiple", origenError)
		}
	}

	destino, destinoError := zenity.SelectFile(
		zenity.Directory())
	if destinoError != nil {
		if destinoError == zenity.ErrCanceled {
			os.Exit(0)
		}
		log.Fatalln("error en el diálogo de destino", destinoError)
	}

	return origen, destino
}

func obtenerArchivos(origen string, extensionesValidas []string) ([]Archivo, error) {
	var archivos []Archivo

	info, infoError := os.Stat(origen)
	if infoError != nil {
		mensajeError := fmt.Sprintf("error crítico comprobando la ruta proporcionada: %v", infoError)
		return archivos, errors.New(mensajeError)
	}
	// Origen es un archivo
	if !info.IsDir() {
		archivos = append(archivos, Archivo{
			Nombre: obtenerNombreDeRuta(origen),
			Tipo:   filepath.Ext(origen),
			Ruta:   origen,
		})
		return archivos, nil
	}
	// Origen es una ruta
	directorioError := filepath.Walk(origen, func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			mensajeError := fmt.Sprintf("error crítico recorriendo el directorio proporcionado: %v", err)
			return errors.New(mensajeError)
		}
		if info.IsDir() {
			return nil
		}
		extension := filepath.Ext(ruta)
		if !contiene(extensionesValidas, extension) {
			return nil
		}
		archivos = append(archivos, Archivo{
			Nombre: obtenerNombreDeRuta(ruta),
			Tipo:   filepath.Ext(ruta),
			Ruta:   ruta,
		})
		return nil
	})
	if directorioError != nil {
		return archivos, directorioError
	}
	return archivos, nil

}

func Configurar(origen, destino string, extensionesValidas []string) Extractor {
	var dialogos bool
	var multiple bool
	if origen == "" || destino == "" {
		origen, destino = configurarMedianteDialogos(extensionesValidas)
		dialogos = true
	}
	archivos, archivosError := obtenerArchivos(origen, extensionesValidas)
	if archivosError != nil {
		if dialogos {
			zenity.Info(archivosError.Error(),
				zenity.Title("ERROR CRÍTICO"),
				zenity.Width(600),
				zenity.ErrorIcon)
		}
		log.Fatalln(archivosError)
	}
	if len(archivos) == 0 {
		mensajeError := "no se han encontrado archivos para procesar"
		if dialogos {
			zenity.Info(mensajeError,
				zenity.Title("ERROR CRÍTICO"),
				zenity.Width(600),
				zenity.ErrorIcon)
		}
		log.Fatalln(mensajeError)
	}
	if len(archivos) > 1 {
		multiple = true
	}

	return Extractor{
		Archivos: archivos,
		Dialogos: dialogos,
		Destino:  destino,
		Multiple: multiple,
	}
}
