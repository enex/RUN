package token

// FileSet holds all the files for the source code
type FileSet struct {
	base  int
	files []*File
}

// NewFileSet creates a new FileSet object
func NewFileSet() *FileSet {
	return &FileSet{base: 1}
}

// Add appends a new file to the fileset
// name: Path to the file
// src: Content of the file
func (fs *FileSet) Add(name, src string) *File {
	f := NewFile(name, fs.base, len(src))
	fs.files = append(fs.files, f)
	fs.base += len(src)
	return f
}

// Position returns the row and column position of the given Pos p
func (fs *FileSet) Position(p Pos) Position {
	var pos Position
	if !p.Valid() {
		panic("invalid position")
	}
	for _, f := range fs.files {
		if p >= Pos(f.Base()) && p < Pos(f.Base()+f.Size()) {
			pos = f.Position(p)
		}
	}
	return pos
}
