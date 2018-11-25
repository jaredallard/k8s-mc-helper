package minecraft

import (
	"fmt"
	"os"
	"strings"

	"github.com/jaredallard/k8s-mc-helper/pkg/ps"
)

// GetPID returns the PID of the minecraft server
func GetPID() (int, *ps.Process, error) {
	procs, err := ps.Find()
	if err != nil {
		return 0, nil, err
	}

	var proc *ps.Process
	for _, p := range procs {
		if strings.Contains(p.Executable, "java") {
			proc = p
		}
	}

	if proc == nil {
		return 0, nil, fmt.Errorf("Failed to find a running minecraft process")
	}

	pid := proc.PID
	return pid, proc, nil
}

// FindServer returns a server object from a found serve
func FindServer() (*Server, error) {
	_, proc, err := GetPID()
	if err != nil {
		return nil, err
	}

	return &Server{
		PID: proc.PID,
		Dir: proc.CWD,
	}, nil
}

// SendCommand sends a command to the server
func (s *Server) SendCommand(msg string) error {
	proc := fmt.Sprintf("/proc/%d/fd/0", s.PID)
	f, err := os.OpenFile(proc, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	_, err = f.WriteString(msg + "\n")
	return err
}
