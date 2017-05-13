package vm

import "testing"

func TestEvalNil(t *testing.T) {
	input := `nil`

	evaluated := testEval(t, input)

	if _, ok := evaluated.(*Null); !ok {
		t.Fatalf("Expect result to be Null. got=%T", evaluated)
	}
}
