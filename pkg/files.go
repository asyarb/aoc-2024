package files

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func OpenRelative(relativePath string) *os.File {
	_, callingFile, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatalf("Unable to find file: %s", callingFile)
	}

	dir := filepath.Dir(callingFile)
	path := filepath.Join(dir, relativePath)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
