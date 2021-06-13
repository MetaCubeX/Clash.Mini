package maybe

type maybe struct {
	object 	interface{}
	result  interface{}
	isNil	bool
}

func Maybe() *maybe {
	return &maybe{}
}

func (m *maybe) IfOk(doFunc func(o interface{})) {
	if m.isNil {
		return
	}
	doFunc(m.object)
}

func (m *maybe) IfOkI(doFunc func(o interface{}) interface{}) interface{} {
	if m.isNil {
		return nil
	}
	return doFunc(m.object)
}

func (m *maybe) IfOkString(doFunc func(o interface{}) string) string {
	if m.isNil {
		return ""
	}
	return doFunc(m.object)
}

func (m *maybe) OfNullable(o interface{}) *maybe {
	m.object = o
	if o == nil {
		m.isNil = true
	}
	return m
}
