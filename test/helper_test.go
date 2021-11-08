package test

import (
	"testing"
)

func TestCompareString(t *testing.T) {
	if !CompareString("123", "123") {
		t.Fail()
	}
	if CompareString("abc", "def") {
		t.Fail()
	}
}

func TestCompareBytes(t *testing.T) {
	if !CompareBytes([]byte("123"), []byte("123")) {
		t.Fail()
	}
	if CompareBytes([]byte("abc"), []byte("def")) {
		t.Fail()
	}
}

func TestCompareJSON(t *testing.T) {
	type JsonTest struct {
		Id string `yaml:"id"`
		Name string `yaml:"name"`
	}
	a := JsonTest{Id: "123", Name: "abc"}
	b := JsonTest{Id: "123", Name: "abc"}
	ok, err := CompareJSON(a, b)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Fail()
	}

	a = JsonTest{Id: "000", Name: "abc"}
	b = JsonTest{Id: "111", Name: "abc"}
	ok, err = CompareJSON(a, b)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Fail()
	}
}

func TestCompareYaml(t *testing.T) {
	a := []byte(`
info:
  link:
    - http://aaa
    - http://bbb
  num:
    - 123
    - 456
`)
	b := []byte(`
info:
  link:
    - http://aaa
    - http://bbb
  num:
    - 123
    - 456
`)
	ok, err := CompareYaml(a, b)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Fail()
	}

	a = []byte(`
info:
  link:
    - http://aaa
    - http://bbb
  num:
    - 000
    - 111
`)
	b = []byte(`
info:
  link:
    - http://aaa
    - http://bbb
  num:
    - 123
    - 456
`)
	ok, err = CompareYaml(a, b)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Fail()
	}
}