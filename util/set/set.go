package set

type void struct{}

type Set struct {
	m 			*map[interface{}]void
	hashFunc 	func(interface{}) string
}

func NewSet(vs... string) *Set {
	s := &Set{}
	s.Clear()
	s.hashFunc = func(i interface{}) string {
		return i.(string)
	}
	if len(vs) > 0 {
		s.add(vs...)
	}
	return s
}

func NewSetWithFunc(hashFunc func(interface{}) string, vs... interface{}) *Set {
	s := &Set{}
	s.Clear()
	s.hashFunc = hashFunc
	if len(vs) > 0 {
		s.Add(vs...)
	}
	return s
}

func (s *Set) Contains(v interface{}) bool {
	_, exists := (*(s.m))[s.hashFunc(v)]
	return exists
}

func (s *Set) Add(vs... interface{}) {
	var hash string
	for _, v := range vs {
		hash = s.hashFunc(v)
		if _, exists := (*(s.m))[hash]; !exists {
			continue
		}
		(*(s.m))[hash] = void{}
	}
}

func (s *Set) add(vs... string) {
	for _, v := range vs {
		(*(s.m))[v] = void{}
	}
}

func (s *Set) Delete(v interface{}) {
	delete(*(s.m), v)
}

func (s *Set) Clear() {
	m := make(map[interface{}]void)
	s.m = &m
}

// exampleFull
//func exampleFull() {
//	s := NewSetWithFunc(func(i interface{}) string {
//		return i.(string)
//	}, "1", "2", "3")
//	fmt.Println(s.Contains("1"))
//	fmt.Println(s.Contains("2"))
//	fmt.Println(s.Contains("5"))
//}
//
// exampleSimple
//func exampleSimple() {
//	s := NewSet("1", "2", "3")
//	fmt.Println(s.Contains("1"))
//	fmt.Println(s.Contains("2"))
//	fmt.Println(s.Contains("5"))
//}

