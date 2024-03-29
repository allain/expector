package expector

import (
	"errors"
	"testing"
)

func TestExpectCanBeCreated(t *testing.T) {
	t2 := &testing.T{}
	Expect := New(t2, Options{BailOnFail: false})
	Expect(10)
}

func TestExpectToEqualInts(t *testing.T) {
	Expect := New(&testing.T{}, Options{BailOnFail: false})
	if Expect(10).ToEqual(10).Failed() {
		t.Error("should not have failed")
	}
	if !Expect(10).ToEqual(20).Failed() {
		t.Error("should have failed")
	}
}

func TestExpectToEqualSlice(t *testing.T) {
	Expect := New(&testing.T{}, Options{BailOnFail: false})

	if Expect([]int{10}).ToEqual([]int{10}).Failed() {
		t.Error("should not have failed")
	}
	if Expect([]int{20}).Not().ToEqual([]int{20, 20}).Failed() {
		t.Error("should have failed")
	}
	if !Expect([]int{20}).ToEqual([]int{30}).Failed() {
		t.Error("should have failed")
	}
}

func TestExpectNotToEqual(t *testing.T) {
	Expect := New(&testing.T{}, Options{BailOnFail: false})

	if Expect(10).Not().ToEqual(20).Failed() {
		t.Error("should not have failed")
	}
	if !Expect(10).Not().ToEqual(10).Failed() {
		t.Error("should not have failed")
	}
}

func TestExpectToBeNil(t *testing.T) {
	Expect := New(&testing.T{}, Options{BailOnFail: false})
	if Expect(nil).ToBeNil().Failed() {
		t.Error("should not have failed")
	}
	
	Expect = New(&testing.T{}, Options{BailOnFail: false})
	if !Expect(10).ToBeNil().Failed() {
		t.Error("should have failed")
	}
	
	Expect = New(&testing.T{}, Options{BailOnFail: false})
	if Expect(10).Not().ToBeNil().Failed() {
		t.Error("should not have failed")
	}
	
	Expect = New(&testing.T{}, Options{BailOnFail: false})
	if !Expect(nil).Not().ToBeNil().Failed() {
		t.Error("should have failed")
	}
}

func TestExpectToBeTrue(t *testing.T) {
  Expect := New(&testing.T{}, Options{BailOnFail: false})
  if Expect(true).ToBeTrue().Failed() {
	  t.Error("should not fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if !Expect(true).Not().ToBeTrue().Failed() {
	  t.Error("should fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if !Expect(nil).ToBeTrue().Failed() {
	  t.Error("should fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if Expect(nil).Not().ToBeTrue().Failed() {
	  t.Error("should not fail")
  }
}

func TestExpectToMatch(t *testing.T) {
  Expect := New(&testing.T{}, Options{BailOnFail: false})
  if Expect("hello").ToMatch("hello").Failed() {
	  t.Error("should not fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if !Expect("hello").Not().ToMatch("hello").Failed() {
	  t.Error("should fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if !Expect("hello").ToMatch("howdy").Failed() {
	  t.Error("should fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if !Expect("hello").ToMatch("(").Failed() {
	  t.Error("should fail because of invalid regex")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if !Expect([]string{"hello"}).ToMatch("hello").Failed() {
	  t.Error("should fail when matching non-string targets")
  }
}

func TestExpectToBeError(t *testing.T) {
  Expect := New(&testing.T{}, Options{BailOnFail: false})
  if Expect(errors.New("testing error")).ToBeError().Failed() {
	  t.Error("should not fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if Expect(nil).Not().ToBeError().Failed() {
	  t.Error("should not fail")
  }
}

func TestExpectToMatchError(t *testing.T) {
  Expect := New(&testing.T{}, Options{BailOnFail: false})
  if Expect(errors.New("testing error")).ToMatchError("testing error").Failed() {
	  t.Error("should not fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if Expect(errors.New("testing")).Not().ToMatchError("pass").Failed() {
	  t.Error("should not fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if !Expect(nil).ToMatchError("(").Failed() {
	  t.Error("should not fail")
  }
  
  Expect = New(&testing.T{}, Options{BailOnFail: false})
  if Expect("pass").Not().ToMatchError("pass").Failed() {
	  t.Error("should fail")
  }
}