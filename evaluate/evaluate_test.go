package lisp

import (
	"testing"

	lisp "lisp/parse"
)

func TestEvaluate(t *testing.T) {
	t.Run("Evaluate 1 we should get back 1", func(t *testing.T) {
		want := lisp.NumberAtom{1}
		error, got := evaluateStr("1")

		if error != nil {
			t.Errorf("Recieved error: %v", error)
		}

		if want != got {
			t.Errorf("Expected %v got %v", want, got)
		}
	})

	t.Run("Evaluate (+ 1 2 3) we get 6", func(t *testing.T) {
		want := lisp.NumberAtom{6}
		error, got := evaluateStr("(+ 1 2 3)")

		if error != nil {
			t.Errorf("Recieved error: %v", error)
		}

		if want != got {
			t.Errorf("Expected %v got %v", want, got)
		}
	})

	t.Run("Evaluate (+ 1 (+ 2 3)) we get 6", func(t *testing.T) {
	})
}
