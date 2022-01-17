package domain

type Video struct {
	Duration string
	file     File
}

func (v *Video) FileName() string {
	return v.file.Name
}

func (v *Video) Size() uint64 {
	return v.file.Size
}

func NewVideo(filename string, size uint64, duration string) *Video {
	return &Video{
		Duration: duration,
		file: File{
			Name: filename,
			Size: size,
		},
	}
}

type File struct {
	Name string
	Size uint64
}

func (f File) NewFile(name string, size uint64) *File {
	return &File{
		Name: name,
		Size: size,
	}
}
