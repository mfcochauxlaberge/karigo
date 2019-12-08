package gold

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// NewRunner returns a *Runner.
//
// dir defines the Runner's directory, which is where all the files will be
// stored. A good option is "testdata".
//
// The update-golden-files flag is parsed and its value will be used by the
// runner. To override this setting, set the Update field on the returned
// *Runner.
//
// NewRunner calls flag.Parse.
func NewRunner(dir string) *Runner {
	update := flag.Bool("update-golden-files", false, "update the golden files")
	// TODO Should this be called here? Maybe let
	// the user call it when appropriate.
	flag.Parse()

	return &Runner{
		Update:    *update,
		Directory: dir,
	}
}

// A Runner stores context information when handling the generated outputs and
// the outputs stored in files.
type Runner struct {
	// Update reports whether or not the runner
	// must run in update mode. If it is the
	// case, files are created instead of being
	// compared.
	Update bool

	// The directory in which the golden files
	// will be stored. If empty, "testdata" is
	// used by default.
	Directory string

	// Filters holds a list of filters to apply
	// on the contents given to Test. They are
	// applied in the order they appear in the
	// slice.
	Filters []Filter
}

// Prepares prepares the runner's directory if in update mode by deleting it
// (including its content) and recreating it.
func (r *Runner) Prepare() error {
	// In update mode, everything is recreated
	// from scratch.
	if r.Update {
		err := os.RemoveAll(r.Directory)
		if err != nil {
			return err
		}

		err = os.Mkdir(filepath.Join(r.Directory), os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// Test takes a path and some contents to check.
//
// An error is returned if the comparison failed or any error that occurred
// during the process. Only nil means the compared contents where the same.
//
// In update mode, the file is created instead of being compared.
func (r *Runner) Test(path string, content []byte) error {
	path = filepath.Join(r.Directory, path)

	for _, filter := range r.Filters {
		content = filter(content)
	}

	if r.Update {
		// Make sure the necessary directories exist.
		dir, _ := filepath.Split(path)

		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("gold: could not create directory: %w", err)
		}

		err = ioutil.WriteFile(path, content, 0644)
		if err != nil {
			return fmt.Errorf("gold: could not write file: %w", err)
		}
	} else {
		// Compare the file with the given content.
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return nil
		}

		if !bytes.Equal(file, content) {
			return ComparisonError{}
		}
	}

	return nil
}
