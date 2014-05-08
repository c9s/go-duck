package duck

import "reflect"

const (
	String = iota
	Boolean
	Float
	Integer
)

type DuckType struct {
	Value    interface{}
	Type     int
	RefType  reflect.Type
	RefValue reflect.Value
}

func NewString(val string) *DuckType {
	d := DuckType{}
	d.Value = val
	d.RefType = reflect.TypeOf(val)
	d.RefValue = reflect.ValueOf(val)
	return &d
}

func New(val interface{}) *DuckType {
	d := DuckType{}
	d.Value = val
	d.RefType = reflect.TypeOf(val)
	d.RefValue = reflect.ValueOf(val)
	return &d
}

func (self *DuckType) AsString() string {
	return ""
}

func (self *DuckType) AsInt() int {
	return 0
}

func (self *DuckType) AsBool() bool {
	return true
}

func (self *DuckType) AsFloat() float64 {
	return 3.1415
}
