package duck

import "reflect"
import "strconv"

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

func StringToInt(str string) int {
	val, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return 0
	}
	return int(val)
}

func StringToInt64(str string) int64 {
	val, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return 0
	}
	return int64(val)
}

func StringToFloat(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return val
}

func StringToBool(str string) bool {
	val, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return val
}
