package extractor

import (
	"log"
	"strings"

	"github.com/ncruces/zenity"
)

func (e *Extractor) Informar() {
	if len(e.Errores) == 0 {
		if e.Dialogos {
			zenity.Info("Proceso de extracción de imágenes realizado satisfactoriamente",
				zenity.Title("FIN"),
				zenity.Width(600),
				zenity.InfoIcon)
		}
		return
	}
	if e.Dialogos {
		mensajeError := strings.Join(e.Errores, "\n")
		zenity.Info(mensajeError,
			zenity.Title("Errores en el proceso de descomprensión"),
			zenity.Width(600),
			zenity.ErrorIcon)
	}
	for _, v := range e.Errores {
		log.Println(v)
	}
}
