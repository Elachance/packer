package rpc

import (
	"github.com/mitchellh/packer/packer"
	"net/rpc"
)

// Registers the appropriate endpoint on an RPC server to serve a
// Packer Build.
func RegisterBuild(s *rpc.Server, b packer.Build) {
	s.RegisterName("Build", &BuildServer{b})
}

// Registers the appropriate endpoint on an RPC server to serve a
// Packer Builder.
func RegisterBuilder(s *rpc.Server, b packer.Builder) {
	s.RegisterName("Builder", &BuilderServer{b})
}

// Registers the appropriate endpoint on an RPC server to serve a
// Packer Command.
func RegisterCommand(s *rpc.Server, c packer.Command) {
	s.RegisterName("Command", &CommandServer{c})
}

// Registers the appropriate endpoint on an RPC server to serve a
// Packer Environment
func RegisterEnvironment(s *rpc.Server, e packer.Environment) {
	s.RegisterName("Environment", &EnvironmentServer{e})
}

// Registers the appropriate endpoint on an RPC server to serve a
// Packer UI
func RegisterUi(s *rpc.Server, ui packer.Ui) {
	s.RegisterName("Ui", &UiServer{ui})
}

func serveSingleConn(s *rpc.Server) string {
	l := netListenerInRange(portRangeMin, portRangeMax)

	// Accept a single connection in a goroutine and then exit
	go func() {
		defer l.Close()
		conn, err := l.Accept()
		if err != nil {
			return
		}

		s.ServeConn(conn)
	}()

	return l.Addr().String()
}