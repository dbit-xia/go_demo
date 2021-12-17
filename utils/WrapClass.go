package utils

type WrapClass struct {
	fn   interface{}
	args []interface{}
}

func (w *WrapClass) Bind(fn interface{}, args ...interface{}) {
	w.fn = fn
	w.args = args
	return
}

func (w *WrapClass) Call() (interface{}, error) {
	return Call(w.fn, w.args...)
}

func Wrap(fn interface{}, args ...interface{}) func() (interface{}, error) {
	var wrapObject WrapClass
	wrapObject.Bind(fn, args...)
	return wrapObject.Call
}
