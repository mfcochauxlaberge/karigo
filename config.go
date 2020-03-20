package karigo

// Config ...
type Config struct {
	// Server
	Port uint

	// Node
	Hosts   []string
	Journal map[string]string
	Sources map[string]map[string]string
}
