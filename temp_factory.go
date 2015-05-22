package main

import (
	"io/ioutil"
	"os"
)

type TempFactory struct {
	path string
	files []string
}

// Create a new temporary file factory.
// defer Cleanup() if you want the files removed.
func NewTempFactory(path string) TempFactory {
	if (path == "") {
		path = DefaultTempPath()
	}
	return TempFactory{path: path}
}

// Default temporary file path
func DefaultTempPath() string {
	return ""
}

// Create a temp file with given value. Returns the path.
func (tf *TempFactory) Push(value string) (string, error) {
	f, err := ioutil.TempFile(tf.path, "cauldron")
	if err != nil {
		return "", err
	}
	defer f.Close()

	f.Write([]byte(value))
	name := f.Name()
	tf.files = append(tf.files, name)
	return name, nil
}

// Remove the temporary files created with this factory.
func (tf *TempFactory) Cleanup() {
	for _, file := range tf.files {
		os.Remove(file)
	}
	tf = nil
}
