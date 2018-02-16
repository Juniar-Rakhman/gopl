// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This is an example that creates interface pollution
// by improperly using an interface when one is not needed.
package main

import "fmt"

// Server defines a contract for tcp servers.
//type Server interface {
//	Start() error
//	Stop() error
//	Wait() error
//}

// Server is our Server implementation.
type Server struct {
	host string

	// PRETEND THERE ARE MORE FIELDS.
}

// NewServer returns an interface value of type Server
// with a Server implementation.
func NewServer(host string) *Server {

	// SMELL - Storing an unexported type pointer in the interface.
	return &Server{host}
}

// Start allows the Server to begin to accept requests.
func (s *Server) Start() error {
	fmt.Println("Server start")
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Stop shuts the Server down.
func (s *Server) Stop() error {
	fmt.Println("Server stop")
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Wait prevents the Server from accepting new connections.
func (s *Server) Wait() error {
	fmt.Println("Server wait")
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

func main() {

	// Create a new Server.
	srv := NewServer("localhost")

	// Use the API.
	srv.Start()
	srv.Stop()
	srv.Wait()
}

// =============================================================================

// NOTES:

// Smells:
//  * The package declares an interface that matches the entire API of its own concrete type.
//  * The interface is exported but the concrete type is unexported.
//  * The factory function returns the interface value with the unexported concrete type value inside.
//  * The interface can be removed and nothing changes for the user of the API.
//  * The interface is not decoupling the API from change.
