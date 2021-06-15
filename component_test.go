package golevel7

import "testing"

func TestCompParse(t *testing.T) {
	val := []rune("v1&v2&v3&&v5")
	seps := NewDelimeters()
	cmp := &Component{Value: val}
	cmp.parse(seps)
	if len(cmp.SubComponents) != 5 {
		t.Errorf("Expected 5 subcomponents got %d\n", len(cmp.SubComponents))
	}
}

func TestCompSet(t *testing.T) {
	seps := NewDelimeters()
	loc := "ZZZ.1.0.5"
	l := NewLocation(loc)
	cmp := &Component{}
	err := cmp.Set(l, "TEST", seps)
	if err != nil {
		t.Error(err)
	}
	if len(cmp.SubComponents) != 6 {
		t.Fatalf("Expected 6 got %d\n", len(cmp.SubComponents))
	}
	if string(cmp.SubComponents[5].Value) != "TEST" {
		t.Errorf("Expected TEST got %s\n", string(cmp.SubComponents[5].Value))
	}
}
