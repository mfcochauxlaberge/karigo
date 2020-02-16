package karigo

// Config ...
type Config struct {
	// Karigo
	Host       string
	Port       uint
	OtherHosts []string

	// Journal
	Journal map[string]string

	// Source
	Sources map[string]map[string]string
}
