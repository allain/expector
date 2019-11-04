package expector

import (
	"regexp"
	"reflect"
	"testing"
)

// Options used 
type Options struct {
	BailOnFail bool
}

// New returns a new instance of Expect that wraps the provided testing object
func New(t *testing.T, options ...Options) func(interface{}) Expect {
	opts := Options{}
	if len(options) > 0 {
		opts.BailOnFail = options[0].BailOnFail
	} else {
		opts.BailOnFail = true
	}

	return func(target interface{}) Expect {
		expect := Expect{t: t, target: target}
		expect.bailOnFail = opts.BailOnFail
		return expect
	}
}

// Expect is a utility for asserting things on values. It is heavily inspired by Facebook's Jest
type Expect struct {
	t       *testing.T
	inverse bool
	target  interface{}
	bailOnFail bool
}

func (e *Expect) result(t *testing.T) *testing.T {
  if t.Failed() && e.bailOnFail {
	  t.FailNow()
  }
  return t
}

// ToEqual fails if the expected value does not match the target
func (e Expect) ToEqual(val interface{}) *testing.T {
	equal := reflect.DeepEqual(e.target, val)
	if e.inverse && equal {
		e.t.Errorf("expected %v not to equal %v", e.target, val)
	} else if !e.inverse && !equal {
		e.t.Errorf("expected %v to equal %v", e.target, val)
	}
	return e.result(e.t)
}

// ToBeNil fails is the target is not nil
func (e Expect) ToBeNil() *testing.T {
	if e.inverse && e.target == nil {
		e.t.Errorf("expected %v not to to be nil but was", e.target)
	} else if !e.inverse && e.target != nil {
		e.t.Errorf("expected nil but got %v", e.target)
	}
	return e.result(e.t)
}

// ToBeTrue fails if target is not true 
func (e Expect) ToBeTrue() *testing.T {
	if e.inverse && e.target == true{
		e.t.Errorf("expected false but got %v", e.target)
	} else if !e.inverse && e.target != true {
		e.t.Errorf("expected true but got %v", e.target)
	}
	return e.result(e.t)
}

// ToMatch throws if the expect target does not match the regex
func (e Expect) ToMatch(pattern string) *testing.T {
  switch e.target.(type) {
  case string:
	matches, err := regexp.MatchString(pattern, e.target.(string))
	if err != nil {
		e.t.Errorf("regex is invalid %w", err)
	} else if e.inverse && matches {
		e.t.Error("matches regex but shouldn't")
	} else if !e.inverse && !matches {
		e.t.Errorf("%v does not match regex %v", e.target, pattern)
	}
  default:
	e.t.Errorf("regex can only match against string but got %v", e.target)
  }
	return e.result(e.t)
}

func (e Expect) ToBeError() *testing.T {
	err, ok := e.target.(error)
	if e.inverse {
        if ok {
			e.t.Errorf("expected no error but got one %w", err)
		}
	} else if !ok {
		e.t.Errorf("expected error but was %v", e.target)
	}

	return e.result(e.t)

}

// ToMatchError fails test if the target is not a matching error
func (e Expect) ToMatchError(regex string) *testing.T {
	err, ok := e.target.(error)
	if e.inverse {
        if ok {
			matches, patternErr := regexp.MatchString(regex, err.Error()) 
			if patternErr != nil {
				e.t.Errorf("regex is invalid: %s", regex)
			} else if matches {
				e.t.Errorf("error matches regex but shouldn't")
			}
		}
	} else if !ok {
		e.t.Errorf("expected error but was %v", e.target)
	} else {
		matches, patternErr := regexp.MatchString(regex, err.Error()) 
		if patternErr != nil {
			e.t.Errorf("regex is invalid: %s", regex)
		} else if !matches {
			e.t.Errorf("error does not match provided regex: %s", err.Error())
		}
	}

	return e.result(e.t)
}

// Not returns an expect that negates the matcher
func (e Expect) Not() Expect {
	return Expect{t: e.t, inverse: !e.inverse, target: e.target}
}