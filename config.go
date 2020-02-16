package karigo

// Config ...
type Config struct {
	// Karigo
	Port  uint
	Hosts []string

	// Journal
	Journal map[string]string

	// Source
	Sources map[string]map[string]string
}
