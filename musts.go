package decimal

import "fmt"

// MustAdd is like [Add] but panics if computing error.
func (d Decimal) MustAdd(e Decimal) Decimal {
	f, err := d.Add(e)
	if err != nil {
		panic(fmt.Sprintf("MustAdd(%v) failed: %v", d, err))
	}
	return f
}

// MustSub is like [Sub] but panics if computing error.
func (d Decimal) MustSub(e Decimal) Decimal {
	f, err := d.Sub(e)
	if err != nil {
		panic(fmt.Sprintf("MustSub(%v) failed: %v", d, err))
	}
	return f
}

// MustMul is like [Mul] but panics if computing error.
func (d Decimal) MustMul(e Decimal) Decimal {
	f, err := d.Mul(e)
	if err != nil {
		panic(fmt.Sprintf("MustMul(%v) failed: %v", d, err))
	}
	return f
}

// MustQuo is like [Quo] but panics if computing error.
func (d Decimal) MustQuo(e Decimal) Decimal {
	f, err := d.Quo(e)
	if err != nil {
		panic(fmt.Sprintf("MustQuo(%v) failed: %v", d, err))
	}
	return f
}

// MustLog is like [Log] but panics if computing error.
func (d Decimal) MustLog(e Decimal) Decimal {
	f, err := e.Log()
	if err != nil {
		panic(fmt.Sprintf("MustLog(%v) failed: %v", d, err))
	}
	return f
}

// MustFloat64 is like [Float64] but panics if computing error.
func (d Decimal) MustFloat64() float64 {
	f, ok := d.Float64()
	if !ok {
		panic(fmt.Sprintf("MustFloat64(%v) failed", d))
	}
	return f
}
