package karigo

// Service ...
type Service interface {
	// Connect instantiates a connection to the service behind
	// the source.
	//
	// Implementations should document the elements they expect
	// in the map.
	Connect(map[string]string) error

	// Ping reports whether there is an active connection or not.
	Ping() bool
}
