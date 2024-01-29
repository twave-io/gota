package series

import (
	"fmt"
	"math"
	"strconv"
)

type int8Element struct {
	e   int8
	nan bool
}

// force int8Element struct to implement Element interface
var _ Element = (*int8Element)(nil)

func (e *int8Element) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case string:
		if val == "NaN" {
			e.nan = true
			return
		}
		// i, err := strconv.Atoi(value.(string))
		i, err := strconv.ParseInt(value.(string), 10, 8)
		if err != nil {
			e.nan = true
			return
		}
		e.e = int8(i)
	case int:
		e.e = int8(val)
	case float64:
		f := val
		if math.IsNaN(f) ||
			math.IsInf(f, 0) ||
			math.IsInf(f, 1) {
			e.nan = true
			return
		}
		e.e = int8(f)
	case bool:
		b := val
		if b {
			e.e = 1
		} else {
			e.e = 0
		}
	case Element:
		v, err := val.Int8()
		if err != nil {
			e.nan = true
			return
		}
		e.e = v
	default:
		e.nan = true
		return
	}
}

func (e int8Element) Copy() Element {
	if e.IsNA() {
		return &int8Element{0, true}
	}
	return &int8Element{e.e, false}
}

func (e int8Element) IsNA() bool {
	return e.nan
}

func (e int8Element) Type() Type {
	return Int
}

func (e int8Element) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return int(e.e)
}

func (e int8Element) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprint(e.e)
}

func (e int8Element) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return int(e.e), nil
}

// TODO: Testing
func (e int8Element) Int8() (int8, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int8")
	}
	return int8(e.e), nil
}
func (e int8Element) Int32() (int32, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int32")
	}
	return int32(e.e), nil
}

func (e int8Element) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	return float64(e.e)
}

func (e int8Element) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	switch e.e {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}
	return false, fmt.Errorf("can't convert Int \"%v\" to bool", e.e)
}

func (e int8Element) Eq(elem Element) bool {
	i, err := elem.Int8()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e == i
}

func (e int8Element) Neq(elem Element) bool {
	i, err := elem.Int8()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e != i
}

func (e int8Element) Less(elem Element) bool {
	i, err := elem.Int8()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e < i
}

func (e int8Element) LessEq(elem Element) bool {
	i, err := elem.Int8()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e <= i
}

func (e int8Element) Greater(elem Element) bool {
	i, err := elem.Int8()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e > i
}

func (e int8Element) GreaterEq(elem Element) bool {
	i, err := elem.Int8()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e >= i
}
