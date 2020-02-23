package karigo

// A Service is the interface to handle external services, whether they are on
// the same host or another one.
//
// This is used for managing journals and sources. If an implementation does not
// represent an external service, the methods can pretend every is fine. For
// example, Connect could always return nil.
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
