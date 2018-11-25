package ps

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mitchellh/go-ps"
)

// Process is a UnixProcess
type Process struct {
	PID        int
	CMDLine    string
	Executable string
	CWD        string
}

// Find all processes running on a system.
func Find() ([]*Process, error) {
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	pids := make([]*Process, 0)
	for _, file := range files {
		// check if an int
		if _, err := strconv.Atoi(file.Name()); err != nil {
			continue
		}

		pid, err := strconv.Atoi(file.Name())
		if err != nil {
			fmt.Printf("failed to convert PID: %s to a int\n", file.Name())
		}
		pidPath := filepath.Join("/proc/" + file.Name())
		b, err := ioutil.ReadFile(filepath.Join(pidPath, "cmdline"))
		if err != nil {
			fmt.Printf("%d: failed to build cmdline\n", pid)
			continue
		}

		cmdline := string(b)
		p, err := ps.FindProcess(pid)
		if err != nil {
			fmt.Printf("%d: failed to get executable\n", pid)
			continue
		}

		c, err := os.Readlink(filepath.Join(pidPath, "cwd"))
		if err != nil {
			//fmt.Printf("%d: failed to read cwd\n", pid)
			continue
		}

		pids = append(pids, &Process{
			PID:        pid,
			CMDLine:    cmdline,
			Executable: p.Executable(),
			CWD:        c,
		})
	}

	return pids, nil
}
