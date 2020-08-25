package utils

type LazyLoadWrapper struct {
	flag     bool
	loadFunc func() interface{}
	data     interface{}
}

func NewLazyLoadWrapper(fn func() interface{}) *LazyLoadWrapper {
	return &LazyLoadWrapper{
		flag:     false,
		loadFunc: fn,
		data:     nil,
	}
}

func (l *LazyLoadWrapper) IsLoaded() bool {
	return l.flag
}

func (l *LazyLoadWrapper) Load() {
	l.data = l.loadFunc()
}

func (l *LazyLoadWrapper) Get() interface{} {
	if !l.IsLoaded() {
		l.Load()
	}
	return l.data
}
