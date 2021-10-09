package confer

// Load get data, and assign to source
//
// source must a pointer to struct
func Load(source interface{}, options ...ImplOption) error {
	r := &Impl{}
	for _, v := range options {
		v(r)
	}
	return r.load(source)
}

type Impl struct {
	loader   []Loader
	transfer []Transfer
}

func (r *Impl) load(source interface{}) error {
	// vv := reflect.ValueOf(source)
	// vt := reflect.TypeOf(source)
	// if vv.Kind() != reflect.Ptr {
	// 	return fmt.Errorf("source need to be a pointer to a struct")
	// }
	// vv = vv.Elem()
	// vt = vt.Elem()
	// if vv.Kind() != reflect.Struct {
	// 	return fmt.Errorf("source need to be a pointer to a struct")
	// }
	// for i := 0; i < vv.NumField(); i++ {
	// 	fv := vv.Field(i)
	// 	ft := vt.Field(i)
	// 	tag, ok := ft.Tag.Lookup(tagName)
	// 	if !ok {
	// 		continue // TODO return error ?
	// 	}
	//
	// 	fmt.Println(fv.Interface())
	// }
	return nil
}

type ImplOption func(r *Impl)

func WithLoader(loaders ...Loader) ImplOption {
	return func(r *Impl) {
		r.loader = append(r.loader, loaders...)
	}
}

func WithTransfer(transfers ...Transfer) ImplOption {
	return func(r *Impl) {
		r.transfer = append(r.transfer, transfers...)
	}
}

type Loader interface {
	Load(args []string) (string, error)
}

type Transfer interface{}

const tagName = "conf"
