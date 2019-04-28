package karigo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"plugin"
)

func compile(code []byte) (Tx, error) {
	if len(code) == 0 {
		return nil, errors.New("karigo: can't compile empty code")
	}

	// Create temporary directory
	dir, err := ioutil.TempDir("", "karigo_pkg_")
	if err != nil {
		return nil, err
	}

	// Create temporary file
	file, err := ioutil.TempFile(dir, "karigo_func_*.go")
	if err != nil {
		return nil, err
	}
	stat, _ := file.Stat()

	// Write code to temporary file
	_, err = file.Write(code)
	if err != nil {
		return nil, err
	}

	cmd := exec.Cmd{
		Path: "/home/mfcl/.gvm/gos/go1.12.3/bin/go",
		Args: []string{
			"go",
			"build",
			"-buildmode",
			"plugin",
			"-o",
			"func.so",
			stat.Name(),
		},
		Dir: dir,
	}
	_, err = cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("karigo: compilation error: %s", err.Error())
	}

	// Load plugin
	plug, err := plugin.Open(filepath.Join(dir, "func.so"))
	if err != nil {
		return nil, err
	}

	// Lookup Tx symbole
	testFunc, err := plug.Lookup("Tx")
	if err != nil {
		return nil, errors.New("karigo: no Tx function compiled")
	}

	tx, ok := testFunc.(Tx)
	if !ok {
		return nil, errors.New("karigo: Tx symbol is not a Tx")
	}

	return tx, nil
}
