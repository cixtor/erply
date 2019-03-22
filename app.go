package main

// Application is the base struct for the whole web API service.
//
// All the API endpoints are attached to this struct as methods, this allows us
// to reference additional resources like the database, cache layer (if there
// is one), and other middlewares.
type Application struct {
	// host is the name or IP address for the web server.
	host string

	// port is the port number where the web server is running.
	port string
}
