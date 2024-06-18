package series

import (
	"fmt"
	"math"
	"strconv"
)

type float32Element struct {
	e   float32
	nan bool
}

// force float32Element struct to implement Element interface
var _ Element = (*float32Element)(nil)

func (e *float32Element) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case string:
		if val == "NaN" {
			e.nan = true
			return
		}
		f, err := strconv.ParseFloat(value.(string), 32)
		if err != nil {
			e.nan = true
			return
		}
		e.e = float32(f)
	case int:
		e.e = float32(val)
	case uint8:
		e.e = float32(val)
	case int64:
		e.e = float32(val)
	case uint16:
		e.e = float32(val)
	case uint32:
		e.e = float32(val)
	case uint64:
		e.e = float32(val)
	case float64:
		e.e = float32(val)
	case float32:
		e.e = float32(val)
	case bool:
		b := val
		if b {
			e.e = 1
		} else {
			e.e = 0
		}
	case Element:
		e.e = val.Float32()
	default:
		e.nan = true
		return
	}
}

func (e float32Element) Copy() Element {
	if e.IsNA() {
		return &float32Element{0.0, true}
	}
	return &float32Element{e.e, false}
}

func (e float32Element) IsNA() bool {
	if e.nan || math.IsNaN(float64(e.e)) {
		return true
	}
	return false
}

func (e float32Element) Type() Type {
	return Float32
}

func (e float32Element) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return float32(e.e)
}

func (e float32Element) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprintf("%f", e.e)
}

func (e float32Element) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	f := e.e
	if math.IsInf(float64(f), 1) || math.IsInf(float64(f), -1) {
		return 0, fmt.Errorf("can't convert Inf to int")
	}
	if f != f { // NaN always returns false for equality
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return int(f), nil
}

func (e float32Element) Uint8() (uint8, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint8")
	}
	f := e.e
	if math.IsInf(float64(f), 1) || math.IsInf(float64(f), -1) {
		return 0, fmt.Errorf("can't convert Inf to uint8")
	}
	if f != f { // NaN always returns false for equality
		return 0, fmt.Errorf("can't convert NaN to uint8")
	}
	return uint8(f), nil
}

func (e float32Element) Uint16() (uint16, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint16")
	}
	f := e.e
	if math.IsInf(float64(f), 1) || math.IsInf(float64(f), -1) {
		return 0, fmt.Errorf("can't convert Inf to uint16")
	}
	if f != f { // NaN always returns false for equality
		return 0, fmt.Errorf("can't convert NaN to uint16")
	}
	return uint16(f), nil
}

func (e float32Element) Uint32() (uint32, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint32")
	}
	f := e.e
	if math.IsInf(float64(f), 1) || math.IsInf(float64(f), -1) {
		return 0, fmt.Errorf("can't convert Inf to uint32")
	}
	if f != f { // NaN always returns false for equality
		return 0, fmt.Errorf("can't convert NaN to uint32")
	}
	return uint32(f), nil
}

func (e float32Element) Int64() (int64, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int64")
	}
	f := e.e
	if math.IsInf(float64(f), 1) || math.IsInf(float64(f), -1) {
		return 0, fmt.Errorf("can't convert Inf to int64")
	}
	if f != f { // NaN always returns false for equality
		return 0, fmt.Errorf("can't convert NaN to int64")
	}
	return int64(f), nil
}

func (e float32Element) Uint64() (uint64, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint64")
	}
	f := e.e
	if math.IsInf(float64(f), 1) || math.IsInf(float64(f), -1) {
		return 0, fmt.Errorf("can't convert Inf to uint64")
	}
	if f != f { // NaN always returns false for equality
		return 0, fmt.Errorf("can't convert NaN to uint64")
	}
	return uint64(f), nil
}

func (e float32Element) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	return float64(e.e)
}

func (e float32Element) Float32() float32 {
	if e.IsNA() {
		return float32(math.NaN())
	}
	return e.e
}

func (e float32Element) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	switch e.e {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}
	return false, fmt.Errorf("can't convert Float32 \"%v\" to bool", e.e)
}

func (e float32Element) Eq(elem Element) bool {
	f := elem.Float32()
	if e.IsNA() || f != f {
		return false
	}
	return e.e == f
}

func (e float32Element) Neq(elem Element) bool {
	f := elem.Float32()
	if e.IsNA() || f != f {
		return false
	}
	return e.e != f
}

func (e float32Element) Less(elem Element) bool {
	f := elem.Float32()
	if e.IsNA() || f != f {
		return false
	}
	return e.e < f
}

func (e float32Element) LessEq(elem Element) bool {
	f := elem.Float32()
	if e.IsNA() || f != f {
		return false
	}
	return e.e <= f
}

func (e float32Element) Greater(elem Element) bool {
	f := elem.Float32()
	if e.IsNA() || f != f {
		return false
	}
	return e.e > f
}

func (e float32Element) GreaterEq(elem Element) bool {
	f := elem.Float32()
	if e.IsNA() || f != f {
		return false
	}
	return e.e >= f
}
