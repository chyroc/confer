package extractor

type File struct{}

func NewFile() *File {
	return &File{}
}

func (r *File) Name() string {
	return "file"
}

func (r *File) Extract(args []string) (string, error) {
	panic("implement me")
}
