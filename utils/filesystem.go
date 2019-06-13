package utils

import "net/http"

// FileSystem custom file system handler
type FileSystem struct {
	// fs http.FileSystem
	FS http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.FS.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
		return nil, err		
	}

	// if s.IsDir() {		
		// index := strings.TrimSuffix(path, "/") + "/index.html"
		// if _, err := fs.fs.Open(index); err != nil {
		// 	return nil, err
		// }
	// }

	return f, nil
}

