package extractor

import (
	"path/filepath"
	"strings"
)

func contiene(coleccion []string, elemento string) bool {
	for _, v := range coleccion {
		if v == elemento {
			return true
		}
	}
	return false
}

func obtenerNombreDeRuta(ruta string) string {
	nombre := filepath.Base(ruta)
	return strings.TrimSuffix(nombre, filepath.Ext(ruta))
}
