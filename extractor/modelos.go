package extractor

type Extractor struct {
	Archivos []Archivo
	Destino  string
	Dialogos bool
	Errores  []string
	Multiple bool
}

type Archivo struct {
	Nombre string
	Tipo   string
	Ruta   string
}
