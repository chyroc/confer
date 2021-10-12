package extractor

type Inherit struct{}

func (r *Inherit) Name() string {
	panic("implement me")
}

func (r *Inherit) Extract(args []string) (string, error) {
	panic("implement me")
}
